package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gufranmirza/offline-otp/generator"
)

var config = generator.Config{
	Secret:              "vjkhvkdfsv8d854vd65f4v65sdf4v65dsf4vv5df64v65d",
	ExpiryTimeInSeconds: 10,
	OTPLength:           4,
}

type req struct {
	OTP int `json:"otp"`
}

func validate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	buff := req{}
	err = json.Unmarshal(body, &buff)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	otpGen := generator.New(&config)
	result := otpGen.Validate(buff.OTP)
	if result {
		fmt.Fprintf(w, "VALID OTP\n")
		return
	}

	fmt.Fprintf(w, "INVALID OTP\n")
}

func main() {
	otpGen := generator.New(&config)
	otp := otpGen.Generate()
	fmt.Println("GENERATED ONE TIME OTP: ", otp)

	http.HandleFunc("/validate-otp", validate)
	http.ListenAndServe(":8090", nil)
}
