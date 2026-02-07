package db

import (
	"context"
	"database/sql"
	"fmt"
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

func (db Database) GetProviderId(ctx context.Context, name string) error {
	row := db.sqldb.QueryRowContext(ctx, sqlGetProviderId, name)
	if row.Err() != nil {
		return row.Err()
	}

	var o Order
	err := row.Scan(&o.Id, &o.ProviderId, &o.Name, &o.Price, &o.OrderDate)
	if err != nil {
		return err
	}

	return nil
}

func (db Database) DeleteOrdersFromProvider(ctx context.Context, name string) error {
	_, err := db.sqldb.ExecContext(ctx, sqlDeleteByProvider)
	if err != nil {
		return err
	}

	return nil
}

func (db Database) InsertOrder(ctx context.Context, o Order) error {
	_, err := db.sqldb.ExecContext(
		ctx,
		sqlInsertOrder,
		o.Id, o.ProviderId, o.Name, o.Price, o.OrderDate.String(),
	)
	if err != nil {
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
