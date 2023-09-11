package addyapi

func (ac *AddyClient) GetRecipients(params *RecipientsParams) (*RecipientsResp, error) {
	var data RecipientsResp

	queryParams := make(map[string]string)
	for key, val := range params.Filter {
		queryParams["filter["+key+"]"] = val
	}

	res, err := ac.Client.R().
		SetQueryParams(queryParams).
		SetResult(&data).
		Get("api/v1/recipients")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, &ErrAddyResponse{res.StatusCode(), res.Status()}
	}
	return &data, nil
}

func (ac *AddyClient) GetRecipient(id string) (*RecipientResp, error) {
	var data RecipientResp
	res, err := ac.Client.R().SetResult(&data).Get("api/v1/recipients/" + id)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, &ErrAddyResponse{res.StatusCode(), res.Status()}
	}
	return &data, nil
}
