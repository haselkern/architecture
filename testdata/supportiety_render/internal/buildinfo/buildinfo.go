// Code generated by golangee/eearc; DO NOT EDIT.

package buildinfo

var (
	// JobID is a unique build id. Usually taken from the CI_JOB_ID
	// environment variable at build time.
	JobID string = "unknown"

	// CommitTag usually refers to the VCS branch or tag name and is probably taken from the
	// CI_COMMIT_TAG environment variable at build time.
	CommitTag string = "unknown"

	// JobStartedAt is hopefully in RFC3339 format and is likely taken from
	// the environment variable CI_JOB_STARTED_AT at build time.
	JobStartedAt string = "unknown"

	// CommitSha refers to the VCS commit hash, probably taken from the
	// CI_COMMIT_SHA environment variable at build time.
	CommitSha string = "unknown"

	// Host refers to the building host name, probably copied from the
	// CI_SERVER_HOST environment variable at build time.
	Host string = "unknown"
)

// buildInfo is a private type to allow an interface based access to the environment variables.
type buildInfo struct{}

// ID returns the current BuildID.
func (_ buildInfo) ID() string {
	return JobID
}

// Tag returns the current BuildTag.
func (_ buildInfo) Tag() string {
	return CommitTag
}

// Build provides the build information about the application.
var Build buildInfo
