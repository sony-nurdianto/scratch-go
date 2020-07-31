package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/sony-nurdianto/scratch-go/graph/middleware1"
	"github.com/sony-nurdianto/scratch-go/graph/model"
)

//CreateProduct to create Product
func (d *Domain) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	currentUser, err := middleware1.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("unautenticated")
	}

	if len(input.Name) < 3 {
		return nil, errors.New("name is to short")
	}

	if len(input.Description) < 3 {
		return nil, errors.New("description is to short")
	}

	product := &model.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		UserID:      currentUser.ID,
	}

	return d.ProductRepo.CreateProduct(product)
}

//UpdateProduct this is controler Product Update
func (d *Domain) UpdateProduct(ctx context.Context, id string, input model.UpdateProduct) (*model.Product, error) {
	currentUser, err := middleware1.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("unautenticated")
	}

	product, err := d.ProductRepo.GetByID(id)
	if err != nil || product == nil {
		return nil, errors.New("product not exist")
	}

	if !product.IsOwner(currentUser) {
		return nil, errors.New("you are not user")
	}

	didUpdate := false

	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, errors.New("name to short")
		}
		product.Name = *input.Name

		didUpdate = true
	}
	if input.Description != nil {
		if len(*input.Description) < 3 {
			return nil, errors.New("Description to short")
		}
		product.Description = *input.Description
		didUpdate = true
	}

	if input.Price != nil {
		if *input.Price <= 0 {
			return nil, errors.New("price canot Zerro")
		}

		product.Price = *input.Price
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("no update done")
	}

	product, err = d.ProductRepo.Update(product)
	if err != nil {
		return nil, fmt.Errorf("error while updateing %v", err)
	}

	return product, nil
}

//DeleteProduct this is controler delete Product
func (d *Domain) DeleteProduct(ctx context.Context, id string) (bool, error) {
	currentUser, err := middleware1.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return false, errors.New("unautenticated")
	}

	product, err := d.ProductRepo.GetByID(id)
	if err != nil || product == nil {
		return false, errors.New("product not exist")
	}

	if !product.IsOwner(currentUser) {
		return false, errors.New("you're not owner")
	}

	if !checkOwnerShip(product, currentUser) {
		return false, errors.New("you're not owner")
	}

	err = d.ProductRepo.Delete(product)
	if err != nil {
		return false, fmt.Errorf("error while deleteing : %v", err)
	}

	return true, nil
}
