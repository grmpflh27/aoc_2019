package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	DECKSIZE = 10
)

type ShuffleOpId int

const (
	INCREMENT ShuffleOpId = 1
	CUT       ShuffleOpId = 2
	NEWSTACK  ShuffleOpId = 3
)

type ShuffleOp struct {
	id    ShuffleOpId
	param int
}

func parseCommands() []ShuffleOp {
	fp, _ := os.Open("./input_22.txt")

	scnr := bufio.NewScanner(fp)

	var cmds []ShuffleOp
	var cur ShuffleOp
	for scnr.Scan() {
		line := scnr.Text()
		if strings.Contains(line, "cut") {
			parts := strings.Split(line, " ")
			param, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

			cur = ShuffleOp{CUT, param}
		} else if strings.Contains(line, "increment") {
			parts := strings.Split(line, " ")
			param, _ := strconv.Atoi(strings.TrimSpace(parts[len(parts)-1]))
			cur = ShuffleOp{INCREMENT, param}
		} else {
			if !strings.Contains(line, "deal into new stack") {
				fmt.Println("!!!!")
			}
			cur = ShuffleOp{NEWSTACK, 0}
		}
		cmds = append(cmds, cur)
	}
	return cmds

}

func shuffle(cmd ShuffleOp, deck []int) []int {
	var shuffled []int
	switch cmd.id {
	case CUT:
		if cmd.param < 0 {
			cmd.param = len(deck) + cmd.param
		}

		for i := cmd.param; i < len(deck); i++ {
			shuffled = append(shuffled, deck[i])
		}
		for i := 0; i < cmd.param; i++ {
			shuffled = append(shuffled, deck[i])
		}
	case NEWSTACK:
		// reverse
		for i := len(deck) - 1; i >= 0; i-- {
			shuffled = append(shuffled, deck[i])
		}
	case INCREMENT:
		shuffled = make([]int, len(deck))
		shuffled[0] = deck[0]
		ptr := cmd.param
		inserted := 1
		for inserted < len(deck) {
			if ptr > len(deck) {
				ptr -= len(deck)
			}
			shuffled[ptr] = deck[inserted]
			ptr += cmd.param
			inserted++
		}
	}

	return shuffled
}

func unshuffle(cmd ShuffleOp, deck []int) []int {
	fmt.Println("Unshuffing", cmd)

	var shuffled []int
	switch cmd.id {
	case CUT:
		if cmd.param < 0 {
			cmd.param *= -1
		}

		for i := len(deck) - cmd.param; i < len(deck); i++ {
			shuffled = append(shuffled, deck[i])
		}

		for i := 0; i < cmd.param; i++ {
			shuffled = append(shuffled, deck[i])
		}
	case NEWSTACK:
		// reverse
		for i := len(deck) - 1; i >= 0; i-- {
			shuffled = append(shuffled, deck[i])
		}
	case INCREMENT:
		shuffled = make([]int, len(deck))
		shuffled[0] = deck[0]
		ptr := cmd.param
		inserted := 1

		for inserted < len(deck) {
			if ptr > len(deck) {
				ptr -= len(deck)
			}
			shuffled[inserted] = deck[ptr]
			ptr += cmd.param
			inserted++
		}
	}

	return shuffled
}

func makeDeck(size int) []int {
	var deck []int

	for i := 0; i < size; i++ {
		deck = append(deck, i)
	}
	return deck
}

func revertWithIncrementPos(length int, stepSize int, pos int) int {
	if pos == 0 {
		return 0
	}
	q, r := divmod(pos, stepSize)
	for r != 0 {
		pos += length
		q, r = divmod(pos, stepSize)
	}
	return q
}

func divmod(numerator int, denominator int) (int, int) {
	quotient := numerator / denominator // integer division, decimals are truncated
	remainder := numerator % denominator
	return quotient, remainder
}

func revertNewStack(length int, pos int) int {
	if pos == length/2 {
		return pos
	}

	return (length - pos) - 1
}

func revertCut(length int, cut int, pos int) int {
	if cut < 0 {
		cut *= -1
	}

	if pos >= length-cut {
		return pos - (length - cut)
	} else {
		return pos + cut
	}

}

func main() {

	var day = 22
	fmt.Println("==========")
	fmt.Println("Day ", day)
	fmt.Println("==========")

	cmds := parseCommands()
	{
		deck := makeDeck(10007)
		for _, cmd := range cmds {
			deck = shuffle(cmd, deck)
		}

		for i, card := range deck {
			if card == 2019 {
				fmt.Println("Position of 2019", i)
			}
		}
	}
	{
		const NEWSIZE = 119315717514047
		const ITERATIONS = 101741582076661

		const TAFGETPOSITION = 2020

		// Answer 2
		// trace back steps

		// reverse apply

		pos := 2020
		iterCnt := 0

		for pos != 2020 || iterCnt == 0 {
			cmdCnt := len(cmds) - 1
			for cmdCnt >= 0 {
				cur := cmds[cmdCnt]
				switch cur.id {
				case INCREMENT:
					pos = revertWithIncrementPos(NEWSIZE, cur.param, pos)
				case CUT:
					pos = revertCut(NEWSIZE, cur.param, pos)
				case NEWSIZE:
					pos = revertNewStack(NEWSIZE, pos)
				}
				cmdCnt--
			}
			iterCnt++
			fmt.Println(iterCnt, pos)
			if pos > NEWSIZE {
				break
			}
		}

	}

}
