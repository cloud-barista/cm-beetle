package cloudmodel

import "time"

// * To avoid circular dependencies, the following structs are copied from the cb-tumblebug framework.
// TODO: When the cb-tumblebug framework is updated, we should synchronize these structs.
// * Version: CB-Tumblebug v0.13.2 (commit: 2a7436583f889cc794ebf37d151362a2e684e871)
// * Synchronized: 2026-09-01

// InfraReq is struct for requirements to create Infra
type InfraReq struct {
	Name string `json:"name" validate:"required" example:"infra01"`

	// InstallMonAgent Option for CB-Dragonfly agent installation ([yes/no] default:yes)
	InstallMonAgent string `json:"installMonAgent" example:"no" default:"no" enums:"yes,no"` // yes or no

	// Label is for describing the object by keywords
	Label map[string]string `json:"label"`

	// SystemLabel is for describing the infra in a keyword (any string can be used) for special System purpose
	SystemLabel string `json:"systemLabel" example:"" default:""`

	PlacementAlgo string `json:"placementAlgo,omitempty"`
	Description   string `json:"description" example:"Made in CB-TB"`

	NodeGroups []CreateNodeGroupReq `json:"nodeGroups" validate:"required"`

	// PostCommands are sequential post-deployment command phases that bootstrap the Nodes
	PostCommands []PostCommandReq `json:"postCommands,omitempty" validate:"omitempty"`

	// PostCommandAsync runs post-deployment commands in the background
	PostCommandAsync bool `json:"postCommandAsync,omitempty" example:"false"`

	// PolicyOnPartialFailure determines how to handle Node creation failures
	// - "continue": Continue with partial Infra creation (default)
	// - "rollback": Cleanup entire Infra when any Node fails
	// - "refine": Mark failed Nodes for refinement
	PolicyOnPartialFailure string `json:"policyOnPartialFailure" example:"continue" default:"continue" enums:"continue,rollback,refine"`
}

// CreateNodeGroupReq is struct to get requirements to create a new server instance
type CreateNodeGroupReq struct {
	// NodeGroup name of Nodes. Actual Node name will be generated with -N postfix.
	Name string `json:"name" validate:"required" example:"g1-1"`

	// CspResourceId is resource identifier managed by CSP (required for option=register)
	CspResourceId string `json:"cspResourceId,omitempty" example:"i-014fa6ede6ada0b2c"`

	// NodeGroupSize is the number of Nodes to create in this NodeGroup. If > 0, nodeGroup will be generated.
	NodeGroupSize int `json:"nodeGroupSize" example:"3"`

	// Label is for describing the object by keywords
	Label map[string]string `json:"label"`

	Description string `json:"description" example:"Description"`

	ConnectionName string `json:"connectionName" validate:"required" example:"testcloud01-seoul"`
	SpecId         string `json:"specId" validate:"required"`
	// ImageType        string   `json:"imageType"`
	ImageId string `json:"imageId" validate:"required"`
	// CspImageName is the CSP-side image identifier pre-resolved by EnsureImageAvailable
	// at nodegroup level (Alibaba/Azure latest-version resolution included). When non-empty
	// and the image is not a custom image, CreateNode skips the redundant per-VM GetImage
	// DB call, significantly reducing concurrent DB load during large infra creation.
	// Custom images always go through the full GetImage path (this field stays empty for them).
	CspImageName string `json:"cspImageName,omitempty"`
	VNetId       string `json:"vNetId" validate:"required"`
	SubnetId     string `json:"subnetId" validate:"required"`
	// SubnetIds, when non-empty, spreads this NodeGroup's VMs across these subnets
	// (round-robin by VM index). SubnetId above is the primary/fallback (first subnet).
	// Populated by dynamic provisioning when DistributeSubnets is requested; empty means
	// all VMs use the single SubnetId (default behavior).
	SubnetIds        []string `json:"subnetIds,omitempty"`
	SecurityGroupIds []string `json:"securityGroupIds" validate:"required"`
	SshKeyId         string   `json:"sshKeyId" validate:"required"`
	NodeUserName     string   `json:"nodeUserName,omitempty"`
	NodeUserPassword string   `json:"nodeUserPassword,omitempty"`
	RootDiskType     string   `json:"rootDiskType,omitempty" example:"default, TYPE1, ..."` // "", "default", "TYPE1", AWS: ["standard", "gp2", "gp3"], Azure: ["PremiumSSD", "StandardSSD", "StandardHDD"], GCP: ["pd-standard", "pd-balanced", "pd-ssd", "pd-extreme"], ALIBABA: ["cloud_efficiency", "cloud", "cloud_ssd"], TENCENT: ["CLOUD_PREMIUM", "CLOUD_SSD"]
	RootDiskSize     int      `json:"rootDiskSize,omitempty" example:"50"`                  // Root disk size in GB. 0 = use CSP default.
	DataDiskIds      []string `json:"dataDiskIds"`
}

// InfraDynamicReq is struct for requirements to create Infra dynamically (with default resource option)
type InfraDynamicReq struct {
	Name string `json:"name" validate:"required" example:"infra01"`

	// PolicyOnPartialFailure determines how to handle Node creation failures
	// - "continue": Continue with partial Infra creation (default)
	// - "rollback": Cleanup entire Infra when any Node fails
	// - "refine": Mark failed Nodes for refinement
	PolicyOnPartialFailure string `json:"policyOnPartialFailure" example:"continue" default:"continue" enums:"continue,rollback,refine"`

	// InstallMonAgent Option for CB-Dragonfly agent installation ([yes/no] default:no)
	InstallMonAgent string `json:"installMonAgent" example:"no" default:"no" enums:"yes,no"` // yes or no

	// NodeGroups is array of Node requests for multi-cloud infrastructure
	// Example: Multiple Node groups across different CSPs
	// [
	//   {
	//     "name": "aws-group",
	//     "nodeGroupSize": "3",
	//     "specId": "aws+ap-northeast-2+t3.nano",
	//     "imageId": "ami-01f71f215b23ba262",
	//     "rootDiskSize": "50",
	//     "label": {"role": "worker", "csp": "aws"}
	//   },
	//   {
	//     "name": "azure-group",
	//     "nodeGroupSize": "2",
	//     "specId": "azure+koreasouth+standard_b1s",
	//     "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts:22.04.202505210",
	//     "rootDiskSize": "50",
	//     "label": {"role": "head", "csp": "azure"}
	//   },
	//   {
	//     "name": "gcp-group",
	//     "nodeGroupSize": "1",
	//     "specId": "gcp+asia-northeast3+g1-small",
	//     "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20250712",
	//     "rootDiskSize": "50",
	//     "label": {"role": "test", "csp": "gcp"}
	//   }
	// ]
	NodeGroups []CreateNodeGroupDynamicReq `json:"nodeGroups" validate:"required"`

	// PostCommands are post-deployment command phases that bootstrap the Nodes.
	// Phases run sequentially; each may target a nodeGroupId, nodeId, or labelSelector.
	// A single command set is simply one phase: [{"command": ["..."]}]
	PostCommands []PostCommandReq `json:"postCommands,omitempty"`

	// PostCommandAsync (default false) returns the response as soon as nodes are
	// provisioned and runs post-deployment commands in the background. The response
	// then carries postCommandStatus="Running" plus postCommandRequestId; observe with
	// GET /ns/{nsId}/stream/cmd/infra/{infraId}?xRequestId={postCommandRequestId}
	// or by polling GET /ns/{nsId}/infra/{infraId}.
	PostCommandAsync bool `json:"postCommandAsync,omitempty" example:"false"`

	// SystemLabel is for describing the infra in a keyword (any string can be used) for special System purpose
	SystemLabel string `json:"systemLabel" example:"" default:""`

	Description string `json:"description" example:"Made in CB-TB"`

	// Label is for describing the object by keywords
	Label map[string]string `json:"label"`

	// VNetTemplateId specifies the vNet template ID (from system namespace) to use when
	// auto-creating shared vNet resources. Propagates to all NodeGroups unless overridden
	// at the NodeGroup level. If empty, the default hard-coded CIDR behavior is used.
	VNetTemplateId string `json:"vNetTemplateId,omitempty" example:"vnet-default"`

	// SgTemplateId specifies the SecurityGroup template ID (from system namespace) to use
	// when auto-creating shared SecurityGroup resources. Propagates to all NodeGroups unless
	// overridden at the NodeGroup level. If empty, the default all-open behavior is used.
	SgTemplateId string `json:"sgTemplateId,omitempty" example:"sg-default"`
}

