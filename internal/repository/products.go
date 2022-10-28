package repository

import (
	"context"

	"github.com/o-sokol-o/as/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// docker run --rm -d --name crud-logs-mongo -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=g0langn1nja -p 27017:27017 mongo:latest

type Repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) InsertLog(ctx context.Context, log domain.LogItem) error {
	_, err := r.db.Collection(domain.Collection_Logs).InsertOne(ctx, log)

	return err
}

func (r *Repository) GetProducts(ctx context.Context, paging domain.Paging, sorting domain.Sorting) (domain.Products, error) {

	opts := options.Find()

	// Полный список без пагинации и сортировки
	if paging.NumPage != 0 && paging.ProductPerPage != 0 {

		// Список с пагинацией и сортировкой
		if sorting.Direction != domain.SortNone {
			field := domain.GetSortFieldName(sorting.SortField)
			opts.SetSort(bson.M{field: sorting.Direction})
		}
		opts.SetSkip(int64(paging.NumPage - 1))
		opts.SetLimit(int64(paging.ProductPerPage))
	}

	cur, err := r.db.Collection(domain.Collection_Products).Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var products []domain.Product
	if err := cur.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *Repository) InsertProducts(ctx context.Context, products domain.Products) error {
	newValue := make([]interface{}, len(products))
	for i := range products {
		newValue[i] = products[i]
	}

	_, err := r.db.Collection(domain.Collection_Products).InsertMany(ctx, newValue)

	return err
}

func (r *Repository) UpdateProducts(ctx context.Context, products domain.Products) error {
	for i := range products {

		filter := bson.D{{Key: "name", Value: products[i].Name}}

		update := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "cost", Value: products[i].Cost},
			}},
			{Key: "$inc", Value: bson.D{
				{Key: "change_count", Value: 1},
			}},
			{Key: "$set", Value: bson.D{
				{Key: "update_at", Value: products[i].UpdateAt},
			}},
		}

		_, err := r.db.Collection(domain.Collection_Products).UpdateOne(ctx, filter, update)
		if err != nil {
			return err
		}
	}

	return nil
}
