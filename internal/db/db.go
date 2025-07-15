package db

import (
	"context"
	"database/sql"
	"time"
)

func Database(Addr string, MaxOpenConns, MaxIdleConns int, MaxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", Addr)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(MaxOpenConns)
	db.SetMaxIdleConns(MaxIdleConns)

	duration, err := time.ParseDuration(MaxIdleTime)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(time.Duration(duration))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return db, err
	}

	return db, err
}
