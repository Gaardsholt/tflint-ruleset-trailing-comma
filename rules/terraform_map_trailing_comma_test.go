package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_TerraformMapTrailingCommaRule(t *testing.T) {
	tests := []struct {
		Name     string
		Content  string
		Files    map[string]string
		Expected helper.Issues
	}{
		{
			Name: "no issues",
			Content: `locals {
  a_dictionary = {
    "one"  = "fish",
    "two"  = "fish",
    "red"  = "fish",
    "blue" = "fish",
  }
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "no issue, one item",
			Content: `locals {
  a_dictionary = {
    "one"  = "fish",
  }
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "match: majority no comma",
			Content: `locals {
  a_dictionary = {
    "one"  = "fish",
    "two"  = "fish"
    "red"  = "fish",
    "blue" = "fish"
  }
}`,
			Expected: helper.Issues{
				{
					Rule:    NewTerraformMapTrailingCommaRule(),
					Message: "match: majority have comma",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 14},
						End:      hcl.Pos{Line: 4, Column: 20},
					},
				},
				{
					Rule:    NewTerraformMapTrailingCommaRule(),
					Message: "match: majority have comma",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 6, Column: 14},
						End:      hcl.Pos{Line: 6, Column: 20},
					},
				},
			},
		},
		{
			Name: "match: majority really no comma",
			Content: `locals {
  a_dictionary = {
    "one"  = "fish",
    "two"  = "fish"
    "red"  = "fish"
  }
}`,
			Expected: helper.Issues{
				{
					Rule:    NewTerraformMapTrailingCommaRule(),
					Message: "match: majority no comma",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 14},
						End:      hcl.Pos{Line: 3, Column: 20},
					},
				},
			},
		},
		{
			Name: "bug: do not remove separator comma on same line",
			Content: `locals {
  a_dictionary = {
    "one" = "fish", "two" = "fish"
    "red" = "fish"
  }
}`,
			Expected: helper.Issues{},
		},
		{
			Name:     "single line map",
			Content:  `b_dictionary = { "one" = "fish", "two" = "fish", "red" = "fish", "blue" = "fish" }`,
			Expected: helper.Issues{},
		},
		{
			Name: "submodule file ignored",
			Content: `
		module "child" {
		  source = "./modules/child"
		}
		`,
			Files: map[string]string{
				"modules/child/main.tf": `locals { a_dict = { "one" = "fish", "two" = "fish" } }`,
			},
			Expected: helper.Issues{},
		},
	}

	rule := NewTerraformMapTrailingCommaRule()

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			files := test.Files
			if files == nil {
				files = map[string]string{"resource.tf": test.Content}
			} else {
				if test.Content != "" {
					files["resource.tf"] = test.Content
				}
			}

			runner := helper.TestRunner(t, files)

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, test.Expected, runner.Issues)
		})
	}
}
