package rules

import (
	"github.com/zricethezav/gitleaks/v8/cmd/generate/config/utils"
	"regexp"

	"github.com/zricethezav/gitleaks/v8/cmd/generate/secrets"
	"github.com/zricethezav/gitleaks/v8/config"
)

func Twilio() *config.Rule {
	// define rule
	r := config.Rule{
		RuleID:      "twilio-api-key",
		Description: "Found a Twilio API Key, posing a risk to communication services and sensitive customer interaction data.",
		Regex:       regexp.MustCompile(`SK[0-9a-fA-F]{32}`),
		Entropy:     3,
		Keywords:    []string{"SK"},
	}

	// validate
	tps := []string{
		"twilioAPIKey := \"SK" + secrets.NewSecret(utils.Hex("32")) + "\"",
	}
	return utils.Validate(r, tps, nil)
}
