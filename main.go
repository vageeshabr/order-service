package main

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	orderHTTP "github.com/vageeshabr/order-service/internal/http/order"
	orderSVC "github.com/vageeshabr/order-service/internal/services/order"
	"github.com/vageeshabr/order-service/internal/stores/order"
	"net/http"
)

func main() {
	db, err := getDB()
	if err != nil {
		panic(err)
	}

	store := order.New(db)
	svc := orderSVC.New(store)
	handler := orderHTTP.New(svc)

	r := mux.NewRouter()
	r.HandleFunc("/orders", handler.GET).Methods("GET")

	http.ListenAndServe(":8080", r)
}

var dbError = errors.New("failed to connect to db")

func getDB() (*sql.DB, error) {
	conf := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "test",
		AllowNativePasswords: true,
	}

	connStr := conf.FormatDSN()
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, dbError
	}

	if err := db.Ping(); err != nil {
		return nil, dbError
	}

	return db, nil
}
