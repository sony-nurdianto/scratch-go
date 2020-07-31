server :
	go run server.go

generate : 
	go run github.com/99designs/gqlgen generate

source :
	source .env