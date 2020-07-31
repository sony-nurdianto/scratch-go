package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/sony-nurdianto/scratch-go/graph/generated"
	"github.com/sony-nurdianto/scratch-go/graph/model"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, input model.RegisterUser) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginUser) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateBucket(ctx context.Context, product string, userName string) (*model.UserBucket, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProduct(ctx context.Context, id string, input model.UpdateProduct) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteProduct(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteBucket(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productResolver) UserID(ctx context.Context, obj *model.Product) (*model.Users, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Products(ctx context.Context, filter *model.FilterProduct, limit *int, offset *int) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserBucket(ctx context.Context) (*model.UserBucket, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, filter *model.FilterUser, limit *int, offset *int) ([]*model.Users, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.Users, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userBucketResolver) Products(ctx context.Context, obj *model.UserBucket) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userBucketResolver) UserName(ctx context.Context, obj *model.UserBucket) (*model.Users, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// UserBucket returns generated.UserBucketResolver implementation.
func (r *Resolver) UserBucket() generated.UserBucketResolver { return &userBucketResolver{r} }

type mutationResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userBucketResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) SingleUpload(ctx context.Context, file graphql.Upload) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *productResolver) Name(ctx context.Context, obj *model.Users) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *productResolver) Description(ctx context.Context, obj *model.Users) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *productResolver) Price(ctx context.Context, obj *model.Users) (int, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *productResolver) Image(ctx context.Context, obj *model.Users) (*graphql.Upload, error) {
	panic(fmt.Errorf("not implemented"))
}
