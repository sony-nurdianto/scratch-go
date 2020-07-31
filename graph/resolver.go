package graph

import "github.com/sony-nurdianto/scratch-go/graph/domain"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//Resolver is to connect all repo
type Resolver struct {
	Domain *domain.Domain
}
