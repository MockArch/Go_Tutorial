package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://eshipz.herokuapp.com/rates"

	payload := strings.NewReader("{\n  \"is_document\": false,\n  \"api_token\" :\"5aa2a44caea23407384f51c5\",\n  \"shipment\": {\n\t\"is_test_booking\" : true,\n\t\"is_cod\": true,\n\t\"collect_on_delivery\":{\n\t\t\"amount\":2000,\n\t\t\"currency\":\"INR\"\n\t},\n\t\"purpose\":\"commercial\",\n    \"ship_from\": {\n      \"contact_name\": \"Contact name\",\n      \"company_name\": \"Testing Company\",\n      \"street1\": \"Testing Street\",\n      \"city\": \"Banglore\",\n      \"state\": \"KA\",\n      \"postal_code\": \"560091\",\n      \"country\": \"IN\",\n      \"type\": \"residential\",\n      \"phone\": \"+919428293450\",\n      \"street2\": \"akjdh\",\n      \"tax_id\": \"adaj\",\n      \"street3\": \"asjhdga\",\n      \"email\": \"aramext@test.com\",\n      \"fax\": \"amjsdaj\"\n    },\n    \"ship_to\": {\n      \"contact_name\": \"Dr. Moises Corwin\",\n      \"phone\": \"1-140-225-6410\",\n      \"email\": \"Giovanna42@yahoo.com\",\n      \"street1\": \"28292 Daugherty Orchard\",\n      \"street2\" : null,\n\t  \"street3\" :null,\n      \"postal_code\": \"133301\",\n      \"city\": \"Dehli\",\n      \"state\": \"DL\",\n      \"country\": \"IN\",\n      \"type\": \"business\"\n    },\n    \"parcels\": [{\n      \"description\": \"Food adhas\",\n      \"box_type\": \"custom\",\n      \"weight\": {\n        \"value\": 4,\n        \"unit\": \"kg\"\n      },\n      \"dimension\": {\n        \"width\": 1,\n        \"height\": 1,\n        \"length\": 40,\n        \"unit\": \"cm\"\n      },\n      \"items\": [{ \n        \"description\": \"Food Bar\",\n        \"origin_country\": \"IN\",\n        \"quantity\": 4,\n        \"price\": {\n          \"amount\": 400,\n          \"currency\": \"INR\"\n        },\n        \"weight\": {\n          \"value\":3,\n          \"unit\": \"kg\"\n        }\n      }]\n    }]\n  }\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "7c8c0e94-fc48-b4c5-0cbc-2f6cc9b1cd70")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
