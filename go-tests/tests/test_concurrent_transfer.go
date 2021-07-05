package tests

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/sisu-network/dragon-tests/assert"
	"github.com/sisu-network/dragon-tests/utils"
)

// Tests
func testConcurrentTransfer(title string) {
	recipient := defaultAccounts[0]
	wg := &sync.WaitGroup{}

	beforeBalance, err := client.BalanceAt(context.Background(), recipient.Address, nil)
	if err != nil {
		panic(err)
	}

	amount := new(big.Int).Set(utils.ONE_ETHER)
	amount = amount.Div(utils.ONE_ETHER, big.NewInt(100)) // 0.01 ether
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(account accounts.Account) {
			transfer(account, recipient, amount, wg)
		}(defaultAccounts[i])
	}

	wg.Wait()

	time.Sleep(time.Second * 6) // Wait for block finality
	afterBalance, err := client.BalanceAt(context.Background(), recipient.Address, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("afterBalance = ", afterBalance)

	diff := afterBalance.Sub(afterBalance, beforeBalance)
	// Diff should be 0.09 ether
	expectedResult := big.NewInt(90000000000000000)
	assert.Equal(expectedResult, diff, "The balance difference should be correct")

	utils.MarkTestPassed(title)
}

func TestConcurrentTransfers(title string, round int) {
	for i := 0; i < round; i++ {
		testConcurrentTransfer(title)
		time.Sleep(time.Second * 2)
	}
}
