# Hello World Go REST API

This project is a simple Go application that sets up an HTTP server with a REST endpoint. When accessed, the endpoint responds with "Hello, World!".

## Project Structure

```
hello-world-go
├── main.go        # Entry point of the application
├── go.mod         # Module definition for the Go project
└── README.md      # Documentation for the project
```

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:

   ```
   git clone <repository-url>
   cd hello-world-go
   ```

2. Initialize the Go module:

   ```
   go mod tidy
   ```

### Running the Application

To run the application, execute the following command:

```
go run main.go
```

The server will start on `localhost:8080`.

### Accessing the Endpoint

Open your web browser or use a tool like `curl` to access the endpoint:

```
http://localhost:8080/hello
```

You should see the response:

```
Hello, World!
```

### License

This project is licensed under the MIT License.