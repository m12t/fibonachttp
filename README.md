# fibonachttp
A native Go webserver that calculates the fibonacci sequence recursively... by sending HTTP requests to itself! All in under 40 lines of code.

## Running the server
1. Build and execute the file in one command with `go run main.go` (use `Ctrl+C` to stop the server)
2. That's it. Now you can send an HTTP `GET` request to the server like `curl localhost:8080/fibonacci/n`, where `n` is the fibonacci number you want to calculate.
    - eg. if you send `curl localhost:8080/fibonacci/11`, you'll get back `{"ans":89}`, which is indeed the correct answer for fib(11)!
