package addyapi

func (ac *AddyClient) GetFailedDeliveries() (*FailedDeliveriesResp, error) {
	var data FailedDeliveriesResp
	res, err := ac.Client.R().SetResult(&data).Get("api/v1/failed-deliveries")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, &ErrAddyResponse{res.StatusCode(), res.Status()}
	}
	return &data, nil
}

func (ac *AddyClient) GetFailedDelivery(id string) (*FailedDeliveryResp, error) {
	var data FailedDeliveryResp
	res, err := ac.Client.R().SetResult(&data).Get("api/v1/failed-deliveries/" + id)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, &ErrAddyResponse{res.StatusCode(), res.Status()}
	}
	return &data, nil
}
