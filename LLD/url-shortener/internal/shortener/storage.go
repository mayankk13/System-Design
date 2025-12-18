package shortener

import "sync"

// map: short - long URL
var urlDB = make(map[string]string)
var dbMutext sync.RWMutex

// check exists
func shortLinkExists(short string) bool {
	dbMutext.RLock()
	defer dbMutext.RUnlock()

	_, exists := urlDB[short]
	return exists
}

// save mapping
func saveShortURL(short, LongURL string) {
	dbMutext.RLock()
	defer dbMutext.RUnlock()

	urlDB[short] = LongURL
}

func getLongURL(short string) (string, bool) {
	dbMutext.RLock()
	defer dbMutext.RUnlock()

	longURL, exists := urlDB[short]

	return longURL, exists
}
