package smoketest

import (
	"log"

	"github.com/dapixio/stress-testing/simulate"
)

// TestBadDomains tests an array of bad addresses from testdata.go to make sure they fail on regaddress
func TestBadDomains() {
	log.Println("********* START: TestBadAddresses")

	var (
		failedCount int = 0
		domainList  string
	)

	// create keys for the owner and users
	newOwner, _, _, err := simulate.NewOrg(1)
	if err != nil {
		return
	}

	// Test characters in the IllegalCharSet
	for i := 0; i < len(IllegalCharSet); i++ {
		badDomain := RandomString(5) + string(IllegalCharSet[i]) + RandomString(5)
		_, err := simulate.RegDomain(newOwner, badDomain)
		//log.Println("Register domain failed: ", badDomain)
		if err == nil {
			//log.Println("Error: ", err.Error())
			failedCount++
			domainList = domainList + "  " + badDomain
		}
	}

	if failedCount > 0 {
		log.Println("********* ERROR: Bad domains successfully registered: " + domainList)
	} else {
		log.Println("********* PASS: TestBadAddresses")
	}

	return
}
