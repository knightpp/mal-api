package mal

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

const baseURL = "https://api.myanimelist.net/v2"

type Client struct {
	token *oauth2.Token
	debug bool
}

func NewClient(token *oauth2.Token) (*Client, error) {
	return &Client{token, false}, nil
}

type reqOpt func(req *http.Request)

func setHeader(k, v string) reqOpt {
	return func(req *http.Request) {
		req.Header.Set(k, v)
	}
}

func (client *Client) SetDebug(b bool) {
	client.debug = b
}

func (client *Client) requestAndDecode(method, u string, body io.Reader, out interface{}, opts ...reqOpt) error {
	resp, err := client.request(method, u, body, opts...)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return client.decode(resp, out)
}

func (client *Client) request(method, u string, body io.Reader, opts ...reqOpt) (*http.Response, error) {
	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		opt(req)
	}

	client.token.SetAuthHeader(req)

	log.Debug("sending request", zap.String("url", u), zap.String("method", method))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Debug("request failed", zap.Error(err))
		return nil, err
	}

	return resp, nil
}

func (client *Client) doGet(u string, out interface{}) error {
	return client.requestAndDecode(http.MethodGet, u, nil, out)
}

func (client *Client) doPatch(u string, body io.Reader, out interface{}) error {
	return client.requestAndDecode(http.MethodPatch, u, body, out,
		setHeader("Content-Type", "application/x-www-form-urlencoded"))
}

func (client *Client) doDelete(u string) error {
	resp, err := client.request(http.MethodDelete, u, nil)
	if err != nil {
		return err
	}

	return client.decode(resp, nil)
}

func (client *Client) decode(resp *http.Response, out interface{}) error {
	var body io.Reader = resp.Body

	if client.debug {
		bodyClone, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		path := strings.ReplaceAll(resp.Request.URL.Path, "/", "_") + ".json"
		log.Debug("request dump", zap.String("path", path))
		err = os.WriteFile(path, bodyClone, 0666)
		if err != nil {
			log.Error("write dump failed", zap.Error(err))
			return err
		}

		body = strings.NewReader(string(bodyClone))
	}

	if resp.StatusCode != 200 {
		var apiErr ApiError
		err := json.NewDecoder(body).Decode(&apiErr)
		if err != nil {
			return err
		}
		return &apiErr
	}

	if out == nil {
		return nil
	}

	err := json.NewDecoder(body).Decode(out)
	if err != nil {
		return err
	}
	return nil
}
