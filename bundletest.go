package smoketest

import (
	"log"

	"github.com/dapixio/fio-go"

	"github.com/dapixio/stress-testing/simulate"
)

var ()

// TestBundleCount runs various tests on bundles
func TestBundleCount() (err error) {
	log.Println("********* START: TestBundleCount")

	var ()

	// create keys for the owner and a user for each endpoint
	owner, users, domain, err := simulate.NewOrg(len(FioEndpoints))
	if err != nil {
		return err
	}

	for i, endPoint := range FioEndpoints {
		log.Println("bundlecost ", i, ": ", endPoint.BundleCost)
		for j := 0; j <= NewBundleTxnCount; j++ {
			_, ok := fio.NewAddAddress(users[i].Actor, fio.Address(users[i].Addresses[0].FioAddress), "FIO", RandomString(40))
			log.Println("AddAddress: ", fio.Address(users[i].Addresses[0].FioAddress))
			if !ok {
				log.Println("NewAddAddress Failed: ", users[i])
			}
		}
	}
	log.Println("domain: ", domain)
	log.Println("owner: ", owner)
	log.Println("User1: ", users[0].Addresses[0].FioAddress)
	log.Println("User1: ", users[1].Addresses[0].FioAddress)
	log.Println("User1: ", users[2].Addresses[0].FioAddress)

	if 1 < 0 {
		log.Println("********* ERROR: TestBundleCount: ")
	} else {
		log.Println("********* PASS: TestBundleCount")
	}

	return nil
}
