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
	"github.com/AcesTerra/dc-labs/challenges/final/controller"
	"github.com/AcesTerra/dc-labs/challenges/final/scheduler"
	"github.com/gin-contrib/static"
	//"github.com/gin-gonic/gin"
)

// Channels used for jobs
var jobs = make(chan scheduler.Job)
var rpcChan = make(chan string)

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
	bearer := c.Request.Header["Authorization"]
	token := bearer[0]
	splitedToken := strings.Split(token, " ")
	t := string(splitedToken[1])
	//var userToken string
	isRegistered := false
	for _,v := range allUsers{
		if (v.token == t){
			isRegistered = true
			//userToken = v.token
		}
	}
	if (isRegistered){
		sampleJob := scheduler.Job{Address: "localhost:50051", RPCName: "test"}
		jobs <- sampleJob
		rpcResponse := <- rpcChan
		//fmt.Println(rpcResponse)
		//time.Sleep(time.Second * 5)
		splitedResponse := strings.Split(rpcResponse, ";")
		c.JSON(http.StatusOK, gin.H{"Workload":splitedResponse[0], "Job ID":splitedResponse[3], "Status":splitedResponse[1], "Usage":splitedResponse[2]})
	}
}

func filter(c *gin.Context){
	bearer := c.Request.Header["Authorization"]
	token := bearer[0]
	splitedToken := strings.Split(token, " ")
	t := string(splitedToken[1])
	//var userToken string
	isRegistered := false
	for _,v := range allUsers{
		if (v.token == t){
			isRegistered = true
			//userToken = v.token
		}
	}
	if (isRegistered){
		workloadId,_ := c.GetPostForm("workload-id")
		filter,_ := c.GetPostForm("filter")
		//file,_ := c.GetPostForm("data")
		//fmt.Println("Entered function")
		file, err := c.FormFile("data")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "Error uploading image", "error": err.Error})
			return
		}
		fileName := filepath.Base(file.Filename)
		//uploadedImage := filepath.Base(file.Filename)
		fileName = "worker/" + filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, fileName); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "Error uploading image", "error": err.Error})
			return
		}
		//fmt.Println(workloadId)
		//fmt.Println(filter)
		jobID := controller.GetWorkloadJobId(workloadId)
		strJobID := strconv.Itoa(jobID)
		//fmt.Println(jobID)
		//sampleJob := scheduler.Job{Address: "localhost:50051", RPCName: "test"}
		//sampleJob := scheduler.Job{Address: "localhost:50051", RPCName: "test"}
		//jobs <- sampleJob
		c.JSON(http.StatusOK, gin.H{"Workload ID": workloadId, "Filter":filter, "Job ID":strJobID, "Status":"Scheduling", "Results":"http://localhost:8080/results/", "File":fileName})
	}
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
	r.POST("/workloads/filter", filter)
	/*r.Use(static.Serve("/results", static.LocalFile("/tmp", false)))
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "test")
	})*/
	//r.Static("/results", "./home/alfredo/dc-labs/challenges/final/results")
	//r.StaticFS("/results", http.Dir("my_file_system"))

	r.Use(static.Serve("/results", static.LocalFile("/home/alfredo/dc-labs/challenges/final/results", true)))
	//router.Use(static.Serve("/", static.LocalFile("/results", true)))

	// Start Controller
	go controller.Start()

	// Start Scheduler
	go scheduler.Start(jobs, rpcChan)

	//Start API
	r.Run(":8080") // Listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
