### Note
This package rely on HRMS V2 service

### Configuration
```sh
export AUTHORIZER_PROVIDER_BASE_URL=https://sit-public-api.freshket.co
```

### Example
- cmd/lambda/authorizer/main.go
```go
func main() {
	lambda.Start(authv2.HandleEvent)
}
```

- cmd/lambda/example/main.go
```go
func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	user, err := authv2.GetUserInfoFromEvent(req)

	if err != nil {
        // handle error
	}

    // type UserInfo struct {
    //      CreatedAt     time.Time
    //      Email         string
    //      EmailVerified bool
    //      FamilyName    string
    //      GivenName     string
    //      Identities    []struct {
    //      	Provider    string
    //      	AccessToken string
    //      	ExpiresIn   int
    //      	UserID      string
    //      	Connection  string
    //      	IsSocial    bool
    //      }
    //      Locale      string
    //      Name        string
    //      Nickname    string
    //      Picture     string
    //      UpdatedAt   time.Time
    //      UserID      string
    //      LastIP      string
    //      LastLogin   time.Time
    //      LoginsCount int
    // }
}
```

