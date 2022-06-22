package services

import (
	"fmt"
	"net/http"

	"encoding/json"
)

type TokenizeRequest struct {
	Pan            string `json:"pan"`
	PanExp         string `json:"panExp"`
	CardHolderName string `json:"cardHolderName"`
	Cvv            string `json:"cvv"`
}

type TokenizeResponse struct {
	TokenReferenceId     string `json:"tokenReferenceId"`
	TokenState           string `json:"tokenState"`
	TokenizationDecision string `json:"tokenizationDecision"`
}

func TokenizeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Tokenize Handler")

	w.Header().Set("Content-Type", "application/json")
	var tokenizeRequest TokenizeRequest
	_ = json.NewDecoder(r.Body).Decode(&tokenizeRequest)

	fmt.Println(tokenizeRequest.Pan, tokenizeRequest.PanExp)

	//Dummy response
	tokenizeReponse := TokenizeResponse{"1", "ACTIVE", "APPROVED"}
	json.NewEncoder(w).Encode(tokenizeReponse)
}

func GetPaymentBundleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetPaymentBundle Handler")
}