// CreateNodeGroupDynamicReq is struct to get requirements to create a new server instance dynamically (with default resource option)
type CreateNodeGroupDynamicReq struct {
	// NodeGroup name, actual Node name will be generated with -N postfix.
	Name string `json:"name" example:"g1"`

	// NodeGroupSize is the number of Nodes to create in this NodeGroup. If > 0, nodeGroup will be generated. Default is 1.
	NodeGroupSize int `json:"nodeGroupSize" example:"3"`

	// Label is for describing the object by keywords
	Label map[string]string `json:"label" example:"{\"role\":\"worker\",\"env\":\"test\"}"`

	Description string `json:"description" example:"Created via CB-Tumblebug"`

	// SpecId is field for id of a spec in common namespace
	SpecId string `json:"specId" validate:"required" example:"aws+ap-northeast-2+t3.nano"`
	// ImageId is field for id of a image in common namespace
	ImageId string `json:"imageId" validate:"required" example:"ami-01f71f215b23ba262"`

	RootDiskType string `json:"rootDiskType,omitempty" example:"gp3" default:"default"` // "", "default", "TYPE1", AWS: ["standard", "gp2", "gp3"], Azure: ["PremiumSSD", "StandardSSD", "StandardHDD"], GCP: ["pd-standard", "pd-balanced", "pd-ssd", "pd-extreme"], ALIBABA: ["cloud_efficiency", "cloud", "cloud_essd"], TENCENT: ["CLOUD_PREMIUM", "CLOUD_SSD"]
	RootDiskSize int    `json:"rootDiskSize,omitempty" example:"50"`                    // Root disk size in GB. 0 = use CSP default.

	// NOTE: A node user password (i.e., NodeUserPassword) is intentionally NOT accepted here. Linux nodes are
	// accessed via SSH key pairs, and a random password is generated internally for the
	// CSP-side requirement (Windows). See CreateNode in core/infra/provisioning.go.

	// if ConnectionName is given, the Node tries to use associated credential.
	// if not, it will use predefined ConnectionName in Spec objects
	ConnectionName string `json:"connectionName,omitempty" example:"aws-ap-northeast-2" default:""`
	// Zone is an optional field to specify the availability zone for Node placement.
	// If specified, subnet will be created in this zone for resources like GPU Nodes
	// that may only be available in specific zones. If empty, auto-selection applies.
	Zone string `json:"zone,omitempty" example:"ap-northeast-2a" default:""`

	// VNetTemplateId overrides the Infra-level VNetTemplateId for this NodeGroup.
	// If empty, inherits the VNetTemplateId from the parent InfraDynamicReq.
	VNetTemplateId string `json:"vNetTemplateId,omitempty" example:""`

	// SgTemplateId overrides the Infra-level SgTemplateId for this NodeGroup.
	// If empty, inherits the SgTemplateId from the parent InfraDynamicReq.
	SgTemplateId string `json:"sgTemplateId,omitempty" example:""`

	// DistributeSubnets, when true, spreads this NodeGroup's VMs across the VNet's subnets
	// (round-robin), which spreads them across availability zones for multi-zone VNets.
	// Best-effort: subnets whose zone lacks the requested spec are excluded so VMs consolidate
	// to zones that have it. Ignored when Zone is set (that pins a single subnet) or when the
	// VNet has a single subnet. Default false (all VMs land in the first subnet).
	DistributeSubnets bool `json:"distributeSubnets,omitempty" example:"false"`
}

// InfraCmdReq is struct for remote command
type InfraCmdReq struct {
	// UserName is the SSH username to use for command execution
	UserName string `json:"userName" example:"cb-user" default:""`

	// Command is the list of commands to execute
	Command []string `json:"command" validate:"required" example:"client_ip=$(echo $SSH_CLIENT | awk '{print $1}'); echo SSH client IP is: $client_ip"`

	// TimeoutMinutes is the timeout for command execution in minutes (default: 30, min: 1, max: 120)
	// If not specified or set to 0, the default timeout (30 minutes) will be used
	TimeoutMinutes int `json:"timeoutMinutes,omitempty" example:"30" default:"30"`
}

// PostCommandReq is struct for post-deployment command with targeting
type PostCommandReq struct {
	InfraCmdReq

	// NodeGroupId limits execution to one nodeGroup
	NodeGroupId string `json:"nodeGroupId,omitempty" example:"g1"`
	// NodeId limits execution to a single node
	NodeId string `json:"nodeId,omitempty" example:"g1-1"`
	// LabelSelector limits execution to nodes matching the selector (e.g. "role=worker")
	LabelSelector string `json:"labelSelector,omitempty" example:"role=worker"`
	// ContinueOnError keeps running the remaining phases when this phase fails (default: false)
	ContinueOnError bool `json:"continueOnError,omitempty" example:"false"`
}

// AddNodeGroupDynamicReq is the request body for adding a NodeGroup to an existing Infra.
// It is CreateNodeGroupDynamicReq plus bootstrap fields; those fields are intentionally
// absent from the nodeGroups[] elements of InfraDynamicReq, where Infra-level
// postCommands (with per-phase targeting) is the single way to bootstrap.
type AddNodeGroupDynamicReq struct {
	CreateNodeGroupDynamicReq

	// PostCommands are sequential bootstrap phases for the newly added nodes.
	// Phases without an explicit target are scoped to this NodeGroup.
	PostCommands []PostCommandReq `json:"postCommands,omitempty" validate:"omitempty"`

	// PostCommandAsync returns the response as soon as the nodes are provisioned and
	// runs the bootstrap commands in the background (observe via streaming/polling)
	PostCommandAsync bool `json:"postCommandAsync,omitempty" example:"false"`
}

// PostCommandStatus summarizes a post-deployment command run across target nodes
type PostCommandStatus string

const (
	// PostCommandStatusNone indicates no post-deployment command was requested
	PostCommandStatusNone PostCommandStatus = "None"
	// PostCommandStatusCompleted indicates all target nodes succeeded
	PostCommandStatusCompleted PostCommandStatus = "Completed"
	// PostCommandStatusCompletedWithErrors indicates some target nodes failed
	PostCommandStatusCompletedWithErrors PostCommandStatus = "CompletedWithErrors"
	// PostCommandStatusFailed indicates all target nodes failed (or execution could not start)
	PostCommandStatusFailed PostCommandStatus = "Failed"
	// PostCommandStatusSkipped indicates the phase did not run (a previous phase failed)
	PostCommandStatusSkipped PostCommandStatus = "Skipped"
	// PostCommandStatusRunning indicates post-deployment commands are still executing
	// (async mode: the creation response returns before they finish)
	PostCommandStatusRunning PostCommandStatus = "Running"
)

// SshCmdResultForAPI is struct for REST API response with error as string
type SshCmdResultForAPI struct {
	InfraId string         `json:"infraId"`
	NodeId  string         `json:"nodeId"`
	NodeIp  string         `json:"nodeIp"`
	Command map[int]string `json:"command"`
	Stdout  map[int]string `json:"stdout"`
	Stderr  map[int]string `json:"stderr"`
	Error   string         `json:"error"` // String representation of error for JSON serialization
}

// InfraSshCmdResultForAPI is struct for multiple SSH command results (API response)
type InfraSshCmdResultForAPI struct {
	Results []SshCmdResultForAPI `json:"results"`
}

// PostCommandPhaseResult is the outcome of a single post-command phase
type PostCommandPhaseResult struct {
	// Phase is the 1-based execution order
	Phase int `json:"phase" example:"1"`
	// Target echoes the scope this phase ran against
	Target string `json:"target" example:"nodeGroupId=control"`
	// Status is the aggregated outcome of this phase (Skipped when a previous phase stopped execution)
	Status PostCommandStatus `json:"status" example:"Completed"`
	// Results holds per-node command results
	Results InfraSshCmdResultForAPI `json:"results"`
}

// ResourcePruneResult represents the outcome of pruning a single resource's orphaned metadata.
type ResourcePruneResult struct {
	// Resource type
	ResourceType string `json:"resourceType" example:"objectStorage"`
	// Resource ID
	ResourceId string `json:"resourceId" example:"os01"`
	// Connection name
	ConnectionName string `json:"connectionName" example:"aws-ap-northeast-2"`
	// Whether the prune operation was successful
	Success bool `json:"success" example:"true"`
	// Descriptive message about the prune outcome
	Message string `json:"message,omitempty" example:"Orphaned metadata for ObjectStorage (os01) pruned successfully"`
	// Error detail if prune failed
	Error string `json:"error,omitempty" example:""`
}

