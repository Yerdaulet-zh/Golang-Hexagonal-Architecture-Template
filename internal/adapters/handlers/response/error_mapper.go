package response

import (
	"net/http"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/domain"
)

// MapError translates Domain errors into HTTP responses
func MapErrorToResponse(w http.ResponseWriter, err error) {
	switch err {
	// Business Errors
	case domain.ErrInvalidEmailFormat, domain.ErrInvalidEmailHost:
		WriteError(w, http.StatusBadRequest, err.Error())
	case domain.ErrInvalidMessageLenght:
		WriteError(w, http.StatusUnprocessableEntity, "Message exceeds 255 characters")

	// Repository Errors
	case domain.ErrDatabaseInternalError:
		WriteError(w, http.StatusInternalServerError, err.Error())

	default:
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
