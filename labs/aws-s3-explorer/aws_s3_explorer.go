package main

import(
	"fmt"
	"os"
	"net/http"
	"encoding/xml"
	"log"
	"io/ioutil"
	"bytes"
	"strconv"
	"strings"
)

//Structure to store the content of XML
type ListBucketResult struct {
	Content	[]Contents `xml:"Contents"`
}

//Structure to get info from XML contents
type Contents struct {
	Key	string `xml:"Key"`
}

//Function to get the extensions and count it
func getContent(content []Contents) (int, string){
	directoriesCntr := 0
	noExtension := 0
	extensionsMap := make(map[string]int)
	for _, item := range content{
		name := string(item.Key)
		if(name[len(name)-1] == '/'){
			directoriesCntr++
		} else{
			extension := strings.Split(name, ".")
			if (len(extension) < 2){
				noExtension++
			} else{
				if value, ok := extensionsMap[extension[len(extension)-1]]; ok {
					value++
					extensionsMap[extension[len(extension)-1]] = value
				} else{
					extensionsMap[extension[len(extension)-1]] = 1
				}
			}
		}
	}
	extensionsMap["no-extension"] = noExtension
	extensions := ""
	for key, value := range extensionsMap{
		extensions = extensions + key + "(" + strconv.Itoa(value) + "), "
	}
	extensions = extensions[0:len(extensions)-2]
	return directoriesCntr, extensions
}

//Function that get the http response and translate it to text
func getXML(url string) (string){
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	rawXML, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(rawXML)
}

//Function to extract information from XML
func decodeXML(data string) (int, int, string){
	buf := bytes.NewBufferString(data)
	list := new(ListBucketResult)
	decoded := xml.NewDecoder(buf)

	err := decoded.Decode(list)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	directories, extensions := getContent(list.Content)
	return len(list.Content), directories, extensions
}

func main(){
	if (len(os.Args) < 3){
		fmt.Println("Usage: go run aws_s3_explorer.go --bucket amazon_S3_file")
		return
	}
	bucket := os.Args[1]
	if (bucket != "--bucket"){
		fmt.Println("Incorrect flag")
		fmt.Println("Usage: go run aws_s3_explorer.go --bucket amazon_S3_file")
		return
	}

	url := "https://" + os.Args[2] + ".s3.amazonaws.com"
	xml := getXML(url)
	objects, directories, extensions := decodeXML(xml)
	fmt.Println("AWS S3 Explorer")
	fmt.Printf("Bucket Name\t\t: %s\n", os.Args[2])
	fmt.Printf("Number of objects\t: %d\n", objects)
	fmt.Printf("Number of directories\t: %d\n", directories)
	fmt.Printf("Extensions\t\t: %s\n", extensions)
}
