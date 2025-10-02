package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5" // New import for the router
	_ "github.com/lib/pq"
)

type Stock struct {
	Ticker string  `json:"ticker"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
}

var db *sql.DB

// GET /stocks
func getWatchlist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, err := db.Query("SELECT ticker, name, price FROM stocks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var stocks []Stock
	for rows.Next() {
		var s Stock
		if err := rows.Scan(&s.Ticker, &s.Name, &s.Price); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		stocks = append(stocks, s)
	}
	json.NewEncoder(w).Encode(stocks)
}

// POST /stocks
func addStock(w http.ResponseWriter, r *http.Request) {
	var newStock Stock
	if err := json.NewDecoder(r.Body).Decode(&newStock); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `INSERT INTO stocks (ticker, name, price) VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlStatement, newStock.Ticker, newStock.Name, newStock.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Launch a goroutine to do a background task
	go fetchLivePrice(newStock.Ticker)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Stock %s added successfully", newStock.Ticker)
}

// PUT /stocks/{ticker}
func updateStock(w http.ResponseWriter, r *http.Request) {
	ticker := chi.URLParam(r, "ticker")
	var stock Stock
	if err := json.NewDecoder(r.Body).Decode(&stock); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `UPDATE stocks SET name = $1, price = $2 WHERE ticker = $3`
	_, err := db.Exec(sqlStatement, stock.Name, stock.Price, ticker)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Stock %s updated successfully", ticker)
}

// DELETE /stocks/{ticker}
func deleteStock(w http.ResponseWriter, r *http.Request) {
	ticker := chi.URLParam(r, "ticker")

	sqlStatement := `DELETE FROM stocks WHERE ticker = $1`
	_, err := db.Exec(sqlStatement, ticker)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Stock %s deleted successfully", ticker)
}

// fetchLivePrice simulates fetching a live price in the background
func fetchLivePrice(ticker string) {
	fmt.Printf("GOROUTINE: Started fetching live price for %s\n", ticker)
	// Simulate a network call that takes time
	time.Sleep(2 * time.Second)
	// In a real app, you would update the database with the new price here
	fmt.Printf("GOROUTINE: Finished fetching live price for %s\n", ticker)
}

func initDB() {
	connStr := os.Getenv("DATABASE_URL")
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	createTableSQL := `CREATE TABLE IF NOT EXISTS stocks (
		ticker TEXT PRIMARY KEY,
		name TEXT,
		price NUMERIC(10, 2)
	);`
	if _, err = db.Exec(createTableSQL); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connection successful and table created.")
}

func main() {
	initDB()

	// Create a new router
	r := chi.NewRouter()

	// Register our endpoints
	r.Get("/stocks", getWatchlist)
	r.Post("/stocks", addStock)
	r.Put("/stocks/{ticker}", updateStock)
	r.Delete("/stocks/{ticker}", deleteStock)

	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
