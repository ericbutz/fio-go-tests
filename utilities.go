package smoketest

import (
	"context"
	"time"

	"github.com/dapixio/fio-go"
	"github.com/dapixio/stress-testing/simulate"
)

func createAccount(userCount int, faucetKey string, url string) (failedCount int, owner *fio.Account, users []*fio.Account, err error) {
	// create keys for the owner and users
	//log.Println("creating keys for users")
	newOwner, users, _, err := simulate.NewOrg(userCount)
	if err != nil {
		return userCount, nil, nil, err
	}

	//log.Println("establishing connection to " + url)
	api, _, err := fio.NewConnection(newOwner.KeyBag, url)
	if err != nil {
		failedCount = userCount
		return
	}

	// closure for waiting on transactions
	wait := func(txid string) bool {
		var block uint32
		blockChan := make(chan uint32)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		go func() {
			block, err = api.WaitForConfirm(api.GetCurrentBlock()-4, txid)
			if err != nil {
				failedCount = userCount
			}
			blockChan <- block
		}()
		for {
			select {
			case block = <-blockChan:
				if block == 1 {
					return false
				}
				//log.Printf("\tfound txid %s in block %d\n", txid, block)
				return true
			case <-ctx.Done():
				return false
			}
		}
	}

	// get cost
	cost, _ := simulate.GetOrgCost(api)

	// fund the new owner:
	//log.Printf("funding owner %s with ᵮ%.9f\n", newOwner.PubKey, cost)
	txid, err := simulate.FundOwner(faucetKey, newOwner.PubKey, cost)
	if err != nil {
		return userCount, nil, nil, err
	}
	if done := wait(txid); !done {
		return
	}
	return 0, newOwner, users, err
}