// ResourcePruneResults represents the aggregated results of a batch resource prune operation.
type ResourcePruneResults struct {
	// Total number of pruned resources
	TotalPruned int `json:"totalPruned" example:"2"`
	// Number of successfully pruned resources
	SuccessCount int `json:"successCount" example:"2"`
	// Number of failed prune attempts
	FailedCount int `json:"failedCount" example:"0"`
	// Individual prune results per resource
	Results []ResourcePruneResult `json:"results"`
}

// CommandExecutionStatus represents the status of command execution
type CommandExecutionStatus string

const (
	// CommandStatusQueued indicates the command has been requested but not started
	CommandStatusQueued CommandExecutionStatus = "Queued"

	// CommandStatusHandling indicates the command is currently being processed
	CommandStatusHandling CommandExecutionStatus = "Handling"

	// CommandStatusCompleted indicates the command execution completed successfully
	CommandStatusCompleted CommandExecutionStatus = "Completed"

	// CommandStatusCompletedWithError indicates that the SSH transport
	// succeeded and the remote command ran to completion, but the command
	// itself exited with a non-zero status. Typical causes: the user's
	// script reported an error, the remote process was killed by the kernel
	// (e.g. OOM-killer), or any non-zero exit from the program. Distinct
	// from CommandStatusFailed (which signals an SSH/transport-level
	// failure — bastion auth, dial timeout, mid-session EOF, etc.).
	CommandStatusCompletedWithError CommandExecutionStatus = "CompletedWithError"

	// CommandStatusFailed indicates the command execution failed
	CommandStatusFailed CommandExecutionStatus = "Failed"

	// CommandStatusTimeout indicates the command execution timed out
	CommandStatusTimeout CommandExecutionStatus = "Timeout"

	// CommandStatusCancelled indicates the command was cancelled by user request
	CommandStatusCancelled CommandExecutionStatus = "Cancelled"

	// CommandStatusInterrupted indicates the command was interrupted (e.g., system restart)
	CommandStatusInterrupted CommandExecutionStatus = "Interrupted"
)

// CommandStatusInfo represents a single remote command execution record
type CommandStatusInfo struct {
	// Index is sequential identifier for this command execution (1, 2, 3, ...)
	Index int `json:"index" example:"1"`

	// XRequestId is the request ID from X-Request-ID header when the command was executed
	XRequestId string `json:"xRequestId,omitempty" example:"req-12345678-abcd-1234-efgh-123456789012"`

	// CommandRequested is the original command as requested by the user
	CommandRequested string `json:"commandRequested" example:"ls -la"`

	// CommandExecuted is the actual SSH command executed on the Node (may be adjusted)
	CommandExecuted string `json:"commandExecuted" example:"ls -la"`

	// Status represents the current status of the command execution
	Status CommandExecutionStatus `json:"status" example:"Completed"`

	// StartedTime is when the command execution started
	StartedTime string `json:"startedTime" example:"2024-01-15 10:30:00" default:""`

	// CompletedTime is when the command execution completed (success or failure)
	CompletedTime string `json:"completedTime,omitempty" example:"2024-01-15 10:30:05"`

	// ElapsedTime is the duration of command execution in seconds
	ElapsedTime int64 `json:"elapsedTime,omitempty" example:"120"`

	// ResultSummary provides a brief summary of the execution result
	ResultSummary string `json:"resultSummary,omitempty" example:"Command executed successfully"`

	// ErrorMessage contains error details if the execution failed
	ErrorMessage string `json:"errorMessage,omitempty" example:"SSH connection failed"`

	// Stdout contains the standard output from command execution (truncated for history)
	Stdout string `json:"stdout,omitempty" example:"total 8\ndrwxr-xr-x 2 user user 4096 Jan 15 10:30 ."`

	// Stderr contains the standard error from command execution (truncated for history)
	Stderr string `json:"stderr,omitempty" example:""`

	// RepeatCount is the number of times this exact command produced this exact
	// outcome (same CommandRequested, Status, ResultSummary, and ErrorMessage) on
	// consecutive attempts. Absent/0 means it has not repeated. Repeats are merged
	// into a single record instead of appended, so retry storms (e.g. a failing
	// install script retried repeatedly) do not grow this VM's history unbounded.
	RepeatCount int `json:"repeatCount,omitempty" example:"3"`

	// LastOccurredTime is when the most recent repeat of this outcome happened.
	// Only set once RepeatCount is greater than 0; StartedTime/CompletedTime keep
	// referring to the first occurrence.
	LastOccurredTime string `json:"lastOccurredTime,omitempty" example:"2024-01-15T10:35:00Z"`
}

// InfraInfo is struct for Infra info
type InfraInfo struct {
	// ResourceType is the type of the resource
	ResourceType string `json:"resourceType"`

	// Id is unique identifier for the object
	Id string `json:"id" example:"aws-ap-southeast-1"`
	// Uid is universally unique identifier for the object, used for labelSelector
	Uid string `json:"uid,omitempty" example:"wef12awefadf1221edcf"`

	// Name is human-readable string to represent the object
	Name string `json:"name" example:"aws-ap-southeast-1"`

	Status       string          `json:"status"`
	StatusCount  StatusCountInfo `json:"statusCount"`
	TargetStatus string          `json:"targetStatus"`
	TargetAction string          `json:"targetAction"`

	// InstallMonAgent Option for CB-Dragonfly agent installation ([yes/no] default:no)
	InstallMonAgent string `json:"installMonAgent" example:"no" default:"no" enums:"yes,no"` // yes or no

	// ConfigureCloudAdaptiveNetwork is an option to configure Cloud Adaptive Network (CLADNet) ([yes/no] default:yes)
	ConfigureCloudAdaptiveNetwork string `json:"configureCloudAdaptiveNetwork" example:"yes" default:"no" enums:"yes,no"` // yes or no

	// Label is for describing the object by keywords
	Label map[string]string `json:"label"`

	// SystemLabel is for describing the infra in a keyword (any string can be used) for special System purpose
	SystemLabel string `json:"systemLabel" example:"Managed by CB-Tumblebug" default:""`

	// Latest system message such as error message
	SystemMessage []string `json:"systemMessage"` // systeam-given string message

	PlacementAlgo string     `json:"placementAlgo,omitempty"`
	Description   string     `json:"description"`
	Node          []NodeInfo `json:"node"`

	// Cluster is the list of implicit clusters synthesized at query-time from Nodes.
	Cluster []InfraClusterInfo `json:"cluster,omitempty"`

	// List of IDs for new nodes. Return IDs if the nodes are newly added. This field should be used for return body only.
	NewNodeList []string `json:"newNodeList"`

	// PostCommands are the requested post-deployment command phases
	PostCommands []PostCommandReq `json:"postCommands,omitempty"`

	// PostCommandAsync echoes whether the commands run in the background
	PostCommandAsync bool `json:"postCommandAsync,omitempty"`

	// PostCommandResults holds per-phase outcomes
	PostCommandResults []PostCommandPhaseResult `json:"postCommandResults,omitempty"`

	// PostCommandStatus summarizes the post-deployment command outcome.
	// "Running" means execution is still in progress (async mode): stream it with
	// GET /ns/{nsId}/stream/cmd/infra/{infraId}?xRequestId={postCommandRequestId}
	// or poll this object until the status becomes terminal.
	PostCommandStatus PostCommandStatus `json:"postCommandStatus,omitempty" example:"Completed"`

	// PostCommandRequestId is the streaming/tracking key of the post-deployment run
	// (always set when post-deployment commands were requested, in both modes)
	PostCommandRequestId string `json:"postCommandRequestId,omitempty" example:"pc-infra01-1a2b3c"`

	// CreationErrors contains information about Node creation failures (if any)
	CreationErrors *InfraCreationErrors `json:"creationErrors,omitempty"`
}

