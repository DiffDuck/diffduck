package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	fmt.Println("Running diffduck")
	if len(os.Args) < 2 {
		fmt.Println("Usage: diffduck <path>")
		os.Exit(1)
	}

	arg := os.Args[1]
	if arg == "-v" || arg == "--version" {
		fmt.Println("DiffDuck v0.1.0")
		os.Exit(0)
	}

	path := filepath.Clean(arg)

	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	if fileInfo.IsDir() {
		fmt.Println("Error: ", path, "is a directory, not a file.")
		os.Exit(1)
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" && line[0:1] != "#" {
			fmt.Println("Commit message is not empty. Skipping.")
			os.Exit(0)
		}
	}

	runWorkflow()
	fmt.Println("Writing commit message to", path)
	if err := os.WriteFile(path, []byte("Hello, DiffDuck!\n"), 0644); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func runWorkflow() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialModel() model {
	return model{
		choices:  []string{"foo", "bar", "baz"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Select a choice with up/down arrows and enter to toggle:\n\n"

	for i, choice := range m.choices {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		checked := " " // not checked
		if _, ok := m.selected[i]; ok {
			checked = "x" // checked!
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}
