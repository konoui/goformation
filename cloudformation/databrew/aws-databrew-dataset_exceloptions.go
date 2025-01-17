// Code generated by "go generate". Please don't change this file directly.

package databrew

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// Dataset_ExcelOptions AWS CloudFormation Resource (AWS::DataBrew::Dataset.ExcelOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-exceloptions.html
type Dataset_ExcelOptions struct {

	// HeaderRow AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-exceloptions.html#cfn-databrew-dataset-exceloptions-headerrow
	HeaderRow *bool `json:"HeaderRow,omitempty"`

	// SheetIndexes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-exceloptions.html#cfn-databrew-dataset-exceloptions-sheetindexes
	SheetIndexes *[]int `json:"SheetIndexes,omitempty"`

	// SheetNames AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-exceloptions.html#cfn-databrew-dataset-exceloptions-sheetnames
	SheetNames *[]string `json:"SheetNames,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Dataset_ExcelOptions) AWSCloudFormationType() string {
	return "AWS::DataBrew::Dataset.ExcelOptions"
}