// InfraClusterInfo is a lightweight, on-demand cluster view synthesized from Infra NodeGroups and Nodes.
// A cluster is implicitly formed by NodeGroups that share the same network boundary (currently VNet-centric grouping).
type InfraClusterInfo struct {
	// Id is a deterministic cluster identifier generated from grouping attributes.
	Id string `json:"id"`

	// Name is a human-readable cluster name. Currently same as Id.
	Name string `json:"name"`

	// InfraId is the parent Infra ID.
	InfraId string `json:"infraId"`

	// VNetId is the shared VNet boundary used for implicit clustering.
	VNetId string `json:"vNetId,omitempty"`

	// ConnectionNames are unique connection names included in this cluster.
	ConnectionNames []string `json:"connectionNames"`

	// ProviderNames are unique CSP providers included in this cluster.
	ProviderNames []string `json:"providerNames"`

	// RegionNames are unique regions included in this cluster.
	RegionNames []string `json:"regionNames"`

	// NodeGroupIds are NodeGroups that belong to this implicit cluster.
	NodeGroupIds []string `json:"nodeGroupIds"`

	// NodeIds are Nodes that belong to this implicit cluster.
	NodeIds []string `json:"nodeIds"`

	// NodeGroupCount is the number of NodeGroups in this cluster.
	NodeGroupCount int `json:"nodeGroupCount"`

	// NodeCount is the number of Nodes in this cluster.
	NodeCount int `json:"nodeCount"`

	// RepresentativeNodeGroupId is a representative NodeGroup ID for quick inspection.
	RepresentativeNodeGroupId string `json:"representativeNodeGroupId,omitempty"`

	// RepresentativeNodeId is a representative Node ID for quick inspection.
	RepresentativeNodeId string `json:"representativeNodeId,omitempty"`
}

// InfraCreationErrors represents errors that occurred during Infra creation
type InfraCreationErrors struct {
	// NodeObjectCreationErrors contains errors from Node object creation phase
	NodeObjectCreationErrors []NodeCreationError `json:"nodeObjectCreationErrors,omitempty"`

	// NodeCreationErrors contains errors from actual Node creation phase
	NodeCreationErrors []NodeCreationError `json:"nodeCreationErrors,omitempty"`

	// TotalNodeCount is the total number of Nodes that were supposed to be created
	TotalNodeCount int `json:"totalNodeCount"`

	// SuccessfulNodeCount is the number of Nodes that were successfully created
	SuccessfulNodeCount int `json:"successfulNodeCount"`

	// FailedNodeCount is the number of Nodes that failed to be created
	FailedNodeCount int `json:"failedNodeCount"`

	// FailureHandlingStrategy indicates how failures were handled
	FailureHandlingStrategy string `json:"failureHandlingStrategy,omitempty"` // "rollback", "refine", "continue"
}

// NodeCreationError represents a single Node creation error
type NodeCreationError struct {
	// NodeName is the name of the Node that failed
	NodeName string `json:"nodeName"`

	// Error is the error message
	Error string `json:"error"`

	// Phase indicates when the error occurred
	Phase string `json:"phase"` // "object_creation", "vm_creation"

	// Timestamp when the error occurred
	Timestamp string `json:"timestamp"`
}

// StatusCountInfo is struct to count the number of Nodes in each status. ex: Running=4, Suspended=8.
type StatusCountInfo struct {

	// CountTotal is for Total Nodes
	CountTotal int `json:"countTotal"`

	// CountCreating is for counting Creating
	CountCreating int `json:"countCreating"`

	// CountRunning is for counting Running
	CountRunning int `json:"countRunning"`

	// CountFailed is for counting Failed
	CountFailed int `json:"countFailed"`

	// CountSuspended is for counting Suspended
	CountSuspended int `json:"countSuspended"`

	// CountRebooting is for counting Rebooting
	CountRebooting int `json:"countRebooting"`

	// CountTerminated is for counting Terminated
	CountTerminated int `json:"countTerminated"`

	// CountSuspending is for counting Suspending
	CountSuspending int `json:"countSuspending"`

	// CountResuming is for counting Resuming
	CountResuming int `json:"countResuming"`

	// CountTerminating is for counting Terminating
	CountTerminating int `json:"countTerminating"`

	// CountRegistering is for counting Registering
	CountRegistering int `json:"countRegistering"`

	// CountReconciling is for counting Reconciling
	CountReconciling int `json:"countReconciling"`

	// CountUndefined is for counting Undefined
	CountUndefined int `json:"countUndefined"`
}

// NodeInfoInNs is a Node together with the Infra it belongs to. Nodes are stored under
// their Infra, so a namespace-wide listing has to carry the parent Infra id for the
// caller to map each Node back to it.
type NodeInfoInNs struct {
	// InfraId is the Infra this Node belongs to
	InfraId string `json:"infraId" example:"infra01"`
	NodeInfo
}

type NodeInfo struct {
	// ResourceType is the type of the resource
	ResourceType string `json:"resourceType"`

	// Id is unique identifier for the object
	Id string `json:"id" example:"aws-ap-southeast-1"`
	// Uid is universally unique identifier for the object, used for labelSelector
	Uid string `json:"uid,omitempty" example:"wef12awefadf1221edcf"`
	// CspResourceName is name assigned to the CSP resource. This name is internally used to handle the resource.
	CspResourceName string `json:"cspResourceName,omitempty" example:"we12fawefadf1221edcf"`
	// CspResourceId is resource identifier managed by CSP
	CspResourceId string `json:"cspResourceId,omitempty" example:"csp-06eb41e14121c550a"`

	// Name is human-readable string to represent the object
	Name string `json:"name" example:"aws-ap-southeast-1"`

	// defined if the Node is in a group
	NodeGroupId string `json:"nodeGroupId"`

	Location Location `json:"location"`

	// Required by CB-Tumblebug
	Status       string `json:"status"`
	TargetStatus string `json:"targetStatus"`
	TargetAction string `json:"targetAction"`

	// Montoring agent status
	MonAgentStatus string `json:"monAgentStatus" example:"[installed, notInstalled, failed]"` // yes or no// installed, notInstalled, failed

	// NetworkAgent status
	NetworkAgentStatus string `json:"networkAgentStatus" example:"[notInstalled, installing, installed, failed]"` // notInstalled, installing, installed, failed

	// Latest system message such as error message
	SystemMessage string `json:"systemMessage" example:"Failed because ..." default:""` // systeam-given string message

	// Failure is the structured form of a creation failure: which class of CSP
	// rejection it was, the zone that was attempted, and whether a retry can
	// plausibly succeed. SystemMessage carries the same failure as one line of
	// text for display; this field is what retry logic reads.
	Failure *ProvisioningFailure `json:"failure,omitempty"`

	// Created time
	CreatedTime string `json:"createdTime" example:"2022-11-10 23:00:00" default:""`

	Label       map[string]string `json:"label"`
	Description string            `json:"description"`

	Region         RegionInfo `json:"region"` // AWS, ex) {us-east1, us-east1-c} or {ap-northeast-2}
	PublicIP       string     `json:"publicIP"`
	SSHPort        int        `json:"sshPort"`
	PublicDNS      string     `json:"publicDNS"`
	PrivateIP      string     `json:"privateIP"`
	PrivateDNS     string     `json:"privateDNS"`
	RootDiskType   string     `json:"rootDiskType"`
	RootDiskSize   int        `json:"rootDiskSize"`
	RootDeviceName string     `json:"RootDeviceName"`

	ConnectionName   string       `json:"connectionName"`
	ConnectionConfig ConnConfig   `json:"connectionConfig"`
	SpecId           string       `json:"specId"`
	CspSpecName      string       `json:"cspSpecName"`
	Spec             SpecSummary  `json:"spec"`
	ImageId          string       `json:"imageId"`
	CspImageName     string       `json:"cspImageName"`
	Image            ImageSummary `json:"image"`
	VNetId           string       `json:"vNetId"`
	CspVNetId        string       `json:"cspVNetId"`
	SubnetId         string       `json:"subnetId"`
	CspSubnetId      string       `json:"cspSubnetId"`
	NetworkInterface string       `json:"networkInterface"`
	SecurityGroupIds []string     `json:"securityGroupIds"`
	DataDiskIds      []string     `json:"dataDiskIds"`
	SshKeyId         string       `json:"sshKeyId"`
	CspSshKeyId      string       `json:"cspSshKeyId"`
	NodeUserName     string       `json:"nodeUserName,omitempty"`
	NodeUserPassword string       `json:"nodeUserPassword,omitempty"`

	// SshHostKeyInfo contains SSH host key information for TOFU (Trust On First Use) verification
	SshHostKeyInfo *SshHostKeyInfo `json:"sshHostKeyInfo,omitempty"`

	// CommandStatus stores the status and history of remote commands executed on this Node
	CommandStatus []CommandStatusInfo `json:"commandStatus,omitempty"`

	AddtionalDetails []KeyValue `json:"addtionalDetails,omitempty"`
}

