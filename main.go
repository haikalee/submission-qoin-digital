package main

import (
	"fmt"

	"daduan.com/util"
)

func main() {
	var player, dice int

	fmt.Print("Masukkan jumlah player: ")
	fmt.Scan(&player)

	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scan(&dice)

	if player < 2 {
		fmt.Println("Jumlah player minimal 2")
		return
	}

	if dice < 1 {
		fmt.Println("Jumlah dadu minimal 1")
		return
	}

	game := util.NewGame(player, dice)
	game.Play()
}
