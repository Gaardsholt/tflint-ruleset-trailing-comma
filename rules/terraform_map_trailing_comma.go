package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// TerraformMapTrailingCommaRule checks whether maps have consistent trailing commas
type TerraformMapTrailingCommaRule struct {
	tflint.DefaultRule
	Config *TerraformMapTrailingCommaRuleConfig
}

type TerraformMapTrailingCommaRuleConfig struct {
	Style string `hclext:"style,optional"`
}

// NewTerraformMapTrailingCommaRule returns a new rule
func NewTerraformMapTrailingCommaRule() *TerraformMapTrailingCommaRule {
	return &TerraformMapTrailingCommaRule{}
}

// Name returns the rule name
func (r *TerraformMapTrailingCommaRule) Name() string {
	return "terraform_map_trailing_comma"
}

// Enabled returns whether the rule is enabled by default
func (r *TerraformMapTrailingCommaRule) Enabled() bool {
	return false
}

// Severity returns the rule severity
func (r *TerraformMapTrailingCommaRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *TerraformMapTrailingCommaRule) Link() string {
	return ""
}

// Check checks whether maps have consistent trailing commas
func (r *TerraformMapTrailingCommaRule) Check(runner tflint.Runner) error {
	config := &TerraformMapTrailingCommaRuleConfig{Style: "match"}
	if err := runner.DecodeRuleConfig(r.Name(), config); err != nil {
		return err
	}

	files, err := runner.GetFiles()
	if err != nil {
		return err
	}

	diags := runner.WalkExpressions(tflint.ExprWalkFunc(func(e hcl.Expression) hcl.Diagnostics {
		filename := e.Range().Filename
		file := files[filename]
		fileLength := len(file.Bytes)

		expr, ok := e.(*hclsyntax.ObjectConsExpr)
		if !ok || len(expr.Items) == 0 {
			return nil
		}

		listRange := expr.Range()
		if listRange.Start.Line == listRange.End.Line {
			return nil
		}

		var itemsWithComma []int
		var itemsWithoutComma []int

		for i, item := range expr.Items {
			valRange := item.ValueExpr.Range()
			commaPos := valRange.End.Byte

			for commaPos < fileLength && isWhitespace(file.Bytes[commaPos]) {
				commaPos++
			}

			if commaPos < fileLength && file.Bytes[commaPos] == ',' {
				itemsWithComma = append(itemsWithComma, i)
			} else {
				itemsWithoutComma = append(itemsWithoutComma, i)
			}
		}

		var wantComma bool
		var message string

		switch config.Style {
		case "all":
			wantComma = true
			message = "all: should have comma"
		case "none":
			wantComma = false
			message = "none: should not have comma"
		case "match":
			if len(itemsWithComma) == 0 || len(itemsWithoutComma) == 0 {
				return nil
			}

			if len(itemsWithComma) >= len(itemsWithoutComma) {
				wantComma = true
				message = "match: majority have comma"
			} else {
				wantComma = false
				message = "match: majority no comma"
			}
		default:
			return nil
		}

		if wantComma {
			for _, i := range itemsWithoutComma {
				item := expr.Items[i]
				runner.EmitIssueWithFix(
					r,
					message,
					item.ValueExpr.Range(),
					func(f tflint.Fixer) error {
						return f.InsertTextAfter(item.ValueExpr.Range(), ",")
					},
				)
			}
		} else {
			for _, i := range itemsWithComma {
				item := expr.Items[i]
				startPos := item.ValueExpr.Range().End
				curr := startPos.Byte

				for curr < fileLength && isWhitespace(file.Bytes[curr]) {
					if file.Bytes[curr] == '\n' {
						startPos.Line++
						startPos.Column = 1
					} else {
						startPos.Column++
					}
					startPos.Byte++
					curr++
				}

				if curr < fileLength && file.Bytes[curr] == ',' {
					endPos := startPos
					endPos.Column++
					endPos.Byte++

					runner.EmitIssueWithFix(
						r,
						message,
						item.ValueExpr.Range(),
						func(f tflint.Fixer) error {
							return f.Remove(hcl.Range{
								Filename: filename,
								Start:    startPos,
								End:      endPos,
							})
						},
					)
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
