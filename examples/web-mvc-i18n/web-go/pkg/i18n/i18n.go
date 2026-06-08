package i18n

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
)

const (
	LangEN      = "en"
	LangID      = "id"
	DefaultLang = LangEN
)

var (
	messages = map[string]map[string]interface{}{}
	mu       sync.RWMutex
)

func Load(dir string) error {
	langs := []string{LangEN, LangID}
	for _, lang := range langs {
		path := fmt.Sprintf("%s/%s.json", dir, lang)
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("i18n: failed to read %s: %w", path, err)
		}
		var m map[string]interface{}
		if err := json.Unmarshal(data, &m); err != nil {
			return fmt.Errorf("i18n: failed to parse %s: %w", path, err)
		}
		mu.Lock()
		messages[lang] = m
		mu.Unlock()
	}
	return nil
}

func T(lang, key string) string {
	mu.RLock()
	defer mu.RUnlock()

	val := lookup(messages[lang], key)
	if val != "" {
		return val
	}
	if lang != DefaultLang {
		val = lookup(messages[DefaultLang], key)
		if val != "" {
			return val
		}
	}
	return key
}

func lookup(m map[string]interface{}, key string) string {
	if m == nil {
		return ""
	}
	parts := strings.SplitN(key, ".", 2)
	v, ok := m[parts[0]]
	if !ok {
		return ""
	}
	if len(parts) == 1 {
		if s, ok := v.(string); ok {
			return s
		}
		return ""
	}
	nested, ok := v.(map[string]interface{})
	if !ok {
		return ""
	}
	return lookup(nested, parts[1])
}

func Translations(lang string) map[string]string {
	mu.RLock()
	m := messages[lang]
	mu.RUnlock()

	out := map[string]string{}
	flatten(m, "", out)
	return out
}

func flatten(m map[string]interface{}, prefix string, out map[string]string) {
	for k, v := range m {
		fullKey := k
		if prefix != "" {
			fullKey = prefix + "." + k
		}
		switch val := v.(type) {
		case string:
			out[fullKey] = val
		case map[string]interface{}:
			flatten(val, fullKey, out)
		}
	}
}

func SupportedLang(lang string) string {
	switch lang {
	case LangEN, LangID:
		return lang
	default:
		return DefaultLang
	}
}
