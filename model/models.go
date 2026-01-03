package model

type Content struct {
	Title         string
	Description   string
	AppStore      string
	GooglePlay    string
	DownloadOn    string
	GetItOn       string
	LangToggle    string
	TargetLang    string
	FooterText    string
	FooterPrivacy string
	FooterTerms   string
	Features      []FeatureData
}

type FeatureData struct {
	Emoji string
	Title string
	Desc  string
}
