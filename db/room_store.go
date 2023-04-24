package db

import (
	"context"

	"github.com/jinheehanaaa/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
}

type MongoRoomStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoRoomStore(client *mongo.Client, dbname string) *MongoRoomStore {
	return &MongoRoomStore{
		client: client,
		coll:   client.Database(dbname).Collection("rooms"),
	}
}

func (s *MongoRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	resp, err := s.coll.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.ID = resp.InsertedID.(primitive.ObjectID)

	// update the hotel with this room id
	//filter := bson.M{"_id": room.HotelID}
	//update := bson.M{"$push": bson.M{"rooms": room.ID}}
	//if err :=

	return room, nil
}
