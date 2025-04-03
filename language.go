package chrono

import (
	"embed"
	"encoding/json"
	"strconv"
	"strings"
	"sync"
)

//go:embed lang
var fs embed.FS

type Language struct {
	dir       string
	locale    string
	resources map[string]string
	rw        *sync.RWMutex
	Error     error
}

func newLanguage() *Language {
	return &Language{
		dir:       "lang/",
		locale:    defaultLocale,
		resources: make(map[string]string),
		rw:        new(sync.RWMutex),
	}
}

// SetLocale sets language locale
func (lang *Language) SetLocale(locale string) *Language {
	if lang == nil || lang.Error != nil {
		return lang
	}

	if locale == "" {
		lang.Error = emptyLocaleError()
		return lang
	}

	lang.rw.Lock()
	defer lang.rw.Unlock()

	lang.locale = locale
	filename := lang.dir + locale + ".json"
	bytes, err := fs.ReadFile(filename)
	if err != nil {
		lang.Error = invalidLocaleError(filename)
		return lang
	}

	_ = json.Unmarshal(bytes, &lang.resources)
	return lang
}

// SetResources sets language resources
func (lang *Language) SetResources(resources map[string]string) *Language {
	if lang == nil || lang.Error != nil {
		return lang
	}

	if resources == nil {
		return lang
	}

	lang.rw.Lock()
	defer lang.rw.Unlock()

	if len(lang.resources) == 0 {
		lang.resources = resources
		return lang
	}

	for k, v := range resources {
		if _, ok := lang.resources[k]; ok {
			lang.resources[k] = v
		}
	}
	return lang
}

// translate string
func (lang *Language) translate(unit string, value int64) string {
	if lang == nil || lang.resources == nil {
		return ""
	}

	lang.rw.Lock()
	defer lang.rw.Unlock()

	if len(lang.resources) == 0 {
		lang.rw.Unlock()
		lang.SetLocale(defaultLocale)
		lang.rw.Lock()
	}

	elem := strings.Split(lang.resources[unit], "|")
	number := getAbsValue(value)

	if len(elem) == 1 {
		return strings.Replace(elem[0], "%d", strconv.FormatInt(value, 10), 1)
	}

	if int64(len(elem)) <= number {
		return strings.Replace(elem[len(elem)-1], "%d", strconv.FormatInt(value, 10), 1)
	}

	if !strings.Contains(elem[number-1], "%d") && value < 0 {
		return "-" + elem[number-1]
	}

	return strings.Replace(elem[number-1], "%d", strconv.FormatInt(value, 10), 1)
}
