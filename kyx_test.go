package kyx

import (
	"testing"
)

var testClint *API

func InitNewAPI() {
	token := ""
	testClint = NewAPI(token)
}

func init() {
	InitNewAPI()
}

func TestKyc(t *testing.T) {
	makeKyc := MakeKyc{
		FirstName: "Ahmet",
		LastName:  "Yılmaz",
	}

	res, err := testClint.CreateKyc(makeKyc) //this function runs a user creation process
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}
	if res.Error != "" {
		t.Fatalf("Error: %v", res.Error)
		return
	}
	t.Log("create kyc res", res)
	updateKyc := UpdateKyc{
		FirstName:                    "Ahmet",
		LastName:                     "Yılmaz",
		BirthYear:                    1990,
		BirthMonth:                   1,
		BirthDay:                     1,
		Nationality:                  "TR",
		NationalID:                   "12345678901",
		CountryCode:                  "TR",
		City:                         "Istanbul",
		Address:                      "Kadıköy",
		PhoneNumber:                  "905555555555",
		DocumentType:                 1,
		DocumentId:                   "1",
		NationalDocumentTypeId:       "1",
		DocumentGivenDate:            "2021-01-01",
		DocumentExpireDate:           "2021-01-01",
		ProvienceID:                  "1",
		ReferenceId:                  "1",
		SectorId:                     "1",
		ZipCode:                      "1",
		TaxResidence:                 "1",
		TaxpayerIdentificationNumber: "1",
		Occupation:                   "1",
		AnnualIncome:                 "1",
		SourceOfWealth:               "1",
		Correspondence:               true,
		Job:                          "1",
		AddressLine1:                 "1",
		AddressLine2:                 "1",
		Gender:                       "1",
		State:                        "1",
	}
	updateRes, err := testClint.UpdateKyc(res.KycId, updateKyc) //this function runs a user update process
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}
	if updateRes.Error != "" {
		t.Fatalf("Error: %v", updateRes.Error)
		return
	}
	t.Log("update kyc res", updateRes)
	r := InfoKyc{
		KycId:       res.KycId,
		ReferenceId: res.ReferenceId,
	}
	infoRes, err := testClint.InfoKyc(r) //this function runs a user info process
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}
	if infoRes.Error != "" {
		t.Fatalf("Error: %v", updateRes.Error)
		return
	}
	t.Log("kyc info", infoRes)

}
