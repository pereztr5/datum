package handlers

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
)

// UserInfo returns the user information for the authenticated user
func (h *Handler) UserInfo(ctx echo.Context) error {
	// setup view context
	context := ctx.Request().Context()
	userCtx := viewer.NewContext(context, viewer.NewUserViewerFromSubject(context))

	userID, err := auth.GetUserIDFromContext(context)
	if err != nil {
		h.Logger.Errorw("unable to get user id from context", "error", err)

		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	// get user from database by subject
	user, err := h.getUserBySub(userCtx, userID)
	if err != nil {
		h.Logger.Errorw("unable to get user by subject", "error", err)

		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, user)
}
