package boardHandler

import "github.com/rahmaniali-ir/checklist-server/internal/http"

type IHandler interface {
	List(*http.GenericRequest) (interface{}, error)
	Create(*http.GenericRequest) (interface{}, error)
	Delete(*http.GenericRequest) (interface{}, error)
	Update(*http.GenericRequest) (interface{}, error)
}
