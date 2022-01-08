package main

import (
	"fmt"

	"github.com/pubnub/go/messaging"
)

func main() {
	subKey := "sub-c-52a9ab50-291b-11e5-baaa-0619f8945a4f"
	pubnub := messaging.NewPubnub("", subKey, "", "", false, "", nil)

	sucCha := make(chan []byte)
	errCha := make(chan []byte)

	go pubnub.Subscribe("lightning_ticker_BTC_USD", "", sucCha, false, errCha)
	for {
		select {
		case res := <-sucCha:
			fmt.Println(string(res))
		case err := <-errCha:
			fmt.Println(string(err))
		case <-messaging.SubscribeTimeout():
			fmt.Println("timeout!")
		}
	}
}
