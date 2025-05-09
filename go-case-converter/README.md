# Go Case Converter

This project provides a simple command-line interface to convert between camel case and snake case formats. It includes two main conversion functions:

1. **ConvertCamelToSnake**: Converts a string from camel case to snake case.
2. **ConvertSnakeToCamel**: Converts a string from snake case to camel case.

## Installation

To install the project, clone the repository and navigate to the project directory:

```bash
git clone https://github.com/yourusername/go-case-converter.git
cd go-case-converter
```

## Usage

### Convert Camel Case to Snake Case

To convert a string from camel case to snake case, use the `ConvertCamelToSnake` function. 

**Example:**

```go
input := "camelCaseString"
output := ConvertCamelToSnake(input)
// output: "camel_case_string"
```

### Convert Snake Case to Camel Case

To convert a string from snake case to camel case, use the `ConvertSnakeToCamel` function.

**Example:**

```go
input := "snake_case_string"
output := ConvertSnakeToCamel(input)
// output: "snakeCaseString"
```

## Running the Application

To run the application, execute the following command in the project directory:

```bash
go run main.go
```

You can then follow the prompts to enter a string for conversion.

## License

This project is licensed under the MIT License.