// FailureClass groups CSP provisioning rejections by what a caller should do
// about them. The CSP-native wording differs per provider (AWS
// "InsufficientInstanceCapacity", GCP "STOCKOUT", Alibaba
// "OperationDenied.NoStock"), so the class — not the message — is what
// retry logic branches on.
const (
	// FailureZoneCapacity is a transient shortage in one zone. Another zone in
	// the same region may succeed right now.
	FailureZoneCapacity = "ZoneCapacity"
	// FailureRegionCapacity is a shortage across the whole region.
	FailureRegionCapacity = "RegionCapacity"
	// FailureAccountQuota is an account limit. Every zone in the region hits it.
	FailureAccountQuota = "AccountQuota"
	// FailureThrottling is CSP API rate limiting; the request itself was fine.
	FailureThrottling = "Throttling"
	// FailureDiskTypeUnavailable is a root-disk category the spec or zone rejects.
	FailureDiskTypeUnavailable = "DiskTypeUnavailable"
	// FailureImageSpecMismatch is an image incompatible with the spec, or an
	// image the CSP cannot find.
	FailureImageSpecMismatch = "ImageSpecMismatch"
	// FailureInvalidRequest is a request the CSP rejected on its own terms —
	// a disk too small, a parameter out of range. Retrying it unchanged fails
	// the same way; the request has to be corrected.
	FailureInvalidRequest = "InvalidRequest"
	// FailureAuth is a credential or permission rejection.
	FailureAuth = "Auth"
	// FailureNetwork is a transport failure reaching the CSP or CB-Spider.
	FailureNetwork = "Network"
	// FailureUnknown is anything not recognized. Callers must still handle it:
	// most providers have no dedicated parser yet.
	FailureUnknown = "Unknown"
)

// RetryHint is the recommended next move for a ProvisioningFailure.
const (
	RetryHintDifferentZone   = "differentZone"
	RetryHintDifferentSpec   = "differentSpec"
	RetryHintDifferentRegion = "differentRegion"
	RetryHintWaitAndRetry    = "waitAndRetry"
	RetryHintSameConfig      = "sameConfig"
	// RetryHintDifferentImage means the spec is fine but this image is not:
	// the CSP cannot find it, or the spec accepts only certain images.
	RetryHintDifferentImage = "differentImage"
	// RetryHintAdjustRequest means a request field has to change — a root disk
	// too small for the flavor, a parameter out of range.
	RetryHintAdjustRequest = "adjustRequest"
	// RetryHintNotRetryable means nothing in this request can be changed to make
	// it succeed: the account lacks quota or permission.
	RetryHintNotRetryable = "notRetryable"
)

// ProvisioningFailure is the structured form of one node-creation failure.
// It replaces the free-form error strings that used to be copied into
// NodeInfo.SystemMessage, NodeCreationError.Error, ProvisioningLog.FailureMessages
// and ProvisioningAttempt.FailureReason, so failures can be aggregated by zone
// and acted on programmatically.
//
// Produced in exactly one place: csp.ClassifyProvisioningFailure.
type ProvisioningFailure struct {
	// Class is one of the Failure* constants.
	Class string `json:"class" example:"ZoneCapacity"`
	// Retryable reports whether retrying can plausibly succeed without the
	// user changing something (quota increase, different spec).
	Retryable bool `json:"retryable" example:"true"`
	// RetryHint is one of the RetryHint* constants.
	RetryHint string `json:"retryHint,omitempty" example:"differentZone"`

	Provider string `json:"provider,omitempty" example:"aws"`
	Region   string `json:"region,omitempty" example:"us-west-2"`
	// CspErrorCode is the provider's own error identifier when one could be
	// extracted (e.g. "InsufficientInstanceCapacity", "STOCKOUT").
	CspErrorCode string `json:"cspErrorCode,omitempty" example:"InsufficientInstanceCapacity"`

	// AttemptedZone is the zone CB-Tumblebug actually requested, recorded from
	// the request itself. Authoritative: several CSPs (e.g. Alibaba) never name
	// the zone in their error text.
	AttemptedZone string `json:"attemptedZone,omitempty" example:"us-west-2a"`
	// ReportedZone is the zone parsed out of the CSP message, when it names one.
	// A mismatch with AttemptedZone means the VM was placed somewhere other than
	// where CB-Tumblebug intended — worth surfacing rather than hiding.
	ReportedZone string `json:"reportedZone,omitempty" example:"us-west-2a"`
	// SuggestedZones are alternatives the CSP itself offered. Only AWS provides
	// these today; empty for every other provider.
	SuggestedZones []string `json:"suggestedZones,omitempty" example:"us-west-2b,us-west-2c"`

	// Message is a single-line human-readable summary.
	Message string `json:"message,omitempty"`
	// RawMessage is the original CSP text after secret redaction, noise
	// trimming and length capping. Kept so an unrecognized failure can still
	// be diagnosed by a human.
	RawMessage string `json:"rawMessage,omitempty"`

	Source     string    `json:"source,omitempty" example:"cb-spider:1024/spider/vm"`
	HttpStatus int       `json:"httpStatus,omitempty" example:"500"`
	RequestId  string    `json:"requestId,omitempty"`
	OccurredAt time.Time `json:"occurredAt,omitempty"`
}

// ZoneCapability answers whether moving a node to another zone is even
// possible for a given connection. Two independent gates must both pass:
// the CSP driver must support zone-based control (CB-Spider's
// ZoneBasedControl capability), and the region must actually have more than
// one zone (Azure has 10 regions with none at all).
type ZoneCapability struct {
	// ZoneControl mirrors CB-Spider's per-driver ZoneBasedControl capability.
	ZoneControl bool `json:"zoneControl"`
	// Zones is the region's zone list from the connection config.
	Zones []string `json:"zones,omitempty"`
	// Shiftable is ZoneControl && len(Zones) >= 2.
	Shiftable bool `json:"shiftable"`
	// Reason explains a false Shiftable in user-facing terms.
	Reason string `json:"reason,omitempty"`
}

// NodeSummary is the per-Node projection returned by the Infra list view
// (GET /ns/{nsId}/infra). It is served from the NodeStatusAgent's in-memory
// store with no per-Node KV read, so it carries live status plus immutable
// config only. Heavy, derived, or sensitive fields — networkAgentStatus,
// commandStatus, sshHostKeyInfo, nodeUserName/Password, addtionalDetails — are
// intentionally omitted; fetch the complete NodeInfo via the single-Node API
// GET /ns/{nsId}/infra/{infraId}/node/{nodeId}.
type NodeSummary struct {
	ResourceType     string            `json:"resourceType"`
	Id               string            `json:"id"`
	Uid              string            `json:"uid,omitempty"`
	Label            map[string]string `json:"label"`
	CspResourceName  string            `json:"cspResourceName,omitempty"`
	CspResourceId    string            `json:"cspResourceId,omitempty"`
	Name             string            `json:"name"`
	NodeGroupId      string            `json:"nodeGroupId"`
	Location         Location          `json:"location"`
	Status           string            `json:"status"`
	TargetStatus     string            `json:"targetStatus"`
	TargetAction     string            `json:"targetAction"`
	MonAgentStatus   string            `json:"monAgentStatus"`
	SystemMessage    string            `json:"systemMessage"`
	CreatedTime      string            `json:"createdTime"`
	Region           RegionInfo        `json:"region"`
	PublicIP         string            `json:"publicIP"`
	SSHPort          int               `json:"sshPort"`
	PublicDNS        string            `json:"publicDNS"`
	PrivateIP        string            `json:"privateIP"`
	PrivateDNS       string            `json:"privateDNS"`
	RootDiskType     string            `json:"rootDiskType"`
	RootDiskSize     int               `json:"rootDiskSize"`
	ConnectionName   string            `json:"connectionName"`
	ConnectionConfig ConnConfig        `json:"connectionConfig"`
	SpecId           string            `json:"specId"`
	CspSpecName      string            `json:"cspSpecName"`
	Spec             SpecSummary       `json:"spec"`
	ImageId          string            `json:"imageId"`
	CspImageName     string            `json:"cspImageName"`
	Image            ImageSummary      `json:"image"`
	VNetId           string            `json:"vNetId"`
	CspVNetId        string            `json:"cspVNetId"`
	SubnetId         string            `json:"subnetId"`
	CspSubnetId      string            `json:"cspSubnetId"`
	NetworkInterface string            `json:"networkInterface"`
	SecurityGroupIds []string          `json:"securityGroupIds"`
	DataDiskIds      []string          `json:"dataDiskIds"`
	SshKeyId         string            `json:"sshKeyId"`
	CspSshKeyId      string            `json:"cspSshKeyId"`
}

