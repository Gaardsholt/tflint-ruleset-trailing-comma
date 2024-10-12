package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_TerraformListsTrailingCommaRule(t *testing.T) {
	tests := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found",
			Content: `resource "vault_generic_endpoint" "user" {
  depends_on = [
    random_password.svc_acc_pass
  ]
}`,
			Expected: helper.Issues{
				{
					Rule:    NewTerraformListsTrailingCommaRule(),
					Message: "Last item in lists should always end with a trailing comma",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 16},
						End:      hcl.Pos{Line: 4, Column: 4},
					},
				},
			},
		},
	}

	rule := NewTerraformListsTrailingCommaRule()

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"resource.tf": test.Content})

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, test.Expected, runner.Issues)
		})
	}
}
