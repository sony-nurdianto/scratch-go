package domain

import (
	"context"
	"errors"
	"log"

	"github.com/sony-nurdianto/scratch-go/graph/model"
)

//RegisterUser func
func (d *Domain) RegisterUser(ctx context.Context, input model.RegisterUser) (*model.AuthResponse, error) {
	_, err := d.UsersRepo.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("email is alredy exist")
	}

	_, err = d.UsersRepo.GetUserByName(input.UserName)
	if err == nil {
		return nil, errors.New("user name is already used")
	}

	user := &model.Users{
		UserName:  input.UserName,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	err = user.HashPassword(input.Password)
	if err != nil {
		log.Printf("error while hasing password : %v", err)
		return nil, errors.New("something went wrong")
	}

	//create verification code

	tx, err := d.UsersRepo.DB.Begin()
	if err != nil {
		log.Printf("erro when creating transaction : %v", err)
		return nil, errors.New("something while wrong ")
	}

	defer tx.Rollback()

	if _, err := d.UsersRepo.CreateUser(tx, user); err != nil {
		log.Printf("erro when creating user : %v", err)
		return nil, errors.New("something wrong")
	}

	if err := tx.Commit(); err != nil {
		log.Printf("erro when commiting %v", err)
		return nil, err
	}

	token, err := user.GenerateToken()
	if err != nil {
		log.Printf("erroe when generating token : %v", err)
		return nil, errors.New("something when wrong")
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

//LoginUser user
func (d *Domain) LoginUser(ctx context.Context, input model.LoginUser) (*model.AuthResponse, error) {
	user, err := d.UsersRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, errors.New("email wrong or doesnt exist")
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		return nil, errors.New("password wrong")
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}
