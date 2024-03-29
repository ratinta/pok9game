package utils

import (
	"math/rand"
	"pok9game/model"
	"strconv"
	"strings"
	"time"
	// "fmt"
)

//GetPoint is calculate point on hand of player
func GetPoint(cardOnhand []string) int {
	sum := 0

	for i := 0; i < len(cardOnhand); i++ {
		removeFirstAlphabet := strings.Split(cardOnhand[i], "")
		joinNum := strings.Join(removeFirstAlphabet[1:], "")
		getNum, err := strconv.Atoi(joinNum)

		if err != nil {
			sum += 0
		}

		sum += getNum
	}

	sumStr := strconv.Itoa(sum)
	getFinaly := strings.Split(sumStr, "")
	almostFinaly := strings.Join(getFinaly[len(getFinaly)-1:len(getFinaly)], "")
	reallyFinaly, err := strconv.Atoi(almostFinaly)

	if err != nil {
		// fmt.Println(err)
		return 0
	}

	return reallyFinaly
}

//DrawCard for ...
func DrawCard(shuffledCards *[]string) string {
	getCard := *shuffledCards
	card := getCard[0]
	*shuffledCards = getCard[1:]

	return card
}

//DealOutCards for deal card each player
func DealOutCards(game *[]model.Player, shuffledCards *[]string) {
	getGame := *game
	for onHand := 0; onHand < 2; onHand++ {
		for idx := 0; idx < len(getGame); idx++ {
			card := &getGame[idx].Card
			*card = append(*card, DrawCard(shuffledCards))
		}
	}
}

//ShuffleCard ...
func ShuffleCard(allCard []string) []string {
	copyAllCard := make([]string, len(allCard))
	copy(copyAllCard, allCard)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(copyAllCard), func(i, j int) {
		copyAllCard[i], copyAllCard[j] = copyAllCard[j], copyAllCard[i]
	})

	return copyAllCard
}
