package main

import (
	"bufio"
	"fmt" 
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" 
)

const (
	host = "localhost"
	port = 5432
	user = "nathanielrand" 
	password = "" 
	dbname = "goblog"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique"`
	Email string `gorm:"not null;unique_index"`
	Orders []Order
	// Password string
}

type Order struct { 
	gorm.Model
	UserID uint 
	Amount int 
	Description string
}

func main() {
	// Format and return results from const as values.
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	
	// Open GORM SQL connections with psqlInfo.
	db, err := gorm.Open("postgres", psqlInfo) 
	if err != nil {
		panic(err) 
	}

	// Defer close of db.
	defer db.Close() 

	// Enable LogMode for GORM.
	db.LogMode(true)

	// Auto-migrate to create tables in db.
	db.AutoMigrate(&User{}, &Order{})	

	// // Get user information and store values.
	// username, email := getInfo()

	// // Create instance of User type passing values from name, email variables.
	// u := &User{
	// 	Username: username,
	// 	Email: email,
	// }

	// // Create a new user via GORM with the User instance values.
	// if err = db.Create(u).Error; err != nil {
	// 	panic(err)
	// }

	// // Print values of User instance.
	// fmt.Printf("%+v\n", u)

	// // Select first user from the db.
	// var uF User 
	// db.First(&uF)
	// if db.Error != nil { 
	// 	panic(db.Error)
	// } 
	// fmt.Println("First Record: ", uF)

	// // Select user with an id of 2 from the db.
	// var u2 User
	// id := 2
	// db.First(&u2, id)
	// if db.Error != nil {
	// 	panic(db.Error) 
	// }
	// fmt.Println("User with ID: ", u2)

	// // Select the first user with an id less than 3 from the db.
	// var u3 User
	// maxId := 3
	// db.Where("id <= ?", maxId).First(&u3) 
	// if db.Error != nil {
	// 	panic(db.Error) 
	// }
	// fmt.Println("User with ID less than 3: ", u3)

	// // Query all records with GORM
	// var users []User 
	// db.Find(&users)
	// if db.Error != nil {
	// 	panic(db.Error) 
	// }
	// fmt.Println("Retrieved", len(users), "users.") 
	// // fmt.Println(users)

	// Assign user variable to the User variable with an id of 2.
	// var user User 
	// db.First(&user, 2)
	// if db.Error != nil {
	// 	panic(db.Error) 
	// }

	// // Create mock order data.
	// createOrder(db, user, 1001, "Fake Description #1")
  	// createOrder(db, user, 9999, "Fake Description #2")
	// createOrder(db, user, 8800, "Fake Description #3")
	

	var user User 
	db.Preload("Orders").First(&user) 
	if db.Error != nil {
		panic(db.Error) 
	}
	fmt.Println("Email:", user.Email) 
	fmt.Println("Number of orders:", len(user.Orders)) 
	fmt.Println("Orders:", user.Orders)


}

func getInfo() (username, email string) { 
	reader := bufio.NewReader(os.Stdin) 

	fmt.Println("What is your username?") 
	username, _ = reader.ReadString('\n') 
	username = strings.TrimSpace(username) 

	fmt.Println("What is your email?") 
	email, _ = reader.ReadString('\n') 
	email = strings.TrimSpace(email) 

	return username, email
}

func createOrder(db *gorm.DB, user User, amount int, desc string) { 
	// Create order in order database
	db.Create(&Order{
		UserID: user.ID, 
		Amount: amount, 
		Description: desc,
	})
	// Error check
	if db.Error != nil {
		panic(db.Error) 
	}
}