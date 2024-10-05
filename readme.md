# Merchant-Bank API

This project is a simple API for handling transactions between merchants and banks. It provides functionality for user authentication, payment processing, and activity logging.

## Features

- User authentication (login/logout) using JWT
- Payment processing between registered users
- Activity logging for all operations
- Data persistence using JSON files

## Project Structure

```
merchant-bank-api/
├── main.go
├── config/
│   ├── config.go
│   └── config.json
├── controllers/
│   └── controllers.go
├── data/
│   ├── customers.json
│   ├── transactions.json
│   └── history.json
├── middlewares/
│   └── auth.go
├── models/
│   └── models.go
├── repositories/
│   └── repositories.go
├── services/
│   └── services.go
└── utils/
    └── utils.go
```

## Prerequisites

- Go (version 1.16 or later)

## Setup

1. Clone the repository:

   ```
   git clone <repository-url>
   cd merchant-bank-api
   ```

2. Initialize the Go module:

   ```
   go mod init merchant-bank-api
   ```

3. Install dependencies:

   ```
   go get github.com/dgrijalva/jwt-go
   go get github.com/gorilla/mux
   ```

4. Create a `config.json` file in the `config` folder:

   ```json
   {
     "port": "8080",
     "jwt_secret": "your_secret_key_here"
   }
   ```

5. Ensure the dummy data JSON files are in the `data` folder:
   - `customers.json`
   - `transactions.json`
   - `history.json`

## Running the Application

To run the application, use the following command in the project root directory:

```
go run main.go
```

The server will start and listen on the port specified in your `config.json` file (default is 8080).

## API Endpoints

1. Login

   - URL: `/login`
   - Method: `POST`
   - Body: `{"username": "user1", "password": "pass1"}`

2. Payment

   - URL: `/payment`
   - Method: `POST`
   - Headers: `Authorization: <JWT_TOKEN>`
   - Body: `{"from": "1", "to": "2", "amount": 100.00}`

3. Logout
   - URL: `/logout`
   - Method: `POST`
   - Headers: `Authorization: <JWT_TOKEN>`

## Security Considerations

- This is a basic implementation and should not be used in production without further security enhancements.
- Passwords in the dummy data are stored in plain text. In a real application, always use secure password hashing.
- The JWT secret in `config.json` should be kept secure and not shared publicly.
- Consider implementing HTTPS for secure communication in a production environment.
