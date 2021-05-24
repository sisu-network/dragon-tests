package tests

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func getBasicInfo() (accounts.Account, *bind.TransactOpts) {
	fromAccount := defaultAccounts[0]

	return fromAccount, getAuth(0)
}

func getAuth(accountIndex int) *bind.TransactOpts {
	account := defaultAccounts[accountIndex]
	privateKey, err := localWallet.PrivateKey(account)
	if err != nil {
		panic(err)
	}

	nonce, err := client.NonceAt(context.Background(), account.Address, nil)
	if err != nil {
		panic(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		panic(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(5000000) // in units
	auth.GasPrice = gasPrice

	return auth
}

func waitForTxReceipt(tx *types.Transaction) (*types.Receipt, error) {
	hash := tx.Hash()

	duration := 5
	waitTime := time.Second * time.Duration(duration)
	expire := time.Now().Add(waitTime)
	for {
		receipt, err := client.TransactionReceipt(context.Background(), hash)
		if err != nil && err.Error() != "not found" {
			panic(err)
		}

		if receipt != nil {
			return receipt, err
		}

		if expire.Before(time.Now()) {
			break
		}

		time.Sleep(time.Second / 2)
	}

	return nil, fmt.Errorf("Cannot get result within %d second", duration)
}

func getBalance(addr common.Address) *big.Int {
	balance, err := client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		panic(err)
	}

	return balance
}

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

func incNonce(auth *bind.TransactOpts) {
	auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
}
