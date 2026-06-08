package middleware

import (
	"agussyahrilmubarok.github.io/web/pkg/i18n"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const langSessionKey = "lang"

func LocaleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)

		if q := c.Query("lang"); q != "" {
			lang := i18n.SupportedLang(q)
			s.Set(langSessionKey, lang)
			s.Save()
		}

		lang := i18n.DefaultLang
		if v := s.Get(langSessionKey); v != nil {
			if l, ok := v.(string); ok {
				lang = i18n.SupportedLang(l)
			}
		}

		c.Set("Lang", lang)
		c.Set("T", i18n.Translations(lang))
		c.Next()
	}
}
