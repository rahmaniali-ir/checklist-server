package board

type Status int
const (
	Status_Undone Status = iota
	Status_In_Progress
	Status_Done
)

type ListItem struct {
	Uid string `json:"uid"`
	ListUid string `json:"listUid"`
	BoardUid string `json:"boardUid"`
	Body string `json:"body"`
	Status Status `json:"status"`
}

type CheckList struct {
	Uid string `json:"uid"`
	BoardUid string `json:"boardUid"`
	Title string `json:"title"`
	Items []ListItem `json:"items"`
}

type Board struct {
	Uid string `json:"uid"`
	Title string `json:"title"`
	Color string `json:"color"`
	Icon string `json:"icon"`
	Image string `json:"image"`
	Lists []CheckList `json:"lists"`
}

type IBoard interface {
	List() []Board
	Get(uid string) (*Board, error)
	Create(board Board) (*Board, error)
	Update(board Board) (*Board, error)
	Delete(uid string) error
}
