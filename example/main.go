package main

import (
	"fmt"
	"os"

	"github.com/dixonwille/wlog"
	"github.com/gen0cide/renum/example/lib"
)

var ui wlog.UI

func init() {
	baseUI := wlog.New(os.Stdin, os.Stdout, os.Stderr)
	prefixedUI := wlog.AddPrefix("[?]", "["+wlog.Cross+"]", "[+]", "", "", "[~]", "["+wlog.Check+"]", "[!]", baseUI)
	colorUI := wlog.AddColor(wlog.BrightCyan, wlog.BrightRed, wlog.BrightWhite, wlog.BrightBlue, wlog.None, wlog.None, wlog.BrightMagenta, wlog.BrightGreen, wlog.BrightYellow, prefixedUI)
	ui = wlog.AddConcurrent(colorUI)
}

func retCode() lib.ErrorCode {
	return lib.ErrorCodeInvalidSQLQuery
}

func main() {
	val := retCode()

	ui.Info("renum.Coder interface")
	ui.Success(fmt.Sprintf("Code() = %d", val.Code()))
	fmt.Println()
	ui.Info("renum.Namespacer interface")
	ui.Success(fmt.Sprintf("Namespace() = %s", val.Namespace()))
	ui.Success(fmt.Sprintf("     Path() = %s", val.Path()))
	fmt.Println()
	ui.Info("renum.Typer interface")
	ui.Success(fmt.Sprintf("       Kind() = %s", val.Kind()))
	ui.Success(fmt.Sprintf("     Source() = %s", val.Source()))
	ui.Success(fmt.Sprintf("PackageName() = %s", val.PackageName()))
	ui.Success(fmt.Sprintf(" ImportPath() = %s", val.ImportPath()))
	fmt.Println()
	ui.Info("renum.Descriptioner interface")
	ui.Success(fmt.Sprintf("Description() = %s", val.Description()))
	fmt.Println()
	ui.Info("fmt.Stringer interface")
	ui.Success(fmt.Sprintf("String() = %s", val.String()))
	fmt.Println()
	ui.Info("error interface")
	ui.Success(fmt.Sprintf("Error() = %s", val.Error()))
}
