package main

import (

    "gorm.io/gorm"
	"gorm.io/driver/mysql"
	"log"

)

// MVC = Model View Controller
// User model
type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"unique"`
    Email    string
}

var db *gorm.DB		// Database

func connectToMariaDB() (*gorm.DB, error) {
    dsn := "root:@tcp(localhost:3306)/mydb_test?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}

func main() {
    db, err := connectToMariaDB()
    if err != nil {
        log.Fatal(err)
    }
    // defer db.Close()

    // Perform database migration
    err = db.AutoMigrate(&User{})
    if err != nil {
        log.Fatal(err)
    }

    // Create a user
    newUser := &User{Username: "john_doe", Email: "john.doe@example.com"}
    err = createUser(db, newUser)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Created User:", newUser)

    // Query user by ID
    userID := newUser.ID
    user, err := getUserByID(db, userID)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("User by ID:", user)

    // Update user
    user.Email = "updated_email@example.com"
    err = updateUser(db, user)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Updated User:", user)

    // Delete user
    err = deleteUser(db, user)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Deleted User:", user)
}


// ฟังก์ชัน GetUsers ใช้สำหรับเรียกข้อมูล user ทั้งหมด
func getUserByID(db *gorm.DB, userID uint) (*User, error) {
    var user User
    result := db.First(&user, userID)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}


// ฟังก์ชัน CreateUser ใช้สำหรับสร้าง user
func createUser(db *gorm.DB, user *User) error {
    result := db.Create(user)
	if result.Error != nil {
    log.Fatal(result.Error)
}
    return nil
}


// ฟังก์ชัน UpdateUser ใช้สำหรับแก้ไข user
func updateUser(db *gorm.DB, user *User) error {
    result := db.Save(user)
    if result.Error != nil {
        return result.Error
    }
    return nil
}


// ฟังก์ชัน DeleteUser ใช้สำหรับลบ user
func deleteUser(db *gorm.DB, user *User) error {
    result := db.Delete(user)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

