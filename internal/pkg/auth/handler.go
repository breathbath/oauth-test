package auth

import (
	"github.com/ory/fosite"
	"github.com/ory/fosite/handler/oauth2"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	defaultScope = "fosite"
	defaultAud   = "admins"
)

type TokenHandler struct {
	prov fosite.OAuth2Provider
}

func (th *TokenHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ctx := fosite.NewContext()

	sess := &oauth2.JWTSession{}
	accessRequest, err := th.prov.NewAccessRequest(ctx, req, sess)
	if err != nil {
		logrus.Infof("Access request failed because: %+v", err)
		logrus.Infof("Request: %+v", accessRequest)
		th.prov.WriteAccessError(rw, accessRequest, err)
		return
	}

	if accessRequest.GetRequestedScopes().Has(defaultScope) {
		accessRequest.GrantScope(defaultScope)
	}

	if accessRequest.GetRequestedAudience().Has(defaultAud) {
		accessRequest.GrantAudience(defaultAud)
	}

	response, err := th.prov.NewAccessResponse(ctx, accessRequest)
	if err != nil {
		logrus.Infof("Access request failed because: %+v", err)
		logrus.Infof("Request: %+v", accessRequest)
		th.prov.WriteAccessError(rw, accessRequest, err)
		return
	}

	th.prov.WriteAccessResponse(rw, accessRequest, response)
}

func NewTokenHandler(prov fosite.OAuth2Provider) (h http.Handler, err error) {
	return &TokenHandler{prov: prov}, nil
}
