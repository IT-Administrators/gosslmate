package gosslmate

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

const apiUri = "https://api.certspotter.com/v1/issuances"

// Validate json object.
func isValidJson(b []byte) bool {
	return json.Valid(b)
}

// Function to issue new get request.
func invokeHttpGet(uri string) []byte {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		log.Fatalf("Error creating HTTP request %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request %v", err)
	}
	// Close connection.
	defer res.Body.Close()
	// Create bytes to return.
	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading HTTP request %v", err)
	}
	if !isValidJson(responseBytes) {
		log.Fatal("Response is not valid json.")
	}
	// Return type is bytes. Can be resolved with string(responsebytes).
	return responseBytes
}

func convertToJson(b []byte) []sslMate {
	// Create sslMate object to save response as json.
	sslm := []sslMate{}
	if strings.HasPrefix(string(b), "{") {
		res := sslMate{}
		e := json.Unmarshal(b, &res)
		if e != nil {
			log.Fatal(e)
		}
		sslm = append(sslm, res)
		return sslm
	} else {
		// var res []sslMate
		e := json.Unmarshal(b, &sslm)

		if e != nil {
			log.Fatal(e)
		}
		return sslm
	}
}

// Main function to query sslmate logs.
func GetCtLogs(sslmq sslMateParam) []sslMate {
	sslmq.buildUri()
	queryres := invokeHttpGet(sslmq.getUriString())
	return convertToJson(queryres)
}
