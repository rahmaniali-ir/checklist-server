package boardService

import model "github.com/rahmaniali-ir/checklist-server/internal/models/board"

type iService struct {
	model *model.IBoard
}

var _ IService = &iService{}

func New(model model.IBoard) *iService {
	return &iService{
		model: &model,
	}
}

func (s *iService) List() []model.Board {
	return (*s.model).List()
}

func (s *iService) Create(title string, color string, icon string) (*model.Board, error) {
	return (*s.model).Create(model.Board{
		Title: title,
		Color: color,
		Icon: icon,
	})
}

func (s *iService) Update(board model.Board) (error) {
	return (*s.model).Update(board)
}

func (s *iService) Delete(uid string) (error) {
	return (*s.model).Delete(uid)
}
