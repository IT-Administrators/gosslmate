package gosslmate

import "net/url"

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
