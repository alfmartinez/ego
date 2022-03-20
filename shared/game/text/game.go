package text

import (
	"bufio"
	"fmt"
	"os"
)

type textGame struct {
}

func (g *textGame) Start() {
	g.scanner()
}

func (g *textGame) Stop() {

}

func (g *textGame) scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	var stopLoop bool
	for !stopLoop && scanner.Scan() {
		input := scanner.Text()
		fmt.Println(input)
		if input == "exit" {
			stopLoop = true
		}
	}
}
