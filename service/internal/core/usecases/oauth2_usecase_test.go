//go:build integration
// +build integration

package usecases

import (
	"github.com/stretchr/testify/mock"
	"learn-oauth2/internal/core/ports"
	"learn-oauth2/internal/core/ports/mocks"
	"testing"
)

func Test_oauth2UseCase_AuthzFacebook(t *testing.T) {
	const accessToken string = "817832142661743|_Vx1271myLT4XIfOBUzSiJWYwRU"
	const accessTokenClient string = "EAALn0GJZAaG8BACKouZCfM5nzflnwsTy2PJWrOZAql0JmYPBYQc4vRKbl2y5flJzNhMJCRm6QZBOnFyZANGDRLifemxrBo9YJ0EAhfiVJwzN5wG51zJDxH1xYGiKmGCBG2HVkNHUrZCqBEIraCadQPtcpUJdNBA7EVIXBPc8VfLq5ZBR81xaLGTphqdJZAdufHgjliFy2rtE1gWFWUzl8fFm"
	mockResponse := map[string]interface{}{
		"data": map[string]interface{}{
			"app_id":                 "817832142661743",
			"type":                   "USER",
			"application":            "myApp",
			"data_access_expires_at": 1680798424,
			"error": map[string]interface{}{
				"code":    190,
				"message": "Error validating access token: Session has expired on Friday, 06-Jan-23 10:00:00 PST. The current time is Saturday, 07-Jan-23 23:33:31 PST.",
				"subcode": 463,
			},
			"expires_at": 1673028000,
			"is_valid":   false,
			"scopes": []string{
				"email",
				"public_profile",
			},
			"user_id": "1818786208488789",
		},
	}

	mockRequestService := new(mocks.IRequest)
	mockRequestService.On("Json", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("map[string]string"), mock.AnythingOfType("map[string]interface {}")).Return(mockResponse, nil)

	type fields struct {
		RequestService ports.IRequest
	}
	type args struct {
		accessTokenClient string
		accessToken       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test integration authorization facebook",
			fields{mockRequestService},
			args{
				accessTokenClient: accessTokenClient,
				accessToken:       accessToken,
			},
			true,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := oauth2UseCase{
				RequestService: tt.fields.RequestService,
			}
			got, err := o.AuthzFacebook(tt.args.accessTokenClient, tt.args.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthzFacebook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AuthzFacebook() got = %v, want %v", got, tt.want)
			}
		})
	}
}
