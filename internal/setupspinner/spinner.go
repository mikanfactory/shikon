package setupspinner

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var accentColor = lipgloss.Color("#89b4fa")

// StatusMsg updates the displayed status text.
type StatusMsg string

// DoneMsg signals that the setup is complete.
type DoneMsg struct {
	Err error
}

// Model is a mini Bubble Tea model that shows a spinner with a status message.
type Model struct {
	spinner spinner.Model
	status  string
	done    bool
	err     error
}

// New creates a new spinner model with the given initial status message.
func New(initialStatus string) Model {
	s := spinner.New()
	s.Spinner = spinner.MiniDot
	s.Style = lipgloss.NewStyle().Foreground(accentColor)
	return Model{
		spinner: s,
		status:  initialStatus,
	}
}

func (m Model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case StatusMsg:
		m.status = string(msg)
		return m, nil
	case DoneMsg:
		m.done = true
		m.err = msg.Err
		return m, tea.Quit
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
		return m, nil
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m Model) View() string {
	if m.done {
		return ""
	}
	return "  " + m.spinner.View() + " " + m.status + "\n"
}

// Result returns the error from DoneMsg, or nil on success.
func (m Model) Result() error {
	return m.err
}
