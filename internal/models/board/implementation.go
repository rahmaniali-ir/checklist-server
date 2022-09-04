package board

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type iBoard struct {
	db *mongo.Database
}

var _ IBoard = &iBoard{}

func New(db *mongo.Database) *iBoard {
	return &iBoard{
		db: db,
	}
}

func (b *iBoard) List() []Board {
	boardsCollection := b.db.Collection("boards")
	ctx, _ := context.WithTimeout((context.Background()), 10*time.Second)
	
	var boards []Board
	cursor, err := boardsCollection.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &boards); err != nil {
		panic(err)
	}
	
	return boards
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
