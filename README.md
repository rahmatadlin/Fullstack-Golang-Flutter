# Fullstack-Golang-Flutter

## SERVER SIDE

1. **Create Server Folder**

   - How to create server folder

   ```bash
   mkdir server
   ```

   - Navigate to server folder

   ```bash
   cd server
   ```

2. **Intialize**

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

   - Or we can use this

   ```
   go get -u gorm.io/gorm
   ```

   - Then copy and paste this

   ```
   go get -u gorm.io/driver/mysql
   ```

3. **Create main.go file**

   - How to add file

   ```
   touch main.go
   ```

    - How to run
    ```
    go run main.go
    ```

4. **Setup PORT**
 
    #### Due to restrictions, localhost cannot be used. Please utilize the IPv4 Address instead.

    - Open cmd and then type this
    ```
    ipconfig
    ```
    - Create .env for your IPv4 Address

    ```
    touch .env
    ```

    - then copy and paste in .env file

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

     - Run again
    ```
    go run main.go
    ```
