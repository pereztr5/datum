package google

import (
	"net/http"

	"golang.org/x/oauth2"
	google "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"

	oauth2Login "github.com/datumforge/datum/internal/providers/oauth2"
	"github.com/datumforge/datum/internal/sessions"
)

const (
	ProviderName = "GOOGLE"
)

// StateHandler checks for a state cookie, if found, adds to context; if missing, a
// random generated value is added to the context and to a (short-lived) state cookie
// issued to the requester - this implements OAuth 2 RFC 6749 10.12 CSRF Protection
func StateHandler(config sessions.CookieConfig, success http.Handler) http.Handler {
	return oauth2Login.StateHandler(config, success)
}

// LoginHandler handles Google login requests by reading the state value from
// the ctx and redirecting requests to the AuthURL with that state value
func LoginHandler(config *oauth2.Config, failure http.Handler) http.Handler {
	return oauth2Login.LoginHandler(config, failure)
}

// CallbackHandler handles Google redirection URI requests and adds the Google
// access token and Userinfo to the ctx
func CallbackHandler(config *oauth2.Config, success, failure http.Handler) http.Handler {
	success = googleHandler(config, success, failure)
	return oauth2Login.CallbackHandler(config, success, failure)
}

// googleHandler is a http.Handler that gets the OAuth2 Token from the ctx
// to get the corresponding Google Userinfo
func googleHandler(config *oauth2.Config, success, failure http.Handler) http.Handler {
	if failure == nil {
		failure = DefaultFailureHandler
	}

	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		token, err := oauth2Login.TokenFromContext(ctx)

		if err != nil {
			ctx = WithError(ctx, err)
			failure.ServeHTTP(w, req.WithContext(ctx))

			return
		}

		httpClient := config.Client(ctx, token)
		googleService, err := google.NewService(ctx, option.WithHTTPClient(httpClient))

		if err != nil {
			ctx = WithError(ctx, err)
			failure.ServeHTTP(w, req.WithContext(ctx))

			return
		}

		userInfoPlus, err := googleService.Userinfo.Get().Do()
		err = validateResponse(userInfoPlus, err)

		if err != nil {
			ctx = WithError(ctx, err)
			failure.ServeHTTP(w, req.WithContext(ctx))

			return
		}

		ctx = WithUser(ctx, userInfoPlus)
		success.ServeHTTP(w, req.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

// validateResponse returns an error if we get unexpected things
func validateResponse(user *google.Userinfo, err error) error {
	if err != nil {
		return ErrUnableToGetGoogleUser
	}

	if user == nil || user.Id == "" {
		return ErrCannotValidateGoogleUser
	}

	return nil
}
