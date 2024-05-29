package kyx

type AmlCheckRequest struct {
	Name        string `json:"name"`
	ReferenceID string `json:"reference_id"`
}
type AmlCheckResponse struct {
	ReferenceNumber    string `json:"reference_number"`
	OutReferenceNumber string `json:"out_reference_number"`
	RiskLevel          string `json:"risk_level"`
	IsWhiteList        bool   `json:"is_white_list"`
	WhiteListMessage   string `json:"white_list_message"`
	RiskLevelID        int64  `json:"risk_level_id"`
	IsSafeList         bool   `json:"is_safe_list"`
	Error              string `json:"error" example:""`
}
