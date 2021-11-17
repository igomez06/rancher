package cloudcredentials

const S3CredentialConfigurationFileKey = "s3CredentialConfig"

type S3CredentialConfig struct {
	AccessKey            string `json:"accessKey"`
	SecretKey            string `json:"secretKey"`
	DefaultBucket        string `json:"defaultBucket,omitempty"`
	DefaultEndpoint      string `json:"defaultEndpoint,omitempty"`
	DefaultEndpointCA    string `json:"defaultEndpointCA,omitempty"`
	DefaultFolder        string `json:"defaultFolder,omitempty"`
	DefaultRegion        string `json:"defaultRegion,omitempty"`
	DefaultSkipSSLVerify string `json:"defaultSkipSSLVerify,omitempty"`
}
