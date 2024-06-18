package util

import (
	"fmt"
	"math/rand"
	"slices"
)

type Player struct {
	Position int
	Name     string
	Score    int
	Dice     []int
	Length   int
}

func NewPlayer(pos, diceLength int) *Player {
	name := fmt.Sprintf("Player %d", pos+1)

	player := Player{Score: 0, Dice: []int{}, Name: name, Length: diceLength, Position: pos}

	for j := 0; j < diceLength; j++ {
		player.Dice = append(player.Dice, rand.Intn(6)+1)
	}

	return &player
}

func (p *Player) Roll() {
	p.Dice = []int{}
	for j := 0; j < p.Length; j++ {
		p.Dice = append(p.Dice, rand.Intn(6)+1)
	}
}

/* fungsi untuk mencari angka 6 */
func (p *Player) Find6Number() {
	if !slices.Contains(p.Dice, 6) {
		return
	}

	/* menambah score jika menemukan angka 6 dan menghapus dari element */
	var temp []int
	for _, dice := range p.Dice {
		if dice == 6 {
			p.Score += 1
		} else {
			temp = append(temp, dice)
		}
	}

	p.Dice = temp
	p.Length = len(temp)
}

/* fungsi untuk mencari angka 1 */
func (p *Player) Find1Number() int {
	if !slices.Contains(p.Dice, 1) {
		return 0
	}

	/* menghapus angka 1 pada array */
	var temp []int
	count := 0
	for _, dice := range p.Dice {
		if dice == 1 {
			count++
			continue
		}
		temp = append(temp, dice)
	}

	p.Dice = temp
	p.Length = len(temp)

	return count
}

func (p *Player) Add1Number(count int) {
	for i := 0; i < count; i++ {
		p.Dice = append(p.Dice, 1)
	}

	p.Length = len(p.Dice)
}
