# Bookstore Microservice

The Bookstore Microservice is a Go application that exposes RESTful APIs to manage a list of books in a store. It provides endpoints to list all books, get details of a specific book, and add a new book to the bookstore.

## Table of Contents
- [Dependencies](#dependencies)
- [How to Run](#how-to-run)
  - [Standalone](#standalone)
  - [Docker Container](#docker-container)

## Dependencies

To run the Bookstore Microservice, you need the following dependencies:

- Go (version 1.17 or later): https://golang.org/dl/
- Docker: https://www.docker.com/get-started

## How to Run

### Standalone

To run the application as a standalone app, follow these steps:

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/anirbanbhat/bookstore_microservice.git
   ```

2. Change into the project directory:

   ```bash
   cd bookstore_microservice
   ```

3. Install the Go Echo framework dependencies:

   ```bash
   go mod download
   ```

4. Build the application:

   ```bash
   go build -o bookstore_microservice
   ```

5. Run the application:

   ```bash
   ./bookstore_microservice
   ```

6. The Bookstore Microservice will now be running locally at `http://localhost:8080`.

### Docker Container

To run the application as a Docker container, ensure you have Docker installed and follow these steps:

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/anirbanbhat/bookstore_microservice.git
   ```

   Replace `anirbanbhat` with your actual GitHub username.

2. Change into the project directory:

   ```bash
   cd bookstore_microservice
   ```

3. Build the Docker image using the provided `Dockerfile`:

   ```bash
   docker build -t golang-bookstore-app .
   ```

4. Run the Docker container:

   ```bash
   docker run -p 8080:8080 golang-bookstore-app
   ```

5. The Bookstore Microservice will now be running in a Docker container at `http://localhost:8080`.

