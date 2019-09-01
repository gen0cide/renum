package plugins

// All returns a pre-populated list of plugins written for renum's code generator.
var All = []Constructor{
	// required
	NewConstPlugin,
	NewHeaderPlugin,
	NewParsePlugin,
	NewSnakeCasePlugin,
	NewStringerPlugin,
	NewNamesPlugin,
	NewValuesPlugin,
	NewTextPlugin,

	// Caser
	NewCamelCasePlugin,
	NewCommandCasePlugin,
	NewPascalCasePlugin,
	NewScreamingCasePlugin,
	NewDottedCasePlugin,
	NewTrainCasePlugin,

	// to include a type definition
	NewDefinitionPlugin,

	// required for renum.Enum, optional for others
	NewCoderPlugin,
	NewDescriptionerPlugin,
	NewTyperPlugin,
	NewSourcerPlugin,
	NewNamespacerPlugin,

	// errors
	NewMessagesPlugin,
	NewErrorsPlugin,

	// serializers plugins
	NewSQLPlugin,
	NewJSONPlugin,
	NewFlagsPlugin,
	NewYAMLPlugin,
	NewCSVPlugin,

	// codes plugins
	NewYARPCCodesPlugin,
	NewHTTPCodesPlugin,
	NewOSExitCodesPlugin,
}
