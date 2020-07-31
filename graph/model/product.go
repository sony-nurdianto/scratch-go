package model

import "github.com/99designs/gqlgen/graphql"

//Product Model
type Product struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       int            `json:"price"`
	Image       graphql.Upload `json:"image"`
	UserID      []*Users       `json:"user_id"`
}
