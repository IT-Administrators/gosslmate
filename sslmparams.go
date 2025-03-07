package gosslmate

import (
	"net/url"
)

// Parameter for query string configuration. The struct properties are public to be able to change them.
// Example: struct.ShowIssuer = false
type sslMateParam struct {
	Domain                   string `Domain:"example.com"`
	SearchSubDomains         bool   `SearchSubDomains:"true"`
	ShowDnsNames             bool   `ShowDnsNames:"true"`
	ShowIssuer               bool   `ShowIssuer:"true"`
	ShowIssuerWebsite        bool   `ShowIssuerWebsite:"false"`
	ShowIssuerCaaDomains     bool   `ShowIssuerCaaDomains:"false"`
	ShowIssuerOperator       bool   `ShowIssuerOperator:"false"`
	ShowIssuerPubkeyDer      bool   `ShowIssuerPubkeyDer:"false"`
	ShowIssuerNameDer        bool   `ShowIssuerNameDer:"false"`
	ShowRevocationInfo       bool   `ShowRevocationInfo:"true"`
	ShowProblemReportingInfo bool   `ShowProblemReportingInfo:"true"`
	ShowCertData             bool   `ShowCertData:"true"`
	MatchWildcards           bool   `MatchWildcards:"false"`
	ShowPubKeyDer            bool   `PubKeyDer:"false"`
	ShowPubKey               bool   `PubKey:"false"`

	uriString string
}

// Create new query parameter.
func NewSslMateQuery(Domain string) *sslMateParam {

	return &sslMateParam{
		Domain:                   Domain,
		SearchSubDomains:         true,
		ShowDnsNames:             true,
		ShowIssuer:               true,
		ShowRevocationInfo:       true,
		ShowProblemReportingInfo: true,
		ShowCertData:             true,
	}
}

// This method is used to get the non public uristing. This contains the uri for the initial query.
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
	if sslm.ShowIssuerWebsite {
		newQ.Add("expand", "issuer.website")
	}
	if sslm.ShowIssuerCaaDomains {
		newQ.Add("expand", "issuer.caa_domains")
	}
	if sslm.ShowIssuerPubkeyDer {
		newQ.Add("expand", "issuer.pubkey_der")
	}
	if sslm.ShowIssuerNameDer {
		newQ.Add("expand", "issuer.name_der")
	}
	if sslm.ShowIssuerOperator {
		newQ.Add("expand", "issuer.operator")
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
	if sslm.MatchWildcards {
		newQ.Add("match_wildcards", "true")
	}
	if sslm.ShowPubKeyDer {
		newQ.Add("expand", "pubkey_der")
	}
	if sslm.ShowPubKey {
		newQ.Add("expand", "pubkey")
	}
	result.RawQuery = newQ.Encode()
	// Write urinstring to self.
	sslm.uriString = result.String()
	return sslm
}
