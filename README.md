# http-server-go
A basic HTTP server written in Go for playing poker.

## Running the application

### Prerequisites
- Go (version 1.23 or newer)
- Make

### Using Make

This project includes a `Makefile` to simplify common tasks.

**To run the web server:**
```sh
make run-webserver
```
The server will start on `http://localhost:5000`.

**To run the command-line interface (CLI) version:**
```sh
make run-cli
```

**To run the tests:**
```sh
make test
```

**To build the binaries:**
```sh
make build
```
This will create `cli` and `webserver` executables in the `build/` directory.

## Server Functionality

The HTTP server provides a set of endpoints to manage player scores and play a game of poker.

### API Endpoints

*   **`GET /league`**
    *   Returns a JSON array of all players, sorted by their number of wins in descending order.

*   **`GET /players/{name}`**
    *   Retrieves the score for a specific player.
    *   Example: `GET /players/Pepper`
    *   Returns the player's score as a plain text response.
    *   Returns a `404 Not Found` if the player does not exist.

*   **`POST /players/{name}`**
    *   Records a win for a specific player.
    *   Example: `POST /players/Pepper`
    *   Returns a `202 Accepted` on success.