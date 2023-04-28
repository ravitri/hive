package aws

// MachinePoolPlatform stores the configuration for a machine pool
// installed on AWS.
type MachinePoolPlatform struct {
	// Zones is list of availability zones that can be used.
	Zones []string `json:"zones,omitempty"`

	// Subnets is the list of subnets to which to attach the machines.
	// There must be exactly one private subnet for each availability zone used.
	// If public subnets are specified, there must be exactly one private and one public subnet specified for each availability zone.
	Subnets []string `json:"subnets,omitempty"`

	// InstanceType defines the ec2 instance type.
	// eg. m4-large
	InstanceType string `json:"type"`

	// EC2RootVolume defines the storage for ec2 instance.
	EC2RootVolume `json:"rootVolume"`

	// SpotMarketOptions allows users to configure instances to be run using AWS Spot instances.
	// +optional
	SpotMarketOptions *SpotMarketOptions `json:"spotMarketOptions,omitempty"`

	// EC2MetadataOptions defines metadata service interaction options for EC2 instances in the machine pool.
	// +optional
	EC2Metadata *EC2Metadata `json:"metadataService,omitempty"`
}

// SpotMarketOptions defines the options available to a user when configuring
// Machines to run on Spot instances.
// Most users should provide an empty struct.
type SpotMarketOptions struct {
	// The maximum price the user is willing to pay for their instances
	// Default: On-Demand price
	// +optional
	MaxPrice *string `json:"maxPrice,omitempty"`
}

// EC2RootVolume defines the storage for an ec2 instance.
type EC2RootVolume struct {
	// IOPS defines the iops for the storage.
	// +optional
	IOPS int `json:"iops,omitempty"`
	// Size defines the size of the storage.
	Size int `json:"size"`
	// Type defines the type of the storage.
	Type string `json:"type"`
	// The KMS key that will be used to encrypt the EBS volume.
	// If no key is provided the default KMS key for the account will be used.
	// https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetEbsDefaultKmsKeyId.html
	// +optional
	KMSKeyARN string `json:"kmsKeyARN,omitempty"`
}

// EC2Metadata defines the metadata service interaction options for an ec2 instance.
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-instance-metadata-service.html
type EC2Metadata struct {
	// Authentication determines whether or not the host requires the use of authentication when interacting with the metadata service.
	// When using authentication, this enforces v2 interaction method (IMDSv2) with the metadata service.
	// When omitted, this means the user has no opinion and the value is left to the platform to choose a good
	// default, which is subject to change over time. The current default is optional.
	// At this point this field represents `HttpTokens` parameter from `InstanceMetadataOptionsRequest` structure in AWS EC2 API
	// https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceMetadataOptionsRequest.html
	// +optional
	Authentication string `json:"authentication,omitempty"`
}
