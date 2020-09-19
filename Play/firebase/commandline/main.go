package main

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/peterh/liner"

	"cloud.google.com/go/firestore"
)

var (
	client *firestore.Client
	ctx    context.Context
	err    error

	collection string
)

func prompt() string {
	return fmt.Sprintf("<%s> ", time.Now().Format("Jan 2 15:04:05.000"))
}

func main() {

	// Initialize commands
	commandsInit()

	// Create firebase client
	//connect([]string{})

	// Start loop with history and completion
	_ = interactiveLoop()
}

func interactiveLoop() error {
	s := liner.NewLiner()
	s.SetTabCompletionStyle(liner.TabPrints)
	s.SetCompleter(func(line string) (ret []string) {
		for _, c := range commandKeys {
			if strings.HasPrefix(c, line) {
				ret = append(ret, c)
			}
		}
		return
	})
	defer func() { _ = s.Close() }()
	for {
		p, err := s.Prompt(prompt())
		if err == io.EOF {
			return nil
		}
		if err != nil {
			panic(err)
		}
		if executeCommand(p) {
			s.AppendHistory(p)
		}
	}
}
