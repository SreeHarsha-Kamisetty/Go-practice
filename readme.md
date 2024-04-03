# Go project

## To initiate a project  

```go
go mod init backend
```

### create a main server file

```bash
touch server.go
```

### declare package name

- example main
- import required modules
- fmt - for formatting
- http - http module for creating server

### write handler functions for different routes

- example

 ```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}
 ```

- Declare main function and handle different routes

 ```go
 func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server started on http://localhost:8080")
    http.ListenAndServe(":8080", nil) // port on which the server should run
}
 ```

### Run command to build the server

```go
go build
```

- This create and .exe file using which we can run our server

```bash
./backend.exe
```
