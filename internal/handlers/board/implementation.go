package boardHandler

import (
	"bytes"
	"encoding/json"

	"github.com/rahmaniali-ir/checklist-server/internal/http"
	"github.com/rahmaniali-ir/checklist-server/internal/models/board"
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

func (h iHandler) Create(req *http.GenericRequest) (interface{}, error) {
	board := board.Board{}
	reader := bytes.NewReader(req.Body)
	json.NewDecoder(reader).Decode(&board)

	return (*h.service).Create(board.Title, board.Color, board.Icon)
}

func (h iHandler) Delete(req *http.GenericRequest) (interface{}, error) {
	return nil, (*h.service).Delete(req.PathParams["uid"])
}

func (h iHandler) Update(req *http.GenericRequest) (interface{}, error) {
	board := board.Board{}
	reader := bytes.NewReader(req.Body)
	json.NewDecoder(reader).Decode(&board)

	return nil, (*h.service).Update(board)
}
