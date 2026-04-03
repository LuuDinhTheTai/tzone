package page

import (
	"github.com/LuuDinhTheTai/tzone/web/page/shared"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func AllBrandPage() g.Node {
	return shared.BaseLayout("List of all mobile phone brands - GSMArena.com",
		h.Div(h.ID("wrapper"), h.Class("l-container"),
			h.Div(h.ID("outer"), h.Class("row"),
				h.Div(h.ID("subHeader3")),
				h.Div(h.ID("body"), h.Class("clearfix"),
					h.Div(h.Class("main main-makers l-box col float-right"),
						h.StyleEl(g.Raw(`
.review-background {
    background-image: url('https://images.unsplash.com/photo-1511707171634-5f897ff02aa9?ixlib=rb-4.0.3&auto=format&fit=crop&w=1200&q=80');
    background-size: cover;
    background-position: center top;
    background-repeat: no-repeat;
    min-height: 250px;
    position: relative;
}
.main-makers .st-text {
    padding: 30px;
}
.main-makers table {
    width: 100%;
}
.main-makers table td {
    width: 50%;
    padding: 0;
    background: transparent;
    border: none;
}
.main-makers table td a {
    display: block;
    font-family: Google-Oswald, Arimo, Arial, sans-serif;
    text-transform: uppercase;
    text-decoration: none;
    font-size: 35px;
    color: #999;
    padding: 15px 10px 10px;
    transition: all 0.2s ease;
}
.main-makers table td a:hover {
    color: #d50000;
    background: #f9f9f9;
}
.main-makers table td a span {
    font-weight: 400;
    font-size: 20px;
    display: block;
    font-family: Arimo, Arial, sans-serif;
    margin-top: 10px;
    text-transform: none;
}
						`)),
						h.Div(h.Style("position: relative; margin-bottom: 80px;"),
							h.Img(h.Src("https://images.unsplash.com/photo-1511707171634-5f897ff02aa9?ixlib=rb-4.0.3&auto=format&fit=crop&w=1200&q=80"), h.Alt("All Brands Banner"), h.Style("width: 100%; height: auto; display: block; object-fit: cover; max-height: 250px;")),
							h.Div(h.Style("position: absolute; bottom: 30px; left: 30px; background: rgba(255,255,255,0.85); padding: 5px 15px;"),
								h.H1(g.Text("All mobile phone brands"), h.Style("color: #333; font-size: 28px; font-weight: normal; margin: 0; text-transform: uppercase; font-family: Google-Oswald, Arimo, Arial, sans-serif;")),
							),
						),
						h.Div(h.Class("st-text"),
							g.Raw(`
<table>
<tr><td><a href="/brand?brand=acer">Acer<br><span>113 devices</span></a></td>
<td><a href="/brand?brand=alcatel">alcatel<br><span>414 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=allview">Allview<br><span>157 devices</span></a></td>
<td><a href="/brand?brand=amazon">Amazon<br><span>25 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=amoi">Amoi<br><span>47 devices</span></a></td>
<td><a href="/brand?brand=apple">Apple<br><span>147 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=archos">Archos<br><span>43 devices</span></a></td>
<td><a href="/brand?brand=asus">Asus<br><span>207 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=at&t">AT&T<br><span>4 devices</span></a></td>
<td><a href="/brand?brand=benefon">Benefon<br><span>9 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=benq">BenQ<br><span>35 devices</span></a></td>
<td><a href="/brand?brand=benq_siemens">BenQ-Siemens<br><span>28 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=bird">Bird<br><span>61 devices</span></a></td>
<td><a href="/brand?brand=blackberry">BlackBerry<br><span>92 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=blackview">Blackview<br><span>113 devices</span></a></td>
<td><a href="/brand?brand=blu">BLU<br><span>369 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=bosch">Bosch<br><span>10 devices</span></a></td>
<td><a href="/brand?brand=bq">BQ<br><span>20 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=casio">Casio<br><span>5 devices</span></a></td>
<td><a href="/brand?brand=cat">Cat<br><span>22 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=celkon">Celkon<br><span>229 devices</span></a></td>
<td><a href="/brand?brand=chea">Chea<br><span>12 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=coolpad">Coolpad<br><span>56 devices</span></a></td>
<td><a href="/brand?brand=cubot">Cubot<br><span>100 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=dell">Dell<br><span>20 devices</span></a></td>
<td><a href="/brand?brand=doogee">Doogee<br><span>157 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=emporia">Emporia<br><span>15 devices</span></a></td>
<td><a href="/brand?brand=energizer">Energizer<br><span>78 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=ericsson">Ericsson<br><span>40 devices</span></a></td>
<td><a href="/brand?brand=eten">Eten<br><span>22 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=fairphone">Fairphone<br><span>5 devices</span></a></td>
<td><a href="/brand?brand=fujitsu_siemens">Fujitsu Siemens<br><span>2 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=garmin_asus">Garmin-Asus<br><span>5 devices</span></a></td>
<td><a href="/brand?brand=gigabyte">Gigabyte<br><span>63 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=gionee">Gionee<br><span>95 devices</span></a></td>
<td><a href="/brand?brand=google">Google<br><span>40 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=haier">Haier<br><span>59 devices</span></a></td>
<td><a href="/brand?brand=hmd">HMD<br><span>36 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=honor">Honor<br><span>303 devices</span></a></td>
<td><a href="/brand?brand=hp">HP<br><span>41 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=htc">HTC<br><span>296 devices</span></a></td>
<td><a href="/brand?brand=huawei">Huawei<br><span>516 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=i_mate">i-mate<br><span>34 devices</span></a></td>
<td><a href="/brand?brand=i_mobile">i-mobile<br><span>37 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=icemobile">Icemobile<br><span>61 devices</span></a></td>
<td><a href="/brand?brand=infinix">Infinix<br><span>169 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=innostream">Innostream<br><span>18 devices</span></a></td>
<td><a href="/brand?brand=inq">iNQ<br><span>5 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=intex">Intex<br><span>15 devices</span></a></td>
<td><a href="/brand?brand=itel">itel<br><span>61 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=jolla">Jolla<br><span>5 devices</span></a></td>
<td><a href="/brand?brand=karbonn">Karbonn<br><span>60 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=kyocera">Kyocera<br><span>28 devices</span></a></td>
<td><a href="/brand?brand=lava">Lava<br><span>177 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=leeco">LeEco<br><span>9 devices</span></a></td>
<td><a href="/brand?brand=lenovo">Lenovo<br><span>259 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=lg">LG<br><span>667 devices</span></a></td>
<td><a href="/brand?brand=maxon">Maxon<br><span>31 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=maxwest">Maxwest<br><span>41 devices</span></a></td>
<td><a href="/brand?brand=meizu">Meizu<br><span>87 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=micromax">Micromax<br><span>289 devices</span></a></td>
<td><a href="/brand?brand=microsoft">Microsoft<br><span>32 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=mitac">Mitac<br><span>12 devices</span></a></td>
<td><a href="/brand?brand=mitsubishi">Mitsubishi<br><span>25 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=modu">Modu<br><span>8 devices</span></a></td>
<td><a href="/brand?brand=motorola">Motorola<br><span>683 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=mwg">MWg<br><span>5 devices</span></a></td>
<td><a href="/brand?brand=nec">NEC<br><span>73 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=neonode">Neonode<br><span>3 devices</span></a></td>
<td><a href="/brand?brand=niu">NIU<br><span>30 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=nokia">Nokia<br><span>592 devices</span></a></td>
<td><a href="/brand?brand=nothing">Nothing<br><span>15 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=nvidia">Nvidia<br><span>3 devices</span></a></td>
<td><a href="/brand?brand=o2">O2<br><span>45 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=oneplus">OnePlus<br><span>105 devices</span></a></td>
<td><a href="/brand?brand=oppo">Oppo<br><span>402 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=orange">Orange<br><span>19 devices</span></a></td>
<td><a href="/brand?brand=oscal">Oscal<br><span>31 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=oukitel">Oukitel<br><span>104 devices</span></a></td>
<td><a href="/brand?brand=palm">Palm<br><span>17 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=panasonic">Panasonic<br><span>123 devices</span></a></td>
<td><a href="/brand?brand=pantech">Pantech<br><span>72 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=parla">Parla<br><span>10 devices</span></a></td>
<td><a href="/brand?brand=philips">Philips<br><span>233 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=plum">Plum<br><span>113 devices</span></a></td>
<td><a href="/brand?brand=posh">Posh<br><span>30 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=prestigio">Prestigio<br><span>56 devices</span></a></td>
<td><a href="/brand?brand=qmobile">QMobile<br><span>90 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=qtek">Qtek<br><span>21 devices</span></a></td>
<td><a href="/brand?brand=razer">Razer<br><span>2 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=realme">Realme<br><span>285 devices</span></a></td>
<td><a href="/brand?brand=rugone">RugOne<br><span>2 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=sagem">Sagem<br><span>120 devices</span></a></td>
<td><a href="/brand?brand=samsung">Samsung<br><span>1454 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=sendo">Sendo<br><span>19 devices</span></a></td>
<td><a href="/brand?brand=sewon">Sewon<br><span>25 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=sharp">Sharp<br><span>82 devices</span></a></td>
<td><a href="/brand?brand=siemens">Siemens<br><span>94 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=sonim">Sonim<br><span>24 devices</span></a></td>
<td><a href="/brand?brand=sony">Sony<br><span>162 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=sony_ericsson">Sony Ericsson<br><span>188 devices</span></a></td>
<td><a href="/brand?brand=spice">Spice<br><span>120 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=t_mobile">T-Mobile<br><span>69 devices</span></a></td>
<td><a href="/brand?brand=tcl">TCL<br><span>100 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=tecno">Tecno<br><span>185 devices</span></a></td>
<td><a href="/brand?brand=tel_me_">Tel.Me.<br><span>7 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=telit">Telit<br><span>30 devices</span></a></td>
<td><a href="/brand?brand=thuraya">Thuraya<br><span>2 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=toshiba">Toshiba<br><span>35 devices</span></a></td>
<td><a href="/brand?brand=ulefone">Ulefone<br><span>145 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=umidigi">Umidigi<br><span>80 devices</span></a></td>
<td><a href="/brand?brand=unnecto">Unnecto<br><span>30 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=vertu">Vertu<br><span>17 devices</span></a></td>
<td><a href="/brand?brand=verykool">verykool<br><span>139 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=vivo">vivo<br><span>583 devices</span></a></td>
<td><a href="/brand?brand=vk_mobile">VK Mobile<br><span>31 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=vodafone">Vodafone<br><span>87 devices</span></a></td>
<td><a href="/brand?brand=wiko">Wiko<br><span>102 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=wnd">WND<br><span>5 devices</span></a></td>
<td><a href="/brand?brand=xcute">XCute<br><span>4 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=xiaomi">Xiaomi<br><span>507 devices</span></a></td>
<td><a href="/brand?brand=xolo">XOLO<br><span>81 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=yezz">Yezz<br><span>113 devices</span></a></td>
<td><a href="/brand?brand=yota">Yota<br><span>3 devices</span></a></td></tr>

<tr><td><a href="/brand?brand=yu">YU<br><span>13 devices</span></a></td>
<td><a href="/brand?brand=zte">ZTE<br><span>430 devices</span></a></td></tr>

</table>
							`),
						),
					),
					Sidebar(),
				),
			),
		),
	)
}
