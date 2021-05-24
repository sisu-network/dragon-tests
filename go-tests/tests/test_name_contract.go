package tests

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	name "github.com/sisu-network/dragon-tests/tests/contracts/name"
	"github.com/sisu-network/dragon-tests/utils"
)

var (
	chainId = big.NewInt(1)
)

func TestChangeName(title string) {
	fromAccount, auth := getBasicInfo()

	_, tx, instance, err := name.DeployName(auth, client)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 2)

	contractAddress, err := bind.WaitDeployed(context.Background(), client, tx)
	fmt.Println("contractAddress = ", contractAddress)

	if err != nil {
		panic(err)
	}

	// Get the name.
	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	name, err := instance.GetName(callOpts)
	if err != nil {
		panic(err)
	}

	utils.Assert("Hello", name, "name is correct")

	// Set name
	nonce, err := client.NonceAt(context.Background(), fromAccount.Address, nil)
	auth.Nonce = big.NewInt(int64(nonce))

	tx, err = instance.SetName(auth, "NewName")
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 1)

	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		panic(err)
	}
	utils.Assert(receipt.Status, uint64(1), "Transaction should be successful")

	utils.MarkTestPassed(title)
}
