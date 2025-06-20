Swagger:
go install GitHub.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
swag init -g server.go -d .,./controller (default is main.go) // init setiap kali annotations di update

http://localhost:8080/swagger/index.html