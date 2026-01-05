package page

import (
	"github.com/LuuDinhTheTai/tzone/web/page/shared"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func IndexPage() Node {
	return shared.Layout(shared.PageProps{
		Title: "GSMArena.com - mobile phone reviews, news, specifications",
		Body: Div(Class("grid grid-cols-1 md:grid-cols-12 gap-6"),

			// --- LEFT COLUMN (CONTENT - 8/12) ---
			Div(Class("md:col-span-8 space-y-8"),

				// Section: Featured Review (Hero)
				Div(Class("bg-white border border-gray-300 p-4 shadow-sm"),
					H3(Class("text-xl font-bold mb-3 text-black hover-underline cursor-pointer"), Text("Best-sounding phones and tablets we tested in 2025")),
					Div(Class("relative h-64 bg-gray-200 w-full mb-2 overflow-hidden group"),
						Img(Src("https://fdn.gsmarena.com/imgroot/reviews/25/best-smartphone-speakers-2025/-728x314/gsmarena_201.jpg"), Class("w-full h-full object-cover transition-transform group-hover:scale-105")),
					),
					Div(Class("flex justify-between text-xs text-gray-500 border-t pt-2"),
						Span(Class("flex items-center gap-1"), Text("🕒 02 Jan 2026")),
						Span(Class("flex items-center gap-1"), Text("💬 61 comments")),
					),
				),

				// Section: Sub Reviews Grid
				Div(Class("grid grid-cols-2 gap-4"),
					ReviewCard("Huawei Mate X7 review", "https://fdn.gsmarena.com/imgroot/reviews/25/huawei-mate-x7/-347x151/gsmarena_001.jpg", "38"),
					ReviewCard("Xiaomi Redmi Note 15 Pro 5G review", "https://fdn.gsmarena.com/imgroot/reviews/25/redmi-note-15-pro/-347x151/gsmarena_001.jpg", "40"),
				),

				// Section: News Feed
				Div(Class("space-y-0"),
					H4(Class("bg-[#f0f0f0] px-2 py-1 font-bold text-gray-600 uppercase text-xs border-b border-gray-300 mb-4"), Text("Latest News")),

					NewsItem("Realme Neo8 to debut in January with Snapdragon 8 Gen 5 SoC",
						"The phone will be available in a Cyber Purple color with an RGB light at the back.",
						"https://fdn.gsmarena.com/imgroot/news/26/01/realme-neo8-launch-design-specs/-344x215/gsmarena_000.jpg",
						"36 minutes ago"),

					NewsItem("Huawei FreeBuds 7i review",
						"Huawei's do-it-all TWS earbuds bring on a rich set of features without breaking the bank.",
						"https://fdn.gsmarena.com/imgroot/news/25/12/huawei-freebuds-7i-review/-344x215/gsmarena_000.jpg",
						"17 hours ago"),

					NewsItem("Top 10 trending phones of week 1",
						"New year, new leader.",
						"https://fdn.gsmarena.com/imgroot/news/19/04/top10-trending-phones/-344x215/gsmarena_003.jpg",
						"21 hours ago"),

					NewsItem("Oppo Find X9 Ultra's camera sensors get detailed in new leak",
						"It's an intriguing setup, that's for sure.",
						"https://fdn.gsmarena.com/imgroot/news/26/01/oppo-find-x9-ultra-camera-sensors/-344x215/gsmarena_000.jpg",
						"04 Jan 2026"),
				),
			),

			// --- RIGHT COLUMN (SIDEBAR - 4/12) ---
			Aside(Class("md:col-span-4 space-y-6"),

				// Phone Finder Widget
				Div(Class("bg-white border border-gray-300 p-4"),
					H4(Class("text-[#a40000] font-bold uppercase mb-4 flex items-center gap-2"),
						Span(Text("Phone Finder")),
					),
					Div(Class("space-y-2"),
						BrandList(),
						A(Href("#"), Class("block text-center text-xs font-bold uppercase text-[#a40000] mt-4 hover:underline"), Text("All brands")),
						A(Href("#"), Class("block text-center text-xs font-bold uppercase text-[#a40000] hover:underline"), Text("Rumor mill")),
					),
				),

				// Latest Devices
				Div(Class("bg-white border border-gray-300"),
					H4(Class("bg-[#f0f0f0] p-2 font-bold text-gray-700 uppercase text-xs border-b border-gray-300"), Text("Latest devices")),
					Div(Class("p-2 max-h-[400px] overflow-y-auto custom-scrollbar"),
						DeviceSmall("Oppo Reno15 Pro Max", "https://fdn2.gsmarena.com/vv/bigpic/oppo-reno15-pro.jpg"),
						DeviceSmall("Xiaomi Poco M8 Pro", "https://fdn2.gsmarena.com/vv/bigpic/xiaomi-poco-m8-pro.jpg"),
						DeviceSmall("Xiaomi 17 Ultra", "https://fdn2.gsmarena.com/vv/bigpic/xiaomi-17-ultra.jpg"),
						DeviceSmall("Vivo X300 Ultra", "https://fdn2.gsmarena.com/vv/bigpic/vivo-x200-ultra.jpg"),
						DeviceSmall("OnePlus 15R", "https://fdn2.gsmarena.com/vv/bigpic/oneplus-15r.jpg"),
					),
				),

				// Top 10 Interest
				Div(Class("bg-white border border-gray-300"),
					H4(Class("bg-[#f0f0f0] p-2 font-bold text-gray-700 uppercase text-xs border-b border-gray-300"), Text("Top 10 by daily interest")),
					Table(Class("w-full text-xs"),
						TBody(
							Top10Row(1, "Xiaomi 17 Ultra", "27,083"),
							Top10Row(2, "Samsung Galaxy A56", "20,085"),
							Top10Row(3, "Samsung S25 Ultra", "19,269"),
							Top10Row(4, "iPhone 17 Pro Max", "17,853"),
							Top10Row(5, "Honor Win", "15,897"),
						),
					),
				),
			),
		),
	})
}

// --- Helper Components cho Index ---

func ReviewCard(title, imgUrl, comments string) Node {
	return Article(Class("bg-white border border-gray-300 shadow-sm"),
		A(Href("#"), Class("block group"),
			Div(Class("p-3"),
				H3(Class("font-bold text-sm leading-tight mb-2 group-hover:underline text-[#a40000]"), Text(title)),
			),
			Img(Src(imgUrl), Class("w-full h-32 object-cover")),
		),
		Div(Class("p-2 border-t text-xs text-gray-500 flex justify-end"),
			Span(Class("flex items-center gap-1"), Text("💬 "+comments)),
		),
	)
}

func NewsItem(title, desc, imgUrl, time string) Node {
	return Div(Class("flex gap-4 mb-6 border-b border-gray-200 pb-4 last:border-0"),
		Div(Class("w-1/3"),
			Img(Src(imgUrl), Class("w-full h-auto object-cover border border-gray-200")),
		),
		Div(Class("w-2/3"),
			H3(Class("font-bold text-md text-[#a40000] hover:underline mb-1 leading-tight"),
				A(Href("#"), Text(title)),
			),
			P(Class("text-gray-600 text-sm mb-2 leading-snug"), Text(desc)),
			Div(Class("text-xs text-gray-400 flex items-center gap-2"),
				Span(Text("🕒 "+time)),
				Span(Text("• 💬 10 comments")),
			),
		),
	)
}

// Xóa hàm BrandList cũ và thay bằng hàm này
func BrandList() Node {
	brands := []string{"Samsung", "Apple", "Huawei", "Nokia", "Sony", "LG", "HTC", "Motorola", "Lenovo", "Xiaomi", "Google", "Honor", "Oppo", "Realme", "OnePlus", "vivo"}

	// Tạo slice chứa các Node
	var items []Node
	for _, b := range brands {
		// Tạo từng thẻ Li
		item := Li(
			A(Href("#"), Class("text-gray-800 hover:text-[#a40000] hover:underline block pl-2"), Text(b)),
		)
		items = append(items, item)
	}

	return Ul(
		Class("grid grid-cols-2 gap-y-1 text-xs"),
		// Sử dụng Group để gom slice []Node thành 1 Node duy nhất
		// Cách này tránh lỗi "invalid use of ..."
		Group(items),
	)
}

func DeviceSmall(name, imgUrl string) Node {
	return A(Href("#"), Class("flex items-center gap-3 p-2 hover:bg-gray-100 border-b border-gray-100 last:border-0"),
		Img(Src(imgUrl), Class("w-10 h-12 object-contain")),
		Span(Class("text-xs font-bold text-gray-700 group-hover:text-[#a40000]"), Text(name)),
	)
}

func Top10Row(rank int, name, hits string) Node {
	return Tr(Class("border-b border-gray-100 last:border-0 hover:bg-[#fff2ee]"),
		Td(Class("p-2 text-gray-400 w-6"), Textf("%d.", rank)),
		Td(Class("p-2 font-bold text-[#a40000]"), A(Href("#"), Text(name))),
		Td(Class("p-2 text-gray-500 text-right"), Text(hits)),
	)
}
