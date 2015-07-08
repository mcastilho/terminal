package terminal

import (
	"fmt"
)

type Printer interface {
	Print(a ...interface{}) (n int, err error)
	Printf(format string, a ...interface{}) (n int, err error)
	Println(a ...interface{}) (n int, err error)
	ForcePrint(a ...interface{}) (n int, err error)
	ForcePrintf(format string, a ...interface{}) (n int, err error)
	ForcePrintln(a ...interface{}) (n int, err error)
	DisableOutput(bool)
}

type printer struct {
	disableOutput bool
}

func NewPrinter() *printer {
	return &printer{}
}

func (p *printer) Print(values ...interface{}) (n int, err error) {
	if p.disableOutput {
		return
	}
	return p.ForcePrint(values...)
}

func (p *printer) Printf(format string, a ...interface{}) (n int, err error) {
	if p.disableOutput {
		return
	}
	return p.ForcePrintf(format, a...)
}

func (p *printer) Println(values ...interface{}) (n int, err error) {
	if p.disableOutput {
		return
	}
	return p.ForcePrintln(values...)
}

func (p *printer) ForcePrint(values ...interface{}) (n int, err error) {
	str := fmt.Sprint(values)
	return fmt.Print(str)
}

func (p *printer) ForcePrintf(format string, a ...interface{}) (n int, err error) {
	str := fmt.Sprintf(format, a...)
	return fmt.Print(str)
}

func (p *printer) ForcePrintln(values ...interface{}) (n int, err error) {
	str := fmt.Sprint(values)
	return fmt.Println(str)
}

func (p *printer) DisableOutput(disable bool) {
	p.disableOutput = disable
}
