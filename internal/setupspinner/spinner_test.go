package setupspinner

import (
	"fmt"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestStatusMsgUpdatesStatus(t *testing.T) {
	m := New("initial status")
	updated, _ := m.Update(StatusMsg("new status"))
	model := updated.(Model)

	if model.status != "new status" {
		t.Errorf("expected status 'new status', got %q", model.status)
	}
}

func TestDoneMsgSuccess(t *testing.T) {
	m := New("working...")
	updated, cmd := m.Update(DoneMsg{})
	model := updated.(Model)

	if !model.done {
		t.Error("expected done to be true")
	}
	if model.err != nil {
		t.Errorf("expected no error, got %v", model.err)
	}
	if cmd == nil {
		t.Fatal("expected quit command")
	}

	// Verify it produces a tea.QuitMsg
	msg := cmd()
	if _, ok := msg.(tea.QuitMsg); !ok {
		t.Errorf("expected tea.QuitMsg, got %T", msg)
	}
}

func TestDoneMsgWithError(t *testing.T) {
	m := New("working...")
	testErr := fmt.Errorf("setup failed")
	updated, cmd := m.Update(DoneMsg{Err: testErr})
	model := updated.(Model)

	if !model.done {
		t.Error("expected done to be true")
	}
	if model.err != testErr {
		t.Errorf("expected error %v, got %v", testErr, model.err)
	}
	if cmd == nil {
		t.Fatal("expected quit command")
	}
}

func TestViewShowsStatusWithSpinner(t *testing.T) {
	m := New("Setting up workspace...")
	view := m.View()

	if !strings.Contains(view, "Setting up workspace...") {
		t.Errorf("expected view to contain status message, got %q", view)
	}
}

func TestViewEmptyWhenDone(t *testing.T) {
	m := New("working...")
	updated, _ := m.Update(DoneMsg{})
	model := updated.(Model)
	view := model.View()

	if view != "" {
		t.Errorf("expected empty view when done, got %q", view)
	}
}

func TestResultReturnsError(t *testing.T) {
	m := New("working...")
	testErr := fmt.Errorf("some error")
	updated, _ := m.Update(DoneMsg{Err: testErr})
	model := updated.(Model)

	if model.Result() != testErr {
		t.Errorf("expected Result() to return %v, got %v", testErr, model.Result())
	}
}

func TestResultReturnsNilOnSuccess(t *testing.T) {
	m := New("working...")
	updated, _ := m.Update(DoneMsg{})
	model := updated.(Model)

	if model.Result() != nil {
		t.Errorf("expected Result() to return nil, got %v", model.Result())
	}
}

func TestCtrlCQuitsProgram(t *testing.T) {
	m := New("working...")
	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyCtrlC})

	if cmd == nil {
		t.Fatal("expected quit command on ctrl+c")
	}

	msg := cmd()
	if _, ok := msg.(tea.QuitMsg); !ok {
		t.Errorf("expected tea.QuitMsg, got %T", msg)
	}
}

func TestNewSetsInitialStatus(t *testing.T) {
	m := New("my initial status")
	if m.status != "my initial status" {
		t.Errorf("expected initial status 'my initial status', got %q", m.status)
	}
}
