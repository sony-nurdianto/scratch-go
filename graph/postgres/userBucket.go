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

//DeleteBucket
func (ub *BucketRepo) DeleteBucket(userbucket *model.UserBucket) error {
	_, err := ub.DB.Model(userbucket).Where("id = ?", userbucket.ID).Delete()
	return err
}
