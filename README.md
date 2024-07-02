# EquityFusion

EquityFusion is a Go-based RESTful API application designed to simulate the workings of a stock market. This project provides a backend system where users can register, buy, and sell stocks using tokens, with orders processed through a simulated stock exchange.

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
  
## Setup and Usage

1. **Clone the Repository:**
   ```bash
   git clone <repository-url>
   cd EquityFusion
   ```

2. **Build and Run Locally:**
   ```bash
   go build
   ./EquityFusion
   ```
   
3. **API Documentation:**
   - Detailed API documentation and usage examples are provided in the source code.

4. **Docker Deployment:**
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
