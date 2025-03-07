package gosslmate

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const apiUri = "https://api.certspotter.com/v1/issuances"

type issuer struct {
	Friendly_name string
}

type revocation struct {
	Time       string
	Reason     string
	Checked_At string
}

// Create struct to consume json. Each parameter from request restult matches struct property.
type sslMate struct {
	ID          string
	Tbs_sha256  string
	Cert_sha256 string
	// If jsonobject is array. "dns_names":["*.example.com","example.com"]
	Dns_names     []string
	Pubkey_sha256 string
	Issuer        issuer
	Not_before    string
	Not_after     string
	Revoked       bool
	// Nested object. "revocation":{"time":null,"reason":null,"checked_at":"2025-01-18T11:06:32Z"},
	Revocation        revocation
	Problem_reporting string
	Cert_der          string
}

// Parameter for query string configuration.
type sslMateParam struct {
	Domain                   string `Domain:"example.com"`
	SearchSubDomains         bool   `SearchSubDomains:"true"`
	ShowDnsNames             bool   `ShowDnsNames:"true"`
	ShowIssuer               bool   `ShowIssuer:"true"`
	ShowRevocationInfo       bool   `ShowRevocationInfo:"true"`
	ShowProblemReportingInfo bool   `ShowProblemReportingInfo:"true"`
	ShowCertData             bool   `ShowCertData:"true"`
	uriString                string
}

// Create new query parameter.
func NewSslMateQuery(Domain string, seachSubDomains bool, showDnsNames bool, showIssuer bool, showRevocationInfo bool, showProblemReportingInfo bool, showCertData bool) *sslMateParam {

	return &sslMateParam{
		Domain:                   Domain,
		SearchSubDomains:         seachSubDomains,
		ShowDnsNames:             showDnsNames,
		ShowIssuer:               showIssuer,
		ShowRevocationInfo:       showRevocationInfo,
		ShowProblemReportingInfo: showProblemReportingInfo,
		ShowCertData:             showCertData,
	}
}

// Get the saved uri string.
func (sslm *sslMateParam) getUriString() string {
	return sslm.uriString
}

// Create the uristring depending on the set parameters in sslMateParam struct.
// Use getUriString() to show string.
func (sslm *sslMateParam) buildUri() *sslMateParam {

	result, _ := url.Parse(apiUri)
	// Create new query.
	newQ := result.Query()
	// Set query parameter. This parameter is mandatory.
	newQ.Set("domain", sslm.Domain)
	if sslm.SearchSubDomains {
		newQ.Add("include_subdomains", "true")
	}
	if sslm.ShowDnsNames {
		newQ.Add("expand", "dns_names")
	}
	if sslm.ShowIssuer {
		newQ.Add("expand", "issuer")
	}
	if sslm.ShowRevocationInfo {
		newQ.Add("expand", "revocation")
	}
	if sslm.ShowProblemReportingInfo {
		newQ.Add("expand", "problem_reporting")
	}
	if sslm.ShowCertData {
		newQ.Add("expand", "cert_der")
	}
	result.RawQuery = newQ.Encode()
	sslm.uriString = result.String()
	return sslm
}

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
