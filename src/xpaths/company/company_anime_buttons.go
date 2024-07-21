package company_xpaths

import (
	"strings"

	"regexp"

	"strconv"

	"github.com/antchfx/htmlquery"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/html"
)

const NUMBER_REGEX string = `\((\d+)\)`

func Company_anime_buttons_totals(anime_buttons_node *html.Node) (totals_int int, err error) {
	const total_xpath string = "/li[1]"
	totals_node, err := htmlquery.Query(anime_buttons_node, total_xpath)
	if err != nil {
		log.Error().Err(err).Msg("Error with totals node")
		return -1, err
	}

	totals_full_string := htmlquery.InnerText(totals_node)
	log.Trace().Str("value", totals_full_string).Msg("Totals full text")

	// validate "All" appears in the string
	if !strings.Contains(totals_full_string, "All") {
		log.Error().Str("actual", totals_full_string).Msg("Totals text does not contain 'All'")
		return -1, err
	}

	// get number only from string
	totals_int, err = company_anime_buttons_get_number(totals_full_string)
	if err != nil {
		log.Error().Msg("Error getting int from totals_full_string (comapny)")
		return -1, err
	}

	return totals_int, nil

}

func company_anime_buttons_tv(anime_buttons_node *html.Node) (tv_int int, err error) {
	const tv_xpath string = "/li[2]"
	tv_node, err := htmlquery.Query(anime_buttons_node, tv_xpath)
	if err != nil {
		log.Error().Err(err).Msg("Error with tv node")
		return -1, err
	}

	tv_full_text := htmlquery.InnerText(tv_node)
	log.Trace().Str("value", tv_full_text).Msg("Tv full text")

	// validate "TV" appears in the string
	if !strings.Contains(tv_full_text, "TV") {
		log.Error().Str("actual", tv_full_text).Msg("TV text does not contain 'TV'")
		return -1, err
	}

	// get number only from string
	tv_int, err = company_anime_buttons_get_number(tv_full_text)
	if err != nil {
		log.Error().Msg("Error getting int from tv_full_string (comapny)")
		return -1, err
	}

	return tv_int, nil
}

func company_anime_buttons_ona(anime_buttons_node *html.Node) (ona_int int, err error) {
	const ona_xpath string = "/li[3]"
	ona_node, err := htmlquery.Query(anime_buttons_node, ona_xpath)
	if err != nil {
		log.Error().Err(err).Msg("Error with ona node")
		return -1, err
	}

	ona_full_text := htmlquery.InnerText(ona_node)
	log.Trace().Str("value", ona_full_text).Msg("ONA full text")

	// validate "ONA" appears in the string
	if !strings.Contains(ona_full_text, "ONA") {
		log.Error().Str("actual", ona_full_text).Msg("ONA text does not contain 'ONA'")
		return -1, err
	}

	// get number only from string
	ona_int, err = company_anime_buttons_get_number(ona_full_text)
	if err != nil {
		log.Error().Msg("Error getting int from ona_full_text (comapny)")
		return -1, err
	}

	return ona_int, nil
}

func company_anime_buttons_ova(anime_buttons_node *html.Node) (ova_int int, err error) {
	const ova_xpath string = "/li[4]"
	ova_node, err := htmlquery.Query(anime_buttons_node, ova_xpath)
	if err != nil {
		log.Error().Err(err).Msg("Error with OVA node")
		return -1, err
	}

	ova_full_text := htmlquery.InnerText(ova_node)
	log.Trace().Str("value", ova_full_text).Msg("OVA full text")

	// validate "OVA" appears in the string
	if !strings.Contains(ova_full_text, "OVA") {
		log.Error().Str("actual", ova_full_text).Msg("OVA text does not contain 'OVA'")
		return -1, err
	}

	// get number only from string
	ova_int, err = company_anime_buttons_get_number(ova_full_text)
	if err != nil {
		log.Error().Msg("Error getting int from ova_full_text (comapny)")
		return -1, err
	}

	return ova_int, nil
}

func company_anime_buttons_movie(anime_buttons_node *html.Node) (movie_int int, err error) {
	const movie_xpath string = "/li[5]"
	movie_node, err := htmlquery.Query(anime_buttons_node, movie_xpath)
	if err != nil {
		log.Error().Err(err).Msg("Error with movie node")
		return -1, err
	}

	movie_full_text := htmlquery.InnerText(movie_node)
	log.Trace().Str("value", movie_full_text).Msg("Movie full text")

	// validate "movie" appears in the string
	if !strings.Contains(movie_full_text, "Movie") {
		log.Error().Str("actual", movie_full_text).Msg("Movie text does not contain 'Movie'")
		return -1, err
	}

	// get number only from string
	movie_int, err = company_anime_buttons_get_number(movie_full_text)
	if err != nil {
		log.Error().Msg("Error getting int from movie_full_text (comapny)")
		return -1, err
	}

	return movie_int, nil
}

func company_anime_buttons_other(anime_buttons_node *html.Node) (other_int int, err error) {
	const other_xpath string = "/li[6]"
	other_node, err := htmlquery.Query(anime_buttons_node, other_xpath)
	if err != nil {
		log.Error().Err(err).Msg("Error with other node")
		return -1, err
	}

	other_full_text := htmlquery.InnerText(other_node)
	log.Trace().Str("value", other_full_text).Msg("Other full text")

	// validate "other" appears in the string
	if !strings.Contains(other_full_text, "Other") {
		log.Error().Str("actual", other_full_text).Msg("Other text does not contain 'Other'")
		return -1, err
	}

	other_int, err = company_anime_buttons_get_number(other_full_text)
	if err != nil {
		log.Error().Msg("Error getting int from other_full_text (comapny)")
		return -1, err
	}

	return other_int, nil
}

func company_anime_buttons_get_number(input_string string) (extracted_int int, err error) {
	// regex
	re := regexp.MustCompile(`\((\d+)\)`)

	// get number
	extracted_number := re.FindStringSubmatch(input_string)

	// cast to int
	extracted_int, err = strconv.Atoi(extracted_number[1])
	if err != nil {
		log.Error().
			Err(err).
			Str("input string", input_string).
			Msg("Error converting string to int")
		return -1, err
	}
	log.Trace().Str("input string", input_string).Int("extracted int", extracted_int).Msg("Int extraction (company/anime buttons)")

	return extracted_int, nil
}
