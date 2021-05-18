package services

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
	"log"
	"os"
)

func LoadVars() (string, string, string, string, string, string, string, string) {

	viper.SetConfigFile("./configs/config.env")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}

	ebsiHost := viper.GetString("EbsiHost")
	didIdentifier := viper.GetString("DIDIdentifier")
	didDocumentVersion := viper.GetString("DIDDocumentVersion")
	timestampData := viper.GetString("TimestampData")
	didVersionMetadata := viper.GetString("DIDVersionMetadata")
	privateKeyHex := viper.GetString("PrivateKeyHex")
	method := viper.GetString("Method")
	trustedIssuerVersion := viper.GetString("TrustedIssuerVersion")

	return ebsiHost, didIdentifier, didDocumentVersion, timestampData, didVersionMetadata, privateKeyHex, method, trustedIssuerVersion
}

func LoadKeys(privateKeyHex string) (ecdsa.PrivateKey, ecdsa.PublicKey, common.Address) {
	privK, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		panic(err.Error())
	}

	//Public key its derived from the privKey
	pubK := privK.PublicKey

	//Address its generated from pubK. Address = Keccak-256(pubK),and then we take the last 40 characters(20 bytes) and add 0x
	address := crypto.PubkeyToAddress(pubK)

	return *privK, pubK, address
}

func WriteOutput(fileName string, message string) {
	f, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(message)

	if err2 != nil {
		log.Fatal(err2)
	}
}
