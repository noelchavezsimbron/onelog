package main

import (
	"encoding/json"
	"github.com/francoispqt/onelog"
	"github.com/francoispqt/onelog/powerlog"
	"log"
	"os"
	"time"
)

type Perro struct {
	Name string `json:"name"`
	Raza string `json:"raza"`
}

var jsonResponse = ` {"duplicated":false,"object":"charge","id":"chr_test_JXvnOjmWQFgfTZH4","creation_date":1605754117000,"amount":4878,"amount_refunded":0,"current_amount":4878,"installments":1,"installments_amount":4878,"currency_code":"PEN","email":"leo.123@gmail.com.pe","description":null,"source":{"object":"token","id":"tkn_test_WxJp4ynSum89TzhQ","type":"card","creation_date":1605754111000,"email":"noel.chavez@interseguro.com.pe","card_number":"411111******1111","last_four":"1111","active":true,"iin":{"object":"iin","bin":"411111","card_brand":"Visa","card_type":"credito","card_category":"Clásica","issuer":{"name":"BBVA","country":"PERU","country_code":"PE","website":null,"phone_number":null},"installments_allowed":[2,4,6,8,10,12,3,5,7,9,24,48]},"client":{"ip":"190.232.61.83","ip_country":"Peru","ip_country_code":"PE","browser":"UNKNOWN","device_fingerprint":null,"device_type":"Escritorio"},"metadata":{}},"outcome":{"type":"venta_exitosa","code":"AUT0000","merchant_message":"La operación de venta ha sido autorizada exitosamente","user_message":"Su compra ha sido exitosa."},"fraud_score":6.0,"dispute":false,"capture":true,"reference_code":"DhdsLreqfq","authorization_code":"74qnom","metadata":{"code":"0780000001"},"total_fee":230,"fee_details":{"fixed_fee":{},"variable_fee":{"currency_code":"PEN","commision":0.0399,"total":195}},"total_fee_taxes":35,"transfer_amount":4648,"paid":false,"statement_descriptor":"CULQI*","transfer_id":null,"capture_date":1605754117000,"items":[{"key":"Item1","val":"Val1"},{"key":"Item2","val":"Val2"}]}`

func main() {

	var logger = powerlog.New(os.Stdout, onelog.ALL).Hook(
		func(e powerlog.Entry) {
			e.Int64("time", time.Now().Unix())
		},
	)

	perro :=Perro{
		Name: "Roy",
		Raza: "Bull Dog",
	}

	responseBody := new(interface{})

	if err := json.Unmarshal([]byte(jsonResponse), responseBody); err != nil {
		log.Fatal(err)
		return
	}

	respTrance := struct {
		TransactionID string            `json:"transaction_id"`
		Action        string            `json:"action"`
		Url           string            `json:"url"`
		Method        string            `json:"method"`
		Headers       map[string]string `json:"headers"`
		Status        string            `json:"status"`
		SatusCode     int               `json:"satus_code"`
		Body          interface{}       `json:"body"`
	}{
		TransactionID: "010010010010101010101",
		Action:        "Received response from \"httpL//fakeapi.com",
		Url:           "httpL//fakeapi.com",
		Method:        "POST",
		Headers:       nil,
		Status:        "200",
		SatusCode:     200,
		Body:          *responseBody,
	}

	logger.InfoWith().
		String("hellow", "wordld").
		ObjectFunc(
			"persona", func(entry powerlog.Entry) {
				entry.String("nombre", "pepito")
				entry.Int64("edad", 20)
			},
		).
		Object("mascota", perro).
		Object("response", respTrance).
		Write()

}
