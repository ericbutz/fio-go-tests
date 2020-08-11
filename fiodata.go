package smoketest

const (
	LegalCharSet      = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	LegalInternalChar = "-"
	IllegalCharSet    = "`~!@#$%^&*()_+={[}]|:;'<,>.?/"
	NewBundleTxnCount = 50
)

var (
	// FioEndpoints is an array of API endpoints
	FioEndpoints = []struct {
		Name       string
		Fee        int // in SUF
		BundleCost int
	}{
		{"add_pub_address", 400000000, 1},
		{"reject_funds_request", 400000000, 1},
		{"vote_producer", 400000000, 1},
		{"proxy_vote", 400000000, 1},
		{"new_funds_request", 800000000, 2},
		{"record_obt_data", 800000000, 2},
	}
)
