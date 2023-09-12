package addyapi

func (ac *AddyClient) GetRules() (*RulesResp, error) {
	var data RulesResp
	res, err := ac.Client.R().SetResult(&data).Get("api/v1/rules")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, &ErrAddyResponse{res.StatusCode(), res.Status()}
	}
	return &data, nil
}

func (ac *AddyClient) GetRule(id string) (*RuleResp, error) {
	var data RuleResp
	res, err := ac.Client.R().SetResult(&data).Get("api/v1/rules/" + id)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, &ErrAddyResponse{res.StatusCode(), res.Status()}
	}
	return &data, nil
}
