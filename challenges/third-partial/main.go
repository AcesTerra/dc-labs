package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strconv"
	//"github.com/CodersSquad/dc-labs/challenges/third-partial/controller"
	//"github.com/CodersSquad/dc-labs/challenges/third-partial/scheduler"
	"github.com/AcesTerra/controller"
	"github.com/AcesTerra/scheduler"
)

// User struct to store users info.
type user struct {
	username string
	password string
	token    string
}

// Global variable users were users are stored
var allUsers []user

// Upload image to server. Returns a JSON with filename and itws size.
func uploadImage(c *gin.Context) {
	bearer := c.Request.Header["Authorization"]
	token := bearer[0]
	splitedToken := strings.Split(token, " ")
	t := string(splitedToken[1])
	var userToken string
	isRegistered := false
	for _,v := range allUsers{
		if (v.token == t){
			isRegistered = true
			userToken = v.token
		}
	}
	if (isRegistered){
		file, err := c.FormFile("data")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "Error uploading image", "error": err.Error})
			return
		}
		t := time.Now()
		time := t.Format("20060102150405")
		uploadedImage := filepath.Base(file.Filename)
		fileName := userToken + "_" + time + "_" + filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, fileName); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "Error uploading image", "error": err.Error})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message":"An image has been successfully uploaded", "filename": uploadedImage, "size": strconv.FormatInt(file.Size, 10) + "bytes"})
	}
}

// Login for users. Returns a JSON with username and token generated.
func login(c *gin.Context) {
	var newUser user
	userRegistered := false
	userName, password, _ := c.Request.BasicAuth()
	token := tokenGenerator()
	for i,v := range allUsers{
		if (v.username == userName){
			allUsers[i].token = token
			c.JSON(http.StatusOK, gin.H{"message":"New token generated", "token": token})
			userRegistered = true
		}
	}
	if (userRegistered == false){
		newUser.username = userName
		newUser.password = password
		newUser.token = token
		allUsers = append(allUsers, newUser)
		c.JSON(http.StatusOK, gin.H{"message":"Hi " + userName + ", welcome to the DPIP System", "token": token})
	}
}

// Revoke user token and delete user from list.
func logout(c *gin.Context) {
	bearer := c.Request.Header["Authorization"]
	token := bearer[0]
	splitedToken := strings.Split(token, " ")
	t := string(splitedToken[1])
	var userIndex int
	var userName string
	isRegistered := false
	for i,v := range allUsers{
		if (v.token == t){
			userIndex = i
			userName = v.username
			isRegistered = true
		}
	}
	if (isRegistered){
		allUsers[userIndex] = allUsers[len(allUsers)-1]
		allUsers = allUsers[:len(allUsers)-1]
		c.JSON(http.StatusOK, gin.H{"message":"Bye " + userName + ", your token has been revoked"})
	}
}

// Status of login. Returns a JSON with username and time.
func status(c *gin.Context) {
	bearer := c.Request.Header["Authorization"]
	token := bearer[0]
	splitedToken := strings.Split(token, " ")
	t := string(splitedToken[1])
	for _,v := range allUsers{
		if (v.token == t){
			current := time.Now()
			c.JSON(http.StatusOK, gin.H{"message":"Hi " + v.username + ", the DPIP System is Up and Running", "time": current.Format("2006-01-02 15:04:05")})
		}
	}
}

// Generate user tokens.
func tokenGenerator() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// Check worker status
func workerStatus(c *gin.Context){
	worker := c.Param("worker")
	tags,status,usage := controller.WorkerStatus(worker)
	c.JSON(http.StatusOK, gin.H{"Worker":worker, "Tags":tags, "Status":status, "Usage":usage})
}

// Test job over worker
func workerTest(c *gin.Context){
	//controller.workerTest()
}

func main() {
	log.Println("Welcome to the Distributed and Parallel Image Processing System")

	r := gin.Default()
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"username":"password",
		"alfredo":"ecole",
	}))

	authorized.GET("/login", login)
	r.GET("/status", status)
	r.GET("/logout", logout)
	r.POST("/upload", uploadImage)
	r.GET("/status/:worker", workerStatus)
	r.GET("/workloads/test", workerTest)

	//Start API
	go r.Run() // Listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// Start Controller
	go controller.Start()

	// Start Scheduler
	jobs := make(chan scheduler.Job)
	go scheduler.Start(jobs)
	// Send sample jobs
	sampleJob := scheduler.Job{Address: "localhost:50051", RPCName: "hello"}

	for {
		b := make([]byte, 4)
		rand.Read(b)
		sampleJob.RPCName = fmt.Sprintf("hello-%v", b)
		jobs <- sampleJob
		time.Sleep(time.Second * 5)
	}
}
