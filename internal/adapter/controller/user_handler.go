package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Jiei-S/boilerplate-clean-architecture/internal/adapter/gateway"
	rest "github.com/Jiei-S/boilerplate-clean-architecture/internal/infrastructure/openapi"
	usecase "github.com/Jiei-S/boilerplate-clean-architecture/internal/usecase"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/uptrace/bun"

	pkgErr "github.com/Jiei-S/boilerplate-clean-architecture/pkg/error"
)

var _ rest.ServerInterface = (*UserHandler)(nil)

type UserHandler struct {
	db      *bun.DB
	usecase usecase.UserUsecase
}

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	req, err := ToDTO(r.Body)
	if err != nil {
		HttpError(w, err)
		return
	}

	result, err := h.usecase.AddUser(r.Context(), req)
	if err != nil {
		HttpError(w, err)
		return
	}

	h.HandleOK(w, FromDTO(result))
}

func (h *UserHandler) FindUser(w http.ResponseWriter, r *http.Request, id string) {
	result, err := h.usecase.FindUser(r.Context(), id)
	if err != nil {
		HttpError(w, err)
		return
	}

	h.HandleOK(w, FromDTO(result))
}

func (h *UserHandler) Health(w http.ResponseWriter, r *http.Request) {
	h.HandleOK(w, rest.Health{
		Status: rest.Healthy,
	})
}

func (h *UserHandler) HandleOK(w http.ResponseWriter, obj interface{}) {
	setHeaderContentType(w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(obj)
}

func setHeaderContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func (h *UserHandler) Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil && err != http.ErrAbortHandler {
				HttpError(w, pkgErr.NewApplicationError(err.(error).Error(), pkgErr.LevelError, pkgErr.CodeInternalServerError))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (h *UserHandler) SetDBMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.db.RunInTx(r.Context(), nil, func(ctx context.Context, tx bun.Tx) error {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			next.ServeHTTP(ww, r.WithContext(context.WithValue(r.Context(), gateway.TX_KEY, &tx)))
			return nil
		})
	})
}

func NewUserHandler(
	db *bun.DB,
	usecase usecase.UserUsecase,
) *UserHandler {
	return &UserHandler{
		db:      db,
		usecase: usecase,
	}
}
