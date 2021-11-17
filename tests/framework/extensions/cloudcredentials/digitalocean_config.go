package cloudcredentials

const DigitalOceanCredentialConfigurationFileKey = "digitalOceanCredentialConfig"

type DigitalOceanCredentialConfig struct {
	AccessToken string `json:"accessToken,omitempty"`
}
