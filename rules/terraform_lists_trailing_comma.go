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
	return tflint.WARNING
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

		if !isFileInCurrentModule(filename) {
			return nil
		}

		if _, ok := files[filename]; !ok {
			return nil
		}

		file := files[filename]
		fileLength := len(file.Bytes)

		if fileLength == 0 {
			return nil
		}

		list, ok := e.(*hclsyntax.TupleConsExpr)
		if !ok || len(list.Exprs) == 0 {
			return nil
		}

		listRange := list.Range()
		lastItem := list.Exprs[len(list.Exprs)-1]
		lastItemRange := lastItem.Range()

		if listRange.Start.Line == lastItemRange.Start.Line {
			return nil
		}

		// Check if there's already a trailing comma after the last item
		// We need to skip whitespace and newlines to handle heredoc cases
		commaPos := lastItemRange.End.Byte

		// Skip whitespace and newlines after the last item to look for a comma
		for commaPos < fileLength && isWhitespace(file.Bytes[commaPos]) {
			commaPos++
		}

		if commaPos < fileLength && file.Bytes[commaPos] == ',' {
			// It already has a trailling comma
			return nil
		}

		insertText := ","
		// Check if the last item is a heredoc.
		// A heredoc is a TemplateExpr with a single LiteralValueExpr part.
		if template, ok := lastItem.(*hclsyntax.TemplateExpr); ok {
			if len(template.Parts) == 1 {
				if _, isLiteral := template.Parts[0].(*hclsyntax.LiteralValueExpr); isLiteral {
					// This is a strong indicator of a heredoc, especially if it spans multiple lines.
					if template.Range().Start.Line != template.Range().End.Line {
						insertText = "\n,"
					}
				}
			}
		}

		if err := runner.EmitIssueWithFix(
			r,
			"Last item in lists should always end with a trailing comma",
			listRange,
			func(f tflint.Fixer) error {
				return f.InsertTextAfter(lastItemRange, insertText)
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

		return nil
	}))

	if diags.HasErrors() {
		return diags
	}

	return nil
}
