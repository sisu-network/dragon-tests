package main

import "github.com/sisu-network/dragon-tests/tests"

func main() {
	tests.TestChangeName("Change name in contract")
	tests.TestConcurrentTransfers("Test concurrent transfer", 1)
	tests.TestErc20("Test minting ERC20")
}
