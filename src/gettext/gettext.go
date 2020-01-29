package gettext

import "github.com/chai2010/gettext-go/gettext"

func Setup(locale string, domain string, dir string) {
	gettext.SetLocale(locale)
	gettext.Textdomain(domain)

	gettext.BindTextdomain(domain, dir, nil)
}

func ChangeLocale(locale string) {
	gettext.SetLocale(locale)
}

func Translate(input string) string {
	return gettext.PGettext("", input)
}
