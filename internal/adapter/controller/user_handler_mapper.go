package controller

import (
	"encoding/json"
	"io"
	"net/http"

	rest "github.com/Jiei-S/boilerplate-clean-architecture/go-rest/internal/infrastructure/openapi"
	usecase "github.com/Jiei-S/boilerplate-clean-architecture/go-rest/internal/usecase"
	pkgErr "github.com/Jiei-S/boilerplate-clean-architecture/go-rest/pkg/error"
)

func ToDTO(
	body io.ReadCloser,
) (*usecase.User, *pkgErr.ApplicationError) {
	var dto usecase.User
	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, pkgErr.NewApplicationError("failed to decode request body", pkgErr.LevelWarn, pkgErr.CodeBadRequest)
	}
	return &usecase.User{
		ID:        dto.ID,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Age:       dto.Age,
	}, nil
}

func FromDTO(
	dto *usecase.User,
) *rest.User {
	return &rest.User{
		Id:        dto.ID,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Age:       int(dto.Age),
	}
}

func NotFoundError(w http.ResponseWriter, err *pkgErr.ApplicationError) {
	setHeaderContentType(w)
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(rest.Error{
		Code:    http.StatusNotFound,
		Message: err.Error(),
	})
}

func BadRequestError(w http.ResponseWriter, err *pkgErr.ApplicationError) {
	setHeaderContentType(w)
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(rest.Error{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	})
}

func DuplicateError(w http.ResponseWriter, err *pkgErr.ApplicationError) {
	setHeaderContentType(w)
	w.WriteHeader(http.StatusConflict)
	json.NewEncoder(w).Encode(rest.Error{
		Code:    http.StatusConflict,
		Message: err.Error(),
	})
}

func InternalServerError(w http.ResponseWriter, err *pkgErr.ApplicationError) {
	setHeaderContentType(w)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(rest.Error{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	})
}

func HttpError(w http.ResponseWriter, err *pkgErr.ApplicationError) {
	switch err.Code() {
	case pkgErr.CodeBadRequest:
		BadRequestError(w, err)
	case pkgErr.CodeNotFound:
		NotFoundError(w, err)
	case pkgErr.CodeDuplicate:
		DuplicateError(w, err)
	default:
		InternalServerError(w, err)
	}
}
