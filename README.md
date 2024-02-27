# Fullstack-Golang-Flutter

SERVER SIDE

1. Create Server Folder
   ```
   mkdir server
   ```
2. Go to server folder
   cd server

- Initialize Golang setup

  ```
  go mod init server
  ```

- Install router package from gin-gonic

  ```
  go get github.com/gin-gonic/gin
  ```

- Install MySQL package from gorm
  ```
  go get -u gorm.io/gorm
  ```
  or
  ```
  go get -u gorm.io/gorm
  ```
  then
  ```
  go get -u gorm.io/driver/mysql
  ```

3. Create main.go file

- How to add file

  ```
  touch main.go
  ```

- How to run
  ```
  go run main.go
  ```

4. Due to restrictions, localhost cannot be used. Please utilize the IPv4 Address instead.

   - Open cmd
     ```
     ipconfig
     ```
   - Create .env for your IPv4 Address

     ```
     touch .env
     ```

     then copy and paste

     ```go
     IPv4_ADDRESS=YOUR IPv4 ADDRESS
     ```

   - Copy and paste the IPv4 Address into the `router.Run()` function.

     ```go
     package main

     import (
         "net/http"

         "github.com/gin-gonic/gin"
     )

     func main() {
         // Create router from gin, will automatically import
         router := gin.Default()
         router.GET("/", func(c *gin.Context) {
             c.JSON(http.StatusOK, "Welcome to Test API")
         })

        // by default 8080
	    // check your ip address in cmd, then type ipconfig
	    // copy & paste IPv4 Address
	    // Make sure put in .env 
	    ipv4Address := os.Getenv("IPv4_ADDRESS")
	    router.Run(ipv4Address + ":8080")
     }
     ```
