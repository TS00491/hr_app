package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type BankStatements struct {
	PaymentMonth      string `json:"paymentMonth"`
	PayslipNumber     int    `json:"payslipNumber"`
	BankStatementDate string `json:"bankStatementDate"`
}

func CreateBankStatement(client *mongo.Client, bankStatement *BankStatements) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bankStatementCollection := client.Database("hrapp").Collection("bankstatements")
	result, err := bankStatementCollection.InsertOne(ctx, bankStatement)

	return result, err
}

func GetAllBankStatements(client *mongo.Client) ([]BankStatements, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bankStatementCollection := client.Database("hrapp").Collection("bankstatements")
	cursor, err := bankStatementCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var bankStatements []BankStatements
	err = cursor.All(ctx, &bankStatements)
	if err != nil {
		return nil, err
	}

	return bankStatements, nil
}

func GetBankStatement(client *mongo.Client, id string) (*BankStatements, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bankStatementCollection := client.Database("hrapp").Collection("bankstatements")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	var bankStatement BankStatements
	err = bankStatementCollection.FindOne(ctx, filter).Decode(&bankStatement)
	if err != nil {
		return nil, err
	}

	return &bankStatement, nil
}
