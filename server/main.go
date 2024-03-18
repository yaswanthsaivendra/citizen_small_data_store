package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yaswanthsaivendra/citizen_small_data_store/server/database"
	"github.com/yaswanthsaivendra/citizen_small_data_store/server/handlers"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading.env file")
	}
}

func main() {
	loadEnv()
	err := database.InitDB(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to MongoDB")

	defer func() {
		err := database.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("disconnected from MongoDB")
	}()

	router := gin.Default()

	router.GET("/citizens", handlers.GetCitizens)
	router.POST("/citizens", handlers.AddCitizen)

	router.GET("/citizens/:id", handlers.GetCitizenById)
	router.PUT("/citizens/:id", handlers.UpdateCitizenById)
	router.DELETE("/citizens/:id", handlers.DeleteCitizenById)

	router.Run(":8000")
	fmt.Println("server running on port 8000")

}
