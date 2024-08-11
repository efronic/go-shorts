# Go Shorts

Go Shorts is a Go application that processes images to generate frames using both single-core and multi-core processing. This project demonstrates the use of concurrency in Go to speed up image processing tasks.

## Project Structure
. ├── .gitignore ├── frames/ ├── go.mod ├── go.sum ├── main_multi.go ├── main_single.go ├── main.go

- `main.go`: Entry point of the application. It calls either the single-core or multi-core processing functions.
- `main_single.go`: Contains the single-core image processing logic.
- `main_multi.go`: Contains the multi-core image processing logic.
- `frames/`: Directory where the generated frames are stored.
- `go.mod` and `go.sum`: Go modules files for dependency management.

## Dependencies

This project uses the following dependencies:

- [`github.com/fogleman/gg`](https://github.com/fogleman/gg): A Go library for rendering 2D graphics.
- [`github.com/golang/freetype`](https://github.com/golang/freetype): A Go library for font rasterization.
- [`golang.org/x/image`](https://pkg.go.dev/golang.org/x/image): Supplementary Go image libraries.

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/efronic/go-shorts.git
    cd go-shorts
    ```

2. Install the dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Place your input image in the project directory and name it [`image.jpg`](./image.jpg) or update the path in the code.

2. Run the application:
    ```sh
    go run .
    ```

By default, the application will use multi-core processing. You can switch to single-core processing by modifying the [`main.go`](./main.go) file:

```go
func main() {
    main_single() // single core
    main_multi() // multi core
}