package terminal

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type UI interface {
	Say(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Ask(prompt string, args ...interface{}) (answer string)
	AskForPassword(prompt string, args ...interface{}) (answer string)
	Confirm(message string, args ...interface{}) bool
	Failed(message string, args ...interface{})
	Wait(duration time.Duration)
	EmptyLine()
	Ok()
}

type terminalUI struct {
	stdin   io.Reader
	printer Printer
}

func NewUI(r io.Reader, printer Printer) UI {
	return &terminalUI{
		stdin:   r,
		printer: printer,
	}
}

func (ui *terminalUI) Say(message string, args ...interface{}) {
	if len(args) == 0 {
		ui.printer.Printf("%s\n", message)
	} else {
		ui.printer.Printf(message+"\n", args...)
	}
}

func (ui *terminalUI) Warn(message string, args ...interface{}) {
	message = fmt.Sprintf(message, args...)
	ui.Say(WarningColor(message))
	return
}

func (ui *terminalUI) Confirm(message string, args ...interface{}) bool {
	response := ui.Ask(message, args...)
	switch strings.ToLower(response) {
	case "y", "yes":
		return true
	}
	return false
}

func (ui *terminalUI) Ask(prompt string, args ...interface{}) (answer string) {
	// ui.EmptyLine()
	fmt.Printf(prompt+PromptColor(">")+" ", args...)

	rd := bufio.NewReader(ui.stdin)
	line, err := rd.ReadString('\n')
	if err == nil {
		return strings.TrimSpace(line)
	}
	return ""
}

func (ui *terminalUI) EmptyLine() {
	fmt.Println("")
}

func (ui *terminalUI) Ok() {
	ui.Say(SuccessColor("OK"))
}

func (ui *terminalUI) Failed(message string, args ...interface{}) {
	message = fmt.Sprintf(message, args...)
	ui.Say(FailureColor("FAILED"))
	ui.Say(message)
	os.Exit(1)
}

func (ui *terminalUI) LoadingIndication() {
	ui.printer.Print(".")
}

func (ui *terminalUI) Wait(duration time.Duration) {
	time.Sleep(duration)
}

func (ui *terminalUI) Table(headers []string) Table {
	return NewTable(ui, headers)
}
