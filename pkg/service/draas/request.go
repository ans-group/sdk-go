package draas

// PatchSolutionRequest represents a request to patch a solution
type PatchSolutionRequest struct {
	Name       string `json:"name,omitempty"`
	IOPSTierID string `json:"iops_tier_id,omitempty"`
}

type ResetBackupServiceCredentialsRequest struct {
	Password string `json:"password"`
}
