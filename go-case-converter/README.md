# Go Case Converter

The `go-case-converter` project provides a REST API for converting strings between camel case and snake case formats. It includes two main endpoints:

1. **Convert Camel Case to Snake Case**: `/convert/camel-to-snake`
2. **Convert Snake Case to Camel Case**: `/convert/snake-to-camel`

## Features

- Convert camel case strings (e.g., `camelCase`) to snake case (e.g., `camel_case`).
- Convert snake case strings (e.g., `snake_case`) to camel case (e.g., `snakeCase`).
- Input validation to ensure only valid single-word strings are processed.

## Installation

### Prerequisites

- Go 1.18 or later
- Docker (optional, for containerized deployment)

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/go-case-converter.git
   cd go-case-converter
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`.

## Usage

### Endpoints

#### 1. Convert Camel Case to Snake Case
- **URL**: `/convert/camel-to-snake`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "input": "camelCaseExample"
  }
  ```
- **Response**:
  ```json
  {
    "result": "camel_case_example"
  }
  ```

#### 2. Convert Snake Case to Camel Case
- **URL**: `/convert/snake-to-camel`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "input": "snake_case_example"
  }
  ```
- **Response**:
  ```json
  {
    "result": "snakeCaseExample"
  }
  ```

### Example Usage with `curl`

#### Convert Camel Case to Snake Case
```bash
curl -X POST http://localhost:8080/convert/camel-to-snake \
-H "Content-Type: application/json" \
-d '{"input": "camelCaseExample"}'
```

#### Convert Snake Case to Camel Case
```bash
curl -X POST http://localhost:8080/convert/snake-to-camel \
-H "Content-Type: application/json" \
-d '{"input": "snake_case_example"}'
```
