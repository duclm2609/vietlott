package service

import (
	"crypto/rand"
	"math/big"
	math_rand "math/rand"
	"sort"
	"time"
)

const seed int64 = 173009122017

const ballsPerTicket = 6
const totalBallNumber = 45

var balls = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45}

type Generator struct {
}

func (g Generator) GenerateMega645() []int {
	res := make([]int, ballsPerTicket)
	remainBalls := make([]int, totalBallNumber)
	copy(remainBalls, balls)
	math_rand.Seed(seed)
	for i := 0; i < ballsPerTicket; i++ {
		bound := totalBallNumber - i
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(bound)))
		j := index.Uint64()
		res[i] = remainBalls[j]
		remainBalls = append(remainBalls[:j], remainBalls[j+1:]...)

		sleep := math_rand.Intn(500)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
	}
	sort.Ints(res)
	return res
}
