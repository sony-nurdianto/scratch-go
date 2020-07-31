package postgres

import (
	"github.com/go-pg/pg/v9"
	"github.com/sony-nurdianto/scratch-go/graph/model"
)

//BucketRepo to collect data userBucket from repo
type BucketRepo struct {
	DB *pg.DB
}

//GetBucket all
func (ub *BucketRepo) GetBucket() ([]*model.UserBucket, error) {

	var buckets []*model.UserBucket

	err := ub.DB.Model(&buckets).Select()
	if err != nil {
		return nil, err
	}

	return buckets, nil
}

//CreateBucket
func (ub *BucketRepo) CreateBucket(userBucket *model.UserBucket) (*model.UserBucket, error) {
	_, err := ub.DB.Model(userBucket).Returning("*").Insert()

	return userBucket, err
}

//DeleteBucket to Delete Bucket
func (ub *BucketRepo) DeleteBucket(userbucket *model.UserBucket) error {
	_, err := ub.DB.Model(userbucket).Where("id = ?", userbucket.ID).Delete()
	return err
}

//GetUserBucket to get user Bucket
func (ub *BucketRepo) GetUserBucket(user *model.Users) ([]*model.UserBucket, error) {
	var users []*model.UserBucket
	err := ub.DB.Model(&user).Where("user = ?", user.ID).Select()
	return users, err
}

//GetProductBucket to get Product Bucket
func (ub *BucketRepo) GetProductBucket(product *model.Product) ([]*model.UserBucket, error) {
	var products []*model.UserBucket
	err := ub.DB.Model(&products).Where("product = ?", product.ID).Select()
	return products, err
}
