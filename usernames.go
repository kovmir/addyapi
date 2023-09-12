package addyapi

func (ac *AddyClient) GetUsernames() (*UsernamesResp, error) {
	var data UsernamesResp
	res, err := ac.Client.R().SetResult(&data).Get("api/v1/usernames")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, &ErrAddyResponse{res.StatusCode(), res.Status()}
	}
	return &data, nil
}

func (ac *AddyClient) GetUsername(id string) (*UsernameResp, error) {
	var data UsernameResp
	res, err := ac.Client.R().SetResult(&data).Get("api/v1/usernames/" + id)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, &ErrAddyResponse{res.StatusCode(), res.Status()}
	}
	return &data, nil
}
