package utils

import "cshop-website/model"

// Translations map for easy management
var Translations = map[string]model.Content{
	"ar": {
		Title:       "كلاسيك",
		Description: "نصائح برمجية صغيرة ومفيدة لمطوري فلاتر ودارت. تصفح أكثر من 250 نصيحة برمجية بكل سهولة.",
		DownloadOn:  "حمل من",
		AppStore:    "متجر التطبيقات",
		GetItOn:     "متوفر على",
		GooglePlay:  "جوجل بلاي",
		LangToggle:  "English",
		TargetLang:  "en",
		Features: []model.FeatureData{
			{
				Emoji: "🚀",
				Title: "+250 نصيحة",
				Desc:  "تصفح أكثر من 250 نصيحة وحيلة برمجية حول تطوير التطبيقات، يتم تحديثها يومياً.",
			},
			{
				Emoji: "💾",
				Title: "بدون اتصال",
				Desc:  "يتم حفظ النصائح المحملة محلياً لتتمكن من الوصول إليها في أي مكان وفي أي وقت.",
			},
			{
				Emoji: "🔍",
				Title: "ميزة البحث",
				Desc:  "ابحث عن نصائح محددة فوراً أو دعنا نختار لك نصيحة عشوائية لتتعلم شيئاً جديداً.",
			},
			{
				Emoji: "🌗",
				Title: "الوضع الليلي",
				Desc:  "يتكيف مظهر التطبيق تلقائياً مع نظام هاتفك لتوفير تجربة قراءة مريحة للعين.",
			},
		},
		FooterText:    "© {{year}} متجر كلاسيك. جميع الحقوق محفوظة.",
		FooterPrivacy: "سياسة الخصوصية",
		FooterTerms:   "شروط الاستخدام",
	},
	"en": {
		Title:       "Classic",
		Description: "Bite-sized tips and tricks about Dart and Flutter development. Browse over 250 tips easily.",
		DownloadOn:  "Download on",
		AppStore:    "App Store",
		GetItOn:     "Get it on",
		GooglePlay:  "Google Play",
		LangToggle:  "العربية",
		TargetLang:  "ar",
		Features: []model.FeatureData{
			{Emoji: "🚀", Title: "250+ Tips", Desc: "Browse over 250 tips and tricks about development, curated daily."},
			{Emoji: "💾", Title: "Offline Mode", Desc: "Downloaded tips are saved locally so you can access them anywhere."},
			{Emoji: "🔍", Title: "Search Feature", Desc: "Search existing tips instantly or let us choose a random tip for you."},
			{Emoji: "🌗", Title: "Dark Mode", Desc: "Automatically adjusts to your system theme for comfortable reading."},
		},
		FooterText:    "© {{year}} Classic Shop. All rights reserved.",
		FooterPrivacy: "Privacy Policy",
		FooterTerms:   "Terms of Use",
	},
}
