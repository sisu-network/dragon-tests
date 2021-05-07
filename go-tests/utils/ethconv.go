package utils

import "math/big"

var (
	ONE_ETHER       = big.NewInt(1000000000000000000) // in wei
	ONE_ETHER_FLOAT = new(big.Float).SetInt(ONE_ETHER)
)

func WeiToEthFloat(wei *big.Int) *big.Float {
	weiFloat := new(big.Float).SetInt(wei)

	return new(big.Float).Quo(weiFloat, ONE_ETHER_FLOAT)
}
