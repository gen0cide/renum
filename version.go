package renum

var (
	// Version describes the version of the library.
	Version = `1.0.1`

	// Build describes the git revision for this build.
	// Getting read of this.
	Build = ``
)

// VersionString returns the renum library version, using semantic versioning (https://semver.org)
// decorating the Version with the Build if a build is provided.
func VersionString() string {
	if Build == "" {
		return Version
	}

	return Version + "-" + Build
}
