package db

import (
	"basel2053/ps-board/ps"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (pg *Postgres) FindRecords(ctx context.Context, filter interface{}) ([]ps.Record, error) {
	query := "SELECT * FROM record"
	rows, _ := pg.db.Query(ctx, query)
	// NOTE: with many returned records may take alot of memory so you may switch
	// to using cursor way of fetching data row by row and sending to channel
	records, err := pgx.CollectRows(rows, pgx.RowToStructByName[ps.Record])
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (pg *Postgres) FindRecordById(ctx context.Context, id string) (*ps.Record, error) {
	query := "SELECT * FROM record WHERE id = @id"
	args := pgx.NamedArgs{
		"id": id,
	}
	rows, _ := pg.db.Query(ctx, query, args)
	record, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[ps.Record])
	if err != nil {
		return nil, fmt.Errorf("unable to scan row: %w", err)
	}
	return record, nil
}

func (pg *Postgres) CreateRecord(ctx context.Context, record ps.Record) error {
	query := "INSERT INTO record (title) VALUES (@title)"
	args := pgx.NamedArgs{
		"title": record.Title,
	}
	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert record: %w", err)
	}
	return nil
}

func (pg *Postgres) RemoveRecord(ctx context.Context, id string) error {
	query := "DELETE FROM record WHERE id = @id"
	args := pgx.NamedArgs{
		"id": id,
	}
	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to remove record: %w", err)
	}

	return nil
}
