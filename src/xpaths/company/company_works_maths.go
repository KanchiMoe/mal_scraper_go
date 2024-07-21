package company_xpaths

import "github.com/rs/zerolog/log"

func Validate_company_totals(expected_total int,
	tv int,
	ona int,
	ova int,
	movie int,
	other int) (totals_match bool) {

	// calculate totals
	calculated_total := tv + ona + ova + movie + other
	log.Trace().
		Int("expected", expected_total).
		Int("actual", calculated_total).
		Msg("Comparing expected total with actual")

	// compare
	if calculated_total != expected_total {
		log.Error().Int("expected", expected_total).Int("actual", calculated_total).Msg("Calculated and expected totals don't match")
		return false
	}

	log.Debug().Int("expected", expected_total).Int("actual", calculated_total).Msg("Totals of works match")
	return true
}
