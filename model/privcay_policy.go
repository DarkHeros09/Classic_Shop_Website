package model

type PrivacyPageData struct {
	MetaTitle   string
	MetaDesc    string
	Title       string
	LastUpdated string
	Intro       string
	Sections    []PrivacySection
	Contact     ContactSection
}

type PrivacySection struct {
	Heading string
	Content string
	List    []string
}

type ContactSection struct {
	Heading string
	Text    string
	Email   string
}
