package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Employee struct {
	Fname             string           `json:"fname"`
	Lname             string           `json:"lname"`
	Address           string           `json:"address"`
	PhoneNumber       string           `json:"phonenumber"`
	Email             string           `json:"email"`
	Position          string           `json:"position"`
	Startdate         string           `json:"startdate"`
	Enddate           string           `json:"enddate"`
	Salary            float32          `json:"salary"`
	RightToWork       bool             `json:"righttowork"`
	PassportImage     string           `json:"passportimage"`
	NINumber          string           `json:"ninumber"`
	CurrentlyEmployed bool             `json:"currently_employed"` //Checks if the employee is a current employee or an ex-employee, if ex, then put in the previous employee list.
	JobDesc           string           `json:"job_desc"`
	Payslip           []Payslip        `json:"payslip"`
	SponsorLicence    []SponsorLicence `json:"sponsorlicence"`
	Absences          []Absences       `json:"absences"`
	BankStatements    []BankStatements `json:"bankStatements"`
}

func NewEmployee() *Employee {
	return &Employee{}
}

func GetAllEmployees(client *mongo.Client) ([]Employee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	employeeCollection := client.Database("hrapp").Collection("employees")

	cursor, err := employeeCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var employees []Employee
	err = cursor.All(ctx, &employees)
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func CreateEmployee(client *mongo.Client, employee *Employee) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	employeeCollection := client.Database("hrapp").Collection("employees")
	result, err := employeeCollection.InsertOne(ctx, employee)

	return result, err
}

func GetEmployeeByID(client *mongo.Client, id string) (*Employee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	employeeCollection := client.Database("hrapp").Collection("employees")
	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var employee Employee
	err = employeeCollection.FindOne(ctx, bson.M{"_id": employeeID}).Decode(&employee)
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func UpdateEmployee(client *mongo.Client, id string, updatedEmployee *Employee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	employeeCollection := client.Database("hrapp").Collection("employees")
	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = employeeCollection.UpdateOne(
		ctx,
		bson.M{"_id": employeeID},
		bson.M{"$set": updatedEmployee},
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteEmployee(client *mongo.Client, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	employeeCollection := client.Database("hrapp").Collection("employees")
	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = employeeCollection.DeleteOne(ctx, bson.M{"_id": employeeID})
	if err != nil {
		return err
	}

	return nil
}
