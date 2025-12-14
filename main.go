package main

import(
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices []string
	cursor int
	selected map[int]struct{}
}

func initialModel() model {
	return model{
		choices: []string{"Bread", "Milk", "Salt", "Sugar", "Flour"},
		selected: (make(map[int]struct{}))
	}
}

// function that takes in a model and performes a command
func(m model) Init() tea.Cmd {
	return nil
}

func(m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg := msg.(type) {
	
	// is it a key press
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
}