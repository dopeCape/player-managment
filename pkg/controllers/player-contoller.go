package controller

import (
	"fmt"
	"net/http"
	"strconv"

	model "github.com/dopeCape/player-managment/pkg/models"
	"github.com/dopeCape/player-managment/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func HandleNewPlayer(c *gin.Context) {
	id, error := uuid.NewGen().NewV1()
	var player model.Player
	if error != nil {
		utils.ResSender(c, error.Error(), "Internal server Error", http.StatusInternalServerError, gin.MIMEJSON)
	}
	name := c.PostForm("name")
	country := c.PostForm("country")
	score, error := strconv.Atoi(c.PostForm("score"))
	if error != nil {
		utils.ResSender(c, "Score is required and should be a valid positive Integer", "Invalid Params", http.StatusBadRequest, gin.MIMEJSON)
		return
	}
	if len(name) < 1 || len(name) > 15 {
		utils.ResSender(c, "Name  is required ,should be a valid string and should have a max length of 15 chars", "Invalid Params", http.StatusBadRequest, gin.MIMEJSON)
		return
	}
	if len(country) != 2 {
		utils.ResSender(c, "Country code is required , should be valid string and should have 2 Characters", "Invalid Params", http.StatusBadRequest, gin.MIMEJSON)
		return
	}
	if score < 0 {
		utils.ResSender(c, "Score is required and should be a valid positive Integer", "Invalid Params", http.StatusBadRequest, gin.MIMEJSON)
		return
	}
	player.Id = id.String()
	player.Name = name
	player.Country = country
	player.Score = score

	err := model.AddPlayer(player)
	if err != nil {
		fmt.Println(err.Error())
		utils.ResSender(c, "Internal server error", "", http.StatusInternalServerError, gin.MIMEJSON)
		return
	}
	utils.ResSender(c, "Ok", id.String(), http.StatusOK, gin.MIMEJSON)
	return
}

func HandleUpadtePlayer(c *gin.Context) {
	playerId := c.Param("id")
	name := c.PostForm("name")
	var score int
	if len(c.PostForm("score")) == 0 && len(name) == 0 {
		utils.ResSender(c, "Eiter name or score is required ", "Invalid Params", http.StatusBadRequest, gin.MIMEJSON)
		return

	}
	if len(c.PostForm("score")) > 0 {
		sc, error := strconv.Atoi(c.PostForm("score"))
		if error != nil {
			utils.ResSender(c, "Score is required and should be a valid positive Integer", "Invalid Params", http.StatusBadRequest, gin.MIMEJSON)
			return
		}
		fmt.Println(sc)
		score = sc
	} else {
		score = -1
	}
	if len(playerId) < 1 {
		utils.ResSender(c, "Player Id is required ", "Invalid played id", http.StatusBadRequest, gin.MIMEJSON)
		return
	}
	player, err := model.UpdatePlayer(playerId, name, score)
	if err != nil {
		utils.ResSender(c, "Internal server error", err.Error(), http.StatusInternalServerError, gin.MIMEJSON)
		return
	}
	utils.PlayerSender(c, "Ok", player, http.StatusOK, gin.MIMEJSON)
}
func HandleDeletePlayer(c *gin.Context) {
	playerId := c.Param("id")
	if len(playerId) < 1 {
		utils.ResSender(c, "Player Id is required ", "Invalid played id", http.StatusBadRequest, gin.MIMEJSON)
		return
	}
	model.DeletePlayer(playerId)

	utils.ResSender(c, "Ok", "", http.StatusOK, gin.MIMEJSON)
}
func HandelGetAllPlayer(c *gin.Context) {
	players, err := model.GetAllPlayers()
	if err != nil {
		utils.ResSender(c, "Interal Server Error", err.Error(), http.StatusInternalServerError, gin.MIMEJSON)
		return
	}

	utils.MultiplePlayerSender(c, "Ok", players, http.StatusOK, gin.MIMEJSON)

}
func HandleGetPlayerWithRank(c *gin.Context) {
	rank, err := strconv.Atoi(c.Param("val"))
	if err != nil {
		utils.ResSender(c, "Rank should be a valid Positive interger", err.Error(), http.StatusBadRequest, gin.MIMEJSON)
		return
	}
	if rank < 0 {
		utils.ResSender(c, "Rank should be a valid Positive interger", "Invalid Params", http.StatusBadRequest, gin.MIMEJSON)
		return
	}
	player, errr := model.GetPlayerFromRanK(rank)
	if errr != nil {
		utils.ResSender(c, "Unexpected Erorr occured ", errr.Error(), http.StatusInternalServerError, gin.MIMEJSON)
		return
	}
	utils.PlayerSender(c, "Ok", player, http.StatusOK, gin.MIMEJSON)
}

func HandelGetRandomPlayer(c *gin.Context) {

	player, err := model.GetRandomPlayer()
	if err != nil {
		utils.ResSender(c, "Internal server Error ", err.Error(), http.StatusInternalServerError, gin.MIMEJSON)
		return
	}

	utils.PlayerSender(c, "Ok", player, http.StatusOK, gin.MIMEJSON)

}
