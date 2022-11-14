package mongo

import (
	"context"
	"metadata/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateDslInfo(ctx context.Context, dslInfo model.DslInfoStruct) error {
	collection := GetMongoDb().Collection("dsl_Info")
	_, err := collection.InsertOne(ctx, dslInfo)
	if err != nil {
		return err
	}
	return nil
}

func ListDslInfo(ctx context.Context, page, size int, path, name, method, content string, id int64, dslInfoList *[]model.DslInfoStruct) (error, int64) {
	collection := GetMongoDb().Collection("dsl_Info")
	var query bson.M

	if path != "" {
		query["path"] = path
	}
	if name != "" {
		query[name] = name
	}
	if method != "" {
		query[method] = method
	}
	if content != "" {
		query[content] = content
	}

	if id != 0 {
		query["_id"] = id
	}

	var count int64
	findoptions := options.Find()
	if size != 0 && page != 0 {
		offset := (page - 1) * size
		findoptions.SetLimit(int64(size))
		findoptions.SetSkip(int64(offset))

	}
	cur, err := collection.Find(ctx, query, findoptions)
	if err != nil {
		return err, 0
	}
	for cur.Next(ctx) {
		var elem model.DslInfoStruct
		err := cur.Decode(&elem)
		if err != nil {
			return err, 0
		}
		*dslInfoList = append(*dslInfoList, elem)
		count++
	}
	defer cur.Close(ctx)
	return nil, count
}
