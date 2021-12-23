package mal

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/oauth2"
)

type OAuth struct {
	config oauth2.Config
}

// 52fbc429dc04c003c905a3b1a44269b7
func NewOauth(clientID string) *OAuth {
	config := oauth2.Config{
		ClientID: clientID,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://myanimelist.net/v1/oauth2/authorize",
			TokenURL: "https://myanimelist.net/v1/oauth2/token",
		},
		Scopes: []string{"write:users"},
	}

	return &OAuth{
		config: config,
	}
}

func (auth *OAuth) NewCodeVerifier() string {
	rnd := make([]byte, 128)
	_, err := rand.Read(rnd)
	if err != nil {
		panic("couldn't generate random")
	}

	var buf bytes.Buffer
	_, err = base64.NewEncoder(base64.RawURLEncoding, &buf).Write(rnd)
	if err != nil {
		panic("couldn't generate base64")
	}

	return buf.String()[:128]
}

func (auth *OAuth) AuthCodeURL(state string, challenge string) string {
	return auth.config.AuthCodeURL(
		"state",
		oauth2.SetAuthURLParam("code_challenge", challenge),
		oauth2.SetAuthURLParam("code_challenge_method", "plain"),
	)
}

func (auth *OAuth) Exchange(ctx context.Context, code, verifier string) (*oauth2.Token, error) {
	return auth.config.Exchange(
		context.Background(),
		code,
		oauth2.SetAuthURLParam("code_verifier", verifier),
	)
}
