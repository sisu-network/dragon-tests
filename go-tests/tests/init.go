package tests

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

const LOCALHOST_MNEMONIC = "draft attract behave allow rib raise puzzle frost neck curtain gentle bless letter parrot hold century diet budget paper fetch hat vanish wonder maximum"

var (
	defaultAccounts []accounts.Account
	localWallet     *hdwallet.Wallet
	client          *ethclient.Client
)

func init() {
	var err error
	client, err = ethclient.Dial("http://localhost:1234")
	if err != nil {
		panic(err)
	}

	fmt.Println("Creating an account....")
	localWallet, err = hdwallet.NewFromMnemonic(LOCALHOST_MNEMONIC)
	defaultAccounts = createAccounts(LOCALHOST_MNEMONIC, 10)
	fmt.Println("Done Creating an account....")
}

func createAccounts(mnemonic string, count int) []accounts.Account {
	var err error
	accounts := make([]accounts.Account, count)

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < count; i++ {
		path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", i))
		account, err := wallet.Derive(path, true)
		if err != nil {
			panic(err)
		}
		accounts[i] = account
	}

	wallet.Close()

	return accounts
}
