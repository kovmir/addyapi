package addyapi

func (ac *AddyClient) GetDomainOptions() (*DomainOptions, error) {
	var data DomainOptions
	res, err := ac.Client.R().SetResult(&data).Get("api/v1/domain-options")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, &ErrAddyResponse{res.StatusCode(), res.Status()}
	}
	return &data, nil
}
