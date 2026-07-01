package cloudmodel

// MigratedNlbResult is the response returned after NLB migration.
type MigratedNlbResult struct {
	Status      string    `json:"status"` // "created" | "partial" | "failed"
	Description string    `json:"description"`
	NlbList     []NLBInfo `json:"nlbList"`
}