package user

import "github.com/rahmaniali-ir/checklist-server/internal/http"

type IHandler interface {
	GetProfile(*http.GenericRequest) (interface{}, error)
	SignUp(*http.GenericRequest) (interface{}, error)
	SignIn(*http.GenericRequest) (interface{}, error)
	SignOut(*http.GenericRequest) (interface{}, error)
}
