package shared

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// PageProps chứa dữ liệu cần thiết cho layout
type PageProps struct {
	Title string
	Body  Node
}

// Layout là khung xương chung cho toàn bộ website
func Layout(props PageProps) Node {
	// Sử dụng Doctype chuẩn + thẻ HTML riêng biệt để tránh lỗi wrapper
	return Group([]Node{
		HTML(
			Lang("en"),
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1.0")),
				TitleEl(Text(props.Title)),
				// Sử dụng TailwindCSS từ CDN
				Script(Src("https://cdn.tailwindcss.com")),
				// Custom style
				StyleEl(Text(`
					body { font-family: Arial, sans-serif; background-color: #f6f6f6; }
					.gsm-red { color: #a40000; }
					.gsm-bg-red { background-color: #a40000; }
					.gsm-bg-purple { background-color: #483d8b; }
					.hover-underline:hover { text-decoration: underline; }
					/* Custom scrollbar cho giống web gốc */
					.custom-scrollbar::-webkit-scrollbar { width: 6px; }
					.custom-scrollbar::-webkit-scrollbar-track { background: #f1f1f1; }
					.custom-scrollbar::-webkit-scrollbar-thumb { background: #888; border-radius: 3px; }
					.custom-scrollbar::-webkit-scrollbar-thumb:hover { background: #555; }
				`)),
			),
			Body(
				Class("text-gray-800 text-sm"),

				// --- HEADER ---
				Header(Class("bg-white border-b-4 border-[#a40000]"),
					Div(Class("container mx-auto max-w-[1060px] px-2"),
						// Top Bar
						Div(Class("flex justify-between items-center py-4"),
							// Logo
							A(Href("/"), Class("flex items-center gap-2"),
								Img(Src("https://fdn.gsmarena.com/vv/assets12/i/logo.svg"), Alt("GSMArena"), Class("h-10")),
								Span(Class("text-2xl font-bold tracking-tighter uppercase hidden sm:block"), Text("GSMArena.com")),
							),
							// Search Bar (Fake)
							Div(Class("flex-1 mx-8 max-w-md hidden md:flex"),
								Input(Type("text"), Placeholder("Search"), Class("w-full border border-gray-300 rounded-l px-3 py-1 focus:outline-none focus:border-gray-500")),
								Button(Class("bg-gray-100 border border-l-0 border-gray-300 rounded-r px-3"), Text("Go")),
							),
							// Social Icons
							Div(Class("flex gap-3 text-gray-500 text-xs text-center items-center"),
								SocialIcon("Tip us", "💡"),
								SocialIcon("YouTube", "▶️"),
								SocialIcon("Insta", "📷"),
								LoginButton(),
							),
						),
					),
					// Navigation Menu
					Div(Class("bg-[#483d8b] text-white shadow-md"),
						Div(Class("container mx-auto max-w-[1060px] px-2"),
							Ul(Class("flex flex-wrap gap-6 py-2.5 font-bold uppercase text-xs tracking-wide"),
								NavItem("Home", true),
								NavItem("News", false),
								NavItem("Reviews", false),
								NavItem("Videos", false),
								NavItem("Phone Finder", false),
								NavItem("Deals", false),
								NavItem("Merch", false),
								NavItem("Coverage", false),
								NavItem("Contact", false),
							),
						),
					),
				),

				// --- MAIN CONTENT ---
				Main(Class("container mx-auto max-w-[1060px] px-2 py-4"),
					props.Body,
				),

				// --- FOOTER ---
				Footer(Class("bg-[#282828] text-gray-400 py-8 mt-8 border-t-4 border-[#a40000]"),
					Div(Class("container mx-auto max-w-[1060px] px-2 text-center text-xs"),
						Div(Class("mb-4 flex justify-center gap-4 font-bold uppercase"),
							A(Href("#"), Text("Home")), A(Href("#"), Text("News")), A(Href("#"), Text("Reviews")), A(Href("#"), Text("Compare")), A(Href("#"), Text("Coverage")),
						),
						P(Text("© 2000-2026 GSMArena.com")),
						P(Class("mt-2"), Text("Mobile version | Android app | Tools | Contact us | Privacy | Terms")),
					),
				),
			),
		),
	})
}

// Helper Components
func NavItem(text string, active bool) Node {
	classes := "hover:text-gray-300 transition-colors"
	if active {
		classes += " text-white"
	} else {
		classes += " text-gray-300"
	}
	return Li(
		A(Href("#"), Class(classes), Text(text)),
	)
}

func SocialIcon(name, icon string) Node {
	return A(Href("#"), Class("hover:text-black flex flex-col items-center"),
		Span(Class("text-lg leading-none mb-1"), Text(icon)),
		Span(Text(name)),
	)
}

func LoginButton() Node {
	return A(Href("#"), Class("flex items-center gap-1 text-gray-600 hover:text-black ml-4"),
		Span(Class("font-bold"), Text("Log in")),
	)
}
