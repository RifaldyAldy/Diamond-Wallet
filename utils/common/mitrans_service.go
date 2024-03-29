package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/midtrans/midtrans-go"
)

type ResponseMidtrans struct {
	Token      string `json:"token"`
	UrlPayment string `json:"redirect_url"`
}
type RequestMidtrans struct {
	TransactionDetails midtrans.TransactionDetails `json:"transaction_details"`
}

func GenerateMidtrans(payload midtrans.TransactionDetails) (ResponseMidtrans, error) {
	var response ResponseMidtrans
	url := "https://app.sandbox.midtrans.com/snap/v1/transactions"
	request := RequestMidtrans{TransactionDetails: payload}
	payloads, _ := json.Marshal(request)

	req, _ := http.NewRequest("POST", url, bytes.NewReader(payloads))

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Basic U0ItTWlkLXNlcnZlci1mM2lXS3ZBOE54MXpVZ3YtbG9QQ05oV0o6")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ResponseMidtrans{}, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	err = json.Unmarshal(body, &response)
	if err != nil {
		return ResponseMidtrans{}, err
	}
	fmt.Println(response)
	return response, nil

}
