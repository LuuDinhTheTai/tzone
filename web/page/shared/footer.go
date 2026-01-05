package shared

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func Footer() g.Node {
	return h.Div(h.ID("footer"),
		h.Div(h.Class("footer-logo"),
			h.Img(h.Src("https://fdn2.gsmarena.com/w/css/logo-gsmarena-com.png"), h.Alt("")),
		),
		h.Div(h.ID("footmenu"),
			h.P(
				h.A(h.Href("/"), g.Text("Home")),
				h.A(h.Href("news.php3"), g.Text("News")),
				h.A(h.Href("reviews.php3"), g.Text("Reviews")),
				h.A(h.Href("compare.php3"), g.Text("Compare")),
				h.A(h.Href("network-bands.php3"), g.Text("Coverage")),
				h.A(h.Href("glossary.php3"), g.Text("Glossary")),
				h.A(h.Href("faq.php3"), g.Text("FAQ")),

				h.A(h.Href("rss-news-reviews.php3"), h.Class("rss-icon"), g.Text("RSS")),
				h.A(h.Target("_blank"), h.Rel("noopener"), h.Href("https://www.youtube.com/channel/UCbLq9tsbo8peV22VxbDAfXA?sub_confirmation=1"), h.Class("yt-icon"), g.Text("Youtube")),
				h.A(h.Target("_blank"), h.Rel("noopener"), h.Href("https://www.instagram.com/gsmarenateam/"), h.Class("ig-icon"), g.Text("Instagram")),
				h.A(h.Target("_blank"), h.Rel("noopener"), h.Href("https://www.tiktok.com/@gsmarenateam"), h.Class("tiktok-icon"), g.Text("TikTok")),
				h.A(h.Target("_blank"), h.Rel("noopener"), h.Href("https://www.facebook.com/GSMArenacom-189627474421/"), h.Class("fb-icon"), g.Text("Facebook")),
				h.A(h.Target("_blank"), h.Rel("noopener"), h.Href("https://twitter.com/gsmarena_com"), h.Class("tw-icon"), g.Text("Twitter")),
			),

			h.P(
				g.Text("© 2000-2026 "),
				h.A(h.Href("team.php3"), g.Text("GSMArena.com")),
				g.Text(" "),
				h.A(h.ID("switch-version"), h.Href("#"), g.Text("Mobile version")),
				g.Text(" "),
				h.A(h.Target("_blank"), h.Rel("noopener"), h.Href("https://play.google.com/store/apps/details?id=com.gsmarena.android"), g.Text("Android app")),
				g.Text(" "),
				h.A(h.Href("tools.php3"), g.Text("Tools")),
				g.Text(" "),
				h.A(h.Href("contact.php3"), g.Text("Contact us")),
				g.Text(" "),
				h.A(h.Href("https://merch.gsmarena.com/"), h.Target("_blank"), g.Text("Merch store")),
				g.Text(" "),
				h.A(h.Href("privacy-policy.php3"), g.Text("Privacy")),
				g.Text(" "),
				h.A(h.Href("terms.php3"), g.Text("Terms of use")),
				g.Text(" "),

				h.A(h.ID("unic-gdpr"), g.Attr("onclick", `__tcfapi("openunic");return false;`), h.Style("display:none;cursor:pointer;"), g.Text("Change Ad Consent")),
				h.A(h.ID("unic-ccpa"), g.Attr("onclick", `window.__uspapi('openunic')`), h.Style("display:none;cursor:pointer;"), g.Text("Do not sell my data")),
			),
		),
	)
}
