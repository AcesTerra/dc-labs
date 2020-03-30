package main

import (
	"crypto/rand"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"image"
	"os"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"time"

)

// User: Structure to store users info
type User struct{
	Username string
	Password string
	Token string
}

// Image: Structure to store image info
type Image struct{
	Name string
	Token string
}

// Global variable user. All functions are able to access to it
var user User

func main() {
	// User and password registered
	user.Username="aidan1"
	user.Password="123abc"
	router := mux.NewRouter()
	//All routes for API
	router.HandleFunc("/Login",login)
	router.HandleFunc("/Logout",logout)
	router.HandleFunc("/Status",status)
	router.HandleFunc("/Upload",addImage)
	router.HandleFunc("/Token",token)
	http.ListenAndServe(":5000",router)
}

// Receive image and return name and size of it
func addImage(w http.ResponseWriter, r *http.Request) {
	var img Image
	json.NewDecoder(r.Body).Decode(&img)
	if img.Token == user.Token {
		json.NewEncoder(w).Encode(struct {
			Message string
			Filename string
			Size int64
		}{Message:"An image has been successfully uploaded", Filename:img.Name, Size:getImageSize(img.Name)	})
	}else{
		json.NewEncoder(w).Encode(struct {
			Message string
		}{Message:"Wrong token"})
	}
}

// Login and show token
func login(w http.ResponseWriter, r *http.Request) {
	var userw User
	json.NewDecoder(r.Body).Decode(&userw)
	if user.Username == userw.Username && user.Password == userw.Password{
		user.Token=tokenGenerator()
		json.NewEncoder(w).Encode(struct {
			Message string
			Token string
	}{Message:"Hi username, welcome to the DPIP System",Token:user.Token})
	}else{
		json.NewEncoder(w).Encode(struct {
			Message string
		}{Message:"wong username or password"})
	}
}

// Logout and erase token
func logout(w http.ResponseWriter, r *http.Request) {
	user.Token=""
	json.NewEncoder(w).Encode(struct {
		Message string
	}{Message:"Bye username, your token has been revoked"})
}

// Retrive time of API request
func status(w http.ResponseWriter, r *http.Request) {
	var userw User
	json.NewDecoder(r.Body).Decode(&userw)
	if userw.Token== user.Token{
		current := time.Now()
		json.NewEncoder(w).Encode(struct {
			Message string
			Time string
		}{Message:"Hi username, the DPIP System is Up and Running",Time: current.Format("2006-01-02 15:04:05")})
	}else{
		json.NewEncoder(w).Encode(struct {
			Message string
		}{Message:"wrong token"})
	}
}
// Generate token
func tokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// Retrieve image size
func getImageSize(imagePath string) (int64){
	fi, err := os.Stat(imagePath);
	if err != nil {
			log.Println(imagePath, err)
	}
	// Get the size
	size := fi.Size()
	return size
}

// Get image dimensions
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
// Test to see generated token
func token(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(struct {
			Token string
		}{Token:user.Token})
}
