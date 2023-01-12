package usecases

import (
	"errors"
	"fmt"
	"learn-oauth2/internal/core/ports"
	"log"
)

type oauth2UseCase struct {
	RequestService ports.IRequest
	LoggerService  ports.ILogger
}

func NewOauth2UseCase(requestService ports.IRequest) ports.Oauth2UseCase {
	return &oauth2UseCase{RequestService: requestService}
}

func (o oauth2UseCase) AuthzFacebook(accessTokenClient string, accessToken string) (string, error) {
	var facebookId string
	var url string = fmt.Sprintf("https://graph.facebook.com/debug_token?input_token=%s&access_token=%s", accessTokenClient, accessToken)
	var method string = "GET"
	var headers map[string]string = map[string]string{
		"Content-type": "application/json",
	}

	res, err := o.RequestService.Json(method, url, headers, nil)
	if err != nil {
		log.Println("err :", err)
		return "", errors.New("error http request")
	}

	if data, exists := res["data"].(map[string]interface{}); exists == false {
		return "", errors.New("error http request [2]")
	} else {
		_, exists := data["error"].(map[string]interface{})
		if exists == true {
			return "", errors.New("error http request [2]")
		} else if data["is_valid"] == false {
			return "", errors.New("error http request [3]")
		} else {
			facebookId = fmt.Sprint(data["user_id"])
		}
	}
	return facebookId, nil
}
