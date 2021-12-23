package mal

import "net/url"

func (client *Client) GetMyUserInformation(opts ...OptFn) (*AnimeStatistics, error) {
	values := make(url.Values)
	for _, opt := range opts {
		opt(values)
	}
	valuesStr := values.Encode()

	u := baseURL + "/users/@me"
	if valuesStr != "" {
		u += "?" + valuesStr
	}

	var statistics AnimeStatistics
	err := client.doGet(u, &statistics)
	if err != nil {
		return nil, err
	}

	return &statistics, nil
}
