package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// Model for business
type Business struct {
	Name     string     `json:"name"`
	Address  string     `json:"address"`
	Employee []Employee `json:"employee"`
}

func NewBusiness() *Business {
	return &Business{}
}

func GetAllBusinesses(client *mongo.Client) ([]Business, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	businessCollection := client.Database("hrapp").Collection("businesses")

	cursor, err := businessCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var businesses []Business
	err = cursor.All(ctx, &businesses)
	if err != nil {
		return nil, err
	}

	return businesses, nil
}
func CreateBusiness(client *mongo.Client, business *Business) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	businessCollection := client.Database("hrapp").Collection("businesses")
	result, err := businessCollection.InsertOne(ctx, business)

	return result, err
}
