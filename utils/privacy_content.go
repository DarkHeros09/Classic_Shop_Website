package utils

import "cshop-website/model"

var PrivacyContent = map[string]model.PrivacyPageData{
	"en": {
		MetaTitle:   "Privacy Policy - TipsApp Fashion",
		MetaDesc:    "Learn how TipsApp Fashion collects, uses, and protects your personal shopping data.",
		Title:       "Privacy Policy",
		LastUpdated: "Last updated: January 3, 2026",
		Intro:       "Welcome to TipsApp Fashion. We value your trust and are committed to protecting your personal information. This Privacy Policy explains how we collect, use, and share your data when you shop for clothing and accessories through our mobile application.",
		Sections: []model.PrivacySection{
			{
				Heading: "1. Information We Collect",
				Content: "To fulfill your orders and provide a personalized shopping experience, we collect various types of information:",
				List: []string{
					"Personal Identity: Name, email address, and phone number for account management.",
					"Shipping Data: Physical address and delivery instructions for courier services.",
					"Payment Information: Credit card tokens (processed securely by third-party gateways).",
					"Fashion Preferences: Clothing sizes, style preferences, and wishlist items.",
				},
			},
			{
				Heading: "2. How We Use Your Information",
				Content: "We use the collected data primarily to process your transactions and improve our store:",
				List: []string{
					"Processing and delivering your orders via logistics partners.",
					"Sending order confirmations, shipping updates, and invoices.",
					"Processing returns and refunds.",
				},
			},
			{
				Heading: "3. Sharing Your Data",
				Content: "We do not sell your personal data. However, we share necessary data with trusted third parties:",
				List: []string{
					"Logistics Partners: We share your name and address with couriers (e.g., DHL, Aramex).",
					"Payment Processors: Financial data is securely handled by Stripe/PayPal.",
				},
			},
			{
				Heading: "4. Cookies & Tracking Technologies",
				Content: "We use cookies to keep items in your shopping cart and analyze site traffic. You can control cookies through your browser settings, but some store features may not function properly without them.",
			},
			{
				Heading: "5. Data Retention",
				Content: "We retain your personal information only for as long as is necessary. We will retain and use your Order History for internal analysis purposes and to comply with tax laws.",
			},
			{
				Heading: "6. Security of Data",
				Content: "We implement industry-standard security measures, including SSL encryption for all data in transit. We tokenize payment data so your full card details never touch our servers.",
			},
			{
				Heading: "7. Your Rights",
				Content: "Depending on your location, you have the right to access, correct, or delete your personal data. You may also opt-out of marketing communications at any time via the 'Unsubscribe' link in our emails.",
			},
		},
		Contact: model.ContactSection{
			Heading: "Contact Us",
			Text:    "If you have questions about your order data or this policy:",
			Email:   "privacy@tipsapp.com",
		},
	},
	"ar": {
		MetaTitle:   "سياسة الخصوصية - تيبس آب للأزياء",
		MetaDesc:    "تعرف على كيفية جمع واستخدام وحماية بيانات التسوق الخاصة بك في تيبس آب.",
		Title:       "سياسة الخصوصية",
		LastUpdated: "آخر تحديث: 3 يناير 2026",
		Intro:       "مرحبًا بك في تيبس آب للأزياء. نحن نقدر ثقتك ونلتزم بحماية معلوماتك الشخصية. تشرح سياسة الخصوصية هذه كيفية جمع بياناتك واستخدامها ومشاركتها عند التسوق لشراء الملابس والإكسسوارات.",
		Sections: []model.PrivacySection{
			{
				Heading: "1. المعلومات التي نجمعها", // English Numbering
				Content: "لتلبية طلباتك وتوفير تجربة تسوق مخصصة، نقوم بجمع أنواع مختلفة من المعلومات:",
				List: []string{
					"الهوية الشخصية: الاسم وعنوان البريد الإلكتروني ورقم الهاتف.",
					"بيانات الشحن: العنوان الفعلي وتعليمات التوصيل.",
					"معلومات الدفع: رموز بطاقة الائتمان (تتم معالجتها بأمان بواسطة بوابات طرف ثالث).",
					"تفضيلات الموضة: مقاسات الملابس وتفضيلات الأسلوب.",
				},
			},
			{
				Heading: "2. كيف نستخدم معلوماتك",
				Content: "نستخدم البيانات التي تم جمعها بشكل أساسي لمعالجة معاملاتك وتحسين متجرنا:",
				List: []string{
					"معالجة وتسليم طلباتك عبر شركاء الخدمات اللوجستية.",
					"إرسال تأكيدات الطلب وتحديثات الشحن والفواتير.",
					"معالجة عمليات الإرجاع واسترداد الأموال.",
				},
			},
			{
				Heading: "3. مشاركة بياناتك",
				Content: "نحن لا نبيع بياناتك الشخصية. ومع ذلك، نشارك البيانات الضرورية مع أطراف ثالثة موثوقة:",
				List: []string{
					"شركاء الخدمات اللوجستية: نشارك اسمك وعنوانك مع شركات الشحن.",
					"معالجو الدفع: يتم التعامل مع البيانات المالية بأمان بواسطة Stripe/PayPal.",
				},
			},
			{
				Heading: "4. ملفات تعريف الارتباط وتقنيات التتبع",
				Content: "نستخدم ملفات تعريف الارتباط للاحتفاظ بالعناصر في عربة التسوق الخاصة بك وتحليل حركة المرور. يمكنك التحكم في ملفات تعريف الارتباط من خلال إعدادات المتصفح.",
			},
			{
				Heading: "5. الاحتفاظ بالبيانات",
				Content: "نحتفظ بمعلوماتك الشخصية فقط طالما كان ذلك ضروريًا. سنحتفظ بسجل الطلبات لأغراض التحليل الداخلي والامتثال لقوانين الضرائب.",
			},
			{
				Heading: "6. أمن البيانات",
				Content: "نحن نطبق تدابير أمنية متوافقة مع معايير الصناعة، بما في ذلك تشفير SSL. نحن نقوم بترميز بيانات الدفع حتى لا تلمس تفاصيل بطاقتك الكاملة خوادمنا.",
			},
			{
				Heading: "7. حقوقك",
				Content: "بناءً على موقعك، لديك الحق في الوصول إلى بياناتك الشخصية أو تصحيحها أو حذفها. يمكنك أيضًا إلغاء الاشتراك في الاتصالات التسويقية في أي وقت.",
			},
		},
		Contact: model.ContactSection{
			Heading: "اتصل بنا",
			Text:    "إذا كانت لديك أسئلة حول بيانات طلبك أو هذه السياسة:",
			Email:   "privacy@tipsapp.com",
		},
	},
}
