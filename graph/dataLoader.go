package graph

import (
	"context"
	"net/http"
	"time"

	"github.com/go-pg/pg"
	"github.com/sony-nurdianto/scratch-go/graph/model"
)

const userLoaderKey = "userloader"
const productLoaderKey = "productloader"
const userBucketLoaderKey = "bucketloader"

//DataLoaderUser for handling loader data user
func DataLoaderUser(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userloader := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*model.Users, []error) {
				var users []*model.Users

				err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()
				if err != nil {
					return nil, []error{err}
				}

				u := make(map[string]*model.Users, len(users))

				for _, user := range users {
					u[user.ID] = user
				}

				result := make([]*model.Users, len(ids))

				for i, id := range ids {
					result[i] = u[id]
				}

				return users, nil
			},
		}

		ctx := context.WithValue(r.Context(), userLoaderKey, &userloader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//DataLoaderProduct for handling loader data product
func DataLoaderProduct(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		productloader := ProductLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*model.Product, []error) {

				var products []*model.Product

				err := db.Model(&products).Where("id = (?)", pg.In(ids)).Select()
				if err != nil {
					return nil, []error{err}
				}

				p := make(map[string]*model.Product, len(products))

				for _, product := range products {
					p[product.ID] = product
				}

				result := make([]*model.Product, len(ids))

				for i, id := range ids {
					result[i] = p[id]
				}

				return products, nil
			},
		}

		ctx := context.WithValue(r.Context(), productLoaderKey, &productloader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//DataLoaderUserBucket for handling userBucket data loader
func DataLoaderUserBucket(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userbucketLoader := UserBucketLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*model.UserBucket, []error) {

				var userbuckets []*model.UserBucket

				err := db.Model(&userbuckets).Where("id = (?)", pg.In(ids)).Select()
				if err != nil {
					return nil, []error{err}
				}

				ub := make(map[string]*model.UserBucket, len(userbuckets))

				for _, userbucket := range userbuckets {
					ub[userbucket.ID] = userbucket
				}

				result := make([]*model.UserBucket, len(ids))

				for i, id := range ids {
					result[i] = ub[id]
				}

				return userbuckets, nil

			},
		}

		ctx := context.WithValue(r.Context(), userBucketLoaderKey, &userbucketLoader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userLoaderKey).(*UserLoader)
}

func getProductLoader(ctx context.Context) *ProductLoader {
	return ctx.Value(productLoaderKey).(*ProductLoader)
}
func getUserBucketLoader(ctx context.Context) *UserBucketLoader {
	return ctx.Value(userBucketLoaderKey).(*UserBucketLoader)
}
