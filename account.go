package addyapi

func (ac *AddyClient) GetAccountDetails() (*AccountDetailsResp, error) {
	var data AccountDetailsResp
	res, err := ac.Client.R().SetResult(&data).Get("api/v1/account-details")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, &ErrAddyResponse{res.StatusCode(), res.Status()}
	}
	return &data, nil
}
