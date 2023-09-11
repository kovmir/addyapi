package addyapi

import "fmt"

func (ac *AddyClient) GetAliases(params *AliasParams) (*AliasesResp, error) {
	var data AliasesResp

	queryParams := make(map[string]string)
	for key, val := range params.Filter {
		queryParams["filter["+key+"]"] = val
	}
	if params.Sort != "" {
		if params.SortRev { // Prepend "-" to reverse sort order.
			params.Sort = "-" + params.Sort
		}
		queryParams["sort"] = string(params.Sort)
	}
	if params.PageNumber > 0 {
		queryParams["page[number]"] = fmt.Sprintf("%v", params.PageNumber)
	}
	if params.PageSize > 0 {
		queryParams["page[size]"] = fmt.Sprintf("%v", params.PageSize)
	}

	res, err := ac.Client.R().
		SetQueryParams(queryParams).
		SetResult(&data).
		Get("api/v1/aliases")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, &ErrAddyResponse{res.StatusCode(), res.Status()}
	}
	return &data, nil
}

func (ac *AddyClient) GetAlias(id string) (*AliasResp, error) {
	var data AliasResp
	res, err := ac.Client.R().SetResult(&data).Get("api/v1/aliases/" + id)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, &ErrAddyResponse{res.StatusCode(), res.Status()}
	}
	return &data, nil
}
