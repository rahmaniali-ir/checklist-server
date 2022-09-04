package board

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		Uid: "",
		Title: "",
		Color: "",
		Icon: "",
		Image: "",
		Lists: []CheckList{},
	}, nil
}

func (b *iBoard) Create(board Board) (*Board, error) {
	boards := b.db.Collection("boards")
	ctx, _ := context.WithTimeout((context.Background()), 10*time.Second)

	result, err := boards.InsertOne(ctx, bson.D{
		{ Key: "title", Value: board.Title },
		{ Key: "color", Value: board.Color },
		{ Key: "icon", Value: board.Icon },
		{ Key: "image", Value: board.Image },
		{ Key: "lists", Value: []CheckList{}},
	})

	if err != nil {
		return nil, err
	}

	if uid, ok := result.InsertedID.(primitive.ObjectID); ok {
		board.Uid = uid.Hex()
	}

	return &board, nil
}

func (b *iBoard) Update(board Board) error {
	boards := b.db.Collection("boards")
	ctx, _ := context.WithTimeout((context.Background()), 10*time.Second)

	objectID, err := primitive.ObjectIDFromHex(board.Uid)
	if err != nil {
		return err
	}

	_, err = boards.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{
		"$set": bson.D{
			{ Key: "title", Value: board.Title },
			{ Key: "color", Value: board.Color },
			{ Key: "icon", Value: board.Icon },
			{ Key: "image", Value: board.Image },
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (b *iBoard) Delete(uid string) error {
	boards := b.db.Collection("boards")
	ctx, _ := context.WithTimeout((context.Background()), 10*time.Second)

	objectID, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		return err
	}

	_, err = boards.DeleteOne(ctx, bson.M{ "_id": objectID })

	return err
}
