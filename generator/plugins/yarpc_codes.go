package plugins

import (
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

var (
	_yarpcCodes = map[int]string{
		0:  "yarpcerrors.CodeOK",
		1:  "yarpcerrors.CodeCancelled",
		2:  "yarpcerrors.CodeUnknown",
		3:  "yarpcerrors.CodeInvalidArgument",
		4:  "yarpcerrors.CodeDeadlineExceeded",
		5:  "yarpcerrors.CodeNotFound",
		6:  "yarpcerrors.CodeAlreadyExists",
		7:  "yarpcerrors.CodePermissionDenied",
		8:  "yarpcerrors.CodeResourceExhausted",
		9:  "yarpcerrors.CodeFailedPrecondition",
		10: "yarpcerrors.CodeAborted",
		11: "yarpcerrors.CodeOutOfRange",
		12: "yarpcerrors.CodeUnimplemented",
		13: "yarpcerrors.CodeInternal",
		14: "yarpcerrors.CodeUnavailable",
		15: "yarpcerrors.CodeDataLoss",
		16: "yarpcerrors.CodeUnauthenticated",
	}

	_defaultYARPCCode = _yarpcCodes[2]
)

// NewYARPCCodesPlugin creates a new renum generator plugin to support the renum.YARPCResponder interface.
func NewYARPCCodesPlugin() Plugin {
	p := &yarpcCodesPlugin{
		base: newBase("yarpc_codes", 10),
	}

	p.addFuncs(template.FuncMap{
		"mapify": p.mapify,
	})

	return p
}

type yarpcCodesPlugin struct {
	base
}

func (p *yarpcCodesPlugin) mapify(c *config.Config) (string, error) {
	return mapifyYARPC(c)
}

// Enabled implements the Plugin interface.
func (p *yarpcCodesPlugin) Enabled(c *config.Config) bool {
	return c.Plugins.Codes.YARPC
}

// Validate implements the Plugin interface.
func (p *yarpcCodesPlugin) Validate(c *config.Config) error {
	return nil
}

func mapifyYARPC(c *config.Config) (string, error) {
	return mapBuilder(c, ValueTypeYARPCError, func(e config.Element) interface{} {
		yarpcVal, ok := _yarpcCodes[e.YARPCCode()]
		if !ok {
			return _defaultYARPCCode
		}

		return yarpcVal
	})
}
