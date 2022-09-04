package boardService

import "github.com/rahmaniali-ir/checklist-server/internal/models/board"

type IService interface {
	List() []board.Board
	Create(title string, color string, icon string) (*board.Board, error)
	Delete(uid string) error
}
