package services

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gataca-io/ebsi/insert-did-document/models"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	SignedTransactionMethod = "signedTransaction"
	JsonRPCVersion          = "2.0"
	Protocol                = "eth" // eth or fabric
)

func CallEbsiJsonRPCFunction(payload []byte, ebsiHost string, registryEndpoint string) string {
	req, err := http.NewRequest(http.MethodPost, ebsiHost+registryEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("error creating http request: ", string(payload), err)
		return ""
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(payload)))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error connecting to EBSI ", ebsiHost, err)
		return ""
	}

	fmt.Println("Status code: ", resp.StatusCode, "\nMessage: ", resp.Status)

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println("error reading response")
		return ""
	}

	response := buf.String()

	return response
}

func GenerateSignedTransaction(parsedUT *models.UnsignedTransaction, privKey ecdsa.PrivateKey) []byte {

	signature, r, s, v := SignEthereumTx(&privKey, parsedUT.To, parsedUT.Data, parsedUT.Nonce, parsedUT.GasLimit, parsedUT.ChainId)

	stp := models.SignedTxParams{
		Protocol:             Protocol,
		UnsignedTransaction:  parsedUT,
		R:                    hexutil.EncodeBig(r),
		S:                    hexutil.EncodeBig(s),
		V:                    hexutil.EncodeBig(v),
		SignedRawTransaction: signature,
	}

	stx := models.SignedTransaction{
		JsonRPC: JsonRPCVersion,
		Method:  SignedTransactionMethod,
		Id:      1,
		Params:  []models.SignedTxParams{stp},
	}

	payload, err := json.Marshal(stx)
	if err != nil {
		os.Exit(-1)
	}

	return payload
}

func ParseToUnsignedTransaction(payload string) *models.UnsignedTransaction {
	var hfr models.HelperFunctionResult
	err := json.Unmarshal([]byte(payload), &hfr)
	if err != nil {
		os.Exit(-1)
	}

	return hfr.Result
}

func SignEthereumTx(privK *ecdsa.PrivateKey, toAddressS string, dataS string, nonceS string, gasLimitS string, chainIDS string) (string, *big.Int, *big.Int, *big.Int) {
	//Parse values
	value := big.NewInt(0)
	gasPrice := big.NewInt(0)
	toAddress := common.HexToAddress(toAddressS)
	data := common.FromHex(dataS)
	chainID, _ := hexutil.DecodeBig(chainIDS)

	cleaned := strings.Replace(gasLimitS, "0x", "", -1) // remove 0x suffix if found in the input string
	result, _ := strconv.ParseUint(cleaned, 16, 64)     // base 16 for hexadecimal
	gasLimit := result

	//Nonce parse -> Caution on this one
	cleaned = strings.Replace(nonceS, "0x", "", -1) // remove 0x suffix if found in the input string
	result, _ = strconv.ParseUint(cleaned, 16, 64)  // base 16 for hexadecimal
	nonce := result                                 // We cant use coz nonce's uto format has leading zero and this will panic -> hexutil.MustDecodeUint64(nonceS)

	//Create Tx object
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	//Inside this method, tx will be hashed and the sign with private key and will set v, r, s values on the Transaction object
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privK)
	if err != nil {
		panic(err.Error())
	}

	v, r, s := signedTx.RawSignatureValues()

	//This marshal the Transaction object to rlp encoding
	rlpSignedTx, err := signedTx.MarshalBinary()
	if err != nil {
		panic(err.Error())
	}
	signature := hexutil.Encode(rlpSignedTx)

	return signature, r, s, v
}
