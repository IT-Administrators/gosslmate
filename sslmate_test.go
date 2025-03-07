package gosslmate

import (
	"fmt"
	"io/fs"
	"log"
	"net/url"
	"os"
	"path"
	"testing"
)

var sslmp = NewSslMateQuery("Test.com", true, true, true, true, true, true)
var testfiles = "./examples"

// Get all files in specified directory.
func listFiles(dir string) []string {
	root := os.DirFS(dir)
	// Get only files ending with "".json".
	mdFiles, err := fs.Glob(root, "*.json")

	if err != nil {
		log.Fatal(err)
	}

	var files []string
	for _, v := range mdFiles {
		files = append(files, path.Join(dir, v))
	}
	return files
}

// Read file and return byte array.
func ReadFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func TestInvokeHttpGet(t *testing.T) {
	// var files = listFiles(testfiles)
	// sslm := []sslMate{}
	cont := ReadFile("./examples/json01.json")
	res := convertToJson(cont)
	fmt.Println(res[0])

	// for file := range files {
	// 	// sslm := []sslMate{}
	// 	cont := ReadFile(files[file])
	// 	res := convertToJson(cont)
	// 	// fmt.Println(sslm)
	// 	fmt.Println(res.ID)
	// }
	// fmt.Println(sslm)
	// sslm.buildUri()
	// res := invokeHttpGet(sslm.uriString)
	// // Unmarshal json result and safe to struct.
	// if string(res) == "" {
	// 	t.Error("No response.")
	// }
	// fmt.Println(string(res))
	// fmt.Println(convertToJson(res))
}

func TestBuildUri(t *testing.T) {
	_, err := url.Parse(sslmp.buildUri().getUriString())
	if err != nil {
		t.Error(err)
	}
}
