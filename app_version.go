package addyapi

func (ac *AddyClient) GetAppVersion() (*AppVersion, error) {
	var data AppVersion
	res, err := ac.Client.R().SetResult(&data).Get("api/v1/app-version")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, &ErrAddyResponse{res.StatusCode(), res.Status()}
	}
	return &data, nil
}
