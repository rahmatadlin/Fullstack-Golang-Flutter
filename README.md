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




























1. Prepare
	1. Demo App
	2. Attention
	3. Tools	
2. Setup Backend
	1. Starting Backend Project
	2. Go Routing Image Access
	3. Model Table Backend
	4. Database Config
	5. Test Migrate
3. User RestFull Api
	1. Login Api
	2. Create Account Api
	3. Delete User Api
	4. Get List Employee Api
4. Task RestFull Api
	1. Create Task Api
	2. Delete Task Api
	3. Delete Cascade Example
	4. Submit Task Api
	5. Reject Task Api
	6. Fix Task Api
	7. Approve Task Api
	8. Find Task Api
	9. Get Need Review Task Api
	10. Get Progress Task Api
	11. Get Statistic Task Api
	12. Get Task Where User and Status Api
5. Setup Flutter
	1. Starting Flutter Project (android, assets, pubspec)
	2. Setup Common Use
	3. Model
	4. User Cubit
	5. Setup main
6. User Source
	1. Login Source
	2. Create Account Source
	3. Delete User Source
	4. Get List Employee Source
7. Task Source
	1. Create Task Source
	2. Delete Task Source
	3. Submit Task Source
	4. Reject Task Source
	5. Fix Task Source
	6. Approve Task Source
	7. Find Task Source
	8. Need Review Task Source
	9. Progress Task Source
	10. Statistic Task Source
	11. Where User and Status Task Source
8. Login
	1. Login Header
	2. Login Input
	3. App Button Primary
	4. Login Cubit
9. Home Admin
	1. Admin Header
	2. Button Add Employee
	3. Need Review BLoC
	4. Build Need Review
	5. Build Item Need Review
	6. Failed UI
	7. Employee BLoC
	8. Build Employee
	9. Build Item Employee
	10. Add Employee
	11. Delete Employee
10. Profile
	1. Profile AppBar
	2. Profile Header
	3. Item Menu and Logout
11. Monitoring Employee
	1. Build Monitoring Header
	2. Build Button Add Task
	3. Stat Employee Cubit
	4. Build Task Menu
	5. Build Item Task Menu
	6. Progress Task BLoC
	7. Build Progress Task
	8. Build Item Progress Task
	9. Monitoring Navigation
12. Add New Task
	1. Add Task Input
	2. Date Time Picker
	3. Add New Task Implementation
13. List Task
	1. List Task BLoC
	2. List Task Header
	3. Build List Task
	4. Build Item Task
14. Home Employee
	1. Employee Header
	2. Build Search
	3. Build Task Menu and Progress Task
15. Detail Task
	1. Detail Task Cubit
	2. Detail Task Builder
	3. Build Detail Menu
	4. Build Status
	5. Build Description and Reason
	6. Build Details
	7. Build Attachment
	8. Handle null date
	9. Build Submit
	10. Build Reject
	11. Delete Task
	12. Submit Task
	13. Reject Task
	14. Fix Task to Queue
	15. Approve Task
16. Fixing Bug
	1. Fix Bug Login