package api

type SSHKey struct {
	SshKeyId          string `json:"ssh_key_id"`
	Name              string `json:"name"`
	PublicKey         string `json:"public_key"`
	Md5Fingerprint    string `json:"md5_fingerprint"`
	Sha256Fingerprint string `json:"sha256_fingerprint"`
}

type GetSSHKeysOkResponse struct {
	SshKeys []SSHKey `json:"ssh_keys"`
}

type CreateSSHKeyRequest struct {
	Name      string `json:"name"`
	PublicKey string `json:"public_key"`
}
