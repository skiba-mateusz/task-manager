package user

import (
	"errors"
	"net/http"

	"github.com/skiba-mateusz/task-manager/config"
	"github.com/skiba-mateusz/task-manager/models"
	"github.com/skiba-mateusz/task-manager/services/auth"

	"github.com/skiba-mateusz/task-manager/utils"
)

type Handler struct {
	store models.UserStore
}

func NewHandler(store models.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var payload models.RegisterUserPayload
	if err := utils.ReadJSON(w, r, &payload); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteJSONError(w, http.StatusBadRequest, errors.New("user with provided email already exists"))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(models.User{
		Username: payload.Username,
		Email: payload.Email,
		Password: hashedPassword,
	})
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.JSONResponse(w, http.StatusCreated, nil); err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err)
	}
}


func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var payload models.LoginUserPayload
	if err := utils.ReadJSON(w, r, &payload); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, errors.New("user not found, invalid email or password"))
		return
	}

	if !auth.ComparePassword(user.Password, payload.Password) {
		utils.WriteJSONError(w, http.StatusBadRequest, errors.New("invalid email or passowrd"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.GenerateJWT(secret, int(user.ID))
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.JSONResponse(w, http.StatusOK, map[string]string{"token": token}); err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *Handler) GetAuthenticatedUser(w http.ResponseWriter, r *http.Request) {
	user := auth.GetUserFromContext(r.Context())
	utils.JSONResponse(w, http.StatusOK, user)
}