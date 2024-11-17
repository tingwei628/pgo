package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/tingwei628/pgo/webapi/internal/entity"
	"log"
	"os"
)

type DB struct {
	pool *pgxpool.Pool
}

func NewDB(dbname string) (*DB, error) {
	workingDir, _ := os.Getwd() // get root directory
	err := godotenv.Load(fmt.Sprintf("%s/webapi/.env", workingDir))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_HOST"),
		5432,
		dbname)

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to database: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("Error pinging database: %w", err)
	}
	return &DB{pool: pool}, nil
}

// DB implement  Repository Todo
func (db *DB) InsertItem(ctx context.Context, item entity.Item) error {
	// $1, $2 to prevent sql injection
	sql := `
		INSERT INTO todo_items (task, status)
		VALUES ($1, $2)
	`
	_, err := db.pool.Exec(ctx, sql, item.Task, item.Status)
	return err
}

// DB implement  Repository Todo
func (db *DB) GetAllItems(ctx context.Context) ([]entity.Item, error) {
	sql := `
		SELECT task, status
		FROM todo_items
	`
	rows, err := db.pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []entity.Item
	for rows.Next() {
		var item entity.Item
		err := rows.Scan(&item.Task, &item.Status)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil

}

func (db *DB) GetItemsByKeyword(ctx context.Context, keyword string) ([]entity.Item, error) {
	// ILIKE case-insensitive
	sql := `
		SELECT task, status
		FROM todo_items
		WHERE task ILIKE $1
	`
	rows, err := db.pool.Query(ctx, sql, "%"+keyword+"%")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []entity.Item
	for rows.Next() {
		var item entity.Item
		err := rows.Scan(&item.Task, &item.Status)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil

}

func (db *DB) Close() {
	db.pool.Close()
}
