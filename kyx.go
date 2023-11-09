package kyx

import (
	"bytes"
	"encoding/json"
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

func (t *API) get(path string, payload interface{}, response interface{}) error {
	url := t.EndPoint + path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	return t.do(req, response)
}
func (t *API) delete(path string, payload interface{}) error {
	url := t.EndPoint + path
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	return t.do(req, nil)
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

func (t *API) CreateKyc(payload KycCreate) (KycCreateResponse, error) {
	var response KycCreateResponse
	err := t.post("/kyc", payload, &response)

	return response, err
}
func (t *API) GetKyc(id string) (KycDto, error) {
	var response KycDto
	err := t.get("/kyc/"+id, nil, &response)
	return response, err
}
func (t *API) DeleteKyc(id string) (string, error) {
	var response KycDeleteResponse
	res := t.delete("/kyc/"+id, &response)
	if res != nil {
		return "", res
	}
	return "Kyc deleted", nil

}
