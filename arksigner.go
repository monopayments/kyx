package kyx

type ContactlessRequestDTO struct {
	TransID                            string `json:"transId"`
	DG1Base64                          string `json:"dg1Base64"`
	DG2Base64                          string `json:"dg2Base64"`
	DG3Base64                          string `json:"dg3Base64"`
	DG4Base64                          string `json:"dg4Base64"`
	DG5Base64                          string `json:"dg5Base64"`
	DG6Base64                          string `json:"dg6Base64"`
	DG7Base64                          string `json:"dg7Base64"`
	DG8Base64                          string `json:"dg8Base64"`
	DG9Base64                          string `json:"dg9Base64"`
	DG10Base64                         string `json:"dg10Base64"`
	DG11Base64                         string `json:"dg11Base64"`
	DG12Base64                         string `json:"dg12Base64"`
	DG13Base64                         string `json:"dg13Base64"`
	DG14Base64                         string `json:"dg14Base64"`
	DG15Base64                         string `json:"dg15Base64"`
	DG16Base64                         string `json:"dg16Base64"`
	EFComBase64                        string `json:"efComBase64"`
	ChallangeBase64                    string `json:"challangeBase64"`
	ActiveAuthenticationResponseBase64 string `json:"activeAuthenticationResponseBase64"`
	MRZString                          string `json:"mrzstring"`
	EFSodBase64                        string `json:"efSodBase64"`
	IsLiveAuth                         bool   `json:"isLiveAuth"`
	CompanyCode                        string `json:"companyCode"`
	NameSurname                        string `json:"nameSurname"`
	IdentityNumber                     string `json:"identityNumber"`
}
type VerifyRequest struct {
	TransID        string `json:"transId"`
	Image          string `json:"image"`
	IsLiveAuth     bool   `json:"isLiveAuth"`
	CompanyCode    string `json:"companyCode"`
	NameSurname    string `json:"nameSurname"`
	IdentityNumber string `json:"identityNumber"`
}
type FaceMatchRequest struct {
	TransID          string   `json:"transId"`
	FirstFaceBase64  string   `json:"firstFaceBase64"`
	OtherFacesBase64 []string `json:"otherFacesBase64"`
	IsLiveAuth       bool     `json:"isLiveAuth"`
	ThresholdOption  int      `json:"thresholdOption"`
	CompanyCode      string   `json:"companyCode"`
}
type ArkSignerResponseMessage struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
