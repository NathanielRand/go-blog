package models


import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// ErrNotFound is returned when a resource cannot be found // in the database.
	ErrNotFound = errors.New("models: resource not found")
)

type User struct {
	Username string `gorm:"not null;unique"`
	Email string `gorm:"not null;unique_index"`
}

// UserService type to provide methods 
// to query, create and update users.
type UserService struct {
	db *gorm.DB
}

// NewUserService
func NewUserService(connectionInfo string) (*UserService, error) {
	// Connect to db via GORM.
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		panic(err)
	}
	// Enable logmode.
	db.LogMode(true)
	// Return a UserService object.
	return &UserService{
		db: db,
	}, nil
}

// Close method to close the UserService connection.
func (us *UserService) Close() error {
	return us.db.Close()
}

// ByID will look up a user with the provided ID.
// If the user is found, we will return a nil error
// If the user is not found, we will return ErrNotFound
func (us *UserService) ByID(id uint) (*User, error) {
	// user variable of type User.
	var user User
	// Query db for record with provided ID via gorm.
	err := us.db.Where("id = ?", id).First(&user).Error
	// Switch for error check. Returns a
	// pointer to the user queired if nil.
	switch err {
		case nil:
			return &user, nil
		case gorm.ErrRecordNotFound:
			return nil, ErrNotFound
		default:
			return nil, err
	}
}

// DestructiveReset drops the user table and rebuilds it
// *Should remove in production.
func (us *UserService) DestructiveReset() { 
	us.db.DropTableIfExists(&User{}) 
	us.db.AutoMigrate(&User{})
}