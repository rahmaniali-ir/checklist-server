package user

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/rahmaniali-ir/checklist-server/internal/http"
	userModel "github.com/rahmaniali-ir/checklist-server/internal/models/user"
	userService "github.com/rahmaniali-ir/checklist-server/internal/services/user"
	"github.com/rahmaniali-ir/checklist-server/pkg/session"
)

type handler struct {
	service userService.IUser
}

var _ IHandler = &handler{}

func NewHandler(service userService.IUser) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetProfile(req *http.GenericRequest) (interface{}, error) {
	if req.Session == nil {
		return nil, errors.New("Unauthorized request!")
	}
	
	return h.service.Get(req.Session.Uid)
}

func (h *handler) SignUp(req *http.GenericRequest) (interface{}, error) {
	user := userModel.User{}
	reader := bytes.NewReader(req.Body)
	err := json.NewDecoder(reader).Decode(&user)
	if err != nil {
		return nil, err
	}

	addedUser, err := h.service.Add(user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	token, err := session.Default.SetSession(addedUser.Uid)
	if err != nil {
		return nil, err
	}

	userWithToken := userModel.UserWithToken{
		User: *addedUser,
		Token: token,
	}

	return userWithToken, nil
}

func (h *handler) SignIn(req *http.GenericRequest) (interface{}, error) {
	credentials := userModel.Credentials{}
	reader := bytes.NewReader(req.Body)
	err := json.NewDecoder(reader).Decode(&credentials)
	if err != nil {
		return nil, err
	}

	user, err := h.service.GetByCredentials(credentials.Email, credentials.Password)
	if err != nil {
		return nil, err
	}

	token, err := session.Default.SetSession(user.Uid)
	if err != nil {
		return nil, err
	}

	userWithToken := userModel.UserWithToken{
		User: *user,
		Token: token,
	}

	return userWithToken, nil
}

func (h *handler) SignOut(req *http.GenericRequest) (interface{}, error) {
	err := session.Default.UnsetSession(req.Session.Token)
	return nil, err
}
