package boardService

import (
	model "github.com/rahmaniali-ir/checklist-server/internal/models/board"
)

type IService interface {
	List() []model.Board
	Create(title string, color string, icon string) (*model.Board, error)
	Delete(uid string) error
	Update(board model.Board) error
}
