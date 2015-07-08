package main

import (
	// "fmt"
	// "bufio"
	"os"

	"github.com/mcastilho/terminal"
)

func main() {

	// reader := bufio.NewReader(os.Stdin)
	printer := terminal.NewPrinter()

	ui := terminal.NewUI(os.Stdin, printer)

	ui.Say("Terminal Package Test App")

	name := ui.Ask("Name ")
	pwd := ui.AskForPassword("Password ")

	ui.EmptyLine()
	confirm := ui.Confirm("Are you sure this is your password ? (type YES) ")
	if !confirm {
		ui.Failed("Wrong password! Goodbye!")
	}
	ui.EmptyLine()

	ui.Say("Hello %s, your password is %s", name, terminal.CommandColor(pwd))

	ui.Warn("Your password is weak")

	ui.Ok()
}
