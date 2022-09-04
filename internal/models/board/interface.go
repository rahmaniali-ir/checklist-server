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
	Uid string `json:"uid" bson:"_id"`
	Title string `json:"title" bson:"title"`
	Color string `json:"color" bson:"color"`
	Icon string `json:"icon" bson:"icon"`
	Image string `json:"image" bson:"image"`
	Lists []CheckList `json:"lists" bson:"lists"`
}

type IBoard interface {
	List() []Board
	Get(uid string) (*Board, error)
	Create(board Board) (*Board, error)
	Update(board Board) error
	Delete(uid string) error
}
