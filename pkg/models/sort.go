package model

import (
	"sort"
)

func SortWRTRank(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		return players[i].Score > players[j].Score
	})
	return players

}
