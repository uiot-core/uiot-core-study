package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	//"encoding/base64"
	//"encoding/json"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	s,_ := ioutil.ReadAll(r.Body)

	/*
	jsonb := make(map[string]string)
	err := json.Unmarshal(s,&jsonb)
	if err != nil {
		fmt.Println(err)
	}
	if _,ok :=jsonb["content"]; ok {
		d,err := base64.StdEncoding.DecodeString(jsonb["content"])
		if err != nil {
			fmt.Println(err)
		}
		hexdump := hex.Dump(d)
		fmt.Print(hexdump)
	}
	*/
	hexdump := hex.Dump(s)
	fmt.Print(hexdump)
}

func main() {
	http.HandleFunc("/receive", IndexHandler)
	http.ListenAndServe(":8900", nil)

}
