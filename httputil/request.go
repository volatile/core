package httputil

import (
	"net/http"
	"regexp"
	"sort"
	"strings"
)

var langQualityRegexp = regexp.MustCompile(";q=([0-9](?:.[0-9])?)")

// AcceptedLanguages parses the Accept-Language header and returns the lowercase language tags sorted by client's preference (the first, the most wanted).
// The result is nil when Accept-Language header isn't set.
func AcceptedLanguages(r *http.Request) (langs []string) {
	h := r.Header.Get("Accept-Language")
	if h == "" {
		return
	}

	// Fill langs with qualityValue;languageTag.
	for _, l := range strings.Split(h, ",") {
		langs = append(langs, parseLangQuality(l)+";"+parseLang(l))
	}

	sort.Sort(sort.Reverse(sort.StringSlice(langs))) // Sort languages with greater quality value first.

	// Only keep the language tags.
	for i, l := range langs {
		langs[i] = strings.SplitN(l, ";", 2)[1]
	}
	return
}

// parseLang returns the lowercase language tag parsed from a language tag with its quality value.
func parseLang(l string) string {
	return strings.TrimSpace(strings.ToLower(strings.SplitN(l, ";", 2)[0]))
}

// parseLangQuality returns the quality value parsed from a language tag with its quality value.
func parseLangQuality(l string) string {
	q := langQualityRegexp.FindStringSubmatch(l)
	if len(q) == 2 {
		return q[1]
	}
	return "1.0"
}
