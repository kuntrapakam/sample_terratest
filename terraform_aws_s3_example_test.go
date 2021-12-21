package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAwsS3Example(t *testing.T) {
	t.Parallel()

	// expectedName := fmt.Sprintf("terratest-aws-s3-example-%s", strings.ToLower(random.UniqueId()))
	expectedName := "bucket_id"

	expectedEnvironment := "Dev"

	awsRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		TerraformDir: ".",

		Vars: map[string]interface{}{
			"tag-bucket-name":        expectedName,
			"tag-bucket-environment": expectedEnvironment,
			"region":                 awsRegion,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	bucketID := terraform.Output(t, terraformOptions, "bucket_id")

	actualStatus := aws.GetS3BucketVersioning(t, awsRegion, bucketID)
	expectedStatus := "Enabled"
	assert.Equal(t, expectedStatus, actualStatus)
}
