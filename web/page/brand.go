package page

import (
	"fmt"

	"github.com/LuuDinhTheTai/tzone/web/page/shared"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Device struct {
	Name   string
	ImgSrc string
}

func BrandPage(brandName string, devices []Device, currentPage int, totalPages int) g.Node {
	title := fmt.Sprintf("All %s phones", brandName)
	return shared.BaseLayout(title,
		h.Div(h.ID("wrapper"), h.Class("l-container"),
			h.Div(h.ID("outer"), h.Class("row"),
				h.Div(h.ID("subHeader3")),
				h.Div(h.ID("body"), h.Class("clearfix"),
					h.Div(h.Class("main right main-maker l-box col"),
						// We embed exactly makers.css properties used.
						h.StyleEl(h.Type("text/css"), g.Raw(`
	.review-background {
		background-image: url('https://fdn.gsmarena.com/imgroot/static/headers/makers/acer-2.jpg');
		background-size: 728px 314px;
	}
/* Original GSMArena makers.css */
@charset "UTF-8";.article-info-meta-link:last-of-type a:after{background-image:none}.center-stage.blur:after{position:absolute;width:100%;height:100%;display:block;top:0;left:0;background:url(//cdn2.gsmarena.com/vv/pics/samsung/samsung-a5-3.jpg);background-size:107%;background-position:0 -45px;-moz-filter:blur(5px);-o-filter:blur(5px);-ms-filter:blur(5px);filter:url("data:image/svg+xml;utf9,<svg%20version='1.1'%20xmlns='http://www.w3.org/2000/svg'><filter%20id='blur'><feGaussianBlur%20stdDeviation='5'%20/></filter></svg>#blur");filter:blur(5px);filter:progid:DXImageTransform.Microsoft.Blur(PixelRadius="5");transform:translateZ(0);opacity:1}.center-stage-background-makers{position:absolute;top:0%;right:0;width:100%;-moz-filter:blur(8px);-o-filter:blur(8px);-ms-filter:blur(8px);filter:url("data:image/svg+xml;utf9,<svg%20version='1.1'%20xmlns='http://www.w3.org/2000/svg'><filter%20id='blur'><feGaussianBlur%20stdDeviation='8'%20/></filter></svg>#blur");filter:blur(8px);filter:progid:DXImageTransform.Microsoft.Blur(PixelRadius="8");transform:translateZ(0);z-index:-1}.article-info-name{background:url(../i/pattern-diag-dark-2.gif) repeat}.article-info-meta li:hover{z-index:1000}.article-info-line .blur{-moz-filter:blur(5px);-o-filter:blur(5px);-ms-filter:blur(5px);filter:url("data:image/svg+xml;utf9,<svg%20version='1.1'%20xmlns='http://www.w3.org/2000/svg'><filter%20id='blur'><feGaussianBlur%20stdDeviation='5'%20/></filter></svg>#blur");filter:blur(5px);filter:progid:DXImageTransform.Microsoft.Blur(PixelRadius="5");transform:translateZ(0);z-index:-1}.article-info-line .blur,.article-info-line .blur:after{position:absolute;top:0px;left:0px;width:100%;height:100%}.article-info-line .blur:after{background-color:rgba(0,0,0,.2);z-index:1;content:""}.makers{position:relative;padding-top:10px;padding-left:5px;left:0;width:100%}.makers ul{overflow:hidden;margin:0;padding:0!important}.makers li,.makers ul{list-style:none!important}.makers li{float:left;height:215px;position:relative;text-align:center;width:143px;margin:0px 0px 10px;padding-top:0}.makers a{text-decoration:none!important;height:100%;background:#fff;overflow:hidden}.makers a,.makers li span{display:block;text-align:center;width:100%}.makers li span{position:absolute;top:46%;-ms-transform:translateY(-50%);transform:translateY(-50%);-webkit-transform:translateY(-50%);-moz-transform:translateY(-50%);left:0;padding:0 15px}*>body .makers img{display:block}.makers img{background:#fff;clear:both;margin:0 auto 0px;padding-top:7px;padding-bottom:7px;width:100px}.makers a:hover img{opacity:0.9}.makers strong{clear:both;display:block;font:700 14px Arimo,Arial,sans-serif;margin:0px auto 0;padding:28px 5px;background:hsla(0,0%,96%,.4);text-decoration:none;color:#777;position:absolute;bottom:12px;width:100%}.makers a:hover strong{color:#fff}.section-body{width:100%}.main-maker.col{padding:0 0px 10px}.main-maker .review-hd{overflow:hidden}.maker-filters-nav{display:-ms-flexbox;display:flex;margin:10px 0 10px 10px;font:12px/1.5 Arimo,Arial,sans-serif}.maker-filters-nav li{display:inline-block;margin-bottom:5px}.maker-filters-nav li a{padding:0px 5px}.maker-filters-nav strong{margin:0 5px}.maker-filters-nav strong a{background:var(--color-button-background,#eee);border-radius:4px;padding:5px}.maker-filters-nav strong a:hover{background:var(--color-button-highlight-background,#d50000);color:var(--color-button-highlight-text,#fff);text-decoration:none}.maker-filters-nav strong a:after{content:"✖";margin-left:5px}.review-nav-v2{background:hsla(0,2%,90%,.8)}.related li{width:25%}.related li img{height:155px;width:auto}.related li:first-child,.related li:nth-child(2){width:50%;height:260px}.related li:first-child img,.related li:nth-child(2) img{height:200px}.related li:first-child strong,.related li:nth-child(2) strong{font:700 15px Arimo,Arial,sans-serif}.related li strong{bottom:0}.compare-help-text{display:none;clear:both;background:var(--color-button-highlight-background,#d50000);color:var(--color-button-highlight-text,#fff);padding:5px 5px 5px 30px;margin-bottom:10px;z-index:1;position:-webkit-sticky;position:sticky;top:0;transition:top 0.25s cubic-bezier(0.55,0,0.1,1)}#header.slide--reset.open~#wrapper .compare-help-text{top:97px}#header.slide--reset~#wrapper .compare-help-text{top:52px}.compare-help-text .icon-rate-down,.compare-help-text .icon-rate-up{font-family:gsmarena;font-size:20px;vertical-align:bottom;line-height:unset}.makers.compare-mode li:after{content:"\e901";font-family:gsmarena;position:absolute;top:0;right:3px;font-size:2em;width:1em;height:1em;cursor:pointer;color:var(--color-button-highlight-background,#d50000);background-color:var(--color-button-highlight-text,#fff);border-radius:3px;line-height:1}.makers.compare-mode li.checked:after{content:"\e941";text-shadow:2px 0 0 var(--color-button-highlight-text,#fff),0 2px 0 var(--color-button-highlight-text,#fff),-2px 0 0 var(--color-button-highlight-text,#fff)}.makers.compare-mode li img{-webkit-user-select:none;-ms-user-select:none;user-select:none}.article-info-meta .article-info-meta-link.compare-button{float:left}.article-info-meta .article-info-meta-link.compare-button .count{display:none;position:absolute;top:5px;right:5px;z-index:20;background:#d50000;color:#fff;text-shadow:none;font-size:15px;line-height:1.4em;width:1.3em;height:1.3em;text-align:center;border-radius:4px}.makers a:hover img{border-color:#d50000}.makers a:hover strong{background-color:#d50000}
						`)),
						h.Div(h.Class("review-header maker-header"),
							h.Div(h.Class("review-hd overflow darken review-background"),
								h.Div(h.Class("article-info"),
									h.Div(h.Class("article-info-line border-bottom"),
										h.Div(h.Class("blur")),
									),
									h.Div(h.Class("center-stage blur color-accent article-accent"),
										h.Div(h.Class("article-hgroup"),
											h.H1(h.Class("article-info-name"), g.Text(fmt.Sprintf("%s phones", brandName))),
										),
									),
									h.Div(h.Class("article-info-line"),
										h.Div(h.Class("blur bottom")),
										h.Ul(h.Class("article-info-meta"),
											h.Li(h.Class("article-info-meta-link help help-sort-popularity"),
												h.A(h.Href("acer-phones-f-59-0-r1-p1.php"), g.Raw(`<i class="head-icon icon-popularity"></i> Popularity`)),
											),
											h.Li(h.Class("article-info-meta-link help help-sort-release sort-active"),
												h.A(h.Href("#"), g.Raw(`<i class="head-icon icon-launched"></i> Time of release`)),
											),
											h.Li(h.Class("article-info-meta-link light large help help-review compare-button"),
												h.A(h.Href("#1"), g.Raw(`<i class="head-icon icon-compare"></i> Compare`)),
											),
											h.Li(h.Class("article-info-meta-link light large help left"),
												h.A(h.Href("#"), g.Raw(fmt.Sprintf(`<i class="head-icon icon-in-the-news"></i> %s news`, brandName))),
											),
										),
									),
								),
							),
						),
						h.Div(h.ID("review-body"), h.Class("section-body"),
							h.Div(h.Class("makers"),
								func() g.Node {
									var nodes []g.Node
									for _, d := range devices {
										nodes = append(nodes, devicePhoneItem(d.ImgSrc, d.Name))
									}
									return h.Ul(h.Style("margin-top: 30px;"), g.Group(nodes))
								}(),
								h.Br(h.Class("clear")),
							),
						),
						h.Br(h.Class("clearfix")),
						paginationControls(brandName, currentPage, totalPages),
					),
					Sidebar(),
				),
			),
			h.Script(g.Raw(`
window.addEventListener("DOMContentLoaded",function(){function e(){const e=n==s,t=l.length<s?"":l.length,o=e||l.length<s?"none":"initial";r.innerText=t,r.style.display=o;const i=n-l.length;i>0?$gsm.isMobile?d.innerHTML='<span class="head-icon icon-rate-down"></span>Tap the checkboxes below to select '+i+(l.length>0?" more ":" ")+"device"+(1==i?"":"s")+(l.length>=s?" or hit the Compare button again":"")+".":d.innerHTML='<span class="head-icon icon-rate-down"></span>Tap the checkboxes below to select up to '+i+(l.length>0?" more ":" ")+"device"+(1==i?"":"s")+(l.length>=s?" or hit the Compare button again":"")+".":d.innerHTML='<span class="head-icon icon-rate-up"></span>Hit the Compare button again to proceed.'}const t=$gsm.isMobile?"cmp-list-m":"cmp-list",n=$gsm.isMobile?2:3,s=2,o="compare.php3?",i=document.querySelector($gsm.isMobile?"#list-brands ul, .general-menu ul":".makers");if(!i)return;const c=i.querySelectorAll("li"),l=[],a=document.querySelector(".compare-button");if(!a)return;const r=$gsm.tag("span",{class:"count"},"",a),h=document.querySelector($gsm.isMobile?".material-card.secondary-menu":".review-header");if(!h)return;const d=$gsm.tag("div",{class:"compare-help-text"},"");h.insertAdjacentElement($gsm.isMobile?"beforeend":"afterend",d),a.onclick=function(n){if(0==l.length){const t=i.classList.toggle("compare-mode");e(),d.style.display=t?"block":""}else{const n=[];for(let e=0;e<l.length;e++){const t=l[e].querySelector("a"),s=(t.href.match(/-(\d+)\.php/)||[0,0])[1];n.push("idPhone"+(e+1)+"="+s)}window.location=o+n.join("&"),e()}n.preventDefault()};for(let t=0;t<c.length;t++)c[t].addEventListener("click",function(t){if(i.classList.contains("compare-mode")){if($gsm.isMobile){if(t.offsetX<.8*window.innerWidth)return}else{const e=84;if(t.offsetY>e||t.offsetX<this.offsetWidth-e)return}if(this.classList.contains("checked")){this.classList.remove("checked");const e=l.indexOf(this);l.splice(e,1)}else for(this.classList.add("checked"),l.push(this);l.length>n;)l[0].classList.remove("checked"),l.shift();e(),t.preventDefault()}})});
		`)),
		),
	)
}

func devicePhoneItem(imgSrc string, name string) g.Node {
	// Bypass hotlink protection via an image proxy OR set referrer policy.
	// We use the weserv proxy here to guarantee it displays even from localhost
	proxySrc := "https://images.weserv.nl/?url=" + imgSrc
	return h.Li(
		h.A(h.Href("#"),
			h.Img(h.Src(proxySrc), h.Alt(name)),
			h.Strong(h.Span(g.Text(name))),
		),
	)
}

func paginationControls(brandName string, current, total int) g.Node {
	if total <= 1 {
		return h.Div(h.Class("review-nav-v2"), h.Div(h.Class("nav-pages"))) // empty
	}

	var nodes []g.Node

	// Prev button
	if current > 1 {
		nodes = append(nodes, h.A(h.Href(fmt.Sprintf("/brand?brand=%s&page=%d", brandName, current-1)), h.Class("prevnextbutton"), h.Title("Previous page"), g.Text("◄")))
	} else {
		nodes = append(nodes, h.A(h.Href("#"), h.Class("prevnextbuttondis"), g.Text("◄")))
	}

	// Pages
	for p := 1; p <= total; p++ {
		if p == current {
			nodes = append(nodes, h.Strong(g.Text(fmt.Sprintf("%d", p))))
		} else {
			nodes = append(nodes, h.A(h.Href(fmt.Sprintf("/brand?brand=%s&page=%d", brandName, p)), g.Text(fmt.Sprintf("%d", p))))
		}
	}

	// Next button
	if current < total {
		nodes = append(nodes, h.A(h.Href(fmt.Sprintf("/brand?brand=%s&page=%d", brandName, current+1)), h.Class("prevnextbutton"), h.Title("Next page"), g.Text("►")))
	} else {
		nodes = append(nodes, h.A(h.Href("#"), h.Class("prevnextbuttondis"), g.Text("►")))
	}

	return h.Div(h.Class("review-nav-v2"),
		h.Div(h.Class("nav-pages"),
			g.Group(nodes),
		),
	)
}
