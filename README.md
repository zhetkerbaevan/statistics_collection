# Statistics Collection Service   
The project is a microservice written in Go, designed to collect statistical data in a database.
PostgreSQL is used for storage, with included migration scripts for setup.
## Setup
1. Clone repository
```sh
git clone https://github.com/zhetkerbaevan/statistics_collection.git
cd statistics-collection
```
2. Install Dependencies
 ```sh
go mod tidy
```
3. Database Configuration  
* Set up PostgreSQL database.  
* Configure connection details in internal/config/env.go  
4. Run Migrations
 ```sh
make migrate-up
```
5. Start application
 ```sh
make run
```
## API Endpoints
GET /api/v1/orderbook/{exchange_name}/{pair} - Retrieve all depth orders.  
POST /api/v1/orderbook - Create a new depth order.  
GET /api/v1/orderhistory/{client_name}/{exchange_name}/{label}/{pair} - Retrieve a specific order history.  
POST /api/v1/orderhistory - Create a new order history.  
   
