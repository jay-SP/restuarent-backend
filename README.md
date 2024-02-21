# Restaurant Management Backend

This project is a backend system for managing a restaurant built using Go language with MongoDB as the database. It follows the Model-View-Controller (MVC) architecture pattern and utilizes JSON Web Tokens (JWT) for authentication.

## Features:
- **RESTful APIs:** Implements RESTful APIs to handle various functionalities of the restaurant backend.
- **MongoDB Integration:** Uses MongoDB as the database to store and manage restaurant data.
- **MVC Architecture:** Organizes the codebase following the MVC design pattern for better maintainability and scalability.
- **JWT Authentication:** Implements JWT-based authentication to secure the APIs.

## Requirements:
- Go programming language
- MongoDB database
- JSON Web Tokens (JWT) library for Go

## Restaurant Backend API Routes:

### User Routes:
- **GET /users:** Retrieve all users.
- **GET /users/:user_id:** Retrieve a specific user by ID.
- **POST /users/signup:** Register a new user.
- **POST /users/login:** Authenticate user login.

### Food Routes:
- **GET /foods:** Retrieve all foods.
- **GET /foods/:food_id:** Retrieve a specific food by ID.
- **POST /foods:** Create a new food.
- **PATCH /foods/:food_id:** Update a specific food by ID.

### Menu Routes:
- **GET /menus:** Retrieve all menus.
- **GET /menus/:food_id:** Retrieve a specific menu by food ID.
- **POST /menus:** Create a new menu.
- **PATCH /menus/:food_id:** Update a specific menu by food ID.

### Table Routes:
- **GET /tables:** Retrieve all tables.
- **GET /tables/:table_id:** Retrieve a specific table by ID.
- **POST /tables:** Create a new table.
- **PATCH /tables/:table_id:** Update a specific table by ID.

### Order Routes:
- **GET /orders:** Retrieve all orders.
- **GET /orders/:order_id:** Retrieve a specific order by ID.
- **POST /orders:** Create a new order.
- **PATCH /orders/:order_id:** Update a specific order by ID.

### Order Item Routes:
- **GET /orderItems:** Retrieve all order items.
- **GET /orderItems/:orderItem_id:** Retrieve a specific order item by ID.
- **GET /orderItems-order/:order_id:** Retrieve all order items by order ID.
- **POST /orderItems:** Create a new order item.
- **PATCH /orderItems/:orderItem_id:** Update a specific order item by ID.

### Invoice Routes:
- **GET /invoices:** Retrieve all invoices.
- **GET /invoices/:invoice_id:** Retrieve a specific invoice by ID.
- **POST /invoices:** Create a new invoice.
- **PATCH /invoices/:invoice_id:** Update a specific invoice by ID.

## MongoDB Aggregation:

This project also performs MongoDB aggregation to fetch aggregated data related to order items.
It creates a context with a timeout, defines aggregation pipeline stages such as matching,
lookups, unwinding arrays, projection, and grouping. Finally, it executes the aggregation
pipeline against the MongoDB collection to retrieve structured data including details about 
food, orders, and tables, grouped by order ID and table ID with calculated payment due.


