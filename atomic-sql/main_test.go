package main

import (
	"context"
	"database/sql"
	"log"
	"sync"

	"testing"
	"time"
)

func TestPerformPurchase(t *testing.T) {

	db, err := sql.Open("mysql", "root:password@/Stock")
	if err != nil {
		log.Fatal(err)
	}

	var mu sync.Mutex


	app := App{Db: db, mu: &mu}

	ctx := context.Background()

	// Launch another buy
	app.Timeout.Store(true)
	go app.PerformPurchase(ctx, 10, 1)

	
	time.Sleep(2 * time.Second)

	app.Timeout.Store(false)
	err = app.PerformPurchase(ctx, 10, 1)
	if err == nil || err.Error() != StockError {
		t.Fatalf("expected error %v, but got %v", StockError, err)
	}

}
