# EquityFusion

EquityFusion is a Go-based RESTful API cum CLI application designed to simulate the workings of a stock market. This project provides a backend system where users can register, buy, and sell stocks using tokens, with orders processed through a simulated stock exchange.

## Features

- **User Registration:** 
  - `/register` generates a token for each user.
  
- **Order Management:**
  - `/sell`: Create a sell order with parameters token/id, symbol, quantity, and price.
  - `/buy`: Create a buy order with parameters token/id, symbol, quantity, and price.
  
- **Stock Exchange Simulation:**
  - Orders are matched based on specific rules:
    - Sell orders are prioritized by lowest price.
    - Buy orders are prioritized by highest price.
    - FIFO order is maintained for orders at the same price.
  
- **Advanced Features:**
  - `/portfolio`: Retrieve stock symbols, quantities, and pending orders for a given token/id.
  - `/tradebook`: Log of all transactions for a specific stock symbol.
  
- **Tech Stack:**
  - Go (Golang)
  - Gin Web Framework
  - CobraCli
  
## Setup and Usage

1. **Clone the Repository:**
   ```bash
   git clone <repository-url>
   cd EquityFusion
   ```

2. **Build and Run Locally:**
   ```bash
   make build
   ./ef
   ```
   
3. **Usecase Options**

   **API Endpoints:**

    | Method | URL | Description | JSON Body (POST) | Query Parameters (GET) |
    |---|---|---|---|---|
    | GET | `/health` | Checks the health of the API. | - | - |
    | GET | `/register` | Registers a new user and returns a token. | - | - |
    | GET | `/sell` | Retrieves a list of all sell orders. | - | - | (Deprecated, use POST for creating orders)
    | GET | `/buy` | Retrieves a list of all buy orders. | - | - | (Deprecated, use POST for creating orders)
    | GET | `/seed` | Seeds the system with dummy orders for testing. | - | `?number=<number_of_orders>` |
    | GET | `/tradebook` | Retrieves the tradebook for a specific user. | - | - |
    | GET | `/portfolio` | Retrieves portfolio information for specific user. | - | `?userId=<user_id>` |
    | POST | `/sell` | Creates a new sell order. | `{ "id": "user_id", "symbol": "stock_symbol", "quantity": <quantity>, "price": <price> }` | - |
    | POST | `/buy` | Creates a new buy order. | `{ "id": "user_id", "symbol": "stock_symbol", "quantity": <quantity>, "price": <price> }` | - |


    **CLI Commands:**

    * **`ef register`:** Register a new user and obtain a token.
    * **`ef order -i <user_id> -s <symbol> -q <quantity> -p <price> -b <buy order>`:** Create a buy or sell order.
    * **`ef portfolio -i <user_id>`:** View portfolio details for a specific user.
    * **`ef tradebook`:** View trade history for a specific user.
    * **`ef seed -n <number>`:** Seed the system with dummy orders for testing.
    * **`ef expose -p <port>`:** Expose the RESTful API on the specified port.

4. **[TBA] Docker Deployment:**
   - Build the Docker image:
     ```bash
     docker build -t equityfusion .
     ```
   - Run the Docker container:
     ```bash
     docker run -p 8080:8080 equityfusion
     ```

## License

This project is licensed under the [MIT License](LICENSE).
