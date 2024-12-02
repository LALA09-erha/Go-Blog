# REST API Documentation

## Table of Contents

- [Introduction](#introduction)
- [Requirements](#requirements)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
  - [Authentication](#authentication)
  - [User Management (Admin Only)](#user-management-admin-only)
  - [Category Management (User Only)](#category-management-user-only)
  - [Article Management (User Only)](#article-management-user-only)
- [Examples](#examples)

## Introduction

This is a REST API built with Golang for managing:

- User Authentication (JWT-based).
- User CRUD (Admin only).
- Category CRUD (User only).
- Article CRUD (User only).

## Requirements

- Golang: v1.20+
- MySQL: v8.0+
- Postman / Curl: For testing endpoints.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-repo/project-name.git
   cd project-name
   ```

2. Create a .env file in the root folder:

   ```plaintext
   DB_USER=root
   DB_PASSWORD=yourpassword
   DB_HOST=127.0.0.1
   DB_PORT=3306
   DB_NAME=rest_api
   JWT_SECRET=your_jwt_secret
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Set up the database:

   ```bash
   go run cmd/main.go migrate
   ```

5. Run the application:

   ```bash
   go run cmd/main.go
   ```

   The server will start at http://localhost:8080.

## Running the Application

Use the following tools to test the API:

- **Postman**: Import the API endpoints and test.
- **Curl**: Use the examples below to test via terminal.

## API Endpoints

### Authentication

| Method | Endpoint       | Description       | Payload                                                                                 |
| ------ | -------------- | ----------------- | --------------------------------------------------------------------------------------- |
| POST   | /auth/login    | Login user        | { "email": "string", "password": "string" }                                             |
| POST   | /auth/register | Register new user | { "username": "string", "email": "string", "password": "string", "role": "ADMIN/USER" } |

### User Management (Admin Only)

| Method | Endpoint   | Description         | Payload                                                                                 |
| ------ | ---------- | ------------------- | --------------------------------------------------------------------------------------- |
| GET    | /users     | Get all users       | -                                                                                       |
| GET    | /users/:id | Get user by ID      | -                                                                                       |
| POST   | /users     | Create a new user   | { "username": "string", "email": "string", "password": "string", "role": "ADMIN/USER" } |
| PUT    | /users/:id | Update user details | { "username": "string", "email": "string", "password": "string", "role": "ADMIN/USER" } |
| DELETE | /users/:id | Delete user         | -                                                                                       |

### Category Management (User Only)

| Method | Endpoint        | Description        | Payload                                       |
| ------ | --------------- | ------------------ | --------------------------------------------- |
| GET    | /categories     | Get all categories | -                                             |
| GET    | /categories/:id | Get category by ID | -                                             |
| POST   | /categories     | Create category    | { "name": "string", "description": "string" } |
| PUT    | /categories/:id | Update category    | { "name": "string", "description": "string" } |
| DELETE | /categories/:id | Delete category    | -                                             |

### Article Management (User Only)

| Method | Endpoint      | Description       | Payload                                                             |
| ------ | ------------- | ----------------- | ------------------------------------------------------------------- |
| GET    | /articles     | Get all articles  | -                                                                   |
| GET    | /articles/:id | Get article by ID | -                                                                   |
| POST   | /articles     | Create article    | { "title": "string", "content": "string", "category_id": "string" } |
| PUT    | /articles/:id | Update article    | { "title": "string", "content": "string", "category_id": "string" } |
| DELETE | /articles/:id | Delete article    | -                                                                   |

## Examples

### Authentication

#### Login:

```bash
curl -X POST http://localhost:8080/auth/login \
-H "Content-Type: application/json" \
-d '{"email": "admin@example.com", "password": "password123"}'

```

#### Register:

```bash
curl -X POST http://localhost:8080/auth/register \
-H "Content-Type: application/json" \
-d '{"username": "admin", "email": "admin@example.com", "password": "password123", "role": "ADMIN"}'

```

### User Management (Admin Only)

#### Get All Users:

```bash
curl -X GET http://localhost:8080/users \
-H "Authorization: Bearer <your_jwt_token>"

```

#### Create User:

```bash
curl -X POST http://localhost:8080/users \
-H "Authorization: Bearer <your_jwt_token>" \
-H "Content-Type: application/json" \
-d '{"username": "testuser", "email": "test@example.com", "password": "password", "role": "USER"}'

```

### Category Management (User Only)

#### Create Category:

```bash
curl -X POST http://localhost:8080/categories \
-H "Authorization: Bearer <your_jwt_token>" \
-H "Content-Type: application/json" \
-d '{"name": "Technology", "description": "All about tech"}'

```
