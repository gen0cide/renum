package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/dixonwille/wlog"
	"github.com/gen0cide/renum"

	"github.com/gen0cide/renum/generator"
	"gopkg.in/urfave/cli.v2"
)

var ui wlog.UI
var fileConfig generator.Config
var cliConfig generator.Config

func init() {
	baseUI := wlog.New(os.Stdin, os.Stdout, os.Stderr)
	prefixedUI := wlog.AddPrefix("[?]", "["+wlog.Cross+"]", "[+]", "", "", "[~]", "["+wlog.Check+"]", "[!]", baseUI)
	colorUI := wlog.AddColor(wlog.BrightCyan, wlog.BrightRed, wlog.BrightWhite, wlog.BrightBlue, wlog.None, wlog.None, wlog.BrightMagenta, wlog.BrightGreen, wlog.BrightYellow, prefixedUI)
	ui = wlog.AddConcurrent(colorUI)
	fc, err := generator.NewConfig()
	if err != nil {
		panic(err)
	}
	cc, err := generator.NewConfig()
	if err != nil {
		panic(err)
	}
	fileConfig = fc
	cliConfig = cc
}

var configPath string
var debug bool

var globalFlags = []cli.Flag{
	&cli.PathFlag{
		Name: "config",
		Aliases: []string{
			"c",
		},
		Usage: "Path to the YAML configuration file.",
		EnvVars: []string{
			"RENUM_CONFIG_PATH",
		},
		Destination: &configPath,
	},
	&cli.BoolFlag{
		Name: "debug",
		Aliases: []string{
			"d",
		},
		Usage: "Enable debugging output.",
		EnvVars: []string{
			"RENUM_DEBUG_ENABLED",
		},
		Destination: &debug,
	},
}

var app = &cli.App{
	Name:        "renum",
	Usage:       "Generates customizable enums for Golang.",
	Version:     renum.VersionString(),
	Description: "Renum generates idomatic const enums for Golang, based on definitions configured in a YAML file.",
	Authors: []*cli.Author{
		&cli.Author{
			Name:  "Alex Levinson",
			Email: "gen0cide.threats@gmail.com",
		},
	},
	Flags:     globalFlags,
	Copyright: "(c) 2019 Alex Levinson",
	Writer:    os.Stdout,
	ErrWriter: os.Stderr,
	Commands: []*cli.Command{
		&cli.Command{
			Name: "generate",
			Aliases: []string{
				"g",
			},
			UsageText:   "renum [global options] generate [command options]",
			Usage:       "Generate the enums based on the provided configuration file.",
			Description: "The generate command is used to perform Go code generation based on the inputs provided.",
			Action:      generate,
			Flags:       flagOverrides(),
		},
		&cli.Command{
			Name: "test",
			Aliases: []string{
				"t",
			},
			UsageText:   "renum [global options] test [command options]",
			Usage:       "Test the provided configuration to ensure the result is valid Go code.",
			Description: "The test command can be used as a precurser to generate in order to verify that the generator will produce valid Go code on your system.",
			Action:      testConfig,
			Flags:       flagOverrides(),
		},
	},
}

func handleCLIParams() {
	if outputDir != "" {
		fileConfig.OutputDir = outputDir
	}
	if goType != "" {
		fileConfig.Go.Type = goType
	}
	if goName != "" {
		fileConfig.Go.Name = goName
	}
	if goPackageName != "" {
		fileConfig.Go.PackageName = goPackageName
	}
	if goPackagePath != "" {
		fileConfig.Go.PackagePath = goPackagePath
	}
	if goFilename != "" {
		fileConfig.Go.Filename = goFilename
	}
	if enableErrors {
		fileConfig.Plugins.Error = true
	}
	if enableText {
		fileConfig.Plugins.Text = true
	}
	if enableJSON {
		fileConfig.Plugins.JSON = true
	}
	if enableYAML {
		fileConfig.Plugins.YAML = true
	}
	if enableSQL {
		fileConfig.Plugins.SQL = true
	}
	if enableFlags {
		fileConfig.Plugins.Flags = true
	}
	if enableDescriptions {
		fileConfig.Plugins.Description = true
	}
	if enableNamespaces {
		fileConfig.Plugins.Namespace.Enabled = true
	}
	if pluginNamespacesPath != "" {
		fileConfig.Plugins.Namespace.Namespace = pluginNamespacesPath
	}
	if enableCodesSimple {
		fileConfig.Plugins.Codes.Simple = true
	}
	if enableCodesYARPC {
		fileConfig.Plugins.Codes.YARPC = true
	}
	if pluginCodesYARPCDefault != 0 {
		fileConfig.Plugins.Codes.Defaults.YARPC = &pluginCodesYARPCDefault
	}
	if enableCodesHTTP {
		fileConfig.Plugins.Codes.HTTP = true
	}
	if pluginCodesHTTPDefault != 0 {
		fileConfig.Plugins.Codes.Defaults.HTTP = &pluginCodesHTTPDefault
	}
	if enableCodesOSExit {
		fileConfig.Plugins.Codes.OSExit = true
	}
	if pluginCodesOSExitDefault != 0 {
		fileConfig.Plugins.Codes.Defaults.OSExit = &pluginCodesOSExitDefault
	}
}

func generate(ctx *cli.Context) error {
	err := readConfigFile()
	if err != nil {
		return err
	}

	ui.Success("parsed configuration")

	g, err := generator.NewGenerator(&fileConfig)
	if err != nil {
		ui.Warn("Error creating generator")
		return err
	}

	err = g.Initialize()
	if err != nil {
		ui.Warn("Error initializing generator")
		return err
	}

	ui.Success("initialized generator")

	data, err := g.GenerateEnums()
	if err != nil {
		ui.Warn("Error generating code")
		return err
	}

	ui.Success("generated Go code")

	fileloc := filepath.Join(g.Config.OutputDir, g.Config.Go.OutputFilename())
	err = ioutil.WriteFile(fileloc, data, 0644)
	if err != nil {
		ui.Warn("Error writing generated code")
		return err
	}

	ui.Success(fmt.Sprintf("successfully wrote code to %s", fileloc))

	return nil
}

func testConfig(ctx *cli.Context) error {
	err := readConfigFile()
	if err != nil {
		return err
	}

	g, err := generator.NewGenerator(&fileConfig)
	if err != nil {
		ui.Warn("Error creating generator")
		return err
	}

	err = g.Initialize()
	if err != nil {
		ui.Warn("Error initializing generator")
		return err
	}

	buf := new(bytes.Buffer)
	enc := yaml.NewEncoder(buf)
	err = enc.Encode(g.Config)
	if err != nil {
		ui.Warn("Error encoding configuration")
		return err
	}

	err = enc.Close()
	if err != nil {
		ui.Warn("Error closing YAML encoding buffer")
		return err
	}

	ui.Success("Encoded config")
	ui.Output(buf.String())
	// pp.Println(fileConfig)
	return nil
}

func readConfigFile() error {
	if _, err := os.Stat(configPath); err != nil {
		ui.Warn("Error locating config file")
		return err
	}

	//nolint:gosec
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		ui.Warn("Error reading config file")
		return err
	}

	err = yaml.UnmarshalStrict(data, &fileConfig)
	if err != nil {
		ui.Warn("Error parsing config file YAML")
		return err
	}

	handleCLIParams()

	return nil
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		ui.Error(err.Error())
		os.Exit(1)
	}
}