// InfraInfoSummary is the Infra projection returned by the list view
// (GET /ns/{nsId}/infra). Infra-level fields are complete (read from the single
// Infra object); Nodes are NodeSummary. For the full per-Node object use the
// single-Infra (GET /ns/{nsId}/infra/{infraId}) or single-Node API.
type InfraInfoSummary struct {
	ResourceType                  string            `json:"resourceType"`
	Id                            string            `json:"id"`
	Uid                           string            `json:"uid,omitempty"`
	Name                          string            `json:"name"`
	Status                        string            `json:"status"`
	StatusCount                   StatusCountInfo   `json:"statusCount"`
	TargetStatus                  string            `json:"targetStatus"`
	TargetAction                  string            `json:"targetAction"`
	InstallMonAgent               string            `json:"installMonAgent"`
	ConfigureCloudAdaptiveNetwork string            `json:"configureCloudAdaptiveNetwork"`
	Label                         map[string]string `json:"label"`
	SystemLabel                   string            `json:"systemLabel"`
	SystemMessage                 []string          `json:"systemMessage"`
	Description                   string            `json:"description"`
	Node                          []NodeSummary     `json:"node"`
}

// InfraSshCmdResult is struct for Set of SshCmd Results in terms of Infra
type InfraSshCmdResult struct {
	Results []SshCmdResult `json:"results"`
}

// SshCmdResult is struct for SshCmd Result
type SshCmdResult struct { // Tumblebug
	InfraId string         `json:"infraId"`
	NodeId  string         `json:"nodeId"`
	NodeIp  string         `json:"nodeIp"`
	Command map[int]string `json:"command"`
	Stdout  map[int]string `json:"stdout"`
	Stderr  map[int]string `json:"stderr"`
	Err     error          `json:"err"`
}

// SshHostKeyInfo is struct for SSH host key information (TOFU verification)
type SshHostKeyInfo struct {
	// HostKey is the SSH host public key (base64 encoded)
	HostKey string `json:"hostKey,omitempty"`
	// KeyType is the type of the SSH host key (e.g., ssh-rsa, ssh-ed25519, ecdsa-sha2-nistp256)
	KeyType string `json:"keyType,omitempty" example:"ssh-ed25519"`
	// Fingerprint is the SHA256 fingerprint of the SSH host key
	Fingerprint string `json:"fingerprint,omitempty" example:"SHA256:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"`
	// FirstUsedAt is the timestamp when the host key was first stored (TOFU moment)
	FirstUsedAt string `json:"firstUsedAt,omitempty" example:"2024-01-15T10:30:00Z"`
}

// Location is structure for location information
type Location struct {
	Display   string  `mapstructure:"display" json:"display"`
	Latitude  float64 `mapstructure:"latitude" json:"latitude"`
	Longitude float64 `mapstructure:"longitude" json:"longitude"`
}

// RegionInfo is struct for region information
type RegionInfo struct {
	Region string `json:"region" example:"us-east-1"`
	Zone   string `json:"zone,omitempty" example:"us-east-1a"`
}

// ConnConfig is struct for containing modified CB-Spider struct for connection config
type ConnConfig struct {
	ConfigName           string         `json:"configName"`
	ProviderName         string         `json:"providerName"`
	DriverName           string         `json:"driverName"`
	CredentialName       string         `json:"credentialName"`
	CredentialHolder     string         `json:"credentialHolder"`
	RegionZoneInfoName   string         `json:"regionZoneInfoName"`
	RegionZoneInfo       RegionZoneInfo `json:"regionZoneInfo" gorm:"type:text;serializer:json"`
	RegionDetail         RegionDetail   `json:"regionDetail" gorm:"type:text;serializer:json"`
	RegionRepresentative bool           `json:"regionRepresentative"`
	Verified             bool           `json:"verified"`
	// VerifiedMessage explains why verification failed (empty when verified)
	VerifiedMessage string `json:"verifiedMessage,omitempty" gorm:"-"`
}

// RegionZoneInfo is struct for containing region struct
type RegionZoneInfo struct {
	AssignedRegion string `json:"assignedRegion"`
	AssignedZone   string `json:"assignedZone"`
}

// RegionDetail is structure for region information
type RegionDetail struct {
	RegionId           string   `mapstructure:"id" json:"regionId"`
	RegionName         string   `mapstructure:"regionName" json:"regionName"`
	Description        string   `mapstructure:"description" json:"description"`
	Location           Location `mapstructure:"location" json:"location"`
	Zones              []string `mapstructure:"zone" json:"zones"`
	RepresentativeZone *string  `mapstructure:"representativeZone" json:"representativeZone,omitempty"`
}

// VNetReq is a struct to handle 'Create vNet' request toward CB-Tumblebug.
type VNetReq struct { // Tumblebug
	Name           string      `json:"name" validate:"required" example:"vnet00"`
	ConnectionName string      `json:"connectionName" validate:"required" example:"aws-ap-northeast-2"`
	CidrBlock      string      `json:"cidrBlock" example:"10.0.0.0/16"`
	SubnetInfoList []SubnetReq `json:"subnetInfoList"`
	Description    string      `json:"description" example:"vnet00 managed by CB-Tumblebug"`
	// todo: restore the tag list later
	// TagList        []KeyValue    `json:"tagList,omitempty"`
}

// SubnetReq is a struct that represents TB subnet object.
type SubnetReq struct { // Tumblebug
	Name        string `json:"name" validate:"required" example:"subnet00"`
	IPv4_CIDR   string `json:"ipv4_CIDR" validate:"required" example:"10.0.1.0/24"`
	Zone        string `json:"zone,omitempty" default:""`
	Description string `json:"description,omitempty" example:"subnet00 managed by CB-Tumblebug"`
	// todo: restore the tag list later
	// TagList     []KeyValue `json:"tagList,omitempty"`
}

// SshKeyReq is a struct to handle 'Create SSH key' request toward CB-Tumblebug.
type SshKeyReq struct {
	Name           string `json:"name" validate:"required"`
	ConnectionName string `json:"connectionName" validate:"required"`
	Description    string `json:"description"`

	// Fields for "Register existing SSH keys" feature
	// CspResourceId is required to register object from CSP (option=register)
	CspResourceId    string `json:"cspResourceId"`
	Fingerprint      string `json:"fingerprint"`
	Username         string `json:"username"`
	VerifiedUsername string `json:"verifiedUsername"`
	PublicKey        string `json:"publicKey"`
	PrivateKey       string `json:"privateKey"`
}

// SpecInfo is a struct that represents TB spec object.
type SpecInfo struct { // Tumblebug
	// Id is unique identifier for the object
	Id string `json:"id" example:"aws+ap-southeast+csp-06eb41e14121c550a" gorm:"primaryKey"`
	// Uid is universally unique identifier for the object, used for labelSelector
	Uid string `json:"uid,omitempty" example:"wef12awefadf1221edcf"`

	// CspSpecName is name of the spec given by CSP
	CspSpecName string `json:"cspSpecName,omitempty" example:"csp-06eb41e14121c550a"`

	// Name is human-readable string to represent the object
	Name            string  `json:"name" example:"aws-ap-southeast-1"`
	Namespace       string  `json:"namespace,omitempty" example:"default" gorm:"primaryKey"`
	ConnectionName  string  `json:"connectionName,omitempty"`
	ProviderName    string  `json:"providerName,omitempty"`
	RegionName      string  `json:"regionName,omitempty"`
	RegionLatitude  float64 `json:"regionLatitude"`
	RegionLongitude float64 `json:"regionLongitude"`
	// InfraType can be one of vm|k8s|kubernetes|container, etc.
	InfraType             string   `json:"infraType,omitempty"`
	Architecture          string   `json:"architecture,omitempty" example:"x86_64"`
	OsType                string   `json:"osType,omitempty"`
	VCPU                  uint16   `json:"vCPU,omitempty"`
	MemoryGiB             float32  `json:"memoryGiB,omitempty"`
	DiskSizeGB            float32  `json:"diskSizeGB,omitempty"`
	MaxTotalStorageTiB    uint16   `json:"maxTotalStorageTiB,omitempty"`
	NetBwGbps             uint16   `json:"netBwGbps,omitempty"`
	AcceleratorModel      string   `json:"acceleratorModel,omitempty"`
	AcceleratorCount      uint8    `json:"acceleratorCount,omitempty"`
	AcceleratorMemoryGB   float32  `json:"acceleratorMemoryGB,omitempty"`
	AcceleratorType       string   `json:"acceleratorType,omitempty"`
	CostPerHour           float32  `json:"costPerHour,omitempty"`
	Description           string   `json:"description,omitempty"`
	OrderInFilteredResult uint16   `json:"orderInFilteredResult,omitempty"`
	EvaluationStatus      string   `json:"evaluationStatus,omitempty"`
	EvaluationScore01     float32  `json:"evaluationScore01"`
	EvaluationScore02     float32  `json:"evaluationScore02"`
	EvaluationScore03     float32  `json:"evaluationScore03"`
	EvaluationScore04     float32  `json:"evaluationScore04"`
	EvaluationScore05     float32  `json:"evaluationScore05"`
	EvaluationScore06     float32  `json:"evaluationScore06"`
	EvaluationScore07     float32  `json:"evaluationScore07"`
	EvaluationScore08     float32  `json:"evaluationScore08"`
	EvaluationScore09     float32  `json:"evaluationScore09"`
	EvaluationScore10     float32  `json:"evaluationScore10"`
	RootDiskType          string   `json:"rootDiskType"`
	RootDiskSize          int      `json:"rootDiskSize"`
	AssociatedObjectList  []string `json:"associatedObjectList,omitempty" gorm:"type:text;serializer:json"`
	IsAutoGenerated       bool     `json:"isAutoGenerated,omitempty"`

	// SystemLabel is for describing the Resource in a keyword (any string can be used) for special System purpose
	SystemLabel string     `json:"systemLabel,omitempty" example:"Managed by CB-Tumblebug" default:""`
	Details     []KeyValue `json:"details" gorm:"type:text;serializer:json"`
}

