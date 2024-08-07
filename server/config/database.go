package config

import (
    "fmt"
    "log"
    "os"
    "server/models"

    "github.com/joho/godotenv"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
}

func DatabaseConnection() *gorm.DB {
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
        user, password, host, port, dbName,
    )

    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    return database
}

func CreateOwnerAccount(db *gorm.DB) {
    hashedPasswordBytes, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
    owner := models.User{
        Role:     "Admin",
        Name:     "Owner",
        Password: string(hashedPasswordBytes),
        Email:    "owner@go.id",
    }

    // Condition if email already exists
    if db.Where("email=?", owner.Email).First(&owner).RowsAffected == 0 {
        db.Create(&owner)
    } else {
        fmt.Println("Owner exists")
    }
}
