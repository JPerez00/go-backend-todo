# Learning & Building a Go Backend for a Todo App

![Image](/gopher-todo.png)

Welcome to my journey of learning Go (Golang) by building a simple Todo app backend. This was my first time diving into Go, I wanted to create something practical to solidify my understanding of the language.

## Project Overview
-    Language: Go (Golang)
-    Framework: Gin
-    ORM: GORM
-    Database: PostgreSQL
-    Purpose: Learning Go by building a RESTful API backend

## Features

- CRUD operations for Todo items
- RESTful API endpoints
- Environment variable management
- Local and production-ready configurations
- CORS support for cross-origin requests

## What I did

1. Set Up the Go Environment: Installed Go and initialized a new Go module for the project.

2. Chose the Gin Web Framework: Opted for Gin to simplify routing and middleware, making the development process smoother.

3. Defined the Todo Model: Created a Todo struct to represent tasks, including fields for ID, Title, and Completed status.

4. Integrated with PostgreSQL Using GORM: Used GORM as an ORM to interact with the PostgreSQL database seamlessly.

5. Managed Environment Variables: Utilized the godotenv package to handle environment variables securely, keeping sensitive information like database credentials out of the codebase.

6. Implemented RESTful API Endpoints:

    GET /todos: Retrieve all todos.
    POST /todos: Create a new todo.
    PUT /todos/:id: Update an existing todo.
    DELETE /todos/:id: Delete a todo.

7. Tested Locally: Verified that the application works locally by setting up a local PostgreSQL database and testing the endpoints using tools like curl and Postman.

8. Prepared for Production Deployment: Added CORS support to handle cross-origin requests, ensuring the backend could communicate with a frontend hosted on a different domain.

## Repo Structure

- main.go: The main application file containing the server setup and route definitions.
- go.mod & go.sum: Go module files for dependency management.
- .env.example: An example environment file illustrating the required environment variables.
- README.md: Project overview and documentation (you're reading it!).