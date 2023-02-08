package main

import (
	"database/sql"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"errors"
	"github.com/go-sql-driver/mysql"
	orderHTTP "github.com/vageeshabr/order-service/internal/http/order"
	orderSVC "github.com/vageeshabr/order-service/internal/services/order"
	"github.com/vageeshabr/order-service/internal/stores/order"
)

func OrderHandler(c *gofr.Context) (interface{}, error) {
	return struct {
		Id       int    `json:"id"`
		Customer string `json:"customer"`
	}{
		Id:       1,
		Customer: "vageesha",
	}, nil
}

func main() {
	app := gofr.New()

	app.Server.ValidateHeaders = false

	store := order.New()
	svc := orderSVC.New(store)
	handler := orderHTTP.New(svc)

	app.REST("/orders", handler)

	app.Start()

	//db, err := getDB()
	//if err != nil {
	//	panic(err)
	//}

	//store := order.New(db)
	//svc := orderSVC.New(store)
	//handler := orderHTTP.New(svc)
	//
	//r := mux.NewRouter()
	//r.HandleFunc("/orders", handler.GET).Methods("GET")
	//
	//http.ListenAndServe(":8080", r)
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
