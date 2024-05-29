package kyx

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

// TapsilatAPI is the main struct for the Tapsilat API
type API struct {
	EndPoint string `json:"end_point"`
	Token    string `json:"token"`
}

// NewAPI creates a new TapsilatAPI struct
func NewAPI(token string) *API {
	return &API{
		EndPoint: "https://kyc.tapsilat.dev/api/public",
		Token:    token,
	}
}

// NewAPIWithEndpoint creates a new TapsilatAPI struct with a custom endpoint
func NewCustomAPI(endpoint, token string) *API {
	return &API{
		EndPoint: endpoint,
		Token:    token,
	}
}

func (t *API) post(path string, payload interface{}, response interface{}) error {
	url := t.EndPoint + path
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	return t.do(req, response)
}

func (t *API) put(path string, payload interface{}, response interface{}) error {
	url := t.EndPoint + path
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	return t.do(req, response)
}

func (t *API) get(path string, payload interface{}, response interface{}) error {
	url := t.EndPoint + path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	return t.do(req, response)
}
func (t *API) delete(path string, payload interface{}, response interface{}) error {
	url := t.EndPoint + path
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	return t.do(req, response)
}

func (t *API) do(req *http.Request, response interface{}) error {
	req.Header.Set("x-mono-org-auth", t.Token)
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	decode := json.NewDecoder(bytes.NewReader(body))
	decode.DisallowUnknownFields()
	decode.UseNumber()
	if err := decode.Decode(response); err != nil {
		return err
	}
	return nil
}

func (t *API) CreateKyc(payload MakeKyc) (MakeKycResponse, error) {
	var response MakeKycResponse
	err := t.post("/kyc/create", payload, &response)
	return response, err
}
func (t *API) InfoKyc(payload InfoKyc) (InfoKycResponse, error) {
	var response InfoKycResponse
	err := t.post("/kyc/info", payload, &response)
	return response, err
}
func (t *API) UpdateKyc(kycId string, payload UpdateKyc) (InfoKycResponse, error) {
	var response InfoKycResponse
	err := t.put("/kyc/"+kycId+"/update", payload, &response)
	return response, err
}

// ark-signer
func (t *API) FaceMatch(payload FaceMatchRequest) (ArkSignerResponseMessage, error) {
	var response ArkSignerResponseMessage
	err := t.post("/arksigner/face", payload, &response)
	return response, err
}
func (t *API) VerifyTCKKFront(payload VerifyRequest) (ArkSignerResponseMessage, error) {
	var response ArkSignerResponseMessage
	err := t.post("/arksigner/tckk/front", payload, &response)
	return response, err
}
func (t *API) VerifyTCKKBack(payload VerifyRequest) (ArkSignerResponseMessage, error) {
	var response ArkSignerResponseMessage
	err := t.post("/arksigner/tckk/back", payload, &response)
	return response, err
}
func (t *API) Contactless(payload ContactlessRequestDTO) (ArkSignerResponseMessage, error) {
	var response ArkSignerResponseMessage
	err := t.post("/arksigner/contactless", payload, &response)
	return response, err
}

// aml
func (t *API) AmlSearchByName(payload AmlCheckRequest) (AmlCheckResponse, error) {
	var response AmlCheckResponse
	apiPath := "/aml/user-aml"
	if payload.Name == "" {
		return response, errors.New("name is required")
	}
	apiPath = "?name=" + payload.Name

	if payload.ReferenceID != "" {
		apiPath = "&reference_id=" + payload.ReferenceID
	}

	err := t.get(apiPath, nil, &response)
	return response, err
}
