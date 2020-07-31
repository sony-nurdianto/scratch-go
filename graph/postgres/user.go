package postgres

import (
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/sony-nurdianto/scratch-go/graph/model"
)

//UsersRepo to collect data users from database
type UsersRepo struct {
	DB *pg.DB
}

/*
Documentation get all users

func (p *ProductRepo)GetProduct()([]*model.Product, error){

	var products []*model.Product

	err := p.DB.Model.Product
	if err != nil {
		return nil ,err
	}

	return products , nil
}

*/

//GetUsers to get user and filter by name
func (u *UsersRepo) GetUsers(filter *model.FilterUser, limit, offset *int) ([]*model.Users, error) {

	var user []*model.Users

	query := u.DB.Model(&user)

	if filter != nil {
		if filter.UserName != nil && *filter.UserName != "" {
			query.Where("user_name ILIKE ? ", fmt.Sprintf("%%%s%%", *filter.UserName))
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

	return user, nil

}

//GetUserByField get user with specific value
func (u *UsersRepo) GetUserByField(field, value string) (*model.Users, error) {
	var user model.Users
	err := u.DB.Model(&user).Where(fmt.Sprintf("%v = ?", field), value).First()

	return &user, err
}

//GetUserByID controleer id
func (u *UsersRepo) GetUserByID(id string) (*model.Users, error) {
	return u.GetUserByField("id", id)
}

//GetUserByEmail like the name
func (u *UsersRepo) GetUserByEmail(email string) (*model.Users, error) {
	return u.GetUserByField("email", email)
}

//GetUserByName like the name
func (u *UsersRepo) GetUserByName(name string) (*model.Users, error) {
	return u.GetUserByField("user_name", name)
}

//CreateUser like the name
func (u *UsersRepo) CreateUser(tx *pg.Tx, user *model.Users) (*model.Users, error) {
	_, err := tx.Model(user).Returning("*").Insert()
	return user, err
}
