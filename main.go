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
    choices: []string{"Buy Carrots", "Buy celery", "Buy kakroo"},
    selected: make(map[int]struct{}),
  }
}

func (m model) Init() tea.Cmd {
  return nil
}


