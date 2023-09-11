package addyapi

import "github.com/go-resty/resty/v2"

const AddyBaseURL = "https://app.addy.io"

type AddyClient struct {
	Client *resty.Client
}

func NewClient(token string) *AddyClient {
	return &AddyClient{
		Client: resty.New().
			SetBaseURL(AddyBaseURL).
			SetAuthToken(token).
			SetHeader("Accept", "application/json")}
}

func NewCustomClient(token string, baseURL string) *AddyClient {
	return &AddyClient{
		Client: resty.New().
			SetBaseURL(baseURL).
			SetAuthToken(token).
			SetHeader("Accept", "application/json")}
}
