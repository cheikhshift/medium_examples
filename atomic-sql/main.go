package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	Db      *sql.DB
	mu      *sync.Mutex
	Timeout atomic.Bool
}

const (
	StockError = "Stock is finished"
)

func main() {

	db, err := sql.Open("mysql", "root:MAhomies69@/Stock")
	if err != nil {
		log.Fatal(err)
	}
	var mu sync.Mutex

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	app := App{Db: db, mu: &mu}

	http.HandleFunc("/buy", app.BuyItem)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (a *App) BuyItem(w http.ResponseWriter, r *http.Request) {

	id, amount, err := GetQuery(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Nope!"))
		return
	}

	ctx, _ := context.WithTimeout(r.Context(), 30*time.Second)

	err = a.PerformPurchase(ctx, id, amount)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Nope!"))

		return
	}

	w.Write([]byte("Ok"))
}

func GetQuery(r *http.Request) (id int, amount int, err error) {

	amount, err = strconv.Atoi(r.FormValue("amount"))

	if err != nil {
		return
	}

	id, err = strconv.Atoi(r.FormValue("id"))

	if err != nil {
		return
	}

	return
}

func (a *App) PerformPurchase(ctx context.Context, id, amount int) error {

	a.mu.Lock()
	defer a.mu.Unlock()



	if a.Timeout.Load() {
		time.Sleep(5 * time.Second)
	}

	rows, err := a.Db.QueryContext(ctx, "SELECT Count FROM Inventory WHERE ProductID=?", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	var stock int

	if rows.Next() {
		if err := rows.Scan(&stock); err != nil {
			// Check for a scan error...
			return err
		}

	}

	if stock <= 0 {
		return errors.New(StockError)
	}

	_, err = a.Db.ExecContext(ctx, "UPDATE Inventory SET Count = Count - ? WHERE ProductID = ?", amount, id)
	if err != nil {
		return err
	}

	return nil

}
