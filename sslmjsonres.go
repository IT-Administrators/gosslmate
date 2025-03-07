package gosslmate

type issuer struct {
	Friendly_name string
	Website       string
	Caa_domains   []string
	Operator      operator
	Pubkey_sha256 string
	Pubkey_der    string
	Name          string
	Name_der      string
}

type operator struct {
	Name    string
	Website string
}

type pubkey struct {
	Type       string
	Bit_length int
	Curve      string
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
	Pubkey_der    string
	Pubkey        pubkey
	Issuer        issuer
	Not_before    string
	Not_after     string
	Revoked       bool
	// Nested object. "revocation":{"time":null,"reason":null,"checked_at":"2025-01-18T11:06:32Z"},
	Revocation        revocation
	Problem_reporting string
	Cert_der          string
}
