package domain

import (
	"github.com/sony-nurdianto/scratch-go/graph/model"
	"github.com/sony-nurdianto/scratch-go/graph/postgres"
)

//Domain to handle repo
type Domain struct {
	UsersRepo   postgres.UsersRepo
	ProductRepo postgres.ProductRepo
	BucketRepo  postgres.BucketRepo
}

//NewDomain to initial new Domain
func NewDomain(usersRepo postgres.UsersRepo, bucketRepo postgres.BucketRepo, productRepo postgres.ProductRepo) *Domain {
	return &Domain{UsersRepo: usersRepo, BucketRepo: bucketRepo, ProductRepo: productRepo}
}

//Ownable this is to store func from model
type Ownable interface {
	IsOwner(user *model.Users) bool
}

func checkOwnerShip(o Ownable, user *model.Users) bool {
	return o.IsOwner(user)
}
