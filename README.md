
# 🏦 Banking API - A Scalable, Hexagonal Architecture-based Banking Application in Golang

This **Banking API** is designed with **hexagonal architecture** to ensure scalability and maintainability. Built in **Golang** and powered by **Gorilla Mux** for routing, this API provides banking functionalities like account management and transactions. 

## 🚀 Getting Started

### Prerequisites

Before running the API, ensure the following are installed on your system:

- **Go** - for application development.
- **MySQL** - database for storing banking data.
- **Docker & Docker Compose** (optional, for containerized database setup).

### Environment Variables

To run the application, define the following environment variables. Default values are specified inside `start.sh`:

- **`SERVER_ADDRESS`**: IP address where the server will run.
- **`SERVER_PORT`**: Port for the server.
- **`DB_USER`**: Username for the database.
- **`DB_PASSWD`**: Password for the database.
- **`DB_ADDR`**: IP address of the database server.
- **`DB_PORT`**: Port of the database server.
- **`DB_NAME`**: Name of the database to connect to.

### Quick Start

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/banking-api.git
   cd banking-api
   ```

2. **Run the Application**
   Use the `start.sh` script to download dependencies and run the application:
   ```bash
   ./start.sh
   ```

   > **Note**: `start.sh` sets up default environment variables. Modify these as needed.

## Database Setup

You have two options for setting up your MySQL database:

1. **Using Docker Compose**  
   A `docker-compose.yml` file is available in `resources/docker`. This file includes an initialization script to create tables and insert default data.

   To start the Docker container, navigate to the `resources/docker` directory and run:
   ```bash
   docker-compose up
   ```

2. **Manual Setup Using SQL Script**  
   Alternatively, you can manually set up the database using `resources/database.sql`. Run this script in your MySQL instance to create tables and insert default data.

## 🧪 Testing and Mocks

### Generate Mocks

To generate mocks for testing, run:
```bash
./generate-mocks.sh
```

### Run Unit Tests

To run unit tests and ensure everything is working as expected, use:
```bash
./run-tests.sh
```
