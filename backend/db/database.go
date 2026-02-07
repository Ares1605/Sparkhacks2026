package db

import (
	"context"
	"database/sql"
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
