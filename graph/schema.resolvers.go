package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/sony-nurdianto/scratch-go/graph/generated"
	"github.com/sony-nurdianto/scratch-go/graph/model"
)

func (r *Resolver) RegisterUser(ctx context.Context, input model.RegisterUser) (*model.AuthResponse, error) {
	isValid := validation(ctx, input)
	if !isValid {
		return nil, errors.New("input errors")
	}

	return r.Domain.RegisterUser(ctx, input)
}

//LoginUser user
func (r *Resolver) LoginUser(ctx context.Context, input model.LoginUser) (*model.AuthResponse, error) {
	isValid := validation(ctx, input)
	if !isValid {
		return nil, errors.New("input errors")
	}

	return r.Domain.LoginUser(ctx, input)
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
