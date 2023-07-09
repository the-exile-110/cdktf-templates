package main

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	awsProvider "github.com/cdktf/cdktf-provider-aws-go/aws/v10/provider"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	awsProvider.NewAwsProvider(stack, jsii.String("AWS"), &awsProvider.AwsProviderConfig{
		Region: jsii.String("us-west-1"),
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)
	stack := NewMyStack(app, "aws_instance")

	cdktf.NewS3Backend(stack, &cdktf.S3BackendConfig{
		Bucket: jsii.String("cdktf-go"),
		Key:    jsii.String("terraform.tfstate"),
		Region: jsii.String("ap-northeast-1"),
	})
	fmt.Println(stack)

	app.Synth()
}
