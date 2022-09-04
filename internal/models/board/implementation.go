package board

type iBoard struct {}

var _ IBoard = &iBoard{}

func New() *iBoard {
	return &iBoard{}
}

func (b *iBoard) List() []Board {
	return []Board{
		{
			Uid: "1",
			Title: "Board 1",
			Color: "",
			Icon: "",
			Image: "",
			Lists: []CheckList{},
		},
		{
			Uid: "2",
			Title: "Board 2",
			Color: "",
			Icon: "",
			Image: "",
			Lists: []CheckList{},
		},
	}
}

func (b *iBoard) Get(uid string) (*Board, error) {
	return &Board{
		Uid: uid,
		Title: "",
		Color: "",
		Icon: "",
		Image: "",
		Lists: []CheckList{},
	}, nil
}

func (b *iBoard) Create(board Board) (*Board, error) {
	return &board, nil
}

func (b *iBoard) Update(board Board) (*Board, error) {
	return &board, nil
}

func (b *iBoard) Delete(uid string) error {
	return nil
}
