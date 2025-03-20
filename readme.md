# Go-JWT Authentication API

## Overview
This project is a **JWT-based User Authentication API** built using **Golang** and the **Gin framework**. It provides secure user authentication, authorization, and user management functionalities using MongoDB as the database.

## Features
- **User Registration (Signup)** with hashed passwords
- **User Authentication (Login)** with JWT-based token generation
- **User Authorization** with role-based access control
- **Retrieve Users** (Admin-only API with pagination)
- **Retrieve a Single User** by ID
- **Secure Password Storage** using bcrypt
- **Token Refresh Mechanism**
- **MongoDB Integration** with query optimizations

## Technologies Used
- **Golang** (Gin framework)
- **MongoDB** (NoSQL database)
- **JWT (JSON Web Tokens)** for authentication
- **bcrypt** for password hashing
- **Validator** for input validation

## Project Structure
```
├── controllers/        # API controllers
├── database/          # Database connection setup
├── helpers/           # Utility functions (JWT handling, validation, etc.)
├── models/            # Data models for MongoDB
├── routes/            # API route definitions
├── main.go            # Entry point of the application
├── go.mod             # Go module dependencies
└── README.md          # Project documentation
```

## Installation & Setup
### Prerequisites
Ensure you have the following installed:
- Go 1.18+
- MongoDB (running locally or using MongoDB Atlas)

### Steps to Run the Project
1. **Clone the repository**
   ```sh
   git clone https://github.com/AvadhutJadhav96/Go-JWT.git
   cd Go-JWT
   ```
2. **Install dependencies**
   ```sh
   go mod tidy
   ```
3. **Set up environment variables** (create a `.env` file and define MongoDB connection URI & JWT secret)
   ```sh
   MONGO_URI=mongodb://localhost:27017
   JWT_SECRET=your_secret_key
   ```
4. **Run the application**
   ```sh
   go run main.go
   ```
5. **API is now running on** `http://localhost:8080`

## API Endpoints
### **Authentication Routes**
| Method | Endpoint      | Description          | Access |
|--------|-------------|----------------------|--------|
| POST   | `/signup`    | Register a new user  | Public |
| POST   | `/login`     | Authenticate user & get token | Public |

### **User Management Routes**
| Method | Endpoint      | Description                      | Access |
|--------|-------------|----------------------------------|--------|
| GET    | `/users`     | Get all users (Paginated)        | Admin  |
| GET    | `/user/:id`  | Get a single user by ID          | Admin  |

## Security & Best Practices
- **JWT Authentication**: Each API request requires a valid JWT token.
- **Role-Based Access Control**: Admin-only routes are secured.
- **Input Validation**: All user inputs are validated to prevent injection attacks.
- **Hashed Passwords**: Passwords are stored securely using bcrypt.

## Contribution
If you’d like to contribute:
1. Fork the repo
2. Create a feature branch (`git checkout -b feature-branch`)
3. Commit your changes (`git commit -m "Add new feature"`)
4. Push to the branch (`git push origin feature-branch`)
5. Open a pull request

## License
This project is licensed under the MIT License.

---
**Maintainer:** [Avadhut Jadhav](https://github.com/AvadhutJadhav96)