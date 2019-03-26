package main

import (
	"fmt"
	"math/rand"
	"time"
)

func newEnemy(scale int) *entity {

	// Set Random Seed
	rand.Seed(time.Now().UnixNano())

	// Create random Stats
	rands := map[string]int{
		"sta": 0,
		"int": 0,
		"dex": 0,
		"str": 0,
	}
	for rands["sta"] <= 1 {
		rands["int"] = rand.Intn(8) + 1
		rands["dex"] = rand.Intn(8) + 1
		rands["str"] = rand.Intn(8) + 1
		rands["sta"] = (20 + scale) - rands["int"] - rands["dex"] - rands["srt"]
	}

	// Create random HP + MaxHP
	fhp := ((rand.Intn(10) + (8 * scale)) * 10)

	e := &entity{
		maxHp:  fhp,
		hp:     fhp,
		mp:     100,
		maxMp:  100,
		player: false,
		sta:    rands["sta"],
		str:    rands["str"],
		dex:    rands["dex"],
		int:    rands["int"],
		lvl:    scale,
		exp:    25 * scale,
		alive:  true,
	}

	// Set Enemy
	class := rand.Intn(5)

	switch class {
	case 0:
		e.name = "Soldier"
		e.class.name = "Humanoid"
	case 1:
		e.name = "Ghoul"
		e.class.name = "Humanoid"
	case 2:
		e.name = "Zombie"
		e.class.name = "Humanoid"
	case 3:
		e.name = "Bat"
		e.class.name = "Animal"
	case 4:
		e.name = "Wolf"
		e.class.name = "Animal"
	default:
		fmt.Print("\"", class, "\" broken random number lol")
	}

	return e
}
