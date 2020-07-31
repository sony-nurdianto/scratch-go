package domain

import (
	"context"

	"github.com/sony-nurdianto/scratch-go/graph/middleware1"
	"github.com/sony-nurdianto/scratch-go/graph/model"
)

//CreateBucket to create user bucket
func (d *Domain) CreateBucket(cnx, ctx context.Context, input model.NewUserBucket) (*model.UserBucket, error) {
	currentUser, _ := middleware1.GetCurrentUserFromCTX(ctx)
	currentProduct, _ := middleware1.GetCurrentProductFromCTX(cnx)

	userBucket := &model.UserBucket{
		Products: currentProduct.ID,
		UserName: currentUser.ID,
	}

	return d.BucketRepo.CreateBucket(userBucket)
}
