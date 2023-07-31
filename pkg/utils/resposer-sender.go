package utils

import (
	"encoding/json"
	"fmt"

	model "github.com/dopeCape/player-managment/pkg/models"
	"github.com/gin-gonic/gin"
)

type Res struct {
	M string `json:"M"`
	D string `json:"D"`
}
type PlayerRes struct {
	M string       `json:"M"`
	D model.Player `json:"D"`
}
type MultiplePlayerRes struct {
	M string         `json:"M"`
	D []model.Player `json:"D"`
}

func ResSender(c *gin.Context, msg string, data string, statusCode int, resType string) {
	res := Res{M: msg, D: data}
	jsonData, err := json.Marshal(res)
	if err != nil {
	}
	fmt.Println(data)
	c.Data(statusCode, resType, jsonData)
}
func PlayerSender(c *gin.Context, msg string, data model.Player, statusCode int, resType string) {
	res := PlayerRes{M: msg, D: data}
	jsonData, err := json.Marshal(res)
	if err != nil {
	}
	fmt.Println(data)
	c.Data(statusCode, resType, jsonData)
}
func MultiplePlayerSender(c *gin.Context, msg string, data []model.Player, statusCode int, resType string) {
	res := MultiplePlayerRes{M: msg, D: data}
	jsonData, err := json.Marshal(res)
	if err != nil {
	}
	fmt.Println(data)
	c.Data(statusCode, resType, jsonData)
}
