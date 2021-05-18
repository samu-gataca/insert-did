package models

type TrustedIssuer struct {
	Context          []string           `json:"@context" example:"https://essif.europa.eu/schemas/tir/2020/v1" description:"" `
	PreferredName    string             `json:"preferredName" example:"University Rovira i Virgili" description:"" `
	Domain           string             `json:"domain" example:"Education" description:"" `
	DID              []string           `json:"did" example:"did:ebsi:2239f1ef4d7727cc47495a5a1a6ff4e1c6a93175f2bd61e7" description:""`
	EidasCertificate []EidasCertificate `json:"eidasCertificate,omitempty" description:""`
	ServiceEndpoints []ServiceEndpoint  `json:"serviceEndpoints,omitempty" description:""`
	OrganizationInfo *OrganizationInfo  `json:"organizationInfo" description:""`
	Proof            *Proof             `json:"proof" description:""`
}

type EidasCertificate struct {
	EidasCertificateIssuerNumber string `json:"eidasCertificateIssuerNumber" example:"123456" description:"" `
	EidasCertificateSerialNumber string `json:"eidasCertificateSerialNumber" example:"123456" description:""`
	EidasCertificatePem          string `json:"eidasCertificatePem" example:"<PEM-ENC-CERT>" description:""`
}

type ServiceEndpoint struct {
	Id              string `json:"id" example:"did:ebsi:223f3196d395b829e4cb1efc27457275f729a9ac#openid" description:"" `
	ServiceType     string `json:"type" example:"OpenIdConnectVersion1.0Service" description:""`
	ServiceEndpoint string `json:"serviceEndpoint" example:"https://agent.example.com/8377464" description:""`
}

type OrganizationInfo struct {
	LegalPersonalIdentifier string `json:"legalPersonalIdentifier" example:"123456789" description:"" `
	LegalName               string `json:"legalName" example:"Example Legal Name" description:"" `
	LegalAddress            string `json:"legalAddress,omitempty" example:"Example Street 42, Vienna, Austria" description:"" `
	VATRegistration         string `json:"VATRegistration,omitempty" example:"ATU12345678" description:"" `
	TaxReference            string `json:"taxReference,omitempty" example:"123456789" description:"" `
	LEI                     string `json:"LEI,omitempty" example:"12341212EXAMPLE34512" description:"" `
	EORI                    string `json:"EORI,omitempty" example:"AT123456789101" description:"" `
	SEED                    string `json:"SEED,omitempty" example:"AT12345678910" description:"" `
	SIC                     string `json:"SIC,omitempty" example:"1234" description:"" `
	DomainName              string `json:"domainName,omitempty" example:"https://example.organization.com" description:"" `
}

type Proof struct {
	ProofType          string          `json:"type" example:"EidasSeal2021" description:"" `
	Created            *TimeWithFormat `json:"created" example:"2019-06-22T14:11:44Z" description:"" `
	ProofPurpose       string          `json:"proofPurpose" example:"assertionMethod" description:"" `
	VerificationMethod string          `json:"verificationMethod" example:"did:ebsi:2239f1ef4d7727cc47495a5a1a6ff4e1c6a93175f2bd61e7#1088321447" description:"" `
	Jws                string          `json:"jws" example:"BD21J4fdlnBvBA+y6D...fnC8Y=" description:"" `
}
