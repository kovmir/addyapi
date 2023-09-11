package addyapi

func (ac *AddyClient) GetDomains() (*DomainsResp, error) {
	var data DomainsResp
	res, err := ac.Client.R().SetResult(&data).Get("api/v1/domains")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, &ErrAddyResponse{res.StatusCode(), res.Status()}
	}
	return &data, nil
}

func (ac *AddyClient) GetDomain(id string) (*DomainResp, error) {
	var data DomainResp
	res, err := ac.Client.R().SetResult(&data).Get("api/v1/domains/" + id)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, &ErrAddyResponse{res.StatusCode(), res.Status()}
	}
	return &data, nil
}