// SpecSummary is a lightweight struct containing essential spec information for VmInfo
type SpecSummary struct {
	CspSpecName         string  `json:"cspSpecName,omitempty" example:"t3.medium"`
	VCPU                uint16  `json:"vCPU,omitempty" example:"2"`
	MemoryGiB           float32 `json:"memoryGiB,omitempty" example:"4"`
	AcceleratorModel    string  `json:"acceleratorModel,omitempty" example:"NVIDIA Tesla V100"`
	AcceleratorCount    uint8   `json:"acceleratorCount,omitempty" example:"1"`
	AcceleratorMemoryGB float32 `json:"acceleratorMemoryGB,omitempty" example:"16"`
	AcceleratorType     string  `json:"acceleratorType,omitempty" example:"GPU"`
	CostPerHour         float32 `json:"costPerHour,omitempty" example:"0.0416"`
}

// KeyValue is struct for key-value pair
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ImageInfo is a struct that represents TB image object.
type ImageInfo struct {

	// ResourceType is the type of the resource
	ResourceType string `json:"resourceType"`

	// Composite primary key
	Namespace    string `json:"namespace" example:"default" gorm:"primaryKey"`
	ProviderName string `json:"providerName" gorm:"primaryKey"`
	CspImageName string `json:"cspImageName" example:"csp-06eb41e14121c550a" gorm:"primaryKey" description:"The name of the CSP image for querying image information."`

	// Array field for supporting multiple regions
	RegionList []string `json:"regionList" gorm:"type:text;serializer:json"`

	Id   string `json:"id" example:"aws-ap-southeast-1"`
	Uid  string `json:"uid,omitempty" example:"wef12awefadf1221edcf"`
	Name string `json:"name" example:"aws-ap-southeast-1"`

	// CspImageId is resource identifier managed by CSP
	CspImageId string `json:"cspImageId,omitempty" example:"ami-0d399fba46a30a310"`
	// SourceNodeUid is the UID of the source Node from which this image was created
	SourceNodeUid string `json:"sourceNodeUid" example:"wef12awefadf1221edcf"`
	// SourceCspImageName is the name of the source CSP image from which this image was created
	SourceCspImageName string `json:"sourceCspImageName" example:"csp-06eb41e14121c550a"`

	ConnectionName string `json:"connectionName"`
	InfraType      string `json:"infraType"` // vm|k8s|kubernetes|container, etc.

	FetchedTime  string `json:"fetchedTime"`
	CreationDate string `json:"creationDate"`

	IsGPUImage        bool `json:"isGPUImage" gorm:"column:is_gpu_image" enum:"true|false" default:"false" description:"Whether the image is GPU-enabled or not."`
	IsKubernetesImage bool `json:"isKubernetesImage" gorm:"column:is_kubernetes_image" enum:"true|false" default:"false" description:"Whether this image can be used to create K8s nodes. For AWS/GCP, only type identifiers registered in cloudimage.csv are true."`
	IsBasicImage      bool `json:"isBasicImage" gorm:"column:is_basic_image" enum:"true|false" default:"false" description:"Whether the image is a basic non-GPU OS image (clean OS install, no pre-installed GPU drivers)."`
	IsBasicGpuImage   bool `json:"isBasicGpuImage" gorm:"column:is_basic_gpu_image" enum:"true|false" default:"false" description:"Whether the image is a basic GPU image (GPU drivers pre-installed, recommended for GPU workloads). Mutually exclusive with isBasicImage."`

	OSType string `json:"osType" gorm:"column:os_type" example:"ubuntu 22.04" description:"Simplified OS name and version string"`

	OSArchitecture OSArchitecture `json:"osArchitecture" gorm:"column:os_architecture" example:"x86_64" description:"The architecture of the operating system of the image."`        // arm64, x86_64 etc.
	OSPlatform     OSPlatform     `json:"osPlatform" gorm:"column:os_platform" example:"Linux/UNIX" description:"The platform of the operating system of the image."`                // Linux/UNIX, Windows, NA
	OSDistribution string         `json:"osDistribution" gorm:"column:os_distribution" example:"Ubuntu 22.04~" description:"The distribution of the operating system of the image."` // Ubuntu 22.04~, CentOS 8 etc.
	OSDiskType     string         `json:"osDiskType" gorm:"column:os_disk_type" example:"HDD" description:"The type of the OS disk of for the VM being created."`                    // ebs, HDD, etc.
	OSDiskSizeGB   float64        `json:"osDiskSizeGB" gorm:"column:os_disk_size_gb" example:"50" description:"The (minimum) OS disk size in GB for the VM being created."`          // 10, 50, 100 etc.
	ImageStatus    ImageStatus    `json:"imageStatus" example:"Available" description:"The status of the image, e.g., Available, Deprecated, NA."`                                   // Available, Deprecated, NA

	Details     []KeyValue `json:"details" gorm:"type:text;serializer:json"`
	SystemLabel string     `json:"systemLabel" example:"Managed by CB-Tumblebug" default:""`
	Description string     `json:"description"`

	SystemMessage string `json:"systemMessage,omitempty"`

	// DeletionRequestedAt (RFC3339) marks a deletion tombstone (customImage only):
	// non-empty means the row is kept until CSP-side removal is confirmed
	DeletionRequestedAt string `json:"deletionRequestedAt,omitempty"`

	// CommandHistory stores the status and history of remote commands executed on this VM
	CommandHistory []ImageSourceCommandHistory `json:"commandHistory" gorm:"type:text;serializer:json"`
}

// ImageSourceCommandHistory represents a single remote command execution record
type ImageSourceCommandHistory struct {
	// Index is sequential identifier for this command execution (1, 2, 3, ...)
	Index int `json:"index" example:"1"`
	// CommandExecuted is the actual SSH command executed on the VM (may be adjusted)
	CommandExecuted string `json:"commandExecuted" example:"ls -la"`
}

// ImageSummary is a lightweight struct containing essential image information for VmInfo
type ImageSummary struct {
	ResourceType   string         `json:"resourceType,omitempty" example:"image" description:"image or customImage"`
	CspImageName   string         `json:"cspImageName,omitempty" example:"ami-0123456789abcdef0"`
	OSType         string         `json:"osType" gorm:"column:os_type" example:"ubuntu 22.04" description:"Simplified OS name and version string"`
	OSArchitecture OSArchitecture `json:"osArchitecture,omitempty" example:"x86_64"`
	OSDistribution string         `json:"osDistribution,omitempty" example:"Ubuntu 22.04"`
}

type OSArchitecture string

const (
	ARM32               OSArchitecture = "arm32"
	ARM64               OSArchitecture = "arm64"
	ARM64_MAC           OSArchitecture = "arm64_mac"
	X86_32              OSArchitecture = "x86_32"
	X86_64              OSArchitecture = "x86_64"
	X86_32_MAC          OSArchitecture = "x86_32_mac"
	X86_64_MAC          OSArchitecture = "x86_64_mac"
	S390X               OSArchitecture = "s390x"
	ArchitectureNA      OSArchitecture = "NA"
	ArchitectureUnknown OSArchitecture = ""
)

