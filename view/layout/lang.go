package layout

func getDir(lang string) string {
	if lang == "ar" {
		return "rtl"
	}
	return "ltr"
}

func getSwitchLocaleURL(currentLang string) string {
	if currentLang == "en" {
		return "/"
	}
	return "/?lang=en"
}
