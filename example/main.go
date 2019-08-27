//go:generate renum -c error_code.yml generate -o lib

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

func main() {
	codes := lib.ErrorCodeValues()
	for _, x := range codes {
		print(x)
	}

	ui.Info("Case Parsing Tests")

	testcase := `invalid_sql_query`
	tcode, err := lib.ParseErrorCode(testcase)
	if err != nil {
		panic(err)
	}
	if tcode.String() != testcase {
		ui.Error("snake_case")
	} else {
		ui.Success("snake_case")
	}

	testcase = `InvalidSQLQuery`
	tcode, err = lib.ParseErrorCode(testcase)
	if err != nil {
		panic(err)
	}
	if tcode.PascalCase() != testcase {
		ui.Error("PascalCase")
	} else {
		ui.Success("PascalCase")
	}

	testcase = `invalidSQLQuery`
	tcode, err = lib.ParseErrorCode(testcase)
	if err != nil {
		panic(err)
	}
	if tcode.CamelCase() != testcase {
		ui.Error("camelCase")
	} else {
		ui.Success("camelCase")
	}

	testcase = `INVALID_SQL_QUERY`
	tcode, err = lib.ParseErrorCode(testcase)
	if err != nil {
		panic(err)
	}
	if tcode.ScreamingCase() != testcase {
		ui.Error("SCREAMING_CASE")
	} else {
		ui.Success("SCREAMING_CASE")
	}

	testcase = `invalid-sql-query`
	tcode, err = lib.ParseErrorCode(testcase)
	if err != nil {
		panic(err)
	}
	if tcode.CommandCase() != testcase {
		ui.Error("command-case")
	} else {
		ui.Success("command-case")
	}
}

func print(val lib.ErrorCode) {
	fmt.Println()
	ui.Running(fmt.Sprintf("ENUM VALUE OUTPUT - %s", val.String()))
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
	ui.Success(fmt.Sprintf("Error()   = %s", val.Error()))
	ui.Success(fmt.Sprintf("Message() = %s", val.Message()))
	fmt.Println()
	ui.Info("string casings")
	ui.Success(fmt.Sprintf("snake_case     = %s", val.String()))
	ui.Success(fmt.Sprintf("PascalCase     = %s", val.PascalCase()))
	ui.Success(fmt.Sprintf("camelCase      = %s", val.CamelCase()))
	ui.Success(fmt.Sprintf("SCREAMING_CASE = %s", val.ScreamingCase()))
	ui.Success(fmt.Sprintf("command-case   = %s", val.CommandCase()))
	fmt.Println()
}
