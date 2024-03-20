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

// func corsMiddleware() gin.HandlerFunc {
// 	return
// }

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
	// router.Use(func(c *gin.Context) {
	// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	// 	if c.Request.Method == "OPTIONS" {
	// 		c.Writer.WriteHeader(http.StatusOK)
	// 		return
	// 	}

	// 	c.Next()
	// })

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173/")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	router.GET("/citizens", handlers.GetCitizens)
	router.POST("/citizens", handlers.AddCitizen)

	router.GET("/citizens/:id", handlers.GetCitizenById)
	router.PUT("/citizens/:id", handlers.UpdateCitizenById)
	router.DELETE("/citizens/:id", handlers.DeleteCitizenById)

	router.Run(":8000")
	fmt.Println("server running on port 8000")

}
