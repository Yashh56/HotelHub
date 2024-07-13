# HotelHub

HotelHub is a comprehensive backend system for managing hotels, rooms, bookings, payments, and reviews. Built with Go, Gorilla Mux, Prisma ORM, and PostgreSQL, it provides a robust and scalable solution for handling hotel operations.

## Features

- User registration and login with JWT authentication
- Role-based access control
- Hotel and room management
- Booking management
- Payment processing
- Review system for hotels

## Technologies Used

- Go
- Gorilla Mux (Router)
- Prisma (ORM)
- PostgreSQL (Database)
- jwt-go (JWT authentication)
- gorilla/csrf (CSRF protection)
- gorilla/sessions (Session management)

## Getting Started

### Prerequisites

- Go installed
- PostgreSQL database
- Git

### Installation

1. Clone the repository

    ```bash
    git clone https://github.com/Yashh56/HotelHub.git
    cd HotelHub
    ```

2. Install dependencies

    ```bash
    go mod tidy
    ```

3. Set up environment variables

    Create a `.env` file in the root directory and add the following variables:

    ```env
    DATABASE_URL="your_database_url"
    PORT="your_port_number"
    JWT_SECRET="your_jwt_secret_key"
    ```

### Running the Server

Run the following command to start the server:

```bash
go run main.go
```

# HotelHub API Endpoints

## Auth

- **POST /register**
  - Register a new user
  - Request Body:
    ```json
    {
      "username": "string",
      "email": "string",
      "password": "string"
    }
    ```

- **POST /login**
  - Login a user
  - Request Body:
    ```json
    {
      "email": "string",
      "password": "string"
    }
    ```

## Hotels

- **POST /hotels/create**
  - Create a new hotel (Authenticated users)
  - Request Body:
    ```json
    {
      "name": "string",
      "location": "string",
      "description": "string",
      "rating": "number",
      "totalRooms": "number"
    }
    ```

- **GET /hotels/{hotelID}**
  - Get details of a single hotel

- **GET /reviews/{hotelId}/all**
  - Get reviews of a single hotel

## Rooms

- **POST /rooms/{hotelID}/create**
  - Create a new room (Authenticated users)
  - Request Body:
    ```json
    {
      "roomNumber": "string",
      "type": "string",
      "price": "number",
      "availability": "boolean",
      "description": "string"
    }
    ```

## Bookings

- **POST /booking/{hotelId}/create**
  - Create a new booking (Authenticated users)
  - Request Body:
    ```json
    {
      "checkInDate": "string",
      "checkOutDate": "string",
      "paymentStatus": "string",
      "roomId": "string"
    }
    ```

- **GET /booking/{bookingID}**
  - Get details of a single booking (Authenticated users)

## Reviews

- **POST /review/{hotelId}/create**
  - Create a new review (Authenticated users)
  - Request Body:
    ```json
    {
      "rating": "number",
      "comment": "string",
    }
    ```

## Payments

- **POST /payment/{bookingId}/create**
  - Create a new payment (Authenticated users)
  - Request Body:
    ```json
    {
      "amount": "number",
      "paymentDate": "string",
      "paymentMethod": "string",
      "status": "string",
    }
    ```

- **GET /payment/success**
  - Get all successful payments (Authenticated users)


