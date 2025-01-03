package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// TerraformListsTrailingCommaRule checks whether ...
type TerraformListsTrailingCommaRule struct {
	tflint.DefaultRule
}

// NewTerraformListsTrailingCommaRule returns a new rule
func NewTerraformListsTrailingCommaRule() *TerraformListsTrailingCommaRule {
	return &TerraformListsTrailingCommaRule{}
}

// Name returns the rule name
func (r *TerraformListsTrailingCommaRule) Name() string {
	return "terraform_lists_trailing_comma"
}

// Enabled returns whether the rule is enabled by default
func (r *TerraformListsTrailingCommaRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *TerraformListsTrailingCommaRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *TerraformListsTrailingCommaRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *TerraformListsTrailingCommaRule) Check(runner tflint.Runner) error {
	files, err := runner.GetFiles()
	if err != nil {
		return err
	}

	diags := runner.WalkExpressions(tflint.ExprWalkFunc(func(e hcl.Expression) hcl.Diagnostics {
		filename := e.Range().Filename
		file := files[filename]

		list, ok := e.(*hclsyntax.TupleConsExpr)
		if !ok || len(list.Exprs) <= 0 {
			return nil
		}

		listRange := list.Range()
		lastItem := list.Exprs[len(list.Exprs)-1]
		lastItemRange := lastItem.Range()

		if listRange.Start.Line == lastItemRange.Start.Line {
			return nil
		}

		if lastItemRange.End.Byte < len(file.Bytes) && file.Bytes[lastItemRange.End.Byte] != ',' {
			if err := runner.EmitIssueWithFix(
				r,
				"Last item in lists should always end with a trailing comma",
				listRange,
				func(f tflint.Fixer) error {
					return f.InsertTextAfter(lastItemRange, ",")
				},
			); err != nil {
				return hcl.Diagnostics{
					{
						Severity: hcl.DiagError,
						Summary:  "failed to call EmitIssueWithFix()",
						Detail:   err.Error(),
					},
				}
			}
		}

		return nil
	}))

	if diags.HasErrors() {
		return diags
	}

	return nil
}
