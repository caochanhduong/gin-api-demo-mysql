package controllers

import (
	"demo-gin-api-with-gomod/database"
	"demo-gin-api-with-gomod/models"

	"github.com/gin-gonic/gin"
)

// GetAllUser reads all information about user in database
func GetAllUser(c *gin.Context) {
	db := database.DBConn()

	results, err := db.Query("SELECT * FROM user")
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	users := make([]models.User, 0)
	for results.Next() {
		var user models.User
		err := results.Scan(&user.ID, &user.Name)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
		users = append(users, models.User{
			ID:   user.ID,
			Name: user.Name,
		})
	}
	c.JSON(200, gin.H{
		"result": users,
	})

	db.Close()
}

// GetUserByID gets information of a user in database
func GetUserByID(c *gin.Context) {
	db := database.DBConn()

	var user models.User
	err := db.QueryRow("SELECT * FROM user WHERE id = ?", c.Param("id")).Scan(&user.ID, &user.Name)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"result": user,
	})

	db.Close()
}

// DeleteUserByID deletes information of a user in database
func DeleteUserByID(c *gin.Context) {
	db := database.DBConn()

	result, err := db.Exec("DELETE FROM user WHERE id = ?", c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		num, err := result.RowsAffected()
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"nums row affected": num,
			})
		}
	}

	db.Close()
}

// CreateUser creates information of a user in database
func CreateUser(c *gin.Context) {
	db := database.DBConn()

	var user models.User
	if err := c.ShouldBindJSON(&user); err == nil {
		post, err := db.Prepare("INSERT INTO user(id, name) VALUES (?,?)")
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
		post.Exec(user.ID, user.Name)
		c.JSON(200, gin.H{
			"messages": "inserted",
		})
	} else {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	db.Close()
}

// UpdateUserByID updates information of a user in database
func UpdateUserByID(c *gin.Context) {
	db := database.DBConn()

	var user models.User
	if err := c.ShouldBindJSON(&user); err == nil {
		post, err := db.Prepare("UPDATE user SET name=? WHERE id=" + c.Param("id"))
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
		post.Exec(user.Name)
		c.JSON(200, gin.H{
			"messages": "updated",
		})
	} else {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	db.Close()
}
