package main

import (
	"fmt"
	"log"
	"os"

	"postgre-handson/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
	type Repository struct {
		DB *gorm.DB
	}
*/

type Update map[string]string

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	//DB Connection

	config := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Payment{}, &models.User{})

	//Server
	app := fiber.New()

	//Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Home")
	})

	app.Post("/api/createPayment", func(c *fiber.Ctx) error {
		jd := new(models.Payment)
		c.Accepts("application/json")
		if err := c.BodyParser(&jd); err != nil {
			fmt.Println(err)
			return c.JSON(fiber.Map{"Error": err.Error()})
		}
		createPayment(db, *jd)
		return c.SendString("POST action completed")
	})
	//partial match
	app.Get("/api/getUserByName", func(c *fiber.Ctx) error {
		m := c.Queries()
		var p []models.User
		if getUserByName(db, &p, m["user_name"]) {
			return c.JSON(p)
		} else {
			return c.JSON(fiber.Map{"message": "Not found"})
		}
	})

	app.Get("/api/getUserById/:id", func(c *fiber.Ctx) error {
		m := c.Params("id")
		p := new(models.User)
		d, _ := uuid.Parse(m)
		if getUserById(db, p, d) {
			return c.JSON(p)
		} else {
			return c.JSON(fiber.Map{"message": "Not found"})
		}
	})

	app.Get("/api/getAllUsers", func(c *fiber.Ctx) error {
		var p []models.User
		m := c.Queries()
		if getAllUsersfilter(db, &p, m) {
			return c.JSON(p)

		} else {
			return c.JSON(fiber.Map{"message": "Error while fetching"})
		}

	})

	app.Get("/api/getAllPayments", func(c *fiber.Ctx) error {
		var p []models.Payment
		m := c.Queries()
		_, isPresent := m["created"]
		if isPresent {
			return c.JSON(fiber.Map{"message": "Invalid query parameter"})
		}
		if getAllPaymentsfilter(db, &p, m) {
			return c.JSON(p)
		} else {
			return c.JSON(fiber.Map{"message": "Error while fetching"})
		}

	})
	app.Listen(":8000")

	//GORM
	//UserCreation
	/*
			for i := 12; i < 16; i++ {
				createUser(db, &models.User{UserName: fmt.Sprint("Qwerty", i)})
			}
		//PaymentCreation
			d, _ := uuid.Parse("696ec5e6-23fd-4573-9f2d-aa9a867583a7")
			payment := models.Payment{UserId: d, PaymentMode: "Cash", SuccessStatus: true}
			createPayment(db, payment)
	*/

	//User Fetch and Update
	/*
		var user models.User
		d, _ := uuid.Parse("8756fb7b-1b33-47ca-b65d-90da8dc24b3f")
		if getUserById(db, &user, d) {
			updateUserName(db, &user, "UpdatedAgain")
		}
	*/
	//User Fetch and Delete
	/*
		var user models.User
		d, _ := uuid.Parse("f7ba2cea-fea1-472a-98c9-f4ab188ac199")
		if getUserById(db, &user, d) {
			deleteUser(db, &user)
		}
	*/
}

// Controllers
func createPayment(db *gorm.DB, payment models.Payment) bool {
	return commonStatus(db.Create(&payment), 1, "Payment Creation")

}

func createUser(db *gorm.DB, user *models.User) bool {
	return commonStatus(db.Create(&user), 1, "User Creation")

}

func getUserById(db *gorm.DB, u *models.User, ui uuid.UUID) bool {
	return commonStatus(db.First(&u, ui), 1, "User Fetch")

}

func getAllUsersfilter(db *gorm.DB, u *([]models.User), f map[string]string) bool {
	d := db.Where(f).Find(&u)
	return d.Error == nil

}
func getAllPaymentsfilter(db *gorm.DB, u *([]models.Payment), f map[string]string) bool {
	d := db.Where(f).Find(&u)
	return d.Error == nil
}

func getUserByName(db *gorm.DB, u *([]models.User), s string) bool {
	d := db.Where("user_name LIKE ?", fmt.Sprintf("%%%s%%", s)).Find(&u)
	return d.Error == nil

}
func updateUserName(db *gorm.DB, u *models.User, s string) bool {
	return commonStatus(db.Model(&u).Update("user_name", s), 1, "UserName Update")

}

func deleteUser(db *gorm.DB, u *models.User) bool {
	return commonStatus(db.Delete(&u), 1, "User Deletion")
}

func commonStatus(tx *gorm.DB, i int, s string) bool {
	if tx.RowsAffected == int64(i) {
		fmt.Println("Successful", s)
		return true
	}
	fmt.Println(s, "is not successful")
	return false
}

//Unused
/*
func getAllUsers(db *gorm.DB, u *([]models.User)) bool {
	d := db.Find(&u)
	return d.Error == nil

}*/
