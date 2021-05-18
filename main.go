package main

import (
	s "github.com/gataca-io/ebsi/insert-did-document/services"
	"log"
)

const (
	HashAlgId           = 1
	DIDRegistryEndpoint = "/did-registry/v2/jsonrpc"
	TIREndpoint         = "/trusted-issuers-registry/v2/jsonrpc"

	InsertDIDDocument = "insertDidDocument"
	UpdateDIDDocument = "updateDidDocument"
	InsertIssuer      = "insertIssuer"
	UpdateIssuer      = "updateIssuer"

	UTPayloadFile    = "ut-payload.json"
	UTResponseFile   = "ut-response.json"
	LastResponseFile = "last-response.json"
)

func main() {
	ebsiHost, didIdentifier, didDocumentVersion, didTimestampData, didVersionMetadata, privateKeyHex, method, trustedIssuerVersion := s.LoadVars()

	privKey, _, fromAddress := s.LoadKeys(privateKeyHex)

	var endpoint string
	var utp []byte
	switch method {
	case InsertDIDDocument:
		log.Println("Insert DID Document")
		utp = s.GenerateInsertDidDocumentPayload(fromAddress, didIdentifier, didDocumentVersion, didTimestampData, didVersionMetadata, HashAlgId)
		if utp == nil {
			panic("Generation from insert Did Document payload failed.")
		}

		endpoint = DIDRegistryEndpoint
	case UpdateDIDDocument:
		log.Println("Update DID Document")
		utp = s.GenerateUpdateDidDocumentPayload(fromAddress, didIdentifier, didDocumentVersion, didTimestampData, didVersionMetadata, HashAlgId)
		if utp == nil {
			panic("Generation from update Did Document payload failed.")
		}

		endpoint = DIDRegistryEndpoint
	case InsertIssuer:
		log.Println("Insert Issuer")
		utp = s.GenerateInsertIssuerPayload(fromAddress, didIdentifier, trustedIssuerVersion, HashAlgId)
		if utp == nil {
			panic("Generation from insert issuer payload failed.")
		}

		endpoint = TIREndpoint
	case UpdateIssuer:
		log.Println("Update Issuer")
		utp = s.GenerateUpdateIssuerPayload(fromAddress, didIdentifier, didDocumentVersion, didTimestampData, didVersionMetadata, HashAlgId)
		if utp == nil {
			panic("Generation from update issuer payload failed.")
		}

		endpoint = TIREndpoint
	default:
		log.Panicf("Operation method is not OK. Method: %s", method)
	}

	s.WriteOutput(UTPayloadFile, string(utp))

	ut := s.CallEbsiJsonRPCFunction(utp, ebsiHost, endpoint)
	if ut == "" {
		panic("Call 1 to EBSI failed.")
	}

	s.WriteOutput(UTResponseFile, ut)

	parsedUT := s.ParseToUnsignedTransaction(ut)
	if parsedUT == nil {
		panic("Call to parse Unsigned transaction failed.")
	}

	st := s.GenerateSignedTransaction(parsedUT, privKey)
	if st == nil {
		panic("Generate Signed Transaction failed.")
	}

	res := s.CallEbsiJsonRPCFunction(st, ebsiHost, endpoint)
	if res == "" {
		panic("Call 2 to EBSI failed")
	}

	s.WriteOutput(LastResponseFile, res)
}
