package model

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/dopeCape/player-managment/pkg/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Player struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Score   int    `json:"score"`
}

var collection *mongo.Collection

func init() {

	app.ConnectDB()

	DB := app.GetDB()
	collection = DB.Database("players").Collection("scores")
}

func AddPlayer(p Player) error {
	_, err := collection.InsertOne(context.Background(), p)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePlayer(playerId string, name string, score int) (Player, error) {
	filter := bson.D{{Key: "id", Value: playerId}}
	var updatedPlayer Player
	if score != -1 {
		update := bson.D{{Key: "$set", Value: bson.D{{Key: "score", Value: score}}}}
		err := collection.FindOneAndUpdate(context.Background(), filter, update).Decode(&updatedPlayer)
		fmt.Println(updatedPlayer)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return Player{}, fmt.Errorf("No player with Id:%v", playerId)
			}

			return Player{}, err
		}
		updatedPlayer.Score = score
	}
	if len(name) != 0 {
		update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: name}}}}
		err := collection.FindOneAndUpdate(context.Background(), filter, update).Decode(&updatedPlayer)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return Player{}, fmt.Errorf("No player with Id:%v", playerId)
			}

			return Player{}, err
		}
		updatedPlayer.Name = name
	}

	return updatedPlayer, nil

}

func DeletePlayer(playerId string) {
	filter := bson.D{{Key: "id", Value: playerId}}
	collection.FindOneAndDelete(context.Background(), filter)
}

func GetAllPlayers() ([]Player, error) {
	var players []Player
	playersFound, err := collection.Find(context.Background(), bson.D{{
		Key:   "",
		Value: nil,
	}})
	if err != nil {
		return players, err
	}
	for playersFound.Next(context.TODO()) {
		var result Player
		if err := playersFound.Decode(&result); err != nil {
			log.Fatal(err)
		}
		players = append(players, result)
	}
	if err := playersFound.Err(); err != nil {
		return players, err
	}

	sortedPlayers := SortWRTRank(players)
	return sortedPlayers, nil
}

func GetPlayerFromRanK(rank int) (Player, error) {
	players, err := GetAllPlayers()
	if err != nil {
		return Player{}, err
	}
	if rank >= len(players) {
		return Player{}, fmt.Errorf("Rank should be less than %v(Total no of players registered )", len(players))
	}
	return players[rank-1], nil
}

func GetRandomPlayer() (Player, error) {
	players, err := GetAllPlayers()
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(players))
	if err != nil {
		return Player{}, nil
	}
	return players[randomIndex], nil
}
