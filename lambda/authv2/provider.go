package authv2

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type userInfoResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		CreatedAt     time.Time `json:"created_at"`
		Email         string    `json:"email"`
		EmailVerified bool      `json:"email_verified"`
		FamilyName    string    `json:"family_name"`
		GivenName     string    `json:"given_name"`
		Identities    []struct {
			Provider    string `json:"provider"`
			AccessToken string `json:"access_token"`
			ExpiresIn   int    `json:"expires_in"`
			UserID      string `json:"user_id"`
			Connection  string `json:"connection"`
			IsSocial    bool   `json:"isSocial"`
		} `json:"identities"`
		Locale      string    `json:"locale"`
		Name        string    `json:"name"`
		Nickname    string    `json:"nickname"`
		Picture     string    `json:"picture"`
		UpdatedAt   time.Time `json:"updated_at"`
		UserID      string    `json:"user_id"`
		LastIP      string    `json:"last_ip"`
		LastLogin   time.Time `json:"last_login"`
		LoginsCount int       `json:"logins_count"`
	} `json:"data"`
}

func getUserInfo(ctx context.Context, token string) (*UserInfo, error) {
	var resppayload userInfoResponse

	resp, err := resty.New().
		R().
		SetContext(ctx).
		SetHeader("x-access-token", token).
		SetResult(&resppayload).
		Get(
			fmt.Sprintf(
				"%s/hrms/v2/userinfo",
				os.Getenv("AUTHORIZER_PROVIDER_BASE_URL"),
			),
		)

	if err != nil || resppayload.Data.UserID == "" {
		log.Println("[ERROR][getUserInfo] err", err)
		log.Println("[ERROR][getUserInfo] resp.String()", resp.String())
		return nil, ErrUnauthorized
	}

	return (*UserInfo)(&resppayload.Data), nil
}
