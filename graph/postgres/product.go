package postgres

import (
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/sony-nurdianto/scratch-go/graph/model"
)

//ProductRepo to collect dataproduct from database
type ProductRepo struct {
	DB *pg.DB
}

/*
Documentasion basic query get all

func (p *ProductRepo)GetProduct()([]*model.Product, error){

	var products []*model.Product

	err := p.DB.Model.Product
	if err != nil {
		return nil ,err
	}

	return products , nil
}

*/

//GetProduct all
func (p *ProductRepo) GetProduct(filter *model.FilterProduct, limit, offset *int) ([]*model.Product, error) {

	var products []*model.Product

	query := p.DB.Model(&products)

	if filter != nil {
		if filter.Name != nil && *filter.Name != "" {
			query.Where("name ILIKE ? ", fmt.Sprintf("%%%s%%", *filter.Name))
		}
	}

	if limit != nil {
		query.Limit(*limit)
	}

	if offset != nil {
		query.Offset(*offset)
	}

	err := query.Select()
	if err != nil {
		return nil, err
	}

	return products, nil

}

//CreateProduct to create Product
func (p *ProductRepo) CreateProduct(product *model.Product) (*model.Product, error) {
	_, err := p.DB.Model(product).Returning("*").Insert()

	return product, err
}

//GetByID Product
func (p *ProductRepo) GetByID(id string) (*model.Product, error) {
	var product model.Product

	err := p.DB.Model(&product).Where("id = ?", id).First()

	return &product, err

}

//Update product
func (p *ProductRepo) Update(product *model.Product) (*model.Product, error) {
	_, err := p.DB.Model(product).Where("id = ? ", product.ID).Update()
	return product, err

}

//Delete product
func (p *ProductRepo) Delete(product *model.Product) error {
	_, err := p.DB.Model(product).Where("id = ?", product.ID).Delete()
	return err
}

//GetUserProduct user have product
func (p *ProductRepo) GetUserProduct(user *model.Users) ([]*model.Product, error) {
	var products []*model.Product
	err := p.DB.Model(&products).Where("user_id = ?", user.ID).Select()
	return products, err
}
