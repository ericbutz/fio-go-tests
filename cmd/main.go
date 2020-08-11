package main

import (
	"github.com/dapixio/smoketest"
)

func main() {

	smoketest.TestBadDomains()
	smoketest.TestBundleCount()

}
