package workflow

const (
	EntrypointTemplateName = "entry"

	HelmReleaseArg                 = "helmrelease"
	TimeoutArg                     = "timeout"
	HelmReleaseExecutorName        = "helmrelease-executor"
	HelmReleaseReverseExecutorName = "helmrelease-reverse-executor"

	ValuesKeyGlobal = "global"
	ChartLabelKey   = "chart"

	DefaultTimeout = "5m"

	ExecutorName     = "executor"
	ExecutorImage    = "azureorkestra/executor"
	ExecutorImageTag = "v0.4.1"

	KeptnExecutor         = "azureorkestra/keptn-executor"
	KeptnExecutorImageTag = "v0.1.0"

	ChartMuseumName = "chartmuseum"

	Project        = "orkestra"
	OwnershipLabel = "owner"
	HeritageLabel  = "heritage"
)
