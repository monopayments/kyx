package kyx

type KycCreate struct {
	FirstName          string `json:"first_name" example:"Ahmet"`
	ReferenceID        string `json:"reference_id" example:"f0a0a1e9-69bd-4bef-b8c6-4e8c0d3a1212"`
	LastName           string `json:"last_name" example:"Yılmaz"`
	BirthYear          uint64 `json:"birth_year" example:"1990"`
	BirthMonth         uint64 `json:"birth_month" example:"01"`
	BirthDay           uint64 `json:"birth_day" example:"01"`
	NationalID         string `json:"national_id" example:"12345678901"`
	CountryCode        string `json:"country_code" example:"TR"`
	City               string `json:"city" example:"Istanbul"`
	ProvienceID        string `json:"provience_id" example:"34"`
	Address            string `json:"address" example:"Kadıköy"`
	PhoneNumber        string `json:"phone_number" example:"905555555555"`
	DocumentType       uint64 `json:"document_type" example:"1"`
	DocumentFrontImage string `json:"document_front_image" example:"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMgAAADICAYAAACtWK6eA"`
	DocumentBackImage  string `json:"document_back_image" example:"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMgAAADICAYAAACtWK6eA"`
} // @name KycCreateDto
type KycCreateResponse struct {
	Message string `json:"message" example:"Kyc created successfully"`
	Error   string `json:"error" example:""`
} // @name KycCreateResponseDto
type KycDto struct {
	ID                 string `json:"id" example:"f0a0a1e9-69bd-4bef-b8c6-4e8c0d3a1212"`
	HashID             string `json:"hash_id" example:"kKE4ZvqVgPl9d8eA8RDAxyGJz2onEC"`
	FirstName          string `json:"first_name" example:"Ahmet"`
	LastName           string `json:"last_name" example:"Yılmaz"`
	ReferenceID        string `json:"reference_id" example:"f0a0a1e9-69bd-4bef-b8c6-4e8c0d3a1212"`
	BirthYear          uint64 `json:"birth_year" example:"1990"`
	BirthMonth         uint64 `json:"birth_month" example:"01"`
	BirthDay           uint64 `json:"birth_day" example:"01"`
	NationalID         string `json:"national_id" example:"12345678901"`
	CountryCode        string `json:"country_code" example:"TR"`
	City               string `json:"city" example:"Istanbul"`
	ProvienceID        string `json:"provience_id" example:"34"`
	Address            string `json:"address" example:"Kadıköy"`
	PhoneNumber        string `json:"phone_number" example:"905555555555"`
	DocumentType       uint64 `json:"document_type" example:"1"`
	Level              uint64 `json:"level" example:"10"`
	State              uint64 `json:"state" example:"1"`
	DocumentFrontImage string `json:"document_front_image" example:"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMgAAADICAYAAACtWK6eA"`
	DocumentBackImage  string `json:"document_back_image" example:"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMgAAADICAYAAACtWK6eA"`

	AccountID      string `json:"account_id" example:"f0a0a1e9-69bd-4bef-b8c6-4e8c0d3a1212"`
	OrganizationID string `json:"organization_id" example:"f0a0a1e9-69bd-4bef-b8c6-4e8c0d3a1212"`
} // @name KycDto
