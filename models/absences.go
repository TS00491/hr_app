package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// Model for Employee Absences incl holidays.
type Absences struct {
	Holidays                    string `json:"holiday"`
	UnauthorisedAbsences        string `json:"unauth_absences"`
	AuthorisedAbsences          string `json:"auth_absences"`
	UnauthorisedAbsencesReasons string `json:"unauth_absences_reasons"`
	Notes                       string `json:"notes"`
}

func CreateAbsence(client *mongo.Client, absence *Absences) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	absencesCollection := client.Database("hrapp").Collection("absences")
	result, err := absencesCollection.InsertOne(ctx, absence)

	return result, err
}

func GetAllAbsences(client *mongo.Client) ([]Absences, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	absencesCollection := client.Database("hrapp").Collection("absences")

	cursor, err := absencesCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var absences []Absences
	err = cursor.All(ctx, &absences)
	if err != nil {
		return nil, err
	}

	return absences, nil
}

func GetAbsence(client *mongo.Client, id string) (*Absences, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	absencesCollection := client.Database("hrapp").Collection("absences")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var absence Absences
	err = absencesCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&absence)
	if err != nil {
		return nil, err
	}

	return &absence, nil
}
