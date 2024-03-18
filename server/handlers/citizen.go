package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaswanthsaivendra/citizen_small_data_store/server/database"
	"github.com/yaswanthsaivendra/citizen_small_data_store/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCitizens(c *gin.Context) {
	cursor, err := database.Citizens.Find(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unalbe to fetch citizens"})
		return
	}

	var citizens []models.CitizenModel
	if err = cursor.All(c, &citizens); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch products"})
		return
	}

	c.JSON(http.StatusOK, citizens)

}

func AddCitizen(c *gin.Context) {
	var body models.CitizenSerializer
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	res, err := database.Citizens.InsertOne(c, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add citizen"})
		return
	}

	citizen := models.CitizenModel{
		ID:          res.InsertedID.(primitive.ObjectID),
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		DateOfBirth: body.DateOfBirth,
		Gender:      body.Gender,
		Address:     body.Address,
		City:        body.City,
		State:       body.State,
		Pincode:     body.Pincode,
	}

	c.JSON(http.StatusCreated, citizen)
}

func GetCitizenById(c *gin.Context) {
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	res := database.Citizens.FindOne(c, primitive.M{"_id": _id})
	citizen := models.CitizenModel{}
	err = res.Decode(&citizen)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to find citizen"})
		return
	}
	c.JSON(http.StatusOK, citizen)

}

func UpdateCitizenById(c *gin.Context) {
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var body models.CitizenSerializer

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	_, err = database.Citizens.UpdateOne(c, bson.M{"_id": _id}, bson.M{"$set": body})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to update citizen"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "citizen updated"})
}

func DeleteCitizenById(c *gin.Context) {
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	res, err := database.Citizens.DeleteOne(c, bson.M{"_id": _id})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to delete citizen"})
		return
	}

	if res.DeletedCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "citizen not found"})
	}
	c.JSON(http.StatusOK, gin.H{"success": "citizen deleted"})
}
