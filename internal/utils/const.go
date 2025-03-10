package utils

const JOB_STATUS_UNKNOWN = "unknown"
const JOB_STATUS_RUN = "run"
const JOB_STATUS_ERROR = "error"
const JOB_STATUS_FINISH = "finish"

const JOB_TYPE_EMBEDDED = "embedded"
const JOB_TYPE_RELATED = "related"
const JOB_TYPE_ABSOLUTE = "absolute"

type EmbeddedActions string

const (
	RebootHost     EmbeddedActions = "RebootHost"
	DockerRestart  EmbeddedActions = "DockerRestart"
	ComposeRestart EmbeddedActions = "ComposeRestart"
)
