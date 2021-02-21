package handler

import (
	"fiber-gorm/database"
	"fiber-gorm/model"

	"github.com/gofiber/fiber/v2"
)

// User struct stands for the user model
type User model.User

// GetUsers func handle request to serve all users
func GetUsers(c *fiber.Ctx) error {
	db, err := database.GetDBConn()
	if err != nil {
		panic(err)
	}

	var users []User
	db.Find(&users)
	if len(users) == 0 {
		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"message": "No Users",
			},
		})
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"users": users,
		},
	})
}

// GetUser func handle request get a single user
func GetUser(c *fiber.Ctx) error {
	db, err := database.GetDBConn()
	if err != nil {
		panic(err)
	}

	id := c.Params("id")
	var user User
	result := db.First(&user, id)
	if result.RowsAffected == 0 {
		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"message": "No User Found With Given ID",
			},
		})
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"user": user,
		},
	})
}

// CreateUser func handle request to create a new user
func CreateUser(c *fiber.Ctx) error {
	db, err := database.GetDBConn()
	if err != nil {
		panic(err)
	}

	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"message": "There is Something Wrong",
			},
		})
	}

	db.Create(&user)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": fiber.Map{
			"message": "Successfully Created a New User",
		},
	})
}

// UpdateUser func handle request to update existing user data
func UpdateUser(c *fiber.Ctx) error {
	db, err := database.GetDBConn()
	if err != nil {
		panic(err)
	}

	data := new(User)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"message": "There is Something Wrong",
			},
		})
	}

	id := c.Params("id")
	var user User
	result := db.First(&user, id)
	if result.RowsAffected == 0 {
		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"message": "No User Found With Given ID",
			},
		})
	}

	user.Firstname = data.Firstname
	user.Lastname = data.Lastname
	user.Email = data.Email

	db.Save(&user)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": fiber.Map{
			"message": "Successfully Updated a User",
		},
	})
}

// DeleteUser func handle request to deleting existing user
func DeleteUser(c *fiber.Ctx) error {
	db, err := database.GetDBConn()
	if err != nil {
		panic(err)
	}

	id := c.Params("id")
	var user User
	result := db.First(&user, id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"message": "No User Found With Given ID",
			},
		})
	}

	db.Delete(&user)
	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"message": "Successfully Deleted",
		},
	})
}
