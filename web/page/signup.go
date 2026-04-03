package page

import (
	"github.com/LuuDinhTheTai/tzone/web/page/shared"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func SignupPage() g.Node {
	return shared.BaseLayout("Sign Up - GSMArena.com",
		h.Div(h.ID("wrapper"), h.Class("l-container"),
			h.StyleEl(g.Raw(`
				.main.float-right::after { display: none !important; content: none !important; }
			`)),
			h.Div(h.ID("outer"), h.Class("row"),
				h.Div(h.ID("subHeader3")),
				h.Div(h.ID("body"),
					h.Div(h.Class("main float-right"), h.Style("background: #fff; padding: 0; min-height: 500px;"),
						// Banner Image
						h.Div(h.Style("position: relative;"),
							h.Img(h.Src("https://fdn.gsmarena.com/imgroot/static/headers/glossary-hlr.jpg"), h.Alt("Sign Up Banner"), h.Style("width: 100%; height: auto; display: block; filter: grayscale(100%);")),
							h.Div(h.Style("position: absolute; bottom: 20px; left: 20px;"),
								h.H1(g.Text("Sign Up"), h.Style("color: white; font-size: 32px; font-weight: normal; margin: 0; text-shadow: 1px 1px 2px rgba(0,0,0,0.8);")),
							),
						),
						// Why register section
						h.Div(h.Style("padding: 20px;"),
							h.H3(g.Text("Why register"), h.Style("color: #d50000; font-size: 18px; margin-bottom: 15px; font-weight: bold;")),
							h.Ul(h.Style("list-style: none; padding: 0; margin: 0; line-height: 2; color: #333; font-size: 14px;"),
								h.Li(g.Raw("&#8250; "), g.Text("Your nickname will be reserved for you only and you will be able to use an avatar")),
								h.Li(g.Raw("&#8250; "), g.Text("Your comments and opinions will be posted immediately")),
								h.Li(g.Raw("&#8250; "), g.Text("You will get additional features like reply notification and device bookmarks")),
								h.Li(g.Raw("&#8250; "), g.Text("We care about protecting your privacy. We won't share your data")),
							),
						),
						// Create account form
						h.Div(h.Style("padding: 20px; margin-top: 130px !important; background: #f9f9f9; border-top: 1px solid #eaeaea;"),
							h.H3(g.Text("Create account"), h.Style("color: #333; font-size: 18px; margin-bottom: 20px; font-weight: bold;")),
							h.Form(h.Action("/signup"), h.Method("post"),
								h.Div(h.Style("margin-bottom: 15px;"),
									h.Label(h.For("nickname"), g.Text("Your Nickname"), h.Style("display: block; margin-bottom: 5px; color: #333; font-size: 13px;")),
									h.Input(h.Type("text"), h.ID("nickname"), h.Name("nickname"), h.Style("width: 100%; max-width: 500px; padding: 8px; border: 1px solid #ccc;")),
								),
								h.Div(h.Style("margin-bottom: 15px;"),
									h.Label(h.For("email"), g.Text("Your Email"), h.Style("display: block; margin-bottom: 5px; color: #333; font-size: 13px;")),
									h.Input(h.Type("email"), h.ID("email"), h.Name("email"), h.Style("width: 100%; max-width: 500px; padding: 8px; border: 1px solid #ccc; background: #f2cwce;")), // slightly reddish as in screenshot
								),
								h.Div(h.Style("margin-bottom: 20px;"),
									h.Label(h.For("password"), g.Text("Password (6 to 20 characters)"), h.Style("display: block; margin-bottom: 5px; color: #333; font-size: 13px;")),
									h.Input(h.Type("password"), h.ID("password"), h.Name("password"), h.Placeholder("Your password"), h.Style("width: 100%; max-width: 500px; padding: 8px; border: 1px solid #ccc;")),
								),
								h.Div(h.Style("margin-bottom: 15px; display: flex; align-items: center; gap: 10px;"),
									h.Input(h.Type("checkbox"), h.ID("agree_terms"), h.Name("agree_terms"), h.Style("width: 20px; height: 10px; accent-color: #d50000;")),
									h.Label(h.For("agree_terms"), g.Text("I agree for GSMArena to store my email address, nickname and password"), h.Style("font-size: 13px; color: #333;")),
								),
								h.Div(h.Style("margin-bottom: 25px; display: flex; align-items: center; gap: 10px;"),
									h.Input(h.Type("checkbox"), h.ID("age_verification"), h.Name("age_verification"), h.Style("width: 20px; height: 10px; accent-color: #d50000;")),
									h.Label(h.For("age_verification"), g.Text("I am at least 16 years old"), h.Style("font-size: 13px; color: #333;")),
								),
								h.Div(h.Style("max-width: 500px; text-align: right;"),
									h.Button(h.Type("submit"), g.Text("SUBMIT"), h.Style("background: #ccc; color: #666; padding: 8px 20px; border: none; font-weight: bold; cursor: not-allowed;")),
								),
							),
						),
						h.Div(h.Style("padding: 15px 20px; font-size: 12px; color: #666; background: #fff; border-top: 1px solid #eaeaea;"),
							g.Text("We have to verify your email address. Once you complete this form you will receive an email with a confirmation link."),
						),
					),
					Sidebar(), // Reusing the sidebar from index.go
				),
			),
		),
	)
}
