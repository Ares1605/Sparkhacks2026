package db

import (
	"context"
	"database/sql"
)

type Database struct {
	sqldb *sql.DB
}

func (db Database) GetProviderId(ctx context.Context, name string) error {
	row := db.sqldb.QueryRowContext(ctx, sqlGetProviderId)
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
