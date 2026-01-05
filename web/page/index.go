package page

import (
	"github.com/LuuDinhTheTai/tzone/web/page/shared"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func HomePage() g.Node {
	return shared.BaseLayout("GSMArena.com - mobile phone reviews, news, specifications and more...",
		h.Div(h.ID("wrapper"), h.Class("l-container"),
			h.Div(h.ID("outer"), h.Class("row"),
				//SubHeader(),
				h.Div(h.ID("subHeader3")),
				h.Div(h.ID("body"),
					MainContent(),
					Sidebar(),
				),
			),
		),
	)
}

func SubHeader() g.Node {
	return h.Div(h.ID("subHeader"), h.Class("col")) // Removed Top Advertisement

}

func MainContent() g.Node {
	return h.Div(h.Class("main float-right has-trending"),
		ReviewSection(),
		FeaturedSection(),
		BuyersGuideNotification(),
		NewsColumn(),
	)
}

func ReviewSection() g.Node {
	return h.Div(h.Class("review-column-index"), h.ID("review-sample"), h.Style("position:relative;"),
		h.Div(h.Class("review-column-list"),
			h.Article(h.Class("review-column-list-item review-column-list-item-main "),
				h.A(h.Href("best_sounding_smartphones_and_tablets_tested_2025-review-2917.php"), h.Class("block-link"),
					h.Header(
						h.H3(g.Text("Best-sounding phones and tablets we tested in 2025")),
					),
					h.Img(h.Class("review-list-item-image"), h.Src("https://fdn.gsmarena.com/imgroot/reviews/25/best-smartphone-speakers-2025/-728x314/gsmarena_201.jpg"), h.Alt("Best-sounding phones and tablets we tested in 2025")),
				),
				h.Footer(h.Class("meta-line has-border"),
					h.Time(h.Class("meta-item-time float-left"),
						h.I(h.Class("head-icon icon-clock2")), g.Text("02 Jan 2026"),
					),
					h.A(h.Class("meta-item-comments float-right"), h.Href("reviewcomm-2917.php"),
						h.I(h.Class("head-icon icon-comments")), g.Text("61"),
					),
				),
			),

			h.Div(h.Class("review-column-list-item review-column-list-item-secondary"),
				h.Article(
					h.A(h.Href("huawei_mate_x7-review-2914.php"), h.Class("block-link"),
						h.Header(
							h.H3(g.Text("Huawei Mate X7 review")),
						),
						h.Img(h.Class("review-list-item-image"), h.Src("https://fdn.gsmarena.com/imgroot/reviews/25/huawei-mate-x7/-347x151/gsmarena_001.jpg"), h.Alt("Huawei Mate X7 review")),
					),
					h.Footer(h.Class("meta-line has-border clearfix"),
						h.A(h.Class("meta-item-comments float-right"), h.Href("reviewcomm-2914.php"),
							h.I(h.Class("head-icon icon-comments")), g.Text("38"),
						),
					),
				),
				h.Article(
					h.A(h.Href("xiaomi_redmi_note_15_pro_5g-review-2915.php"), h.Class("block-link"),
						h.Header(
							h.H3(g.Text("Xiaomi Redmi Note 15 Pro 5G review")),
						),
						h.Img(h.Class("review-list-item-image"), h.Src("https://fdn.gsmarena.com/imgroot/reviews/25/redmi-note-15-pro/-347x151/gsmarena_001.jpg"), h.Alt("Xiaomi Redmi Note 15 Pro 5G review")),
					),
					h.Footer(h.Class("meta-line"),
						h.A(h.Class("meta-item-comments float-right"), h.Href("reviewcomm-2915.php"),
							h.I(h.Class("head-icon icon-comments")), g.Text("40"),
						),
					),
				),
				// MORE REVIEWS TOGGLE
				h.Div(h.Class("more-reviews-list-container"),
					h.H4(g.Text("More reviews "),
						h.A(h.Href("reviews.php3"), h.Class("right"), g.Text("View all")),
					),
					h.Ul(h.Class("more-reviews-list"),
						h.Li(h.A(h.Href("redmi_note_pro_plus-review-2913.php"), h.Class("block-link"), g.Text("Xiaomi Redmi Note 15 Pro+ review"))),
						h.Li(h.A(h.Href("oneplus_15r_handson-review-2911.php"), h.Class("block-link"), g.Text("OnePlus 15R hands-on review"))),
						h.Li(h.A(h.Href("huawei_mate_x7_handson-review-2910.php"), h.Class("block-link"), g.Text("Huawei Mate X7 hands-on review"))),
						h.Li(h.A(h.Href("poco_f8_ultra-review-2906.php"), h.Class("block-link"), g.Text("Poco F8 Ultra review"))),
						h.Li(h.A(h.Href("poco_f8_pro-review-2907.php"), h.Class("block-link"), g.Text("Poco F8 Pro review"))),
						h.Li(h.A(h.Href("honor_magic8_lite-review-2908.php"), h.Class("block-link"), g.Text("Honor Magic8 Lite review"))),
					),
				),
				h.A(h.Href("#"), h.Class("more-reviews-list-toggle block-link"), g.Raw("&#9650;")),
			),
		),
	)
}

func FeaturedSection() g.Node {
	return h.Div(h.Class("feat"),
		h.Section(h.Class("feat-container flex-container"),
			h.Div(h.Class("flex-column flex-size-2"),
				h.Article(h.Class("feat-item flex-item"),
					h.Div(h.Class("meta-line"),
						h.Span(h.Class("meta-item-time"), h.I(h.Class("head-icon icon-clock2")), g.Text("17 hours ago")),
						h.A(h.Class("meta-item-comments"), h.Href("newscomm-70895.php"), h.I(h.Class("head-icon icon-comments")), g.Text("10")),
					),
					h.Header(h.Class("feat-item-header"), h.H3(h.Class("feat-item-title"), g.Text("Huawei FreeBuds 7i review"))),
					h.Span(h.Class("feat-item-image"), h.Style("background-image:url('https://fdn.gsmarena.com/imgroot/news/25/12/huawei-freebuds-7i-review/-728x314/gsmarena_000.jpg');")),
					h.A(h.Href("huawei_freebuds_7i_review-news-70895.php"), h.Class("feat-item-link link-faux")),
				),
			),
			h.Div(h.Class("flex-column"),
				h.Article(h.Class("feat-item flex-item"),
					h.Div(h.Class("meta-line"),
						h.Span(h.Class("meta-item-time"), h.I(h.Class("head-icon icon-clock2")), g.Text("04 Jan 2026")),
						h.A(h.Class("meta-item-comments"), h.Href("newscomm-70844.php"), h.I(h.Class("head-icon icon-comments")), g.Text("78")),
					),
					h.Header(h.Class("feat-item-header"), h.H3(h.Class("feat-item-title"), g.Text("2025 Winners and losers: OnePlus"))),
					h.Span(h.Class("feat-item-image"), h.Style("background-image:url('https://fdn.gsmarena.com/imgroot/news/25/12/oneplus-2025/-344x215/gsmarena_000.jpg');")),
					h.A(h.Href("2025_winners_and_losers_oneplus-news-70844.php"), h.Class("feat-item-link link-faux")),
				),
				h.Article(h.Class("feat-item flex-item"),
					h.Div(h.Class("meta-line"),
						h.Span(h.Class("meta-item-time"), h.I(h.Class("head-icon icon-clock2")), g.Text("03 Jan 2026")),
						h.A(h.Class("meta-item-comments"), h.Href("newscomm-70831.php"), h.I(h.Class("head-icon icon-comments")), g.Text("62")),
					),
					h.Header(h.Class("feat-item-header"), h.H3(h.Class("feat-item-title"), g.Text("Top phones of 2025: Camera"))),
					h.Span(h.Class("feat-item-image"), h.Style("background-image:url('https://fdn.gsmarena.com/imgroot/news/25/12/top-phones-2025-camera/-344x215/gsmarena_000.jpg');")),
					h.A(h.Href("top_phones_of_2025_camera-news-70831.php"), h.Class("feat-item-link link-faux")),
				),
			),
		),
	)
}

func BuyersGuideNotification() g.Node {
	return h.A(h.Class("buyers-guide-notification"), h.Href("smartphone_buyers_guide-review-2036.php"), h.Style("border-left-color: #17819f;"),
		h.H3(g.Text("New here? In a hurry?")),
		h.Span(g.Text("See the best phones right now in our all-new "), h.U(g.Text("buyer's guide"))),
		h.I(h.Class("notification"), g.Raw("&#xe962;")),
	)
}

func NewsColumn() g.Node {
	return h.Div(h.Class("news-column-index"),
		h.Div(h.Class("floating-title"),
			NewsItem("realme_neo8_to_debut_in_january_with_snapdragon_8_gen_5_soc-news-70944.php", "https://fdn.gsmarena.com/imgroot/news/26/01/realme-neo8-launch-design-specs/-344x215/gsmarena_000.jpg", "Realme Neo8 to debut in January with Snapdragon 8 Gen 5 SoC", "The phone will be available in a Cyber Purple color with an RGB light at the back.", "36 minutes ago", "newscomm-70944.php", "1"),
			PostItem("asus_rog_xbox_ally_x_review-news-70943.php", "https://fdn.gsmarena.com/imgroot/news/26/01/rog-xbox-ally-x/-728x314/gsmarena_000.jpg", "Asus ROG Xbox Ally X review", "8 hours ago", "newscomm-70943.php", "9"),
			NewsItem("huawei_freebuds_7i_review-news-70895.php", "https://fdn.gsmarena.com/imgroot/news/25/12/huawei-freebuds-7i-review/-344x215/gsmarena_000.jpg", "Huawei FreeBuds 7i review", "Huawei's do-it-all TWS earbuds bring on a rich set of features without breaking the bank.", "17 hours ago", "newscomm-70895.php", "10"),
			NewsItem("top_10_trending_phones_of_week_1-news-70942.php", "https://fdn.gsmarena.com/imgroot/news/19/04/top10-trending-phones/-344x215/gsmarena_003.jpg", "Top 10 trending phones of week 1", "New year, new leader.", "21 hours ago", "newscomm-70942.php", "8"),
			PostItem("2025_winners_and_losers_oneplus-news-70844.php", "https://fdn.gsmarena.com/imgroot/news/25/12/oneplus-2025/-728x314/gsmarena_000.jpg", "2025 Winners and losers: OnePlus", "04 Jan 2026", "newscomm-70844.php", "78"),
			// Removed Advertisement
			NewsItem("oppo_find_x9_ultras_camera_sensors_get_detailed_in_new_leak-news-70938.php", "https://fdn.gsmarena.com/imgroot/news/26/01/oppo-find-x9-ultra-camera-sensors/-344x215/gsmarena_000.jpg", "Oppo Find X9 Ultra's camera sensors get detailed in new leak", "It's an intriguing setup, that's for sure.", "04 Jan 2026", "newscomm-70938.php", "61"),
			PostItem("top_phones_of_2025_camera-news-70831.php", "https://fdn.gsmarena.com/imgroot/news/25/12/top-phones-2025-camera/-728x314/gsmarena_000.jpg", "Top phones of 2025: Camera", "03 Jan 2026", "newscomm-70831.php", "62"),
			NewsItem("pebble_round_2_is_official_with_epaper_touchscreen_twoweek_battery_life-news-70940.php", "https://fdn.gsmarena.com/imgroot/news/26/01/pebble-round-2/-344x215/gsmarena_000.jpg", "Pebble Round 2 is official with e-paper touchscreen, two-week battery life", "Pre-orders have already started in three colors.", "03 Jan 2026", "newscomm-70940.php", "9"),
			PostItem("my_top_phones_of_2025_sagar-news-70929.php", "https://fdn.gsmarena.com/imgroot/news/26/01/my-top-4-phones-2025-sagar/-728x314/gsmarena_001.jpg", "My top 4 phones of 2025 - Sagar", "03 Jan 2026", "newscomm-70929.php", "18"),
			NewsItem("oneplus_showcases_the_turbo_6_in_live_images-news-70941.php", "https://fdn.gsmarena.com/imgroot/news/26/01/oneplus-turbo-6-live-images/-344x215/gsmarena_000.jpg", "OnePlus showcases the Turbo 6 in live images", "Two of its upcoming colorways are featured.", "03 Jan 2026", "newscomm-70941.php", "18"),
			PostItem("best_tablets_of_2025-news-70836.php", "https://fdn.gsmarena.com/imgroot/news/25/12/best-tablets-2025/-728x314/gsmarena_000.jpg", "The best tablets of 2025", "03 Jan 2026", "newscomm-70836.php", "47"),
			NewsItem("there_will_be_no_new_asus_phones_this_year_new_report_claims-news-70939.php", "https://fdn.gsmarena.com/imgroot/news/26/01/asus-no-smartphone-launches-this-year/-344x215/gsmarena_000.jpg", "There will be no new Asus phones this year, report claims", "The info allegedly comes from Asus itself.", "03 Jan 2026", "newscomm-70939.php", "119"),
			NewsItem("samsung_galaxy_s26_series_galaxy_z_fold8_galaxy_z_flip8_usa_price_tipped-news-70931.php", "https://fdn.gsmarena.com/imgroot/news/25/09/samsung-galaxy-s26-ultra-new-renders/-344x215/gsmarena_000.jpg", "Latest Samsung Galaxy S26 series rumor brings good news for potential buyers", "But only for those living in the US.", "03 Jan 2026", "newscomm-70931.php", "60"),
			NewsItem("poco_officially_reveals_more_m8_specs-news-70925.php", "https://fdn.gsmarena.com/imgroot/news/26/01/poco-reveals-more-m8-specs/-344x215/gsmarena_000.jpg", "Poco officially reveals more M8 specs", "It's getting fully official on January 8.", "03 Jan 2026", "newscomm-70925.php", "20"),
			PostItem("2025_winners_and_losers_vivo-news-70819.php", "https://fdn.gsmarena.com/imgroot/news/25/12/vivo-winners-and-losers/-728x314/gsmarena_000.jpg", "2025 Winners and losers: vivo", "03 Jan 2026", "newscomm-70819.php", "28"),
			NewsItem("motorola_teases_special_world_cup_razr_heres_when_its_launching-news-70937.php", "https://fdn.gsmarena.com/imgroot/news/26/01/motorola-razr-fifa-world-cup-2026-teaser-launch-date/-344x215/gsmarena_001.jpg", "Motorola teases special World Cup Razr, here's when it's launching", "Only a few days left to wait now.", "03 Jan 2026", "newscomm-70937.php", "1"),
			// Removed Advertisement2
			NewsItem("mediatek_dimensity_7100_soc_is_official-news-70936.php", "https://fdn.gsmarena.com/imgroot/news/26/01/mediatek-dimensity-7100-official/-344x215/gsmarena_000.jpg", "MediaTek Dimensity 7100 SoC is official", "It will be featured in upcoming devices from Infinix and Tecno.", "02 Jan 2026", "newscomm-70936.php", "56"),
			NewsItem("samsung_unveils_freestyle_portable_projector-news-70935.php", "https://fdn.gsmarena.com/imgroot/news/26/01/samsung-freestyle-plus-announced/-344x215/gsmarena_000.jpg", "Samsung unveils Freestyle+ portable projector", "It's going to be released before the end of June.", "02 Jan 2026", "newscomm-70935.php", "5"),
			NewsItem("motorola_x70_air_pro_passes_through_tenaa_ahead_of_its_official_announcement-news-70927.php", "https://fdn.gsmarena.com/imgroot/news/26/01/motorola-x70-air-pro-tenaa/-344x215/gsmarena_000.jpg", "Motorola X70 Air Pro passes through TENAA ahead of its official announcement", "This could be the phone launching internationally as the Signature.", "02 Jan 2026", "newscomm-70927.php", "10"),
			PostItem("my_top_phones_of_2025__ro-news-70829.php", "https://fdn.gsmarena.com/imgroot/news/25/12/top-phones-2025-ro/-728x314/gsmarena_000.jpg", "My top 4 phones of 2025: Ro", "02 Jan 2026", "newscomm-70829.php", "27"),
			h.Div(h.Class("review-nav-v2 simple")),
		),
		h.Br(h.Class("clear")),
	)
}

func NewsItem(link, imgSrc, title, desc, time, commentsLink, commentsCount string) g.Node {
	return h.Div(h.Class("news-item"),
		h.A(h.Href(link),
			h.Div(h.Class("news-item-media-wrap left"),
				h.Img(h.Src(imgSrc), h.Alt(title)),
			),
			h.H3(g.Text(title)),
			h.P(g.Text(desc)),
		),
		h.Div(h.Class("meta-line"),
			h.Span(h.Class("meta-item-time"), h.I(h.Class("head-icon icon-clock2")), g.Text(time)),
			h.A(h.Class("meta-item-comments"), h.Href(commentsLink), h.I(h.Class("head-icon icon-comments")), g.Text(commentsCount)),
		),
	)
}

func PostItem(link, imgSrc, title, time, commentsLink, commentsCount string) g.Node {
	return h.Div(h.Class("post"),
		h.Article(h.Class("feat-item flex-item"),
			h.Div(h.Class("meta-line"),
				h.Span(h.Class("meta-item-time"), h.I(h.Class("head-icon icon-clock2")), g.Text(time)),
				h.A(h.Class("meta-item-comments"), h.Href(commentsLink), h.I(h.Class("head-icon icon-comments")), g.Text(commentsCount)),
			),
			h.Header(h.Class("feat-item-header"),
				h.H3(h.Class("feat-item-title"), g.Text(title)), h.Br(),
			),
			h.Span(h.Class("feat-item-image"), h.Style("background-image:url('"+imgSrc+"');")),
			h.A(h.Href(link), h.Class("feat-item-link link-faux")),
		),
	)
}

func Sidebar() g.Node {
	return h.Aside(h.Class("sidebar col left"),
		BrandMenu(),
		// Removed SidebarAdvertisement
		LatestDevices(),
		InStoresNow(),
		Top10DailyInterest(),
		Top10ByFans(),
		PopularComparisons(),
		ElectricVehicles(),
		// Removed SidebarStickyAdvertisement
	)
}

func BrandMenu() g.Node {
	return h.Div(h.Class("brandmenu-v2 light l-box clearfix"),
		h.P(h.Class("pad"),
			h.A(h.Href("search.php3"), h.Class("pad-single pad-finder"),
				h.I(h.Class("head-icon icon-search-right")),
				h.Span(g.Text("Phone finder")),
			),
		),
		h.Ul(
			h.Li(h.A(h.Href("samsung-phones-9.php"), g.Text("Samsung"))),
			h.Li(h.A(h.Href("apple-phones-48.php"), g.Text("Apple"))),
			h.Li(h.A(h.Href("huawei-phones-58.php"), g.Text("Huawei"))),
			h.Li(h.A(h.Href("nokia-phones-1.php"), g.Text("Nokia"))),
			h.Li(h.A(h.Href("sony-phones-7.php"), g.Text("Sony"))),
			h.Li(h.A(h.Href("lg-phones-20.php"), g.Text("LG"))),
			h.Li(h.A(h.Href("htc-phones-45.php"), g.Text("HTC"))),
			h.Li(h.A(h.Href("motorola-phones-4.php"), g.Text("Motorola"))),
			h.Li(h.A(h.Href("lenovo-phones-73.php"), g.Text("Lenovo"))),
			h.Li(h.A(h.Href("xiaomi-phones-80.php"), g.Text("Xiaomi"))),
			h.Li(h.A(h.Href("google-phones-107.php"), g.Text("Google"))),
			h.Li(h.A(h.Href("honor-phones-121.php"), g.Text("Honor"))),
			h.Li(h.A(h.Href("oppo-phones-82.php"), g.Text("Oppo"))),
			h.Li(h.A(h.Href("realme-phones-118.php"), g.Text("Realme"))),
			h.Li(h.A(h.Href("oneplus-phones-95.php"), g.Text("OnePlus"))),
			h.Li(h.A(h.Href("nothing-phones-128.php"), g.Text("Nothing"))),
			h.Li(h.A(h.Href("vivo-phones-98.php"), g.Text("vivo"))),
			h.Li(h.A(h.Href("meizu-phones-74.php"), g.Text("Meizu"))),
			h.Li(h.A(h.Href("asus-phones-46.php"), g.Text("Asus"))),
			h.Li(h.A(h.Href("alcatel-phones-5.php"), g.Text("Alcatel"))),
			h.Li(h.A(h.Href("zte-phones-62.php"), g.Text("ZTE"))),
			h.Li(h.A(h.Href("rugone-phones-136.php"), g.Text("RugOne"))),
			h.Li(h.A(h.Href("umidigi-phones-135.php"), g.Text("Umidigi"))),
			h.Li(h.A(h.Href("coolpad-phones-105.php"), g.Text("Coolpad"))),
			h.Li(h.A(h.Href("oscal-phones-134.php"), g.Text("Oscal"))),
			h.Li(h.A(h.Href("sharp-phones-23.php"), g.Text("Sharp"))),
			h.Li(h.A(h.Href("micromax-phones-66.php"), g.Text("Micromax"))),
			h.Li(h.A(h.Href("infinix-phones-119.php"), g.Text("Infinix"))),
			h.Li(h.A(h.Href("ulefone_-phones-124.php"), g.Text("Ulefone"))),
			h.Li(h.A(h.Href("tecno-phones-120.php"), g.Text("Tecno"))),
			h.Li(h.A(h.Href("doogee-phones-129.php"), g.Text("Doogee"))),
			h.Li(h.A(h.Href("blackview-phones-116.php"), g.Text("Blackview"))),
			h.Li(h.A(h.Href("cubot-phones-130.php"), g.Text("Cubot"))),
			h.Li(h.A(h.Href("oukitel-phones-132.php"), g.Text("Oukitel"))),
			h.Li(h.A(h.Href("itel-phones-131.php"), g.Text("Itel"))),
			h.Li(h.A(h.Href("tcl-phones-123.php"), g.Text("TCL"))),
		),
		h.P(h.Class("pad"),
			h.A(h.Href("makers.php3"), h.Class("pad-multiple pad-allbrands"),
				h.I(h.Class("head-icon icon-mobile-phone231")),
				h.Span(g.Text("All brands")),
			),
			h.A(h.Href("rumored.php3"), h.Class("pad-multiple pad-rumormill"),
				h.I(h.Class("head-icon icon-rumored")),
				h.Span(g.Text("Rumor mill")),
			),
		),
	)
}

func LatestDevices() g.Node {
	return h.Div(h.Class("module module-phones module-latest"),
		h.H4(h.Class("section-heading"), g.Text("Latest devices")),
		h.Div(h.Style("overflow:hidden;overflow-y:auto;max-height: 390px;"),
			h.A(h.Href("oppo_reno15_pro_max_5g-14220.php"), h.Class("module-phones-link"), h.Img(h.Src("https://fdn2.gsmarena.com/vv/bigpic/oppo-reno15-pro.jpg")), h.Br(), g.Text("Oppo Reno15 Pro Max")),
			h.A(h.Href("oppo_reno15_pro_5g_(global)-14394.php"), h.Class("module-phones-link"), h.Img(h.Src("https://fdn2.gsmarena.com/vv/bigpic/oppo-reno15-pro-global.jpg")), h.Br(), g.Text("Oppo Reno15 Pro")),
			h.A(h.Href("oppo_reno15_5g_(global)-14395.php"), h.Class("module-phones-link"), h.Img(h.Src("https://fdn2.gsmarena.com/vv/bigpic/oppo-reno15-global.jpg")), h.Br(), g.Text("Oppo Reno15")),
			h.A(h.Href("oppo_reno15_f_5g-14396.php"), h.Class("module-phones-link"), h.Img(h.Src("https://fdn2.gsmarena.com/vv/bigpic/oppo-reno15-f.jpg")), h.Br(), g.Text("Oppo Reno15 F")),
			h.A(h.Href("oppo_reno15_pro_mini_5g-14393.php"), h.Class("module-phones-link"), h.Img(h.Src("https://fdn2.gsmarena.com/vv/bigpic/oppo-reno15-pro-global.jpg")), h.Br(), g.Text("Oppo Reno15 Pro Mini")),
			h.A(h.Href("xiaomi_poco_m8_pro_5g-14390.php"), h.Class("module-phones-link"), h.Img(h.Src("https://fdn2.gsmarena.com/vv/bigpic/xiaomi-poco-m8-pro.jpg")), h.Br(), g.Text("Xiaomi Poco M8 Pro")),
			h.A(h.Href("xiaomi_poco_m8_5g-14391.php"), h.Class("module-phones-link"), h.Img(h.Src("https://fdn2.gsmarena.com/vv/bigpic/xiaomi-poco-m8.jpg")), h.Br(), g.Text("Xiaomi Poco M8")),
			h.A(h.Href("xiaomi_17_ultra_5g-14380.php"), h.Class("module-phones-link"), h.Img(h.Src("https://fdn2.gsmarena.com/vv/bigpic/xiaomi-17-ultra.jpg")), h.Br(), g.Text("Xiaomi 17 Ultra")),
			h.A(h.Href("vivo_x300_ultra_5g-14388.php"), h.Class("module-phones-link"), h.Img(h.Src("https://fdn2.gsmarena.com/vv/bigpic/vivo-x200-ultra.jpg")), h.Br(), g.Text("vivo X300 Ultra")),
		),
	)
}

func InStoresNow() g.Node {
	return h.Div(h.Class("module module-phones module-latest"),
		h.H4(h.Class("section-heading"), g.Text("In stores now")),
		h.Div(h.Style("overflow:hidden;overflow-y:auto;max-height: 390px;"),
			h.A(h.Href("xiaomi_17_ultra_5g-14380.php"), h.Class("module-phones-link"), h.Img(h.Src("https://fdn2.gsmarena.com/vv/bigpic/xiaomi-17-ultra.jpg")), h.Br(), g.Text("Xiaomi 17 Ultra")),
			h.A(h.Href("xiaomi_redmi_note_15_pro+_5g_(global)-14326.php"), h.Class("module-phones-link"), h.Img(h.Src("https://fdn2.gsmarena.com/vv/bigpic/xiaomi-redmi-note-15-pro-plus-global.jpg")), h.Br(), g.Text("Xiaomi Redmi Note 15 Pro+")),
			h.A(h.Href("oneplus_15r_5g-14358.php"), h.Class("module-phones-link"), h.Img(h.Src("https://fdn2.gsmarena.com/vv/bigpic/oneplus-15r.jpg")), h.Br(), g.Text("OnePlus 15R")),
			h.A(h.Href("oneplus_15_5g-14206.php"), h.Class("module-phones-link"), h.Img(h.Src("https://fdn2.gsmarena.com/vv/bigpic/oneplus-15.jpg")), h.Br(), g.Text("OnePlus 15")),
			h.A(h.Href("xiaomi_poco_c85_5g-14337.php"), h.Class("module-phones-link"), h.Img(h.Src("https://fdn2.gsmarena.com/vv/bigpic/xiaomi-poco-c85-5g.jpg")), h.Br(), g.Text("Xiaomi Poco C85")),
			h.A(h.Href("xiaomi_poco_f8_ultra_5g-14301.php"), h.Class("module-phones-link"), h.Img(h.Src("https://fdn2.gsmarena.com/vv/bigpic/xiaomi-poco-f8-ultra-.jpg")), h.Br(), g.Text("Xiaomi Poco F8 Ultra")),
			h.A(h.Href("xiaomi_poco_f8_pro_5g-14303.php"), h.Class("module-phones-link"), h.Img(h.Src("https://fdn2.gsmarena.com/vv/bigpic/xiaomi-poco-f8-pro.jpg")), h.Br(), g.Text("Xiaomi Poco F8 Pro")),
			h.A(h.Href("vivo_x300_pro_5g-14225.php"), h.Class("module-phones-link"), h.Img(h.Src("https://fdn2.gsmarena.com/vv/bigpic/vivo-x300-pro.jpg")), h.Br(), g.Text("vivo X300 Pro")),
			h.A(h.Href("rugone_xever_7_pro_5g-14386.php"), h.Class("module-phones-link"), h.Img(h.Src("https://fdn2.gsmarena.com/vv/bigpic/rugone-xever-7-pro.jpg")), h.Br(), g.Text("RugOne Xever 7 Pro")),
		),
	)
}

func Top10DailyInterest() g.Node {
	return h.Div(h.Class("module module-rankings s3"),
		h.H4(h.Class("section-heading"), g.Text("Top 10 by daily interest")),
		h.Table(h.Class("module-fit green"), g.Attr("cellspacing", "0"),
			h.ColGroup(
				h.Col(h.Class("numb")),
				h.Col(h.Class("phon")),
				h.Col(h.Class("hits")),
				h.Col(h.Class("dummy")),
			),
			h.THead(
				h.Tr(
					h.Th(h.Scope("col"), h.ID("th3a"), g.Attr("abbr", ""), g.Raw("&nbsp;")),
					h.Th(h.Scope("col"), h.ID("th3b"), g.Attr("abbr", ""), g.Text("Device")),
					h.Th(h.Scope("col"), h.ID("th3c"), g.Attr("abbr", ""), g.Text("Daily hits")),
					h.Th(h.Scope("col"), g.Raw("&nbsp;")),
				),
			),
			h.TBody(
				h.Tr(h.Td(g.Attr("headers", "th3a"), g.Text("1.")), h.Th(g.Attr("headers", "th3b"), h.A(h.Href("xiaomi_17_ultra-14380.php"), g.Raw("<nobr>Xiaomi 17 Ultra</nobr>"))), h.Td(g.Attr("headers", "th3c"), g.Text("27,083")), h.Td()),
				h.Tr(h.Td(g.Attr("headers", "th3a"), g.Text("2.")), h.Th(g.Attr("headers", "th3b"), h.A(h.Href("samsung_galaxy_a56-13603.php"), g.Raw("<nobr>Samsung Galaxy A56</nobr>"))), h.Td(g.Attr("headers", "th3c"), g.Text("20,085")), h.Td()),
				h.Tr(h.Td(g.Attr("headers", "th3a"), g.Text("3.")), h.Th(g.Attr("headers", "th3b"), h.A(h.Href("samsung_galaxy_s25_ultra-13322.php"), g.Raw("<nobr>Samsung Galaxy S25 Ultra</nobr>"))), h.Td(g.Attr("headers", "th3c"), g.Text("19,269")), h.Td()),
				h.Tr(h.Td(g.Attr("headers", "th3a"), g.Text("4.")), h.Th(g.Attr("headers", "th3b"), h.A(h.Href("apple_iphone_17_pro_max-13964.php"), g.Raw("<nobr>Apple iPhone 17 Pro Max</nobr>"))), h.Td(g.Attr("headers", "th3c"), g.Text("17,853")), h.Td()),
				h.Tr(h.Td(g.Attr("headers", "th3a"), g.Text("5.")), h.Th(g.Attr("headers", "th3b"), h.A(h.Href("honor_win-14377.php"), g.Raw("<nobr>Honor Win</nobr>"))), h.Td(g.Attr("headers", "th3c"), g.Text("15,897")), h.Td()),
				h.Tr(h.Td(g.Attr("headers", "th3a"), g.Text("6.")), h.Th(g.Attr("headers", "th3b"), h.A(h.Href("samsung_galaxy_s26_ultra-14320.php"), g.Raw("<nobr>Samsung Galaxy S26 Ultra</nobr>"))), h.Td(g.Attr("headers", "th3c"), g.Text("15,622")), h.Td()),
				h.Tr(h.Td(g.Attr("headers", "th3a"), g.Text("7.")), h.Th(g.Attr("headers", "th3b"), h.A(h.Href("samsung_galaxy_a17-14041.php"), g.Raw("<nobr>Samsung Galaxy A17</nobr>"))), h.Td(g.Attr("headers", "th3c"), g.Text("14,824")), h.Td()),
				h.Tr(h.Td(g.Attr("headers", "th3a"), g.Text("8.")), h.Th(g.Attr("headers", "th3b"), h.A(h.Href("oneplus_15-14206.php"), g.Raw("<nobr>OnePlus 15</nobr>"))), h.Td(g.Attr("headers", "th3c"), g.Text("14,499")), h.Td()),
				h.Tr(h.Td(g.Attr("headers", "th3a"), g.Text("9.")), h.Th(g.Attr("headers", "th3b"), h.A(h.Href("apple_iphone_17-14050.php"), g.Raw("<nobr>Apple iPhone 17</nobr>"))), h.Td(g.Attr("headers", "th3c"), g.Text("12,395")), h.Td()),
				h.Tr(h.Td(g.Attr("headers", "th3a"), g.Text("10.")), h.Th(g.Attr("headers", "th3b"), h.A(h.Href("xiaomi_17_pro_max-14181.php"), g.Raw("<nobr>Xiaomi 17 Pro Max</nobr>"))), h.Td(g.Attr("headers", "th3c"), g.Text("12,387")), h.Td()),
			),
		),
	)
}

func Top10ByFans() g.Node {
	return h.Div(h.Class("module module-rankings s3"),
		h.H4(h.Class("section-heading"), g.Text("Top 10 by fans")),
		h.Table(h.Class("module-fit blue"), g.Attr("cellspacing", "0"),
			h.ColGroup(
				h.Col(h.Class("numb")),
				h.Col(h.Class("phon")),
				h.Col(h.Class("hits")),
				h.Col(h.Class("dummy")),
			),
			h.THead(
				h.Tr(
					h.Th(h.Scope("col"), h.ID("th3a"), g.Attr("abbr", ""), g.Raw("&nbsp;")),
					h.Th(h.Scope("col"), h.ID("th3b"), g.Attr("abbr", ""), g.Text("Device")),
					h.Th(h.Scope("col"), h.ID("th3c"), g.Attr("abbr", ""), g.Text("Favorites")),
					h.Th(h.Scope("col"), g.Raw("&nbsp;")),
				),
			),
			h.TBody(
				h.Tr(h.Td(g.Attr("headers", "th3a"), g.Text("1.")), h.Th(g.Attr("headers", "th3b"), h.A(h.Href("samsung_galaxy_s25_ultra-13322.php"), g.Raw("<nobr>Samsung Galaxy S25 Ultra</nobr>"))), h.Td(g.Attr("headers", "th3c"), g.Text("1,318")), h.Td()),
				h.Tr(h.Td(g.Attr("headers", "th3a"), g.Text("2.")), h.Th(g.Attr("headers", "th3b"), h.A(h.Href("oneplus_13-13477.php"), g.Raw("<nobr>OnePlus 13</nobr>"))), h.Td(g.Attr("headers", "th3c"), g.Text("767")), h.Td()),
				h.Tr(h.Td(g.Attr("headers", "th3a"), g.Text("3.")), h.Th(g.Attr("headers", "th3b"), h.A(h.Href("xiaomi_15_ultra-13657.php"), g.Raw("<nobr>Xiaomi 15 Ultra</nobr>"))), h.Td(g.Attr("headers", "th3c"), g.Text("658")), h.Td()),
				h.Tr(h.Td(g.Attr("headers", "th3a"), g.Text("4.")), h.Th(g.Attr("headers", "th3b"), h.A(h.Href("samsung_galaxy_a56-13603.php"), g.Raw("<nobr>Samsung Galaxy A56</nobr>"))), h.Td(g.Attr("headers", "th3c"), g.Text("632")), h.Td()),
				h.Tr(h.Td(g.Attr("headers", "th3a"), g.Text("5.")), h.Th(g.Attr("headers", "th3b"), h.A(h.Href("xiaomi_poco_x7_pro-13582.php"), g.Raw("<nobr>Xiaomi Poco X7 Pro</nobr>"))), h.Td(g.Attr("headers", "th3c"), g.Text("606")), h.Td()),
				h.Tr(h.Td(g.Attr("headers", "th3a"), g.Text("6.")), h.Th(g.Attr("headers", "th3b"), h.A(h.Href("vivo_x200_pro-13410.php"), g.Raw("<nobr>vivo X200 Pro</nobr>"))), h.Td(g.Attr("headers", "th3c"), g.Text("583")), h.Td()),
				h.Tr(h.Td(g.Attr("headers", "th3a"), g.Text("7.")), h.Th(g.Attr("headers", "th3b"), h.A(h.Href("xiaomi_17_pro_max-14181.php"), g.Raw("<nobr>Xiaomi 17 Pro Max</nobr>"))), h.Td(g.Attr("headers", "th3c"), g.Text("556")), h.Td()),
				h.Tr(h.Td(g.Attr("headers", "th3a"), g.Text("8.")), h.Th(g.Attr("headers", "th3b"), h.A(h.Href("samsung_galaxy_s25-13610.php"), g.Raw("<nobr>Samsung Galaxy S25</nobr>"))), h.Td(g.Attr("headers", "th3c"), g.Text("502")), h.Td()),
				h.Tr(h.Td(g.Attr("headers", "th3a"), g.Text("9.")), h.Th(g.Attr("headers", "th3b"), h.A(h.Href("honor_magic7_pro-13480.php"), g.Raw("<nobr>Honor Magic7 Pro</nobr>"))), h.Td(g.Attr("headers", "th3c"), g.Text("471")), h.Td()),
				h.Tr(h.Td(g.Attr("headers", "th3a"), g.Text("10.")), h.Th(g.Attr("headers", "th3b"), h.A(h.Href("xiaomi_15-13472.php"), g.Raw("<nobr>Xiaomi 15</nobr>"))), h.Td(g.Attr("headers", "th3c"), g.Text("461")), h.Td()),
			),
		),
		h.StyleEl(g.Raw(`
.module-fit.purple { 
    width: 100%; 
    table-layout: fixed;
}
.module-fit.purple td {
    text-align: left;
}

.module-rankings .purple tbody tr:nth-child(odd) {
    background-color: #fff2ee;
}
`)),
	)
}

func PopularComparisons() g.Node {
	return h.Div(h.Class("module module-rankings s3"),
		h.H4(h.Class("section-heading"), g.Text("Popular comparisons")),
		h.Table(h.Class("module-fit purple"), g.Attr("cellspacing", "0"),
			h.TBody(
				h.Tr(h.Th(h.A(h.Href("compare.php3?idPhone1=13497&idPhone2=13603"), g.Text("Samsung Galaxy A36 vs. Galaxy A56")))),
				h.Tr(h.Th(h.A(h.Href("compare.php3?idPhone1=13322&idPhone2=13964"), g.Text("Samsung Galaxy S25 Ultra vs. iPhone 17 Pro Max")))),
				h.Tr(h.Th(h.A(h.Href("compare.php3?idPhone1=13317&idPhone2=14050"), g.Text("Apple iPhone 16 vs. Apple iPhone 17")))),
				h.Tr(h.Th(h.A(h.Href("compare.php3?idPhone1=12771&idPhone2=13322"), g.Text("Samsung Galaxy S24 Ultra vs. Galaxy S25 Ultra")))),
				h.Tr(h.Th(h.A(h.Href("compare.php3?idPhone1=13610&idPhone2=14042"), g.Text("Samsung Galaxy S25 vs. Galaxy S25 FE 5G")))),
				h.Tr(h.Th(h.A(h.Href("compare.php3?idPhone1=13123&idPhone2=13964"), g.Text("Apple iPhone 16 Pro Max vs. iPhone 17 Pro Max")))),
				h.Tr(h.Th(h.A(h.Href("compare.php3?idPhone1=14049&idPhone2=14050"), g.Text("Apple iPhone 17 Pro vs. Apple iPhone 17")))),
				h.Tr(h.Th(h.A(h.Href("compare.php3?idPhone1=14206&idPhone2=14358"), g.Text("OnePlus 15 5G vs. OnePlus 15R 5G")))),
				h.Tr(h.Th(h.A(h.Href("compare.php3?idPhone1=14224&idPhone2=14225"), g.Text("vivo X300 5G vs. vivo X300 Pro 5G")))),
				h.Tr(h.Th(h.A(h.Href("compare.php3?idPhone1=13964&idPhone2=14049"), g.Text("Apple iPhone 17 Pro Max vs. iPhone 17 Pro")))),
				h.Tr(h.Th(h.A(h.Href("compare.php3?idPhone1=9616&idPhone2=9721"), g.Text("Samsung Galaxy A60 vs. Galaxy M40")))),
			),
		),
	)
}

func ElectricVehicles() g.Node {
	return h.Div(h.Class("module module-news-l"),
		h.H4(h.Class("section-heading"), h.A(h.Href("https://www.arenaev.com/"), h.Target("_blank"), g.Text("Electric vehicles"))),
		h.Ul(
			h.Li(h.A(h.Href("https://www.arenaev.com/xpeng_g9_review_interior_design-news-5445.php"), h.Target("_blank"),
				h.Img(h.Src("https://st.arenaev.com/news/25/12/xpeng-g9-review/-344x215/arenaev_000.jpg"), h.Alt("2026 XPeng G9 interior, design and features review")),
				h.Span(g.Text("2026 XPeng G9 interior, design and features review")),
			)),
			h.Li(h.A(h.Href("https://www.arenaev.com/tesla_loses_top_spot_as_electric_car_sales_continue_to_slide-news-5467.php"), h.Target("_blank"),
				h.Img(h.Src("https://st.arenaev.com/news/26/01/tesla-loses-top-spot-as-electric-car-sales-slide/-344x215/arenaev_000.jpg"), h.Alt("Tesla loses top spot as electric car sales continue to slide")),
				h.Span(g.Text("Tesla loses top spot as electric car sales continue to slide")),
			)),
			h.Li(h.A(h.Href("https://www.arenaev.com/bmw_alpina_is_now_a_standalone_brand_focused_on_performance_and_bespoke_customization-news-5466.php"), h.Target("_blank"),
				h.Img(h.Src("https://st.arenaev.com/news/26/01/bmw-alpina/-344x215/arenaev_000.jpg"), h.Alt("BMW Alpina is now a standalone brand focused on performance and bespoke customization")),
				h.Span(g.Text("BMW Alpina is now a standalone brand focused on performance and bespoke customization")),
			)),
			h.Li(h.A(h.Href("https://www.arenaev.com/comparison_of_nedc_epa_and_wltp_cycles-news-419.php"), h.Target("_blank"),
				h.Img(h.Src("https://st.arenaev.com/news/22/06/comparison-of-nedc-epa-wltp-cycles/-344x215/arenaev_001.jpg"), h.Alt("Comparison of NEDC, EPA and WLTP cycles")),
				h.Span(g.Text("Comparison of NEDC, EPA and WLTP cycles")),
			)),
		),
	)
}
