# final-project-4
Toko Belanja

## Endpoints
Users
1. POST - http://localhost:8080/users/register - Register a user
2. POST - http://localhost:8080/users/login - Login
3. PATCH - http://localhost:8080/users/topup - Topup the balance of a user

Categories
1. POST - http://localhost:8080/categories - Create a category
2. GET - http://localhost:8080/categories - Get all categories
3. PATCH - http://localhost:8080/categories/{categoryId} - Update a category's type
4. DELETE - http://localhost:8080/categories/{categoryId} - Delete a category

Products
1. POST - http://localhost:8080/products - Create a product
2. GET - http://localhost:8080/products - Get all products
3. PUT - http://localhost:8080/products/{productId} - Update a product
4. DELETE - http://localhost:8080/products/{productId} - Delete a product

TransactionHistories
1. POST - http://localhost:8080/transactions - Create a transaction
2. GET - http://localhost:8080/transactions/my-transactions - Get all transactions from the current logged in user
3. GET - http://localhost:8080/transactions/user-transactions - Get all transactions from all users

## Group 3
1. Prinata Rakha Santoso - GLNG-KS06-005
2. Iqbal Hasanu Hamdani - GLNG-KS06-001
3. Angga Anugerah Saputro - GLNG-KS06-019

## Pembagian Tugas
Prinata Rakha Santoso
- Initialized App and Models
- Users (Register, Login, Topup)
- Product (Create, Get all)
- TransactionHistories (Create)

Iqbal Hasanu Hamdani
- TransactionHistories (Get my-transactions, Get user-transactions)

Angga Anugerah Saputro
- Categories (Create, Get all, Update, Delete)
- Product (Update, Delete)