package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	// "crypto/ecdsa"
    // "crypto/elliptic"
    // "crypto/rand"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
)

func main() {
	// 替换以下变量值
	rpcURL := "https://bsc-testnet-rpc.publicnode.com"
	privateKeyHex := "d5dea133cca06eaf5676b13eb8213db2663c719aa0800548d86ccd22edc99fcd"
	recipientAddressHex := "0xE7b1C5d28312d0c742e70A711fc6e26dDDbaEc62"
	tokenContractAddressHex := "0x75929C42E385661469F563b5EcfE52AE72505e47"
	amountToSend := big.NewInt(1000000000000000000) // 例如发送1个代币（假设代币有18位小数）

	// 创建客户端
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// 加载私钥
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Unable to load private key: %v\n", err)
	}
	fmt.Println("privateKey:",privateKey)

	// 获取chainID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Unable to fetch chain ID: %v\n", err)
	}


	// 获取发起人的地址
	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*crypto.PublicKey)
	// if !ok {
	// 	log.Fatal("Error casting public key to ECDSA")
	// }
	// fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// fmt.Println("fromAddress:",fromAddress)

	// 构建auth
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Unable to create transaction author: %v\n", err)
	}

	// 构造转账数据
	tokenContractAddress := common.HexToAddress(tokenContractAddressHex)
	recipientAddress := common.HexToAddress(recipientAddressHex)
	data := []byte(nil)
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := crypto.Keccak256Hash(transferFnSignature).Bytes()[:4]
	data = append(data, hash...)
	data = append(data, common.LeftPadBytes(recipientAddress.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(amountToSend.Bytes(), 32)...)

	// 估算Gas
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &tokenContractAddress,
		Data: data,
	})
	if err != nil {
		log.Fatalf("Unable to estimate gas: %v\n", err)
	}

	// 获取Gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Unable to suggest gas price: %v\n", err)
	}

	// 创建交易
	tx := types.NewTransaction(auth.Nonce.Uint64(), tokenContractAddress, big.NewInt(0), gasLimit, gasPrice, data)

	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalf("Unable to sign transaction: %v\n", err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Unable to send transaction:")
	}
}




