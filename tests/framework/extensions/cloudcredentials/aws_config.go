package cloudcredentials

const AmazonEC2CredentialConfigurationFileKey = "amazonec2credentialConfig"

type AmazonEC2CredentialConfig struct {
	AccessKey     string `json:"accessKey" yaml:"accessKey"`
	SecretKey     string `json:"secretKey" yaml:"secretKey"`
	DefaultRegion string `json:"defaultRegion,omitempty" yaml:"defaultRegion"`
}
