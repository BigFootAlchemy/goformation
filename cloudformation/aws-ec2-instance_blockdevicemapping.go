package cloudformation

// AWSEC2Instance_BlockDeviceMapping AWS CloudFormation Resource (AWS::EC2::Instance.BlockDeviceMapping)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-mapping.html
type AWSEC2Instance_BlockDeviceMapping struct {

	// DeviceName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-mapping.html#cfn-ec2-blockdev-mapping-devicename
	DeviceName string `json:"DeviceName,omitempty"`

	// Ebs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-mapping.html#cfn-ec2-blockdev-mapping-ebs
	Ebs *AWSEC2Instance_Ebs `json:"Ebs,omitempty"`

	// NoDevice AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-mapping.html#cfn-ec2-blockdev-mapping-nodevice
	NoDevice *AWSEC2Instance_NoDevice `json:"NoDevice,omitempty"`

	// VirtualName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-mapping.html#cfn-ec2-blockdev-mapping-virtualname
	VirtualName string `json:"VirtualName,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2Instance_BlockDeviceMapping) AWSCloudFormationType() string {
	return "AWS::EC2::Instance.BlockDeviceMapping"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2Instance_BlockDeviceMapping) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
