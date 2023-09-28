package main 

import (
	whisp "github.com/patbcole117/haunt/whispers"
	"github.com/patbcole117/haunt/ghost"	
)

func main() {
	h := whisp.NewHTTPHollow("127.0.0.1:80")
	h.Listen()

	g := ghost.NewGhost("127.0.0.1:80", "http")
	g.Haunt()
}