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
		{
			Name: "heredoc without trailing comma",
			Content: `resource "terraform_data" "test" {
  input = [
    "test",
    <<-HERE
      Lorem ipsum
    HERE
  ]
}`,
			Expected: helper.Issues{
				{
					Rule:    NewTerraformListsTrailingCommaRule(),
					Message: "Last item in lists should always end with a trailing comma",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 11},
						End:      hcl.Pos{Line: 7, Column: 4},
					},
				},
			},
		},
		{
			Name: "heredoc with trailing comma",
			Content: `resource "terraform_data" "test" {
  input = [
    "test",
    <<-HERE
      Lorem ipsum
    HERE
    ,
  ]
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "function call issue found",
			Content: `locals {
  test = merge(
    local.a,
    local.b,
    local.c
  )
}`,
			Expected: helper.Issues{
				{
					Rule:    NewTerraformListsTrailingCommaRule(),
					Message: "Last item in lists should always end with a trailing comma",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 10},
						End:      hcl.Pos{Line: 6, Column: 4},
					},
				},
			},
		},
		{
			Name: "function call no issue with trailing comma",
			Content: `locals {
  test = merge(
    local.a,
    local.b,
    local.c,
  )
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "function call no issue single line",
			Content: `locals {
  test = merge(local.a, local.b, local.c)
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "function call heredoc without trailing comma",
			Content: `locals {
  test = merge(
    local.a,
    <<-HERE
      Lorem ipsum
    HERE
  )
}`,
			Expected: helper.Issues{
				{
					Rule:    NewTerraformListsTrailingCommaRule(),
					Message: "Last item in lists should always end with a trailing comma",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 10},
						End:      hcl.Pos{Line: 7, Column: 4},
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
