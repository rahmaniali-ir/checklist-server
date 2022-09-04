package boardService

import "github.com/rahmaniali-ir/checklist-server/internal/models/board"

type IService interface {
	List() []board.Board
}
