package auth

import (
	"crypto/rsa"
	"github.com/ory/fosite"
	"github.com/ory/fosite/compose"
	"github.com/ory/fosite/handler/oauth2"
	"github.com/ory/fosite/storage"
	"github.com/ory/fosite/token/hmac"
	"github.com/ory/fosite/token/jwt"
	"time"
)

var (
	accessTokenLifespan = time.Hour
	defaultIssuer     = "breathbath-test-app"
	tokenURL          = "http://localhost"
)

var hmacStrategy = &oauth2.HMACSHAStrategy{
	Enigma: &hmac.HMACStrategy{
		GlobalSecret: []byte("some-super-cool-secret-that-nobody-knows"),
	},
	AccessTokenLifespan: accessTokenLifespan,
}

type RSAKeyProvider func() (key *rsa.PrivateKey, err error)

func createJWTStrategy(privateKey *rsa.PrivateKey) jwt.JWTStrategy {
	return &oauth2.DefaultJWTStrategy{
		JWTStrategy: &jwt.RS256JWTStrategy{
			PrivateKey: privateKey,
		},
		HMACSHAStrategy: hmacStrategy,
		Issuer:          defaultIssuer,
	}
}

func createStore() fosite.Storage {
	return &storage.MemoryStore{
		Clients: map[string]fosite.Client{
			"admin": &fosite.DefaultClient{
				ID:            "admin",
				Secret:        []byte(`$2a$10$IxMdI6d.LIRZPpSfEwNoeu4rY3FhDREsxFJXikcgdRRAStxUlsuEO`), // = "foobar"
				ResponseTypes: []string{"token"},
				GrantTypes:    []string{"implicit", "refresh_token", "authorization_code", "password", "client_credentials"},
				Scopes:        []string{defaultScope},
				Audience:      []string{defaultAud},
			},
		},
		BlacklistedJTIs:        map[string]time.Time{},
		AuthorizeCodes:         map[string]storage.StoreAuthorizeCode{},
		PKCES:                  map[string]fosite.Requester{},
		AccessTokens:           map[string]fosite.Requester{},
		RefreshTokens:          map[string]fosite.Requester{},
		IDSessions:             map[string]fosite.Requester{},
		AccessTokenRequestIDs:  map[string]string{},
		RefreshTokenRequestIDs: map[string]string{},
	}
}

func NewOauthProvider(keyProvider RSAKeyProvider) (prov fosite.OAuth2Provider, err error) {
	key, err := keyProvider()
	if err != nil {
		return nil, err
	}

	fositeStore := createStore()

	jwtStrategy := createJWTStrategy(key)
	provider := compose.Compose(
		&compose.Config{
			GrantTypeJWTBearerIDOptional:         true,
			GrantTypeJWTBearerIssuedDateOptional: false,
			TokenURL:                             tokenURL,
		},
		fositeStore,
		jwtStrategy,
		nil,
		compose.OAuth2ClientCredentialsGrantFactory,
		compose.RFC7523AssertionGrantFactory,
	)

	return provider, nil
}