type OSPlatform string

const (
	Linux_UNIX OSPlatform = "Linux/UNIX"
	Windows    OSPlatform = "Windows"
	PlatformNA OSPlatform = "NA"
)

type ImageStatus string

const (
	// ImageCreating indicates the image is being created (e.g., snapshot in progress)
	// This is a CB-Tumblebug managed state, independent of CB-Spider's status
	ImageCreating ImageStatus = "Creating"

	// ImageAvailable indicates the image is ready and can be used
	ImageAvailable ImageStatus = "Available"

	// ImageFailed indicates the image creation failed
	// This is a terminal state - no further status updates needed
	ImageFailed ImageStatus = "Failed"

	// ImageUnavailable indicates the image is temporarily unavailable
	// This may transition to Available or Failed
	ImageUnavailable ImageStatus = "Unavailable"

	// ImageDeleting indicates the image is being deleted
	ImageDeleting ImageStatus = "Deleting"

	// ImageDeprecated indicates the image is deprecated and should not be used
	ImageDeprecated ImageStatus = "Deprecated"

	// ImageNA indicates the status is not applicable or unknown
	ImageNA ImageStatus = "NA"
)

// SecurityGroupReq is a struct to handle 'Create security group' request toward CB-Tumblebug.
type SecurityGroupReq struct { // Tumblebug
	Name           string             `json:"name" validate:"required"`
	ConnectionName string             `json:"connectionName" validate:"required"`
	VNetId         string             `json:"vNetId"` // Optional for registration: some CSPs (e.g., Azure, Tencent, NHN) don't bind SG to VPC
	Description    string             `json:"description"`
	FirewallRules  *[]FirewallRuleReq `json:"firewallRules"` // validate:"required"`

	// CspResourceId is required to register object from CSP (option=register)
	CspResourceId string `json:"cspResourceId" example:"required for option=register only. ex: csp-06eb41e14121c550a"`
}

// FirewallRuleReq is a struct to get a request for firewall rule info of CB-Tumblebug.
type FirewallRuleReq struct {
	// Ports is to get multiple ports or port ranges as a string (e.g. "22,900-1000,2000-3000")
	// This allows flexibility in specifying single ports or ranges in a comma-separated format.
	// This field is used to handle both single ports and port ranges in a unified way.
	// It can accept a single port (e.g. "22"), a range (e.g. "900-1000"), or multiple ports/ranges (e.g. "22,900-1000,2000-3000").
	Ports string `json:"Ports" example:"22,900-1000,2000-3000"`
	// Protocol is the protocol type for the rule (TCP, UDP, ICMP). Don't use ALL here.
	Protocol string `validate:"required" json:"Protocol" example:"TCP" enums:"TCP,UDP,ICMP"`
	// Direction is the direction of the rule (inbound or outbound)
	Direction string `validate:"required" json:"Direction" example:"inbound" enums:"inbound,outbound"`
	// CIDR is the allowed IP range (e.g. 0.0.0.0/0, 10.0.0/8)
	CIDR string `json:"CIDR" example:"0.0.0.0/0"`
}

// ============================================================================
// NLB types (CB-Tumblebug v0.12.15)
// ============================================================================

// NlbReq maps to CB-Tumblebug's NLBReq for cloud NLB creation.
// nsId and infraId are supplied as path parameters at migration time.
type NlbReq struct {
	CspResourceId string              `json:"cspResourceId,omitempty"`
	Description   string              `json:"description,omitempty"`
	Type          string              `json:"type"`  // PUBLIC | INTERNAL
	Scope         string              `json:"scope"` // REGION | GLOBAL
	Listener      NlbListenerReq      `json:"listener"`
	TargetGroup   NlbTargetGroupReq   `json:"targetGroup"`
	HealthChecker NlbHealthCheckerReq `json:"healthChecker"`
}

// NlbListenerReq mirrors CB-Tumblebug's NLBListenerReq.
type NlbListenerReq struct {
	Protocol string `json:"protocol"` // TCP | UDP
	Port     string `json:"port"`     // "1"–"65535"
}

// NlbTargetGroupReq mirrors CB-Tumblebug's NLBTargetGroupReq.
type NlbTargetGroupReq struct {
	Protocol    string `json:"protocol"`    // TCP | HTTP | HTTPS
	Port        string `json:"port"`        // Backend port
	NodeGroupId string `json:"nodeGroupId"` // NodeGroup ID in the target Infra
}

// NlbHealthCheckerReq mirrors CB-Tumblebug's NLBHealthCheckerReq.
// Note: Tumblebug accepts only interval, threshold, timeout — not protocol or port.
type NlbHealthCheckerReq struct {
	Interval  int `json:"interval"`  // Health check interval in seconds
	Threshold int `json:"threshold"` // Unhealthy threshold count
	Timeout   int `json:"timeout"`   // Health check timeout in seconds
}

// NLBInfoInNs is an NLB together with the Infra it belongs to. NLBs are stored under
// their Infra, so a namespace-wide listing has to carry the parent Infra id for the
// caller to map each NLB back to it.
type NLBInfoInNs struct {
	// InfraId is the Infra this NLB belongs to
	InfraId string `json:"infraId" example:"infra01"`
	NLBInfo
}

// NLBInfo mirrors CB-Tumblebug's NLBInfo, returned after NLB creation/query.
// Synchronized with NLBInfo from CB-Tumblebug v0.12.19.
// Note: json tags added to Type, Scope, CreatedTime which are missing in the original.
type NLBInfo struct {
	ResourceType         string               `json:"resourceType"`
	Id                   string               `json:"id"`
	Uid                  string               `json:"uid,omitempty"`
	CspResourceName      string               `json:"cspResourceName,omitempty"`
	CspResourceId        string               `json:"cspResourceId,omitempty"`
	Name                 string               `json:"name"`
	ConnectionName       string               `json:"connectionName"`
	ConnectionConfig     ConnConfig           `json:"connectionConfig"`
	Type                 string               `json:"type"`
	Scope                string               `json:"scope"`
	Listener             NLBListenerInfo      `json:"listener"`
	TargetGroup          NLBTargetGroupInfo   `json:"targetGroup"`
	HealthChecker        NLBHealthCheckerInfo `json:"healthChecker"`
	CreatedTime          time.Time            `json:"createdTime,omitempty"`
	Description          string               `json:"description"`
	Status               string               `json:"status"`
	SystemMessage        string               `json:"systemMessage,omitempty"`
	KeyValueList         []KeyValue           `json:"keyValueList,omitempty"`
	AssociatedObjectList []string             `json:"associatedObjectList,omitempty"`
	IsAutoGenerated      bool                 `json:"isAutoGenerated"`
	Location             Location             `json:"location"`

	// DeletionRequestedAt (RFC3339) marks a deletion tombstone: non-empty means the
	// record is kept until CSP-side removal is confirmed
	DeletionRequestedAt string `json:"deletionRequestedAt,omitempty"`

	SystemLabel string `json:"systemLabel,omitempty"`
}

// NLBListenerInfo mirrors CB-Tumblebug's NLBListenerInfo.
type NLBListenerInfo struct {
	Protocol     string     `json:"protocol"`
	IP           string     `json:"ip,omitempty"`
	Port         string     `json:"port"`
	DNSName      string     `json:"dnsName,omitempty"`
	KeyValueList []KeyValue `json:"keyValueList,omitempty"`
}

// NLBTargetGroupInfo mirrors CB-Tumblebug's NLBTargetGroupInfo.
type NLBTargetGroupInfo struct {
	Protocol     string     `json:"protocol"`
	Port         string     `json:"port"`
	NodeGroupId  string     `json:"nodeGroupId,omitempty"`
	Nodes        []string   `json:"nodes,omitempty"`
	KeyValueList []KeyValue `json:"keyValueList,omitempty"`
}

// NLBHealthCheckerInfo mirrors CB-Tumblebug's NLBHealthCheckerInfo.
type NLBHealthCheckerInfo struct {
	Protocol     string     `json:"protocol,omitempty"`
	Port         string     `json:"port,omitempty"`
	Interval     int        `json:"interval"`
	Threshold    int        `json:"threshold"`
	Timeout      int        `json:"timeout"`
	KeyValueList []KeyValue `json:"keyValueList,omitempty"`
}
