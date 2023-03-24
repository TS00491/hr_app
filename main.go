package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hrapp/models"
	"hrapp/routes"
	"os"
	"time"
)

func main() {
	// Connect to MongoDB
	mongoURI := os.Getenv("MONGO_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		fmt.Println("Error creating MongoDB client:", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer client.Disconnect(ctx)

	business := models.NewBusiness()
	employee := models.NewEmployee()
	payslip := models.NewPayslip()
	sponsorlicence := models.NewSponsorLicence()

	fmt.Println(business, employee, payslip, sponsorlicence)

	router := gin.Default()
	routes.SetupRoutes(router, client)

	fmt.Println("Server running on http://localhost:8080")
	err = router.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
