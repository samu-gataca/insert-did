package services

import (
	"crypto/sha256"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gataca-io/ebsi/insert-did-document/models"
	"os"
)

const (
	insertDidDocumentMethod = "insertDidDocument"
	updateDidDocumentMethod = "updateDidDocument"
)

func GenerateInsertDidDocumentPayload(fromAddress common.Address, didIdentifier, didDocumentVersion, timestampData, didVersionMetadata string, hashAlgId int) []byte {

	identifier := hexutil.Encode([]byte(didIdentifier))

	h := sha256.New()
	h.Write([]byte(didDocumentVersion))
	hash := hexutil.Encode(h.Sum(nil))

	DidIV := hexutil.Encode([]byte(didDocumentVersion))
	tsd := hexutil.Encode([]byte(timestampData))
	dvm := hexutil.Encode([]byte(didVersionMetadata))

	ep := models.DIDRParams{
		From:               fromAddress.Hex(),
		Identifier:         identifier,
		HashAlgorithmId:    hashAlgId,
		HashValue:          hash,
		DidVersionInfo:     DidIV,
		TimestampData:      tsd,
		DidVersionMetadata: dvm,
	}

	idd := models.InsertDIDDocument{
		JsonRPC: JsonRPCVersion,
		Method:  insertDidDocumentMethod,
		Params:  []models.DIDRParams{ep},
		Id:      0,
	}

	body, err := json.Marshal(idd)
	if err != nil {
		os.Exit(-1)
	}

	return body
}

func GenerateUpdateDidDocumentPayload(fromAddress common.Address, didIdentifier, didDocumentVersion, timestampData, didVersionMetadata string, hashAlgId int) []byte {

	identifier := hexutil.Encode([]byte(didIdentifier))

	h := sha256.New()
	h.Write([]byte(didDocumentVersion))
	hash := hexutil.Encode(h.Sum(nil))

	DidIV := hexutil.Encode([]byte(didDocumentVersion))
	tsd := hexutil.Encode([]byte(timestampData))
	dvm := hexutil.Encode([]byte(didVersionMetadata))

	ep := models.DIDRParams{
		From:               fromAddress.Hex(),
		Identifier:         identifier,
		HashAlgorithmId:    hashAlgId,
		HashValue:          hash,
		DidVersionInfo:     DidIV,
		TimestampData:      tsd,
		DidVersionMetadata: dvm,
	}

	idd := models.InsertDIDDocument{
		JsonRPC: JsonRPCVersion,
		Method:  updateDidDocumentMethod,
		Params:  []models.DIDRParams{ep},
		Id:      0,
	}

	body, err := json.Marshal(idd)
	if err != nil {
		os.Exit(-1)
	}

	return body
}
