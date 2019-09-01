package main

import (
	"gopkg.in/urfave/cli.v2"
)

var (
	outputDir                string
	goType                   string
	goName                   string
	goPackageName            string
	goPackagePath            string
	goFilename               string
	enableErrors             bool
	enableText               bool
	enableJSON               bool
	enableYAML               bool
	enableSQL                bool
	enableFlags              bool
	enableDescriptions       bool
	enableCodesYARPC         bool
	pluginCodesYARPCDefault  int
	enableCodesHTTP          bool
	pluginCodesHTTPDefault   int
	enableCodesOSExit        bool
	pluginCodesOSExitDefault int
)

func flagOverrides() []cli.Flag {
	return []cli.Flag{
		&cli.PathFlag{
			Name: "output-dir",
			Aliases: []string{
				"o",
			},
			Usage: "Directory to write generated files to.",
			EnvVars: []string{
				"RENUM_OUTPUT_DIR",
			},
			Destination: &outputDir,
		},
		&cli.StringFlag{
			Name:  "go-name",
			Usage: "Name of the Go type you wish to create for the Enum's type.",
			EnvVars: []string{
				"RENUM_GO_NAME",
			},
			Destination: &goName,
		},
		&cli.StringFlag{
			Name:  "go-package-name",
			Usage: "Go package name of the resulting file.",
			EnvVars: []string{
				"RENUM_GO_PACKAGE_NAME",
			},
			Destination: &goPackageName,
		},
		&cli.StringFlag{
			Name:  "go-package-import-path",
			Usage: "Full import path of the resulting Go file's package.",
			EnvVars: []string{
				"RENUM_GO_PACKAGE_IMPORT_PATH",
			},
			Destination: &goPackagePath,
		},
		&cli.StringFlag{
			Name:  "go-filename",
			Usage: `Filename the results will be written to. Defaults to 'generated_%s.go' where %s is a snake case pluralized version of the --go-name.`,
			EnvVars: []string{
				"RENUM_GO_FILENAME",
			},
			Destination: &goFilename,
		},

		// Errors
		&cli.BoolFlag{
			Name:  "enable-errors",
			Usage: "Enable the errors plugin to generate functions for the go error interface.",
			EnvVars: []string{
				"RENUM_ENABLE_ERRORS",
			},
			Destination: &enableErrors,
		},

		// Text
		&cli.BoolFlag{
			Name:  "enable-text",
			Usage: "Enable generation of text.Marshaler and text.Unmarshaler interfaces.",
			EnvVars: []string{
				"RENUM_ENABLE_TEXT",
			},
			Destination: &enableText,
		},

		// JSON
		&cli.BoolFlag{
			Name:  "enable-json",
			Usage: "Enable generation of the json.Marshaler and json.Unmarshaler interfaces.",
			EnvVars: []string{
				"RENUM_ENABLE_JSON",
			},
			Destination: &enableJSON,
		},

		// YAML
		&cli.BoolFlag{
			Name:  "enable-yaml",
			Usage: "Enable generation of the yaml.Marshaler and yaml.Unmarshaler interfaces.",
			EnvVars: []string{
				"RENUM_ENABLE_YAML",
			},
			Destination: &enableYAML,
		},

		// SQL
		&cli.BoolFlag{
			Name:  "enable-sql",
			Usage: "Enable generation of the sql.Scanner and driver.Value interfaces (for database/sql).",
			EnvVars: []string{
				"RENUM_ENABLE_SQL",
			},
			Destination: &enableSQL,
		},

		// Flags
		&cli.BoolFlag{
			Name:  "enable-flags",
			Usage: "Enable generation of the github.com/spf13/pflag.Value interface methods.",
			EnvVars: []string{
				"RENUM_ENABLE_FLAGS",
			},
			Destination: &enableFlags,
		},

		// TODO: Implement pascal case lookup.
		// Pascal
		// &cli.BoolFlag{
		// 	Name:  "--enable-pascal-looku",
		// 	Usage: "Enable debugging output.",
		// 	EnvVars: []string{
		// 		"RENUM_ENABLE_NAMESPACES",
		// 	},
		// 	Destination: &c.Plugins.Namespace.Enabled,
		// },

		// TODO: Implement screaming case lookup.
		// Screaming
		// &cli.BoolFlag{
		// 	Name:  "--enable-namespaces",
		// 	Usage: "Enable debugging output.",
		// 	EnvVars: []string{
		// 		"RENUM_ENABLE_NAMESPACES",
		// 	},
		// 	Destination: &c.Plugins.Namespace.Enabled,
		// },

		// Description
		&cli.BoolFlag{
			Name:  "enable-descriptions",
			Usage: "Enable generation of a Description() string method that returns a longer description of the particular value.",
			EnvVars: []string{
				"RENUM_ENABLE_DESCRIPTIONS",
			},
			Destination: &enableDescriptions,
		},

		// Codes - YARPC
		&cli.BoolFlag{
			Name:  "enable-codes-yarpc",
			Usage: "Enable generation of a ToYARPC() yarpcerrors.Code method that returns a custom yarpcerrors.Code unique to that value. (or default otherwise).",
			EnvVars: []string{
				"RENUM_ENABLE_CODES_YARPC",
			},
			Destination: &enableCodesYARPC,
		},

		&cli.IntFlag{
			Name:  "plugin-codes-yarpc-default",
			Usage: "Set a default value for the YARPC error code returned by enums (can be overridden individually).",
			EnvVars: []string{
				"RENUM_PLUGIN_CODES_YARPC_DEFAULT",
			},
			Destination: &pluginCodesYARPCDefault,
		},

		// Codes - HTTP
		&cli.BoolFlag{
			Name:  "enable-codes-http",
			Usage: "Enable generation of a ToHTTP() int method that returns a custom HTTP Response status code unique to that value. (or default otherwise).",
			EnvVars: []string{
				"RENUM_ENABLE_CODES_HTTP",
			},
			Destination: &enableCodesHTTP,
		},

		&cli.IntFlag{
			Name:  "plugin-codes-http-default",
			Usage: "Set a default value for the HTTP status code returned by enums (can be overridden individually).",
			EnvVars: []string{
				"RENUM_PLUGIN_CODES_HTTP_DEFAULT",
			},
			Destination: &pluginCodesHTTPDefault,
		},

		// Codes - OS Exit
		&cli.BoolFlag{
			Name:  "enable-codes-os-exit",
			Usage: "Enable generation of a ToOSExit() int method that returns a custom os.Exit code unique to that value. (or default otherwise).",
			EnvVars: []string{
				"RENUM_ENABLE_CODES_OS_EXIT",
			},
			Destination: &enableCodesOSExit,
		},

		&cli.IntFlag{
			Name:  "plugin-codes-os-exit-default",
			Usage: "Set a default value for the OS exit code returned by enums (can be overridden individually).",
			EnvVars: []string{
				"RENUM_PLUGIN_CODES_OS_EXIT_DEFAULT",
			},
			Destination: &pluginCodesOSExitDefault,
		},

		// TODO: Implement docs
		// Docs
		// &cli.BoolFlag{
		// 	Name:  "--enable-namespaces",
		// 	Usage: "Enable debugging output.",
		// 	EnvVars: []string{
		// 		"RENUM_ENABLE_NAMESPACES",
		// 	},
		// 	Destination: &c.Plugins.Namespace.Enabled,
		// },
	}
}
