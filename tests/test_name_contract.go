package tests

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	name "github.com/sisu-network/dragon-tests/tests/contracts/name"
	"github.com/sisu-network/dragon-tests/utils"
)

func TestChangeName() {
	fromAccount := localAccounts[0]
	privateKey, err := localWallet.PrivateKey(fromAccount)
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := client.NonceAt(context.Background(), fromAccount.Address, nil)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	_, tx, instance, err := name.DeployName(auth, client)
	if err != nil {
		panic(err)
	}
	_, err = bind.WaitDeployed(context.Background(), client, tx)

	// Get the name.
	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	name, err := instance.GetName(callOpts)
	if err != nil {
		panic(err)
	}

	utils.Assert("Hello", name, "name is correct")

	// Set name
	nonce, err = client.NonceAt(context.Background(), fromAccount.Address, nil)
	auth.Nonce = big.NewInt(int64(nonce))

	tx, err = instance.SetName(auth, "NewName")
	if err != nil {
		panic(err)
	}
	fmt.Println("Tx = ", tx.Hash())

	time.Sleep(time.Second * 2)

	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		panic(err)
	}
	utils.Assert(receipt.Status, uint64(1), "Transaction should be successful")

	fmt.Println("Test Passed")
}
