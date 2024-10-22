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

<img src="https://github.com/Suraj18893/goLang-Ecommerce/blob/master/diagram-export-10-21-2024-11_43_01-PM.svg" >


### Admin Endpoints

```mermaid
flowchart TD
A1[POST /admin/adminlogin] -->|Logs in the Admin| A2[Admin Dashboard]
```

```mermaid
flowchart TD
B1[Admin Dashboard] -->|Navigate to| B2[Manage Categories]
B2 -->|Add a Category| B3[POST /admin/category/add]
B2 -->|Delete a Category| B4[DELETE /admin/category/delete]
B2 -->|Update a Category| B5[PATCH /admin/category/update]
```

```mermaid
flowchart TD
C1[Admin Dashboard] -->|Navigate to| C2[Manage Coupons]
C2 -->|List Coupons| C3[GET /admin/coupons]
C2 -->|Add a Coupon| C4[POST /admin/coupons/create]
C2 -->|Expire a Coupon| C5[POST /admin/coupons/expire]
```

```mermaid
flowchart TD
D1[Admin Dashboard] -->|Navigate to| D2[Manage Inventories]
D2 -->|Add Inventory| D3[POST /admin/inventories/add]
D2 -->|Add Inventory Image| D4[POST /admin/inventories/add-image]
D2 -->|Delete Inventory| D5[DELETE /admin/inventories/delete]
D2 -->|Delete Inventory Image| D6[DELETE /admin/inventories/delete-image]
D2 -->|Update Inventory| D7[PATCH /admin/inventories/update]
D2 -->|Update Inventory Image| D8[PATCH /admin/inventories/update-image]
```

```mermaid
flowchart TD
E1[Admin Dashboard] -->|Navigate to| E2[Manage Offers]
E2 -->|List Offers| E3[GET /admin/offers]
E2 -->|Add Offer| E4[POST /admin/offers/create]
E2 -->|Expire Offer| E5[POST /admin/offers/expire]
```

```mermaid

flowchart TD
F1[Admin Dashboard] -->|Navigate to| F2[Manage Orders]
F2 -->|List Admin Orders| F3[GET /admin/orders]
F2 -->|Update Payment Status| F4[PATCH /admin/orders/edit/mark-as-paid]
F2 -->|Update Order Status| F5[PATCH /admin/orders/edit/status]
```

```mermaid

flowchart TD
G1[Admin Dashboard] -->|Navigate to| G2[Manage Payment Methods]
G2 -->|List Payment Methods| G3[GET /admin/paymentmethods]
G2 -->|Add New Payment Method| G4[POST /admin/paymentmethods/add]
G2 -->|Remove Payment Method| G5[DELETE /admin/paymentmethods/remove]
```

```mermaid

flowchart TD
H1[Admin Dashboard] -->|Navigate to| H2[Manage Products]
H2 -->|List Products| H3[GET /admin/products]
```

```mermaid

flowchart TD
I1[Admin Dashboard] -->|Navigate to| I2[View Sales Reports]
I2 -->|Annual Sales Report| I3[GET /admin/sales/annual]
I2 -->|Custom Sales Report| I4[POST /admin/sales/custom]
I2 -->|Daily Sales Report| I5[GET /admin/sales/daily]
I2 -->|Monthly Sales Report| I6[GET /admin/sales/monthly]
I2 -->|Weekly Sales Report| I7[GET /admin/sales/weekly]
```

```mermaid

flowchart TD
J1[Admin Dashboard] -->|Navigate to| J2[Manage Users]
J2 -->|Block User| J3[POST /admin/users/block]
J2 -->|Get Users| J4[GET /admin/users/getusers]
J2 -->|Unblock User| J5[POST /admin/users/unblock]
```


### User Endpoints

```mermaid
flowchart TD
    A1[POST /users/login] -->|Logs in the User| A2[User Dashboard]
```
```mermaid

flowchart TD
    B1[User Dashboard] -->|Navigate to| B2[Browse Products]
    B2 -->|List Products| B3[GET /products]
    B2 -->|Filter Products by Category| B4[GET /products/category]
    B2 -->|Show Product Details| B5[GET /products/details]
    B2 -->|Search Products| B6[GET /products/search]
```

```mermaid

flowchart TD
    C1[User Dashboard] -->|Navigate to| C2[Manage Cart]
    C2 -->|Get Cart| C3[GET /users/cart]
    C2 -->|Remove from Cart| C4[DELETE /users/cart/remove]
    C2 -->|Decrease Quantity in Cart| C5[POST /users/cart/updateQuantity/minus]
    C2 -->|Increase Quantity in Cart| C6[POST /users/cart/updateQuantity/plus]
    C2 -->|Checkout| C7[GET /users/check-out]
    C7 -->|Place Order| C8[POST /users/check-out/order]
    C7 -->|Download Invoice PDF| C9[GET /users/check-out/order/download-invoice]
```

```mermaid

flowchart TD
    D1[User Dashboard] -->|Navigate to| D2[Manage Wishlist]
    D2 -->|Add Product to Wishlist| D3[POST /users/home/add-to-wishlist]
    D2 -->|Get Wishlist| D4[GET /users/wishlist]
    D2 -->|Remove Product from Wishlist| D5[DELETE /users/wishlist/remove]
```

```mermaid

flowchart TD
    E1[User Dashboard] -->|Navigate to| E2[Manage Profile]
    E2 -->|Get User Addresses| E3[GET /users/profile/address]
    E2 -->|Add User Address| E4[POST /users/profile/address/add]
    E2 -->|Get User Profile Details| E5[GET /users/profile/details]
    E2 -->|Edit User Profile| E6[PATCH /users/profile/edit]
    E2 -->|Get User Orders| E7[GET /users/profile/orders]
    E2 -->|Cancel Order| E8[POST /users/profile/orders/cancel]
    E2 -->|Return Order| E9[POST /users/profile/orders/return]
    E2 -->|Change User Password| E10[PATCH /users/profile/security/change-password]
```

```mermaid

flowchart TD
    F1[POST /users/signup] -->|Sign up the User| F2[Verify OTP]
    F2 -->|Send OTP| F3[POST /users/otplogin]
    F2 -->|Verify OTP| F4[POST /users/verifyotp]
```

 
