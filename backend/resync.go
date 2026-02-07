package main

import (
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
	"time"
	"x/db"
)

const (
	amazonProviderID    = "1"
	amazonProviderRowID = 1
	amazonProviderName  = "amazon"
)

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
	credentials, exists, err := database.GetProviderCredentialsByID(r.Context(), amazonProviderRowID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to load Amazon credentials from database.")
		return
	}

	if !exists || credentials.Username == "" || strings.TrimSpace(credentials.Password) == "" {
		writeError(w, http.StatusBadRequest, "Amazon credentials are not set. Call /set-amazon-credentials first.")
		return
	}

	cmd := exec.CommandContext(
		r.Context(),
		python_executable,
		"sync-amazon-data.py",
		credentials.Username,
		credentials.Password,
	)
	cmd.Dir = "scripts"

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	output, err := cmd.Output()
	if err != nil {
		msg := "Resync failed"
		trimmed := strings.TrimSpace(stderr.String())
		if trimmed == "" {
			trimmed = strings.TrimSpace(string(output))
		}
		if trimmed != "" {
			msg = msg + ": " + trimTo(trimmed, 1000)
		}

		writeError(w, http.StatusInternalServerError, msg)
		return
	}

	importedOrders, importedItems, err := persistScriptOutput(r.Context(), output, credentials.Username)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, fmt.Sprintf("Resync completed. Imported %d orders and %d items.", importedOrders, importedItems))
}

func persistScriptOutput(ctx context.Context, output []byte, providerUsername string) (importedOrders int, importedItems int, err error) {
	parsedOrders, err := parseScriptOutput(output)
	if err != nil {
		return 0, 0, err
	}

	dbOrders, importedItems, err := buildOrders(parsedOrders)
	if err != nil {
		return 0, 0, err
	}

	lastSync := time.Now().UTC().Format(time.RFC3339)

	var username *string
	if raw := strings.TrimSpace(providerUsername); raw != "" {
		username = &raw
	}

	if err := database.ReplaceOrdersForProvider(
		ctx,
		amazonProviderID,
		amazonProviderRowID,
		amazonProviderName,
		username,
		lastSync,
		dbOrders,
	); err != nil {
		return 0, 0, fmt.Errorf("failed writing orders to database: %w", err)
	}

	return len(parsedOrders), importedItems, nil
}

func parseScriptOutput(output []byte) ([]syncScriptOrder, error) {
	trimmedOutput := bytes.TrimSpace(output)
	if len(trimmedOutput) == 0 {
		return nil, fmt.Errorf("script output is empty")
	}
	if trimmedOutput[0] != '[' {
		return nil, fmt.Errorf("script output must be a JSON array; output: %s", trimTo(string(trimmedOutput), 1000))
	}

	var parsedOrders []syncScriptOrder
	if err := json.Unmarshal(trimmedOutput, &parsedOrders); err != nil {
		return nil, fmt.Errorf("script output is not valid JSON array: %w; output: %s", err, trimTo(string(trimmedOutput), 1000))
	}

	return parsedOrders, nil
}

func buildOrders(parsedOrders []syncScriptOrder) ([]db.Order, int, error) {
	orders := make([]db.Order, 0)
	seenIDs := make(map[int]struct{})
	importedItems := 0

	for orderIndex, parsed := range parsedOrders {
		orderDate := parsed.OrderPlacedDate
		if split := strings.SplitN(orderDate, "T", 2); len(split) > 0 {
			orderDate = split[0]
		}

		for idx, item := range parsed.Items {
			price, err := parsePrice(item.Price)
			if err != nil {
				return nil, 0, fmt.Errorf("order %s item %d has invalid price: %w", parsed.OrderNumber, idx, err)
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
				return nil, 0, fmt.Errorf("order %s has invalid order date %q: %w", parsed.OrderNumber, parsed.OrderPlacedDate, err)
			}
			if _, found := seenIDs[order.Id]; found {
				return nil, 0, fmt.Errorf("duplicate generated order id in script payload at order index %d", orderIndex)
			}
			seenIDs[order.Id] = struct{}{}

			orders = append(orders, order)
			importedItems++
		}
	}

	return orders, importedItems, nil
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
