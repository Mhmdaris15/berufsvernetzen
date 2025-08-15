# Berufsvernetzen Backend

This is the backend for the Berufsvernetzen application. It is built with Go and uses the Gin framework. It connects to a MongoDB database and uses MeiliSearch for search functionality.

## Prerequisites

Before you begin, ensure you have the following installed:
- [Go](https://golang.org/doc/install) (version 1.21 or higher)
- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)

## Environment Variables

This project uses environment variables for configuration. Create a `.env` file in the `backend` directory and add the following variables. You can use the `backend/.env.example` file as a template.

```
PORT=3000
MONGO_URI=mongodb://localhost:27017
DATABASE_NAME=berufsvernetzen
GRPC_PORT=9090
PASETO_SYMMETRIC_KEY=a_very_secret_key_of_32_chars
ACCESS_TOKEN_DURATION=1h
MEILISEARCH_MASTER_KEY=your_meilisearch_master_key
GCLOUD_PROJECT_ID=your_gcloud_project_id
```

## Getting Started

1.  **Clone the repository:**
    ```bash
    git clone <repository-url>
    cd berufsvernetzen/backend
    ```

2.  **Install dependencies:**
    ```bash
    go mod download
    ```

3.  **Run the application:**
    ```bash
    go run cmd/server/main.go
    ```
    The server will start on the port specified in your `.env` file (default: 3000).

## Running Tests

To run the test suite, execute the following command from the `backend` directory:

```bash
go test ./...
```

## Docker Usage

The application can also be run using Docker Compose.

1.  **Build and start the services:**
    From the `backend` directory, run:
    ```bash
    docker-compose up --build
    ```
This will build the Go application image and start the container. The `docker-compose.yml` is configured to expose ports 3000 (for the API) and 9090 (for gRPC).
