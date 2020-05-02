package main

import (
	"crypto/rand"
	//
	//"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
)

// User struct to store users info
type User struct {
	Username string
	Password string
	Token    string
}

// Image is
type Image struct {
	Name  string
	Token string
}

// Global variable users were users are stored
//var users []User
var user User

func main() {
	r := gin.Default()
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"username":"password",
		"alfredo":"ecole",
	}))

	authorized.GET("/login", login)
	r.GET("/status", status)
	r.GET("/logout", logout)
	r.POST("/upload", addImage)
	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// agrega imagen y devulve el nombre y tamaño de la imagen
func addImage(c *gin.Context) {
	token := c.Request.Header["Authorization"]
	t := token[0]
	splitToken := strings.Split(t, "Bearer")
	t = string(splitToken[1])
	//if t==" \u003c"+user.Token+"\u003e" {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	//	}

}

// hace login y muestra el token
func login(c *gin.Context) {
	userName, password, _ := c.Request.BasicAuth()
	fmt.Println(userName)
	fmt.Println(password)
	user.Token = tokenGenerator()
	c.JSON(http.StatusOK, gin.H{"message": "Hi username, welcome to the DPIP System", "token": user.Token})

}

// borra el token
func logout(c *gin.Context) {
	token := c.Request.Header["Authorization"]
	t := token[0]
	splitToken := strings.Split(t, " ")
	t = string(splitToken[1])
	if t == user.Token {
		user.Token = ""
		c.JSON(http.StatusOK, gin.H{"message": "Bye username, your token has been revoked"})
	}
}

// devulve la hora en la que se hace la consulta
func status(c *gin.Context) {
	token := c.Request.Header["Authorization"]
	t := token[0]
	splitedToken := strings.Split(t, " ")
	t = string(splitedToken[1])
	if t == user.Token {
		current := time.Now()
		c.JSON(http.StatusOK, gin.H{"message": "Hi username, the DPIP System is Up and Running", "time": current.Format("2006-01-02 15:04:05")})

	}
}

// genera el token
func tokenGenerator() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// devuelve las dimensiones de la imagen
func getImageSize(imagePath string) int64 {
	fi, err := os.Stat(imagePath)
	if err != nil {
		log.Println(imagePath, err)
	}
	// Get the size
	size := fi.Size()
	return size
}

// nos confunidos y primero sacamos las dimencion de la imagen
func getImageDimension(imagePath string) (int, int) {
	file, err := os.Open(imagePath)
	defer file.Close()
	if err != nil {
		log.Println(imagePath, err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	return image.Width, image.Height
}
