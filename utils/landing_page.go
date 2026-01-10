package utils

import "cshop-website/model"

// LandingPage map for easy management
var LandingPage = map[string]model.Content{
	"ar": {
		Title:       "كلاسيك",
		Description: "أناقة بلا مجهود، تصلكِ أينما كنتِ. اكتشفي مجموعتنا المختارة من الأزياء النسائية الراقية واستمتعي بتجربة تسوق ذكية تجمع بين الجمال والسهولة.",
		LangToggle:  "English",
		TargetLang:  "en",
		Features: []model.FeatureData{
			{
				Emoji: "👗",
				Title: "مجموعات مختارة",
				Desc:  "تصفحي تشكيلات مميزة من الملابس المختارة بعناية لتناسب جميع الأذواق والمناسبات.",
			},
			{
				Emoji: "🔍",
				Title: "ميزة البحث",
				Desc:  "اعثري على إطلالتك القادمة فوراً باستخدام نظام البحث المتقدم والفلاتر الذكية.",
			},
			{
				Emoji: "💳",
				Title: "دفع آمن",
				Desc:  "تسوقي بكل ثقة مع بوابات دفع مشفرة وآمنة تماماً تدعم جميع البطاقات والمحافظ الرقمية.",
			},
			{
				Emoji: "⚡",
				Title: "توصيل سريع",
				Desc:  "لا داعي للانتظار الطويل، نوفر لكِ خدمة توصيل سريعة وموثوقة حتى باب منزلك.",
			},
			{
				Emoji: "💾",
				Title: "بدون اتصال",
				Desc:  "يمكنكِ تصفح المنتجات المحملة مسبقاً والوصول إلى قائمة أمنياتك في أي وقت حتى بدون إنترنت.",
			},
			{
				Emoji: "🌗",
				Title: "الوضع الليلي",
				Desc:  "واجهة مريحة للعين تتكيف تلقائياً مع نظام هاتفك لتجربة تسوق ليلية هادئة.",
			},
		},
		FooterText:    "© {{year}} متجر كلاسيك. جميع الحقوق محفوظة.",
		FooterPrivacy: "سياسة الخصوصية",
		FooterTerms:   "شروط الاستخدام",
	},
	"en": {
		Title:       "Classic",
		Description: "Effortless Elegance, Delivered. Browse our curated collection of premium women's fashion and enjoy a smart, seamless shopping experience.",
		LangToggle:  "العربية",
		TargetLang:  "ar",
		Features: []model.FeatureData{
			{
				Emoji: "👗",
				Title: "Curated Collection",
				Desc:  "Browse hand-picked selections of outfits designed to suit every style and occasion.",
			},
			{
				Emoji: "🔍",
				Title: "Search Feature",
				Desc:  "Find your next look instantly with our advanced search and smart filtering system.",
			},
			{
				Emoji: "💳",
				Title: "Secure Payment",
				Desc:  "Shop with total confidence using fully encrypted gateways supporting all major cards.",
			},
			{
				Emoji: "⚡",
				Title: "Fast Delivery",
				Desc:  "No more long waits. We provide fast and reliable shipping straight to your doorstep.",
			},
			{
				Emoji: "💾",
				Title: "Offline Mode",
				Desc:  "Access pre-loaded products and your wishlist anywhere, even without an internet connection.",
			},
			{
				Emoji: "🌗",
				Title: "Dark Mode",
				Desc:  "Automatically adjusts to your system theme for a comfortable late-night shopping experience.",
			},
		},
		FooterText:    "© {{year}} Classic Shop. All rights reserved.",
		FooterPrivacy: "Privacy Policy",
		FooterTerms:   "Terms of Use",
	},
}
