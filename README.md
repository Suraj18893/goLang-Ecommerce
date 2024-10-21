# GoBazaarX - eCommerce API

GoBazaarX is a comprehensive eCommerce API built with Go, following a clean architecture design. It provides a variety of features for both administrators and users to manage products, orders, coupons, categories, and much more.

## Features

- **Admin Panel**: Manage categories, products, orders, coupons, and inventories.
- **User-facing Endpoints**: Manage cart, checkout, wishlist, and profile updates.
- **Authentication**: Admin and User authentication with login, signup, and OTP verification.

## Tech Stack

- **Backend**: Go (Gin Framework)
- **Database**: PostgreSQL
- **Documentation**: Swagger
- **Others**: Twilio for OTP, Unidoc for PDF generation

## Getting Started

### Prerequisites

- Go 1.18+ installed
- PostgreSQL database
- Environment variables set in a `.env` file

### Clone the Repository

```bash
git clone https://github.com/Suraj18893/goLang-Ecommerce.git
```

### Setup the Environment
Create a .env file in the root directory and fill it with the following variables:

```bash
DB_HOST=your_db_host
DB_NAME=your_db_name
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_PORT=your_db_port
DB_SSLMODE=require
ACCOUNTSID=your_account_sid
SERVICEID=your_service_id
AUTHTOKEN=your_auth_token
UNIDOC_LICENSE_API_KEY=your_unidoc_license_key
PORT=8082
```

### Install Dependencies

Run the following command to install the necessary Go modules:

```
go mod tidy
```

Run the Application
```
go run main.go
```

### Access the Application
The application should now be running on http://localhost:8082.
Swagger documentation can be accessed at http://localhost:8082/swagger/index.html.

## API Documentation

### Admin Endpoints

- **POST** `/admin/adminlogin`  
  - **Description**: Admin Login.

- **POST** `/admin/category/add`  
  - **Description**: Add a new category.

- **DELETE** `/admin/category/delete`  
  - **Description**: Delete an existing category.

- **PATCH** `/admin/category/update`  
  - **Description**: Update an existing category.

- **GET** `/admin/coupons`  
  - **Description**: List all available coupons.

- **POST** `/admin/coupons/create`  
  - **Description**: Create a new coupon.

- **POST** `/admin/coupons/expire`  
  - **Description**: Make a coupon invalid/expired.

- **POST** `/admin/inventories/add`  
  - **Description**: Add a new inventory item.

- **POST** `/admin/inventories/add-image`  
  - **Description**: Add an image to an inventory item.

- **DELETE** `/admin/inventories/delete`  
  - **Description**: Delete an existing inventory item.

- **DELETE** `/admin/inventories/delete-image`  
  - **Description**: Delete an image from an inventory item.

- **PATCH** `/admin/inventories/update`  
  - **Description**: Update an inventory item.

- **PATCH** `/admin/inventories/update-image`  
  - **Description**: Update an image for an inventory item.

- **GET** `/admin/offers`  
  - **Description**: List all offers.

- **POST** `/admin/offers/create`  
  - **Description**: Create a new offer.

- **POST** `/admin/offers/expire`  
  - **Description**: Expire an offer.

- **GET** `/admin/orders`  
  - **Description**: View all admin orders.

- **PATCH** `/admin/orders/edit/mark-as-paid`  
  - **Description**: Update the payment status of an order.

- **PATCH** `/admin/orders/edit/status`  
  - **Description**: Update the status of an order.

- **GET** `/admin/paymentmethods`  
  - **Description**: View all payment methods.

- **POST** `/admin/paymentmethods/add`  
  - **Description**: Add a new payment method.

- **DELETE** `/admin/paymentmethods/remove`  
  - **Description**: Remove an existing payment method.

- **GET** `/admin/products`  
  - **Description**: List all products.

- **GET** `/admin/sales/annual`  
  - **Description**: View the annual sales report.

- **POST** `/admin/sales/custom`  
  - **Description**: Generate a custom sales report.

- **GET** `/admin/sales/daily`  
  - **Description**: View the daily sales report.

- **GET** `/admin/sales/monthly`  
  - **Description**: View the monthly sales report.

- **GET** `/admin/sales/weekly`  
  - **Description**: View the weekly sales report.

- **POST** `/admin/users/block`  
  - **Description**: Block a user.

- **GET** `/admin/users/getusers`  
  - **Description**: View all users.

- **POST** `/admin/users/unblock`  
  - **Description**: Unblock a user.

### User Endpoints

- **GET** `/products`  
  - **Description**: List all products.

- **GET** `/products/category`  
  - **Description**: Filter products by category.

- **GET** `/products/details`  
  - **Description**: Show product details.

- **GET** `/products/search`  
  - **Description**: Search for products.

- **GET** `/users/cart`  
  - **Description**: Get the user's cart.

- **DELETE** `/users/cart/remove`  
  - **Description**: Remove a product from the cart.

- **POST** `/users/cart/updateQuantity/minus`  
  - **Description**: Decrease the quantity of a product in the cart.

- **POST** `/users/cart/updateQuantity/plus`  
  - **Description**: Increase the quantity of a product in the cart.

- **GET** `/users/check-out`  
  - **Description**: Proceed to checkout.

- **POST** `/users/check-out/order`  
  - **Description**: Place an order.

- **GET** `/users/check-out/order/download-invoice`  
  - **Description**: Download the invoice in PDF format.

- **POST** `/users/home/add-to-cart`  
  - **Description**: Add a product to the cart.

- **POST** `/users/home/add-to-wishlist`  
  - **Description**: Add a product to the wishlist.

- **POST** `/users/login`  
  - **Description**: User login.

- **POST** `/users/otplogin`  
  - **Description**: Send OTP for login.

- **GET** `/users/profile/address`  
  - **Description**: Get the user's saved addresses.

- **POST** `/users/profile/address/add`  
  - **Description**: Add a new address to the user's profile.

- **GET** `/users/profile/details`  
  - **Description**: Get the user's profile details.

- **PATCH** `/users/profile/edit`  
  - **Description**: Edit the user's profile information.

- **GET** `/users/profile/orders`  
  - **Description**: View the user's orders.

- **POST** `/users/profile/orders/cancel`  
  - **Description**: Cancel an order.

- **POST** `/users/profile/orders/return`  
  - **Description**: Return an order.

- **PATCH** `/users/profile/security/change-password`  
  - **Description**: Change the user's password.

- **POST** `/users/signup`  
  - **Description**: User signup.

- **POST** `/users/verifyotp`  
  - **Description**: Verify OTP for signup.

- **GET** `/users/wishlist`  
  - **Description**: View the user's wishlist.

- **DELETE** `/users/wishlist/remove`  
  - **Description**: Remove a product from the wishlist.










