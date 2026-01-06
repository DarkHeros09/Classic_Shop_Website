package utils

import "cshop-website/model"

var TermsContent = map[string]model.TermsPageData{
	"en": {
		MetaTitle:   "Terms of Use - Classic Fashion",
		MetaDesc:    "Read the Terms of Service for using the Classic app for shopping women's fashion.",
		Title:       "Terms of Service",
		LastUpdated: "Last updated: January 4, 2026",
		Intro:       "Welcome to Classic. By downloading our app or visiting our website, you agree to be bound by the following terms and conditions. Classic is a platform dedicated to providing high-quality women's fashion. Please read these terms carefully before placing an order.",
		Sections: []model.TermsSection{
			{
				Heading: "1. Account Registration",
				Content: "To access certain features (like Wishlist or Order Tracking), you may need to register an account. You agree to provide accurate, current, and complete information. You are solely responsible for protecting your account password and for all activities that occur under your account.",
			},
			{
				Heading: "2. Purchases and Payment",
				Content: "By placing an order on Classic, you confirm that you are legally capable of entering into binding contracts.",
				List: []string{
					"All prices are shown in the local currency and include applicable taxes unless stated otherwise.",
					"We reserve the right to refuse or cancel your order at any time for reasons including product availability, errors in the price, or payment issues.",
					"Payment is securely processed via third-party gateways. We do not store your full card details.",
				},
			},
			{
				Heading: "3. Shipping and Delivery",
				Content: "We aim to deliver your items within the estimated timeframes. However, delays may occur due to logistics partners or customs processing.",
				List: []string{
					"Shipping costs are calculated at checkout based on your location and package weight.",
					"Risk of loss and title for items purchased pass to you upon delivery of the items to the carrier.",
				},
			},
			{
				Heading: "4. Returns and Refunds",
				Content: "We want you to love your Classic purchase. If you are not satisfied, you may return items under the following conditions:",
				List: []string{
					"Items must be returned within 14 days of receipt.",
					"Items must be unworn, unwashed, and with original tags attached.",
					"Intimate apparel (lingerie, swimwear) is non-returnable for hygiene reasons.",
					"Refunds will be processed to the original payment method within 7-10 business days after inspection.",
				},
			},
			{
				Heading: "5. Intellectual Property",
				Content: "The 'Classic' brand, logo, and all content (images, text, designs) on the app are the exclusive property of Classic Fashion Ltd. You may not use, reproduce, or distribute our content without written permission.",
			},
			{
				Heading: "6. Prohibited Activities",
				Content: "You agree not to engage in any of the following prohibited activities:",
				List: []string{
					"Using the app for any illegal purpose or to solicit others to perform or participate in any unlawful acts.",
					"Attempting to decompile, reverse engineer, or hack the application.",
					"Harassing, abusing, or harming another person via our reviews or community features.",
				},
			},
			{
				Heading: "7. Limitation of Liability",
				Content: "Classic shall not be liable for any indirect, incidental, or consequential damages arising from your use of the service or any products procured using the service.",
			},
		},
		Contact: model.ContactSection{
			Heading: "Contact Support",
			Text:    "If you have questions regarding these Terms, please reach out:",
			Email:   "legal@classic-app.com",
		},
	},
	"ar": {
		MetaTitle:   "شروط الاستخدام - كلاسيك للأزياء",
		MetaDesc:    "اقرأ شروط الخدمة لاستخدام تطبيق كلاسيك لتسوق أزياء النساء.",
		Title:       "شروط الخدمة",
		LastUpdated: "آخر تحديث: 4 يناير 2026",
		Intro:       "مرحبًا بك في كلاسيك. عن طريق تنزيل تطبيقنا أو زيارة موقعنا، فإنك توافق على الالتزام بالشروط والأحكام التالية. كلاسيك هي منصة مخصصة لتقديم أزياء نسائية عالية الجودة. يرجى قراءة هذه الشروط بعناية قبل تقديم الطلب.",
		Sections: []model.TermsSection{
			{
				Heading: "1. تسجيل الحساب",
				Content: "للوصول إلى ميزات معينة (مثل قائمة الرغبات أو تتبع الطلبات)، قد تحتاج إلى تسجيل حساب. أنت توافق على تقديم معلومات دقيقة وحديثة وكاملة. أنت مسؤول وحدك عن حماية كلمة مرور حسابك وعن جميع الأنشطة التي تحدث تحت حسابك.",
			},
			{
				Heading: "2. المشتريات والدفع",
				Content: "من خلال تقديم طلب على كلاسيك، فإنك تؤكد أنك مؤهل قانونًا لإبرام عقود ملزمة.",
				List: []string{
					"يتم عرض جميع الأسعار بالعملة المحلية وتشمل الضرائب المطبقة ما لم يُنص على خلاف ذلك.",
					"نحتفظ بالحق في رفض أو إلغاء طلبك في أي وقت لأسباب تشمل توفر المنتج، أو أخطاء في السعر، أو مشاكل الدفع.",
					"تتم معالجة الدفع بأمان عبر بوابات طرف ثالث. نحن لا نقوم بتخزين تفاصيل بطاقتك الكاملة.",
				},
			},
			{
				Heading: "3. الشحن والتوصيل",
				Content: "نحن نهدف إلى توصيل منتجاتك خلال الأطر الزمنية المقدرة. ومع ذلك، قد يحدث تأخير بسبب شركاء الخدمات اللوجستية أو إجراءات الجمارك.",
				List: []string{
					"يتم احتساب تكاليف الشحن عند الدفع بناءً على موقعك ووزن الطرد.",
					"تنتقل مخاطر الخسارة وملكية العناصر المشتراة إليك عند تسليم العناصر إلى شركة الشحن.",
				},
			},
			{
				Heading: "4. الإرجاع والاسترداد",
				Content: "نريدك أن تحب مشترياتك من كلاسيك. إذا لم تكن راضيًا، يمكنك إرجاع العناصر وفقًا للشروط التالية:",
				List: []string{
					"يجب إرجاع العناصر في غضون 14 يومًا من الاستلام.",
					"يجب أن تكون العناصر غير ملبوسة، وغير مغسولة، ومع العلامات الأصلية مرفقة.",
					"الملابس الداخلية وملابس السباحة غير قابلة للإرجاع لأسباب صحية.",
					"ستتم معالجة عمليات الاسترداد إلى طريقة الدفع الأصلية في غضون 7-10 أيام عمل بعد الفحص.",
				},
			},
			{
				Heading: "5. الملكية الفكرية",
				Content: "العلامة التجارية 'كلاسيك'، والشعار، وجميع المحتويات (الصور، النصوص، التصاميم) على التطبيق هي ملكية حصرية لشركة كلاسيك للأزياء. لا يجوز لك استخدام أو إعادة إنتاج أو توزيع المحتوى الخاص بنا دون إذن كتابي.",
			},
			{
				Heading: "6. الأنشطة المحظورة",
				Content: "أنت توافق على عدم المشاركة في أي من الأنشطة المحظورة التالية:",
				List: []string{
					"استخدام التطبيق لأي غرض غير قانوني أو لتحريض الآخرين على القيام بأي أعمال غير قانونية.",
					"محاولة فك تشفير أو عكس هندسة أو اختراق التطبيق.",
					"مضايقة أو إساءة أو إيذاء شخص آخر عبر مراجعاتنا أو ميزات المجتمع.",
				},
			},
			{
				Heading: "7. حدود المسؤولية",
				Content: "لن تكون كلاسيك مسؤولة عن أي أضرار غير مباشرة أو عرضية أو تبعية تنشأ عن استخدامك للخدمة أو أي منتجات تم شراؤها باستخدام الخدمة.",
			},
		},
		Contact: model.ContactSection{
			Heading: "دعم العملاء",
			Text:    "إذا كانت لديك أسئلة بخصوص هذه الشروط، يرجى التواصل معنا:",
			Email:   "legal@classic-app.com",
		},
	},
}
