package boardHandler

import (
	"github.com/rahmaniali-ir/checklist-server/internal/http"
	boardService "github.com/rahmaniali-ir/checklist-server/internal/services/board"
)

type iHandler struct {
	service *boardService.IService
}

var _ IHandler = &iHandler{}

func New(service boardService.IService) *iHandler {
	return &iHandler{
		service: &service,
	}
}

func (h iHandler) List(req *http.GenericRequest) (interface{}, error) {
	return (*h.service).List(), nil
}
