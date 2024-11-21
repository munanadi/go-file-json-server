# JSON File Server

## Overview

This Go application serves JSON files from a specified directory over HTTP. It allows you to organize your JSON files in a directory structure and access them via a web server running on port 9999. The server responds to GET requests with the contents of the JSON files.

## Features

- Serve JSON files from a specified directory.
- Support for subdirectories, allowing for a structured file organization.
- Simple HTTP server running on port 9999.
- Validates JSON files before serving them.

## Getting Started

### Prerequisites

- Go 1.22.2 or later
- A directory containing JSON files

### Installation

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. Build the application:

   ```bash
   go build -o jsonfileserver main.go
   ```

### Usage

1. Prepare your directory structure. For example:

   ```
   dabba/
       ├── users.json
       └── test.json
   ```

   Example content for `users.json`:

   ```json
   [
       {
           "name": "alpha",
           "age": 34
       },
       {
           "name": "beta",
           "age": 23
       }
   ]
   ```

   Example content for `test.json`:

   ```json
   {
       "message": "hello world"
   }
   ```

2. Run the server, pointing to your directory:

   ```bash
   ./jsonfileserver dabba
   ```

3. Access the JSON files via your web browser or a tool like `curl`:

   - For `users.json`: [http://localhost:9999/users](http://localhost:9999/users)
   - For `test.json`: [http://localhost:9999/test](http://localhost:9999/test)

### Improvements

- **Subdirectory Support**: The server can serve JSON files located in subdirectories, allowing for a more organized file structure.
- **HTTPS Support**: Consider implementing HTTPS for secure communication.

## License

This project is licensed under the MIT License
