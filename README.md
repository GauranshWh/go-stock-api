# Go Stock Watchlist API

A simple and robust backend microservice built with Go for managing a personal stock watchlist. This project demonstrates a complete, containerized, and persistent RESTful API service, showcasing modern backend development practices.

---

## Features

- **Full CRUD Functionality:** Create, Read, Update, and Delete stocks from your watchlist.
- **RESTful API:** A clean, logical API design using standard HTTP methods and JSON data format.
- **Persistent Storage:** Data is saved in a PostgreSQL database, managed with Docker Compose.
- **Asynchronous Tasks:** Utilizes Goroutines for non-blocking background tasks (e.g., fetching live price data).
- **Fully Containerized:** The entire application stack (Go API + Database) is managed with Docker for easy setup and consistent deployment.

---

## Tech Stack

- **Backend:** Go (Golang)
- **Database:** PostgreSQL
- **API:** REST (using `net/http` and `chi` router)
- **Containerization:** Docker & Docker Compose

---

## Prerequisites

Make sure you have the following installed on your system:
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

---

## 🚀 Getting Started

Follow these steps to get the application running locally.

**1. Clone the repository:**

bash
git clone [https://github.com/GauranshWh/go-stock-api.git]
cd go-stock-api

2. Create your environment file:
Create a .env file by copying the example file. This file contains the necessary credentials for the database.

Bash

cp .env.example .env
(No changes are needed in this file for local development.)

3. Run the application:
Use Docker Compose to build the images and start the containers. The --build flag is only needed the first time or when you make code changes.

Bash

docker-compose up --build
The API will now be running and available at http://localhost:8080.

API Endpoints
You can interact with the API using a tool like Postman or curl.

Method	Endpoint	Body (JSON) Example	Description
GET	/stocks	(none)	Get all stocks in the watchlist.
POST	/stocks	{ "ticker": "TSLA", "name": "Tesla Inc.", "price": 200.50 }	Add a new stock to the watchlist.
PUT	/stocks/{ticker}	{ "name": "Tesla, Inc.", "price": 205.75 }	Update an existing stock by its ticker.
DELETE	/stocks/{ticker}	(none)	Delete a stock from the watchlist.

Export to Sheets
Example curl Commands:
Add a stock:

Bash

curl -X POST -H "Content-Type: application/json" -d "{\"ticker\": \"NVDA\", \"name\": \"NVIDIA Corp.\", \"price\": 130.75}" http://localhost:8080/stocks
Update a stock:

Bash

curl -X PUT -H "Content-Type: application/json" -d "{\"name\": \"NVIDIA Corporation\", \"price\": 135.00}" http://localhost:8080/stocks/NVDA
Delete a stock:

Bash

curl -X DELETE http://localhost:8080/stocks/NVDA

Future Improvements
[ ] Implement user authentication and authorization using JWT.

[ ] Add comprehensive unit and integration tests.

[ ] Integrate a real-time price feed using Websockets.


