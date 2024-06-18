package util

import (
	"fmt"
)

type Game struct {
	Player      int
	Dice        int
	Players     map[string]*Player
	PlayerNames []string

	IsContinue bool
	Count      int
}

func NewGame(player, dice int) *Game {
	return &Game{
		Player:      player,
		Dice:        dice,
		Players:     map[string]*Player{},
		PlayerNames: []string{},

		IsContinue: true,
		Count:      1,
	}
}

func (g *Game) Play() {
	for g.IsContinue {
		g.roll()

		g.Count++
	}

	/* menampilkan data skor */
	fmt.Println("======================================")
	fmt.Println("Hasil akhir:")
	for name, player := range g.Players {
		fmt.Printf("%s: %d\n", name, player.Score)
	}
}

/* fungsi untuk membuat dadu secara acak sesuai dengan size */
func (g *Game) roll() {
	fmt.Println("======================================")
	fmt.Printf("Giliran %v lempar dadu:\n", g.Count)

	for i := 0; i < g.Player; i++ {
		playerName := fmt.Sprintf("Player %d", i+1)

		var isPlayerExist bool
		if _, ok := g.Players[playerName]; ok {
			isPlayerExist = true
		}

		/* buat player jika player belum ada */
		if !isPlayerExist {
			g.PlayerNames = append(g.PlayerNames, playerName)
			g.Players[playerName] = NewPlayer(i, g.Dice)
		}

		/* jika player sudah ada, acak dadunya */
		if isPlayerExist {
			g.Players[playerName].Roll()
		}
	}

	g.evaluate()
}

/* fungsi untuk mengevaluasi hasil dari pengacakan */
func (g *Game) evaluate() {
	g.print(true)

	foundNumber1 := map[string]int{}

	emptyDice := 0
	for pos, name := range g.PlayerNames {
		/* evaluasi dadu player */
		g.Players[name].Find6Number()

		var nextPlayer string
		if pos == g.Player-1 {
			nextPlayer = g.PlayerNames[0]
		} else {
			nextPlayer = g.PlayerNames[pos+1]
		}

		foundNumber1[nextPlayer] = g.Players[name].Find1Number()

		/* menghitung data player yang sudah tidak memiliki dadu */
		if len(g.Players[name].Dice) == 0 {
			emptyDice++
		}
	}

	/* tambah angka 1 sebanyan n pada next player */
	for name, count := range foundNumber1 {
		if count == 0 {
			continue
		}

		g.Players[name].Add1Number(count)
	}

	if emptyDice >= g.Player-1 {
		g.IsContinue = false
	}

	g.print(false)
}

/* fungsi untuk menampilkan dadu player */
func (g *Game) print(isBefore bool) {
	if isBefore {
		fmt.Println("================sebelum===============")
	} else {
		fmt.Println("================sesudah===============")

	}
	for i := 0; i < g.Player; i++ {
		playerName := fmt.Sprintf("Player %d", i+1)
		fmt.Printf("%s (%v): %v\n", playerName, g.Players[playerName].Score, g.Players[playerName].Dice)
	}
	fmt.Println("======================================")
}
