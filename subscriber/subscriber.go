package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const subConfrmType = "SubscriptionConfirmation"
const notificationType = "Notification"

//confirmSubscription Confirms Subscription by makeing get request to Subscribe URL
func confirmSubscription(subcribeURL string) {
	response, err := http.Get(subcribeURL)
	if err != nil {
		fmt.Printf("Unbale to confirm subscriptions")
	} else {
		fmt.Printf("Subscription Confirmed sucessfully. %d", response.StatusCode)
	}

}

//handler processes messages sent by SNS
func handler(w http.ResponseWriter, r *http.Request) {
	var f interface{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Unable to Parse Body")
	}
	fmt.Printf(string(body))
	err = json.Unmarshal(body, &f)
	if err != nil {
		fmt.Printf("Unable to Unmarshal request")
	}

	data := f.(map[string]interface{})
	fmt.Println(data["Type"].(string))

	if data["Type"].(string) == subConfrmType {
		subcribeURL := data["SubscribeURL"].(string)
		go confirmSubscription(subcribeURL)
	} else if data["Type"].(string) == notificationType {
		fmt.Println("Recieved this message : ", data["Message"].(string))
	}

	fmt.Fprintf(w, "Sucess")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
