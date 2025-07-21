package config

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB wraps the pgxpool.Pool connection pool
type DbConnect struct {
	pool *pgxpool.Pool
}

var DbConn *DbConnect

// NewDB creates a new database connection using pgxpool
func InitDB(databaseURL string) error {
	ctx := context.Background()

	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return fmt.Errorf("failed to parse database URL: %w", err)
	}

	config.MaxConns = AppConfig.MaxConns
	config.MinConns = AppConfig.MinConns
	config.MaxConnLifetime = time.Duration(AppConfig.MaxConnLifetime) * time.Second

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Ping the database to verify the connection.
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return fmt.Errorf("failed to ping database: %w", err)
	}

	DbConn = &DbConnect{pool: pool}
	fmt.Println("Database connection established successfully!")

	return nil
}

// Close closes the database connection
func (db *DbConnect) Close() {
	db.pool.Close()
}

// GetPool returns the underlying pgxpool.Pool
func (db *DbConnect) GetPool() *pgxpool.Pool {
	// Check if the pool is nil or closed
	if db == nil || db.pool == nil {
		log.Panic("Database pool is not initialized")
		return nil
	}
	if err := db.pool.Ping(context.Background()); err != nil {
		fmt.Println("Database pool is not available:", err)
	}

	// Return the pool if it's available
	return db.pool
}

// TestDBCleaner initializes the test database connection and returns a cleanup function.
func TestDBCleaner(m *testing.M) func() {
	connectionString := "postgres://postgres:password@127.0.0.1:5432/jwt_auth_test?sslmode=disable"
	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		log.Panic("failed to connect to test database: " + err.Error())
	}

	DbConn = &DbConnect{pool: pool}

	// Clean tables before running tests
	_, err = DbConn.GetPool().Exec(context.Background(), "TRUNCATE TABLE refresh_tokens, users RESTART IDENTITY CASCADE")
	if err != nil {
		panic("failed to truncate tables: " + err.Error())
	}

	return func() {
		DbConn.Close()
	}
}
