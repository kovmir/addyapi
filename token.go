package addyapi

func (ac *AddyClient) GetAPITokenDetails() (*APITokenDetails, error) {
	var data APITokenDetails
	res, err := ac.Client.R().SetResult(&data).Get("api/v1/api-token-details")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, &ErrAddyResponse{res.StatusCode(), res.Status()}
	}
	return &data, nil
}
