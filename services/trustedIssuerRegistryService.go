package services

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gataca-io/ebsi/insert-did-document/models"
	"os"
)

const (
	insertIssuerMethod = "insertIssuer"
	updateIssuerMethod = "updateIssuer"
)

func GenerateInsertIssuerPayload(fromAddress common.Address, didIdentifier, trustedIssuerVersion string, hashAlgId int) []byte {

	//identifier := hexutil.Encode([]byte(didIdentifier))

	h := sha256.New()
	h.Write([]byte(trustedIssuerVersion))
	hash := hexutil.Encode(h.Sum(nil))

	att := models.TIRAttribute{
		Body: base64.URLEncoding.EncodeToString([]byte(trustedIssuerVersion)), //hexutil.Encode([]byte(trustedIssuerVersion)),
		Hash: hash,
	}

	tp := models.TIRParams{
		From:      fromAddress.Hex(),
		Did:       didIdentifier,
		Attribute: &att,
	}

	idd := models.InsertIssuer{
		JsonRPC: JsonRPCVersion,
		Method:  insertIssuerMethod,
		Params:  []models.TIRParams{tp},
		Id:      0,
	}

	body, err := json.Marshal(idd)
	if err != nil {
		os.Exit(-1)
	}

	return body
}

func GenerateUpdateIssuerPayload(fromAddress common.Address, didIdentifier, didDocumentVersion, timestampData, didVersionMetadata string, hashAlgId int) []byte {

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
		Method:  updateIssuerMethod,
		Params:  []models.DIDRParams{ep},
		Id:      0,
	}

	body, err := json.Marshal(idd)
	if err != nil {
		os.Exit(-1)
	}

	return body
}
