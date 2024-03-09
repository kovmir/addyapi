package addyapi

func (c *Client) AppGetVersion() (*AppVersion, error) {
	return get[AppVersion](c, "api/v1/app-version")
}
