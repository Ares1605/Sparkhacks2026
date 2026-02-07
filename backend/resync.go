package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"x/db"
)

const amazonProviderID = "1"

type syncScriptOrder struct {
	OrderNumber     string           `json:"order_number"`
	OrderPlacedDate string           `json:"order_placed_date"`
	Items           []syncScriptItem `json:"items"`
}

type syncScriptItem struct {
	Price any    `json:"price"`
	Title string `json:"title"`
}

var priceCleaner = regexp.MustCompile(`[^0-9.-]`)

func resync_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cmd := exec.CommandContext(r.Context(), python_executable, "sync-amazon-data.py")
	cmd.Dir = "scripts"

	output, err := cmd.CombinedOutput()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := "Resync failed"
		trimmed := strings.TrimSpace(string(output))
		if trimmed != "" {
			msg = msg + ": " + trimTo(trimmed, 1000)
		}
		json.NewEncoder(w).Encode(&ErrorResponse{Err: msg})
		return
	}

	importedOrders, importedItems, err := persistScriptOutput(r.Context(), output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&ErrorResponse{Err: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&OkResponse{
		Data: fmt.Sprintf("Resync completed. Imported %d orders and %d items.", importedOrders, importedItems),
	})
}

func persistScriptOutput(ctx context.Context, output []byte) (importedOrders int, importedItems int, err error) {
	_ = database.DeleteOrdersFromProvider(ctx, amazonProviderID)

	scanner := bufio.NewScanner(bytes.NewReader(output))
	scanner.Buffer(make([]byte, 64*1024), 10*1024*1024)
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var parsed syncScriptOrder
		if err := json.Unmarshal([]byte(line), &parsed); err != nil {
			return importedOrders, importedItems, fmt.Errorf("script output line %d is not valid JSON: %w", lineNo, err)
		}

		if err := persistOrder(ctx, parsed); err != nil {
			return importedOrders, importedItems, fmt.Errorf("failed importing order on line %d: %w", lineNo, err)
		}

		importedOrders++
		importedItems += len(parsed.Items)
	}
	if err := scanner.Err(); err != nil {
		return importedOrders, importedItems, fmt.Errorf("failed reading script output: %w", err)
	}

	return importedOrders, importedItems, nil
}

func persistOrder(ctx context.Context, parsed syncScriptOrder) error {
	orderDate := parsed.OrderPlacedDate
	if split := strings.SplitN(orderDate, "T", 2); len(split) > 0 {
		orderDate = split[0]
	}

	for idx, item := range parsed.Items {
		price, err := parsePrice(item.Price)
		if err != nil {
			return fmt.Errorf("order %s item %d has invalid price: %w", parsed.OrderNumber, idx, err)
		}

		order := db.Order{
			Id:         buildOrderID(parsed.OrderNumber, idx, item.Title),
			ProviderId: amazonProviderID,
			Name:       strings.TrimSpace(item.Title),
			Price:      price,
		}
		if order.Name == "" {
			order.Name = parsed.OrderNumber
		}
		if err := order.OrderDate.From(orderDate); err != nil {
			return fmt.Errorf("order %s has invalid order date %q: %w", parsed.OrderNumber, parsed.OrderPlacedDate, err)
		}

		if err := database.InsertOrder(ctx, order); err != nil {
			if strings.Contains(err.Error(), "UNIQUE constraint failed: Orders.Id") {
				continue
			}
			return err
		}
	}

	return nil
}

func parsePrice(raw any) (float32, error) {
	switch v := raw.(type) {
	case nil:
		return 0, nil
	case float64:
		return float32(v), nil
	case string:
		cleaned := priceCleaner.ReplaceAllString(strings.TrimSpace(v), "")
		if cleaned == "" || cleaned == "-" || cleaned == "." || cleaned == "-." {
			return 0, nil
		}
		parsed, err := strconv.ParseFloat(cleaned, 32)
		if err != nil {
			return 0, err
		}
		return float32(parsed), nil
	default:
		return 0, fmt.Errorf("unsupported type %T", raw)
	}
}

func buildOrderID(orderNumber string, itemIndex int, itemTitle string) int {
	key := fmt.Sprintf("%s:%d:%s", orderNumber, itemIndex, itemTitle)
	return int(crc32.ChecksumIEEE([]byte(key)))
}

func trimTo(data string, max int) string {
	if len(data) <= max {
		return data
	}
	return data[:max] + "..."
}
