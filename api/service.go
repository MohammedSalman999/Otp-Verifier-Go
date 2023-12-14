package api

import (
	"errors"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
	
)

var client *twilio.RestClient  = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username : envACCOUNTSID(),
	Password : envAUTHTOKEN(),
})

func (app *Config) twilioSendOTP(phoneNumber string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")
	resp , err := client.VerifyV2.CreateVerification(envSERVICEID(),params)
	if err != nil {
		return "" , err
	}

	return *resp.Sid , nil
}

func (app *Config) twilioVerifyOTP(phoneNumber string, code string) ( bool, error) {
	checkParams := &twilioApi.CreateVerificationCheckParams{}
	checkParams.SetTo(phoneNumber)
	checkParams.SetCode(code)

	resp , err := client.VerifyV2.CreateVerificationCheck(envSERVICEID(),checkParams)
	if err != nil {
		return false, err
	}

	if *resp.Status != "approved" {
		return false, errors.New("not a Valid code")
	}

	return true, nil

}
