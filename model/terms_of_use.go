package model

// TermsPageData represents the structure of the Terms of Use page
type TermsPageData struct {
	MetaTitle   string
	MetaDesc    string
	Title       string
	LastUpdated string
	Intro       string
	Sections    []TermsSection
	Contact     ContactSection
}

// TermsSection handles each legal clause
type TermsSection struct {
	Heading string
	Content string
	// Optional list for things like "Prohibited Activities" or "Return Conditions"
	List []string
}

// Re-using the contact structure is fine, or define it here if models are separate
type TermsContactSection struct {
	Heading string
	Text    string
	Email   string
}
