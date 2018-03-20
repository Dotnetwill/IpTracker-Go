package main

import ( 
	"net/http" 
  "time"
  "fmt"
	"strings"
	"os"
	"io/ioutil"
	
  "github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	)

func main() {
	
	if len(os.Args) < 2 {
		fmt.Println("usage: %s <dropbox_token>", os.Args[0])
		return
	}
	
	token := os.Args[1]
	
	// Check with amazon webservices the current IP
	// http://checkip.amazonaws.com
	
	resp, err := http.Get("http://checkip.amazonaws.com")
	if err != nil {
		fmt.Println(os.Stderr, "failed to print IP: %s", err)
		return
	}
	
	defer resp.Body.Close()
	
	ip, err := ioutil.ReadAll(resp.Body)
	
	config := dropbox.Config{
		Token: token,
	}
	
	commitInfo := files.NewCommitInfo("/home-ip.txt")
	commitInfo.Mode.Tag = "overwrite"

	// The Dropbox API only accepts timestamps in UTC with second precision.
	commitInfo.ClientModified = time.Now().UTC().Round(time.Second)

	dbx := files.New(config)
	fileContents := fmt.Sprintf("%s \r\n %s", ip, time.Now())
	
	if _, err := dbx.Upload(commitInfo, strings.NewReader(fileContents)); err != nil {
		fmt.Println(os.Stderr, "failed to upload %s", err)
	}
	
}