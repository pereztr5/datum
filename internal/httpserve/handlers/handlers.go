package handlers

import (
	"github.com/lestrrat-go/jwx/v2/jwk"
	"go.uber.org/zap"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/tokens"
)

// Handler contains configuration options for handlers
type Handler struct {
	// DBClient to interact with the generated ent schema
	DBClient *ent.Client
	// TM contains the token manager in order to validate auth requests
	TM *tokens.TokenManager
	// CookieDomain is the domain set in cookie for authenticated requests, defaults to datum.net
	CookieDomain string
	// Logger provides the zap logger to do logging things from the handlers
	Logger *zap.SugaredLogger
	// ReadyChecks is a set of checkFuncs to determine if the application is "ready" upon startup
	ReadyChecks Checks
	// JWTKeys contains the set of valid JWT authentication key
	JWTKeys jwk.Set
}

type Response struct {
	Message string `json:"message"`
}
