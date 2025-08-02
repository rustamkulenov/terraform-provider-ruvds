package api

import (
	"strconv"
	"strings"
)

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

// GetOSCode generates a code for the OS based on its id and lowercase name with space replaced with hyphen. ( and ) are removed.
func (os *OS) GetCode() string {
	code := os.Name
	if os.ID > 0 {
		code = strconv.Itoa(int(os.ID)) + "-" + code
	}
	code = strings.ToLower(code)
	code = strings.ReplaceAll(code, " ", "-")
	code = strings.ReplaceAll(code, "(", "")
	code = strings.ReplaceAll(code, ")", "")
	return code
}
