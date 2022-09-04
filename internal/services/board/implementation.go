package boardService

import "github.com/rahmaniali-ir/checklist-server/internal/models/board"

type iService struct {
	model *board.IBoard
}

var _ IService = &iService{}

func New(model board.IBoard) *iService {
	return &iService{
		model: &model,
	}
}

func (s *iService) List() []board.Board {
	return (*s.model).List()
}
