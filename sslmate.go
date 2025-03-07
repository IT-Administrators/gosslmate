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

// Function to issue new get request to apiUri.
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

// Convert the []byte response from the invokeHttpGet method to json object.
// This way the result can be used with obj.parameter.
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

// Query CT logs. The response is an array of []ssLMate even when the rest response is not a json array.
// Use res[i].paramter to get the parameter value.
func GetCtLogs(sslmq sslMateParam) []sslMate {
	// Build uri from parameter.
	sslmq.buildUri()
	// Invoke request.
	queryres := invokeHttpGet(sslmq.getUriString())
	// Return json object.
	return convertToJson(queryres)
}
