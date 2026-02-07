package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Database struct {
	sqldb *sql.DB
}

func (db Database) GetAllProviders(ctx context.Context) ([]Provider, error) {
	rows, err := db.sqldb.QueryContext(ctx, sqlGetAllProviders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var providers []Provider

	for rows.Next() {
		var p Provider
		if err := rows.Scan(&p.Id, &p.Name, &p.LastSync, &p.Username, &p.Password); err != nil {
			return nil, err
		}

		providers = append(providers, p)
	}

	return providers, nil
}

func (db Database) GetProviderStatusByID(ctx context.Context, providerID int) (ProviderStatus, bool, error) {
	row := db.sqldb.QueryRowContext(ctx, sqlGetProviderStatusByID, providerID)

	var username sql.NullString
	var lastSync sql.NullString
	if err := row.Scan(&username, &lastSync); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ProviderStatus{}, false, nil
		}
		return ProviderStatus{}, false, err
	}

	status := ProviderStatus{}
	if username.Valid {
		value := strings.TrimSpace(username.String)
		status.Username = &value
	}
	if lastSync.Valid {
		value := strings.TrimSpace(lastSync.String)
		if value != "" {
			status.LastSync = &value
		}
	}

	return status, true, nil
}

func (db Database) GetProviderCredentialsByID(ctx context.Context, providerID int) (ProviderCredentials, bool, error) {
	row := db.sqldb.QueryRowContext(ctx, sqlGetProviderCredentialsByID, providerID)

	var username sql.NullString
	var password sql.NullString
	if err := row.Scan(&username, &password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ProviderCredentials{}, false, nil
		}
		return ProviderCredentials{}, false, err
	}

	return ProviderCredentials{
		Username: strings.TrimSpace(username.String),
		Password: password.String,
	}, true, nil
}

func (db Database) UpsertProviderCredentials(
	ctx context.Context,
	providerID int,
	providerName string,
	username string,
	password string,
) error {
	_, err := db.sqldb.ExecContext(
		ctx,
		sqlUpsertProviderCredentials,
		providerID,
		providerName,
		username,
		password,
	)
	return err
}

func (db Database) ReplaceOrdersForProvider(
	ctx context.Context,
	providerID string,
	providerRecordID int,
	providerName string,
	providerUsername *string,
	lastSync string,
	orders []Order,
) (err error) {
	tx, err := db.sqldb.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	if _, err = tx.ExecContext(ctx, sqlDeleteByProvider, providerID); err != nil {
		return err
	}

	for _, o := range orders {
		if _, err = tx.ExecContext(
			ctx,
			sqlInsertOrder,
			o.Id, o.ProviderId, o.Name, o.Price, o.OrderDate,
		); err != nil {
			return err
		}
	}

	var usernameValue any
	if providerUsername != nil {
		usernameValue = strings.TrimSpace(*providerUsername)
	}

	if _, err = tx.ExecContext(
		ctx,
		sqlUpsertProviderSync,
		providerRecordID,
		providerName,
		lastSync,
		usernameValue,
		providerRecordID,
	); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (db Database) GetAllOrder(ctx context.Context) ([]Order, error) {
	rows, err := db.sqldb.QueryContext(ctx, sqlGetAllOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var o Order
		if err := rows.Scan(&o.Id, &o.ProviderId, &o.Name, &o.Price, &o.OrderDate); err != nil {
			return nil, err
		}

		orders = append(orders, o)
	}

	return orders, nil
}

func (db Database) InsertChatSession(ctx context.Context, sessionUuid string) error {
	_, err := db.sqldb.ExecContext(
		ctx,
		sqlInsertChatSession,
		sessionUuid,
	)
	if err != nil {
		return err
	}

	return nil
}

func (db Database) InsertChatMessage(ctx context.Context, chatMessage ChatMessage) error {
	var roleStr string
	switch chatMessage.Role {
	case UserMessage:
		roleStr = "user"
	case ServerMessage:
		roleStr = "server"
	default:
		return fmt.Errorf("unknown role ID (%d)", chatMessage.Role)
	}

	_, err := db.sqldb.ExecContext(
		ctx,
		sqlInsertChatMessage,
		chatMessage.Message,
		roleStr,
		chatMessage.SessionUuid,
		chatMessage.MessageDate.Format(time.RFC3339Nano),
	)
	if err != nil {
		return err
	}

	return nil
}

func (db Database) GetChatHistory(ctx context.Context, sessionUuid string) ([]DBChatMessage, error) {
	rows, err := db.sqldb.QueryContext(ctx,
		"SELECT Id, Message, Role, MessageDate FROM ChatMessages WHERE SessionUuid=?",
		sessionUuid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	chatHistory := make([]DBChatMessage, 0)
	for rows.Next() {
		var m DBChatMessage
		var roleStr string
		var dateStr string
		if err := rows.Scan(&m.Id, &m.Message, &roleStr, &dateStr); err != nil {
			return nil, err
		}
		switch roleStr {
		case "user":
			m.Role = UserMessage
		case "server":
			m.Role = ServerMessage
		default:
			return nil, fmt.Errorf("unknown role %q", roleStr)
		}
		m.MessageDate, err = time.Parse(time.RFC3339Nano, dateStr)
		if err != nil {
			return nil, err
		}
		chatHistory = append(chatHistory, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return chatHistory, nil
}
