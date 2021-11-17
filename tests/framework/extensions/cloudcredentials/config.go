package cloudcredentials

import (
	"github.com/rancher/norman/types"
)

const CloudCredentialType = "cloudCredential"

type CloudCredential struct {
	types.Resource
	Annotations                  map[string]string             `json:"annotations,omitempty"`
	Created                      string                        `json:"created,omitempty"`
	CreatorID                    string                        `json:"creatorId,omitempty"`
	Description                  string                        `json:"description,omitempty"`
	Labels                       map[string]string             `json:"labels,omitempty"`
	Name                         string                        `json:"name,omitempty"`
	Removed                      string                        `json:"removed,omitempty"`
	S3CredentialConfig           *S3CredentialConfig           `json:"s3credentialConfig,omitempty"`
	DigitalOceanCredentialConfig *DigitalOceanCredentialConfig `json:"digitaloceancredentialConfig,omitempty"`
	UUID                         string                        `json:"uuid,omitempty"`
}
