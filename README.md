# Product CRUD API

This is a simple CRUD API for managing products and variants, built using Golang and MySQL. It provides endpoints for creating, reading, updating, and deleting products and variants, along with pagination, filtering, and sorting options.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Endpoints](#endpoints)

## Features

- Read products, brands, and variants along with pagination, sorting, and filtering.
- Update existing products variant.
- Create a new product.
- Soft and Hard Delete products.

## Installation

1. Clone the repository:

```
git clone https://github.com/mucha-fauzy/crud_mysql_api.git
```

2. Install the required dependencies:

```
go mod download
```

3. Create and seed the table in MySQL (the file are inside the migrations folder and seeders folder)

4. Run the application:

```
go run main.go
```

The API will be available at http://localhost:8080.


## Endpoints

### Read products

Send a GET request to `/api/v1/product` to retrieve all products. You can use the following query parameters for filtering, sorting, and pagination:

* productName: Filter products by product name (optional).
* brandName: Filter products by brand name (optional).
* variantName: Filter products by variant name (optional).
* status: Filter products by status (optional).
* sortBy: Sort products by a specific field (optional).
* page: Page number for pagination (optional, default: 1).
* size: Number of items per page (optional, default: 10).

### Update a product variant

Send a PUT request to `/api/v1/variant/{variantID}` with the variant ID in the URL and a JSON payload containing the updated variant details.

### Create a new product
Send a POST request to `/api/v1/product` with a JSON payload containing the product details.

### Soft delete a product
Send a DELETE request to `/api/v1/product/{productID}` with the product ID in the URL and a JSON payload containing `deleted_by` to soft delete the product.

### Hard delete a product
Send a DELETE request to `/api/v1/product/hard/{productID}` with the product ID in the URL to hard delete the product. Beware the data will be lost.
Ideally administrated in IAM by DB admin assign each role for users