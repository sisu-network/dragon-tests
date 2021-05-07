package main

import "github.com/sisu-network/dragon-tests/tests"

func main() {
	tests.TestChangeName("Change name in contract")
	tests.TestConcurrentTransfers("Test concurrent transfer")
}
