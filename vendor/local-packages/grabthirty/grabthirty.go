package grabthirty

import (
	"math/rand"
	"time"
)

var InfoMap map[string]GameInfo

type GameInfo struct {
	AnswerNum int
}

var EndNum int = 10

func CreateRandomNumber(endNum int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(endNum)
}

func New_game(user_ip string, n int) string {
	return "New Game"
}

/* GrabThirty */
func GrabThirty(user_ip string, argv []string) string {
	return "Running GrabThirty"
}
