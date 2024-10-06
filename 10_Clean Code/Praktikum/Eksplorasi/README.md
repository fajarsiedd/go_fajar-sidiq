# go-math-api

Simple API featuring basic calculations including addition, subtraction, multiplication and division.

## Notes

This project needs refactoring to ensure the application follows clean code principle and best practices for REST API application. You can restructure this project in MVC or clean architecture style.

## How to Use

1. Clone this repository.

2. Run the application.

```sh
go run main.go
```

3. Test the application.

```sh
# add
curl -XPOST -H "Content-type: application/json" -d '{"A":1,"B":2}' 'http://localhost:1323/api/add'

# subtract
curl -XPOST -H "Content-type: application/json" -d '{"A":1,"B":2}' 'http://localhost:1323/api/subtract'

# multiply
curl -XPOST -H "Content-type: application/json" -d '{"A":1,"B":2}' 'http://localhost:1323/api/multiply'

# division
curl -XPOST -H "Content-type: application/json" -d '{"A":1,"B":2}' 'http://localhost:1323/api/div'
```

You can choose REST API client like Postman to test the application.
