package main

import (
	"github.com/Gaardsholt/tflint-ruleset-trailing-comma/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

var (
	version = "dev"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "trailing-comma",
			Version: version,
			Rules: []tflint.Rule{
				rules.NewTerraformListsTrailingCommaRule(),
				rules.NewTerraformMapTrailingCommaRule(),
			},
		},
	})
}
