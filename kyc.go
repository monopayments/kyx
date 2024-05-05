package kyx

type MakeKyc struct {
	FirstName                    string `json:"first_name" example:"Ahmet"`
	LastName                     string `json:"last_name" example:"Yılmaz"`
	BirthYear                    uint64 `json:"birth_year" example:"1990"`
	BirthMonth                   uint64 `json:"birth_month" example:"01"`
	BirthDay                     uint64 `json:"birth_day" example:"01"`
	Nationality                  string `json:"nationality" example:"TR"`
	NationalID                   string `json:"national_id" example:"12345678901"`
	CountryCode                  string `json:"country_code" example:"TR"`
	City                         string `json:"city" example:"Istanbul"`
	Address                      string `json:"address" example:"Kadıköy"`
	PhoneNumber                  string `json:"phone_number" example:"905555555555"`
	DocumentType                 uint64 `json:"document_type" example:"1"`
	DocumentId                   string `json:"document_id" example:"1"`
	NationalDocumentTypeId       string `json:"national_document_type_id" example:"1"`
	DocumentGivenDate            string `json:"document_given_date" example:"2021-01-01"`
	DocumentExpireDate           string `json:"document_expire_date" example:"2021-01-01"`
	DocumentFrontImage           string `json:"document_front_image" example:"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMgAAADICAYAAACtWK6eA"`
	DocumentBackImage            string `json:"document_back_image" example:"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMgAAADICAYAAACtWK6eA"`
	ProvienceID                  string `json:"provience_id" example:"1"`
	ReferenceId                  string `json:"reference_id" example:"1"`
	SectorId                     string `json:"sector_id" example:"1"`
	ZipCode                      string `json:"zip_code" example:"1"`
	TaxResidence                 string `json:"tax_residence" example:"1"`
	TaxpayerIdentificationNumber string `json:"taxpayer_identification_number" example:"1"`
	Occupation                   string `json:"occupation" example:"1"`
	AnnualIncome                 string `json:"annual_income" example:"1"`
	SourceOfWealth               string `json:"source_of_wealth" example:"1"`
	Correspondence               bool   `json:"correspondence" example:"1"`
	Job                          string `json:"job" example:"1"`
	AddressLine1                 string `json:"address_line_1" example:"1"`
	AddressLine2                 string `json:"address_line_2" example:"1"`
	Gender                       string `json:"gender"`
	State                        string `json:"state"`
	ResidenceProofDocument       string `json:"residence_proof_document" example:"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMgAAADICAYAAACtWK6eA"`
}
type UpdateKyc struct {
	FirstName                    string `json:"first_name" example:"Ahmet"`
	LastName                     string `json:"last_name" example:"Yılmaz"`
	BirthYear                    uint64 `json:"birth_year" example:"1990"`
	BirthMonth                   uint64 `json:"birth_month" example:"01"`
	BirthDay                     uint64 `json:"birth_day" example:"01"`
	Nationality                  string `json:"nationality" example:"TR"`
	NationalID                   string `json:"national_id" example:"12345678901"`
	CountryCode                  string `json:"country_code" example:"TR"`
	City                         string `json:"city" example:"Istanbul"`
	Address                      string `json:"address" example:"Kadıköy"`
	PhoneNumber                  string `json:"phone_number" example:"905555555555"`
	DocumentType                 uint64 `json:"document_type" example:"1"`
	DocumentId                   string `json:"document_id" example:"1"`
	NationalDocumentTypeId       string `json:"national_document_type_id" example:"1"`
	DocumentGivenDate            string `json:"document_given_date" example:"2021-01-01"`
	DocumentExpireDate           string `json:"document_expire_date" example:"2021-01-01"`
	DocumentFrontImage           string `json:"document_front_image" example:"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMgAAADICAYAAACtWK6eA"`
	DocumentBackImage            string `json:"document_back_image" example:"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMgAAADICAYAAACtWK6eA"`
	ProvienceID                  string `json:"provience_id" example:"1"`
	ReferenceId                  string `json:"reference_id" example:"1"`
	SectorId                     string `json:"sector_id" example:"1"`
	ZipCode                      string `json:"zip_code" example:"1"`
	TaxResidence                 string `json:"tax_residence" example:"1"`
	TaxpayerIdentificationNumber string `json:"taxpayer_identification_number" example:"1"`
	Occupation                   string `json:"occupation" example:"1"`
	AnnualIncome                 string `json:"annual_income" example:"1"`
	SourceOfWealth               string `json:"source_of_wealth" example:"1"`
	Correspondence               bool   `json:"correspondence" example:"1"`
	Job                          string `json:"job" example:"1"`
	AddressLine1                 string `json:"address_line_1" example:"1"`
	AddressLine2                 string `json:"address_line_2" example:"1"`
	Gender                       string `json:"gender"`
	State                        string `json:"state"`
	ResidenceProofDocument       string `json:"residence_proof_document" example:"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMgAAADICAYAAACtWK6eA"`
}
type MakeKycResponse struct {
	Message     string `json:"message" example:"Kyc created successfully"`
	ReferenceId string `json:"reference_id" example:"f0a0a1e9-69bd-4bef-b8c6-4e8c0d3a1212"`
	KycId       string `json:"kyc_id" example:"f0a0a1e9-69bd-4bef-b8c6-4e8c0d3a1212"`
	KycState    string `json:"kyc_state" example:"1"`
	KycLevel    uint64 `json:"kyc_level" example:"10"`
	Error       string `json:"error" example:""`
}
type InfoKyc struct {
	KycProvider string `json:"kyc_provider" example:"1"`
	ReferenceId string `json:"reference_id" example:"f0a0a1e9-69bd-4bef-b8c6-4e8c0d3a1212"`
	KycId       string `json:"kyc_id" example:"f0a0a1e9-69bd-4bef-b8c6-4e8c0d3a1212"`
}
type InfoKycResponse struct {
	Message                string `json:"message" example:"Kyc created successfully"`
	ReferenceId            string `json:"reference_id" example:"f0a0a1e9-69bd-4bef-b8c6-4e8c0d3a1212"`
	KycId                  string `json:"kyc_id" example:"f0a0a1e9-69bd-4bef-b8c6-4e8c0d3a1212"`
	KycState               string `json:"kyc_state" example:"1"`
	KycLevel               uint64 `json:"kyc_level" example:"10"`
	Error                  string `json:"error" example:""`
	RejectReason           uint64 `json:"reject_reason"`
	RejectNote             string `json:"reject_note"`
	ResidenceProofDocument string `json:"residence_proof_document" example:"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMgAAADICAYAAACtWK6eA"`
}
