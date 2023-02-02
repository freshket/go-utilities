package authv2

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

var (
	ErrTokenIsRequired = errors.New("token is required")
	ErrUnauthorized    = errors.New("unauthorized")
)

type policyEffect string

var (
	policyEffectAllow policyEffect = "Allow"
	policyEffectDeny  policyEffect = "Deny"
)

const (
	policyPayloadKey = "x-user-info"
)

func generatePolicy(event events.APIGatewayCustomAuthorizerRequest, effect policyEffect, payload string) events.APIGatewayCustomAuthorizerResponse {
	r := events.APIGatewayCustomAuthorizerResponse{
		PrincipalID: "user",
		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   string(effect),
					Resource: []string{event.MethodArn},
				},
			},
		},
	}

	if len(payload) > 0 {
		r.Context = map[string]interface{}{
			policyPayloadKey: payload,
		}
	}

	return r
}

func HandleEvent(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	token := strings.Replace(strings.TrimSpace(event.AuthorizationToken), "Bearer ", "", 1)

	if len(strings.TrimSpace(token)) == 0 {
		return generatePolicy(event, policyEffectDeny, ""), nil
	}

	u, err := getUserInfo(ctx, token)

	if err != nil {
		log.Println("[DEBUG][HandleEvent] err", err)
		return generatePolicy(event, policyEffectDeny, ""), nil
	}

	uj, err := json.Marshal(u)

	if err != nil {
		log.Println("[DEBUG][HandleEvent] err", err)
		return generatePolicy(event, policyEffectDeny, ""), nil
	}

	return generatePolicy(event, policyEffectAllow, string(uj)), nil
}

func GetUserInfoFromEvent(req events.APIGatewayProxyRequest) (*UserInfo, error) {
	userjson, ok := req.RequestContext.Authorizer[policyPayloadKey].(string)

	if !ok {
		return nil, ErrTokenIsRequired
	}

	var user UserInfo

	err := json.Unmarshal([]byte(userjson), &user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

type UserInfo struct {
	CreatedAt     time.Time
	Email         string
	EmailVerified bool
	FamilyName    string
	GivenName     string
	Identities    []struct {
		Provider    string
		AccessToken string
		ExpiresIn   int
		UserID      string
		Connection  string
		IsSocial    bool
	}
	Locale      string
	Name        string
	Nickname    string
	Picture     string
	UpdatedAt   time.Time
	UserID      string
	LastIP      string
	LastLogin   time.Time
	LoginsCount int
}
