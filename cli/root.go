package cli

import (
	"github.com/breathbath/oauth-test/internal/pkg/auth"
	"github.com/breathbath/oauth-test/internal/pkg/enc"
	"github.com/breathbath/oauth-test/internal/pkg/server"
	"net/http"
)

const (
	tokenAPIPath = "/token"
)

func StartServer() error {
	prov, err := auth.NewOauthProvider(enc.DefaultKeyProvider)
	if err != nil {
		return err
	}

	tokenHandler, err := auth.NewTokenHandler(prov)
	if err != nil {
		return err
	}

	handlers := map[string]http.Handler{
		tokenAPIPath: tokenHandler,
	}

	return server.Start(handlers)
}
