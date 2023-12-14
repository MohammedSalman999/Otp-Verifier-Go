package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/MohammedSalman999/go-otp-varification/data"
	"github.com/gin-gonic/gin"
)
const appTimeout = time.Second * 10
func (app *Config) sendSMS() gin.HandlerFunc{
	return func (c *gin.Context)  {

		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.OTPData
		defer cancel()

		app.validateBody(c, &payload)

		newData := data.OTPData {
			PhoneNumber: payload.PhoneNumber,
		}

		_,err := app.twilioSendOTP(newData.PhoneNumber)
		if err!= nil {
			app.errorJSON(c, err)
			return
		}

		app.writeJSON(c , http.StatusAccepted , "OTP  Send Successfully")
		
	}
}


func (app *Config) verifySMS() gin.HandlerFunc{
	return func (c *gin.Context)  {

		_, cancel := context.WithTimeout(context.Background() , appTimeout)
		var payload data.VarifyData
		defer cancel()

		app.validateBody(c, &payload)

		newData := data.VarifyData {
			User: payload.User,
			Code: payload.Code,
		}

		_,err := app.twilioVerifyOTP(newData.User.PhoneNumber , newData.Code)
		fmt.Println("err:",err)
		if err!= nil {
			app.errorJSON(c, err)
			return
		}

		app.writeJSON(c , http.StatusAccepted , "OTP  Send Successfully")
		
	}
}