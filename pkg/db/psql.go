package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AliAstanov/my_blogs/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func getConfig(cfg config.PgConfig) *pgxpool.Config {
	const defaultMaxConns = int32(4)
	const defaultMinConns = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	var databaseUrl string = fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DatabaseName,
	)

	dbConfig, err := pgxpool.ParseConfig(databaseUrl)
	fmt.Println("go")
	if err != nil {
		fmt.Println("err:", err)
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		log.Println("Before acquiring the connection poolto the database!!")
		return true
	}

	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		log.Println("After releasing the connection pool to the databasse!!")
		return true
	}

	dbConfig.BeforeClose = func(c *pgx.Conn) {
		log.Println("clothing the connection pool to the database!!")
	}

	return dbConfig

}

func ConnDb(cfg config.PgConfig) (*pgxpool.Pool, error) {
	connPool, err := pgxpool.NewWithConfig(context.Background(), getConfig(cfg))
	if err != nil {
		log.Fatal("Error while creating connection to the database!!", err)

	}

	connection, err := connPool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error while asquiring connection to the database pool !!", err)
	}
	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		log.Fatal("Could not ping database", err)
	}

	fmt.Println("connected to the database!!")
	return connPool, nil
}
