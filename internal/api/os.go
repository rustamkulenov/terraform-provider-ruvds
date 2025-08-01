package api

type OSRequirements struct {
	CPU   int32   `json:"cpu"`
	RAM   float32 `json:"ram"`
	Drive int32   `json:"drive"`
}

type OS struct {
	ID               int32          `json:"id"`
	Name             string         `json:"name"`
	IsActive         bool           `json:"is_active"`
	Type             string         `json:"type"`
	SshKeysSupported bool           `json:"ssh_keys_supported"`
	Requirements     OSRequirements `json:"os_requirements"`
}

type OSResponse struct {
	Items []OS `json:"os"`
}
