# REST API in Go

This is a simple REST API built with Go using the standard `net/http` package. The project demonstrates a clean layered structure with routes, handlers, services, repositories, middleware, models, and JSON response utilities.

The API manages book records with basic CRUD operations. Data is stored in memory, so it is useful for learning, testing with Thunder Client/Postman, and understanding REST API development in Go.

## Features

- Create, read, update, and delete book records
- JSON request and response handling
- In-memory data storage
- Logger middleware
- Simple auth middleware placeholder
- Clean project structure using handlers, services, and repositories
- Configurable server port using the `PORT` environment variable

## Tech Stack

- Go
- net/http
- JSON

## API Endpoints

| Method | Endpoint | Description |
| --- | --- | --- |
| GET | `/api/v1/books` | Get all books |
| POST | `/api/v1/books` | Create a new book |
| GET | `/api/v1/books/{id}` | Get a book by ID |
| PUT | `/api/v1/books/{id}` | Update a book |
| DELETE | `/api/v1/books/{id}` | Delete a book |

## Sample Request Body

```json
{
  "name": "Charan",
  "roll_number": 4937
}
