package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"hrapp/models"
	"net/http"
)

func SetupRoutes(router *gin.Engine, client *mongo.Client) {
	//Business
	router.POST("/business", func(c *gin.Context) { createBusiness(c, client) })
	router.GET("/businesses", func(c *gin.Context) { getAllBusinesses(c, client) }) // New route
	router.GET("/business/:id", func(c *gin.Context) { getBusiness(c, client) })

	//Employees
	router.POST("/employee", func(c *gin.Context) { createEmployee(c, client) })
	router.GET("/employees", func(c *gin.Context) { getAllEmployees(c, client) }) // New route
	router.GET("/employee/:id", func(c *gin.Context) { getEmployee(c, client) })

	//Payslips
	router.POST("/payslip", createPayslip)
	router.GET("/payslips", getAllPayslips) // New route
	router.GET("/payslip/:id", getPayslip)

	// Absences routes
	router.POST("/absence", func(c *gin.Context) { createAbsence(c, client) })
	router.GET("/absences", func(c *gin.Context) { getAllAbsences(c, client) })
	router.GET("/absence/:id", func(c *gin.Context) { getAbsence(c, client) })

	// BankStatements routes
	router.POST("/bankstatement", func(c *gin.Context) { createBankStatement(c, client) })
	router.GET("/bankstatements", func(c *gin.Context) { getAllBankStatements(c, client) })
	router.GET("/bankstatement/:id", func(c *gin.Context) { getBankStatement(c, client) })
}

func createBusiness(c *gin.Context, client *mongo.Client) {
	var business models.Business
	if err := c.ShouldBindJSON(&business); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Received business data:", business) // Print the received data

	result, err := models.CreateBusiness(client, &business)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Inserted business with ID:", result.InsertedID) // Print the result of the insertion

	c.JSON(http.StatusOK, gin.H{"message": "Business created"})
}

func getAllBusinesses(c *gin.Context, client *mongo.Client) {
	businesses, err := models.GetAllBusinesses(client)
	if err != nil {
		fmt.Println("Error fetching businesses:", err) // Print the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(businesses) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No businesses found"})
		return
	}

	fmt.Println("Fetched businesses:", businesses) // Print the fetched businesses
	c.JSON(http.StatusOK, businesses)
}

func getBusiness(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"id":      id,
		"message": "Get business",
	})
}

// Employees
// Employees
func createEmployee(c *gin.Context, client *mongo.Client) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := models.CreateEmployee(client, &employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee created", "id": result.InsertedID})
}

func getAllEmployees(c *gin.Context, client *mongo.Client) {
	employees, err := models.GetAllEmployees(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employees)
}

func getEmployee(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")

	employee, err := models.GetEmployeeByID(client, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

// Payslips
func getAllPayslips(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get all payslips",
	})
}

func createPayslip(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create payslip",
	})
}

func getPayslip(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"id":      id,
		"message": "Get payslip",
	})
}

// Absences
func createAbsence(c *gin.Context, client *mongo.Client) {
	var absence models.Absences
	if err := c.ShouldBindJSON(&absence); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := models.CreateAbsence(client, &absence)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Absence created", "id": result.InsertedID})
}

func getAllAbsences(c *gin.Context, client *mongo.Client) {
	absences, err := models.GetAllAbsences(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, absences)
}

func getAbsence(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")

	absence, err := models.GetAbsence(client, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, absence)
}

// Bank Statements
func createBankStatement(c *gin.Context, client *mongo.Client) {
	var bankStatement models.BankStatements
	if err := c.ShouldBindJSON(&bankStatement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := models.CreateBankStatement(client, &bankStatement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bank statement created", "id": result.InsertedID})
}

func getAllBankStatements(c *gin.Context, client *mongo.Client) {
	bankStatements, err := models.GetAllBankStatements(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bankStatements)
}

func getBankStatement(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")

	bankStatement, err := models.GetBankStatement(client, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bankStatement)
}
