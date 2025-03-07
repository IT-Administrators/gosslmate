package gosslmate

import (
	"io/fs"
	"log"
	"net/url"
	"os"
	"path"
	"testing"
)

var sslmp = NewSslMateQuery("Test.com")
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
func readSslMateFiles(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
func TestBuildUri(t *testing.T) {
	_, err := url.Parse(sslmp.buildUri().getUriString())
	if err != nil {
		t.Error(err)
	}
}

func TestInvokeHttpGet(t *testing.T) {
	res := invokeHttpGet(sslmp.getUriString())
	if len(res) == 0 {
		t.Error(res)
	}
}

func TestConvertToJson(t *testing.T) {
	res := invokeHttpGet(sslmp.getUriString())
	jsonres := convertToJson(res)
	if len(jsonres) == 0 {
		t.Error(jsonres)
	}
}

func TestGetCtLogs(t *testing.T) {
	jsonres := GetCtLogs(*sslmp)
	if len(jsonres) == 0 {
		t.Error(jsonres)
	}
}
