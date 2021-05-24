package tests

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sisu-network/dragon-tests/assert"
	"github.com/sisu-network/dragon-tests/tests/contracts/erc20"
	"github.com/sisu-network/dragon-tests/utils"
)

func TestErc20(title string) {
	_, auth := getBasicInfo()
	contractAddr, tx, instance, err := erc20.DeployErc20(auth, client, "Test Token", "TEST")
	auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
	if err != nil {
		panic(err)
	}
	_, err = bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		panic(err)
	}
	utils.LogDebug("contractAddr = ", contractAddr)

	// mint tokens for 10 known default accounts
	for _, account := range defaultAccounts {
		instance.Mint(auth, account.Address, big.NewInt(10))
		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
	}

	utils.LogDebug("Done minting")

	// Helper function to get ERC20 balance
	getBalance := func(addr common.Address) *big.Int {
		callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
		balance, err := instance.BalanceOf(callOpts, addr)
		if err != nil {
			panic(err)
		}

		return balance
	}

	// 10 random accounts, each send to 10 recipients.
	recipients := createAccounts("indoor dish desk flag debris potato excuse depart ticket judge file exit", 10)

	// Send ERC20 tokens from each default account to 10 different recipients (100 tx in total).
	for i := 0; i < len(defaultAccounts); i++ {
		go func(index int) {
			utils.LogDebug("Transfer token from account index", index)

			var accountAuth *bind.TransactOpts
			if index == 0 {
				accountAuth = auth
			} else {
				accountAuth = getAuth(index)
			}

			for _, recipient := range recipients {
				_, err := instance.Transfer(accountAuth, recipient.Address, big.NewInt(1))
				if err != nil {
					panic(err)
				}

				accountAuth.Nonce = accountAuth.Nonce.Add(accountAuth.Nonce, big.NewInt(1))
			}
		}(i)
	}

	utils.LogDebug("All requests sent")
	time.Sleep(time.Second * 10)
	utils.LogDebug("Done transfering")

	// Check the final balance.
	for _, recipient := range recipients {
		balance := getBalance(recipient.Address)
		assert.Equal(big.NewInt(10), balance, "Balance of each recipient is correct")
	}
}
