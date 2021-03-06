package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSApplicationAutoScalingScalingPolicy AWS CloudFormation Resource (AWS::ApplicationAutoScaling::ScalingPolicy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html
type AWSApplicationAutoScalingScalingPolicy struct {

	// PolicyName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-policyname
	PolicyName string `json:"PolicyName,omitempty"`

	// PolicyType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-policytype
	PolicyType string `json:"PolicyType,omitempty"`

	// ResourceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-resourceid
	ResourceId string `json:"ResourceId,omitempty"`

	// ScalableDimension AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-scalabledimension
	ScalableDimension string `json:"ScalableDimension,omitempty"`

	// ScalingTargetId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-scalingtargetid
	ScalingTargetId string `json:"ScalingTargetId,omitempty"`

	// ServiceNamespace AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-servicenamespace
	ServiceNamespace string `json:"ServiceNamespace,omitempty"`

	// StepScalingPolicyConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration
	StepScalingPolicyConfiguration *AWSApplicationAutoScalingScalingPolicy_StepScalingPolicyConfiguration `json:"StepScalingPolicyConfiguration,omitempty"`

	// TargetTrackingScalingPolicyConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration
	TargetTrackingScalingPolicyConfiguration *AWSApplicationAutoScalingScalingPolicy_TargetTrackingScalingPolicyConfiguration `json:"TargetTrackingScalingPolicyConfiguration,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApplicationAutoScalingScalingPolicy) AWSCloudFormationType() string {
	return "AWS::ApplicationAutoScaling::ScalingPolicy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApplicationAutoScalingScalingPolicy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSApplicationAutoScalingScalingPolicy) MarshalJSON() ([]byte, error) {
	type Properties AWSApplicationAutoScalingScalingPolicy
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSApplicationAutoScalingScalingPolicy) UnmarshalJSON(b []byte) error {
	type Properties AWSApplicationAutoScalingScalingPolicy
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSApplicationAutoScalingScalingPolicy(*res.Properties)
	}

	return nil
}

// GetAllAWSApplicationAutoScalingScalingPolicyResources retrieves all AWSApplicationAutoScalingScalingPolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSApplicationAutoScalingScalingPolicyResources() map[string]AWSApplicationAutoScalingScalingPolicy {
	results := map[string]AWSApplicationAutoScalingScalingPolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSApplicationAutoScalingScalingPolicy:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ApplicationAutoScaling::ScalingPolicy" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSApplicationAutoScalingScalingPolicy
						if err := json.Unmarshal(b, &result); err == nil {
							results[name] = result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSApplicationAutoScalingScalingPolicyWithName retrieves all AWSApplicationAutoScalingScalingPolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApplicationAutoScalingScalingPolicyWithName(name string) (AWSApplicationAutoScalingScalingPolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSApplicationAutoScalingScalingPolicy:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ApplicationAutoScaling::ScalingPolicy" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSApplicationAutoScalingScalingPolicy
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSApplicationAutoScalingScalingPolicy{}, errors.New("resource not found")
}
