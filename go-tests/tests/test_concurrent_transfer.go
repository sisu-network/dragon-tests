package tests

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sisu-network/dragon-tests/assert"
	"github.com/sisu-network/dragon-tests/utils"
)

func transfer(fromAccount accounts.Account, toAddress accounts.Account, amount *big.Int, wg *sync.WaitGroup) {
	nonce, err := client.NonceAt(context.Background(), fromAccount.Address, nil)
	if err != nil {
		panic(err)
	}

	gasLimit := uint64(210000) // in units
	gasPrice := big.NewInt(75)

	var data []byte
	txData := &types.AccessListTx{
		ChainID:  chainId,
		Nonce:    nonce,
		To:       &toAddress.Address,
		Value:    amount,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	}

	prvKey, err := localWallet.PrivateKey(fromAccount)
	if err != nil {
		panic(err)
	}

	signedTx, err := types.SignNewTx(prvKey, types.NewEIP2930Signer(chainId), txData)
	if err != nil {
		panic(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		panic(err)
	}

	if wg != nil {
		wg.Done()
	}
}

// Tests
func TestConcurrentTransfers(title string) {
	recipient := localAccounts[0]
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
		}(localAccounts[i])
	}

	wg.Wait()

	time.Sleep(time.Second * 1) // Wait for block finality
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
