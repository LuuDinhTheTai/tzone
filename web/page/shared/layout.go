package shared

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func BaseLayout(title string, children g.Node) g.Node {
	return h.Doctype(
		h.HTML(g.Attr("xmlns", "http://www.w3.org/1999/xhtml"), g.Attr("xml:lang", "en-US"), h.Lang("en-US"),
			h.Head(
				h.TitleEl(g.Text(title)),
				h.Script(g.Raw(`
DESKTOP_BASE_URL = "https://www.gsmarena.com/";
MOBILE_BASE_URL = "https://m.gsmarena.com/";
ASSETS_BASE_URL  = "https://fdn.gsmarena.com/vv/assets12/";
CDN_BASE_URL = "//fdn.gsmarena.com/";
CDN2_BASE_URL = "//fdn2.gsmarena.com/";
`)),
				h.Meta(h.Charset("utf-8")),
				h.Meta(h.Name("viewport"), h.Content("width=1060, initial-scale=1.0")),
				h.Link(h.Rel("stylesheet"), h.Href("https://fdn.gsmarena.com/vv/assets12/css/gsmarena.css?v=169")),
				h.Link(h.Rel("shortcut icon"), h.Href("https://fdn.gsmarena.com/imgroot/static/favicon.ico")),

				h.Script(g.Raw(`
window["pgGlobalSettings"] = {
    "global": {
        "strategy": "include"
    },
    "adUnits": []
};

function addAdUnit(adUnitCode) {
    window.pgGlobalSettings.adUnits.push({ "adUnitCode": adUnitCode });
};

	addAdUnit("/8095840,14566801/.2_A.35723.3_gsmarena.com_tier1");
	addAdUnit("/8095840,14566801/.2_A.34909.4_gsmarena.com_tier1");
	addAdUnit("/8095840,14566801/.2_A.34911.7_gsmarena.com_tier1");
	addAdUnit("/8095840,14566801/.2_A.40299.4_gsmarena.com_tier1");
	addAdUnit("/8095840,14566801/.2_A.40300.4_gsmarena.com_tier1");
	addAdUnit("/8095840,14566801/.2_A.38135.3_gsmarena.com_tier1");
	addAdUnit("/8095840,14566801/.2_A.46956.3_gsmarena.com_tier1");
`)),
				h.Script(g.Raw(` window.pbjs = {que: []}; `)),
				h.Script(h.Type("text/javascript"), g.Raw(`
!function(){var i,r,o;i="__tcfapiLocator",r=[],(o=window.frames[i])||(function e(){var t=window.document,a=!!o;if(!a)if(t.body){var n=t.createElement("iframe");n.style.cssText="display:none",n.name=i,t.body.appendChild(n)}else setTimeout(e,5);return!a}(),window.__tcfapi=function(){for(var e,t=[],a=0;a<arguments.length;a++)t[a]=arguments[a];if(!t.length)return r;if("setGdprApplies"===t[0])3<t.length&&2===parseInt(t[1],10)&&"boolean"==typeof t[3]&&(e=t[3],"function"==typeof t[2]&&t[2]("set",!0));else if("ping"===t[0]){var n={gdprApplies:e,cmpLoaded:!1,cmpStatus:"stub"};"function"==typeof t[2]&&t[2](n,!0)}else r.push(t)},window.addEventListener("message",function(n){var i="string"==typeof n.data,e={};try{e=i?JSON.parse(n.data):n.data}catch(e){}var r=e.__tcfapiCall;r&&window.__tcfapi(r.command,r.version,function(e,t){var a={__tcfapiReturn:{returnValue:e,success:t,callId:r.callId}};i&&(a=JSON.stringify(a)),n.source.postMessage(a,"*")},r.parameter)},!1))}();
!function(){var i,n,s;i="__uspapiLocator",n=[],(s=window.frames[i])||(function a(){var e=window.document,n=!!s;if(!s)if(e.body){var t=e.createElement("iframe");t.style.cssText="display:none",t.name=i,e.body.appendChild(t)}else setTimeout(a,5);return!n}(),window.__uspapi=function(){for(var a=[],e=0;e<arguments.length;e++)a[e]=arguments[e];if(!a.length)return n;"ping"===a[0]?"function"==typeof a[2]&&a[2]({cmpLoaded:!1,cmpStatus:"stub"},!0):n.push(a)},window.addEventListener("message",function(t){var i="string"==typeof t.data,a={};try{a=i?JSON.parse(t.data):t.data}catch(a){}var s=a.__uspapiCall;s&&window.__uspapi(s.command,s.version,function(a,e){var n={__uspapiReturn:{returnValue:a,success:e,callId:s.callId}};i&&(n=JSON.stringify(n)),t.source.postMessage(n,"*")},s.parameter)},!1))}();
window.__gpp_addFrame=function(e){if(!window.frames[e])if(document.body){var p=document.createElement("iframe");p.style.cssText="display:none",p.name=e,document.body.appendChild(p)}else window.setTimeout(window.__gppaddFrame,10,e)},window.__gpp_stub=function(){var e=arguments;if(__gpp.queue=__gpp.queue||[],!e.length)return __gpp.queue;var p=e[0],t=1<e.length?e[1]:null,n=2<e.length?e[2]:null;if("ping"===p)return{gppVersion:"1.0",cmpStatus:"stub",cmpDisplayStatus:"hidden",apiSupport:["tcfeuv2","tcfcav1","uspv1","uspnatv1","uspcav1","uspvav1","uspcov1","usputv1","uspctv1"],currentAPI:"",cmpId:68};if("addEventListener"===p){__gpp.events=__gpp.events||[],"lastId"in __gpp||(__gpp.lastId=0),__gpp.lastId++;var a=__gpp.lastId;return __gpp.events.push({id:a,callback:t,parameter:n}),{eventName:"listenerRegistered",listenerId:a,data:!0}}if("removeEventListener"===p){var _=!1;__gpp.events=__gpp.events||[];for(var s=0;s<__gpp.events.length;s++)if(__gpp.events[s].id==n){__gpp.events[s].splice(s,1),_=!0;break}return{eventName:"listenerRemoved",listenerId:n,data:_}}if("hasSection"===p||"getSection"===p||"getField"===p||"getGPPData"===p)return null;__gpp.queue.push([].slice.apply(e))},window.__gpp_msghandler=function(n){var a="string"==typeof n.data;try{var p=a?JSON.parse(n.data):n.data}catch(e){p=null}if("object"==typeof p&&null!==p&&"__gppCall"in p){var _=p.__gppCall;window.__gpp(_.command,function(e,p){var t={__gppReturn:{returnValue:e,success:p,callId:_.callId}};n.source.postMessage(a?JSON.stringify(t):t,"*")},_.parameter)}},"__gpp"in window&&"function"==typeof window.__gpp||(window.__gpp=window.__gpp_stub,window.addEventListener("message",window.__gpp_msghandler,!1),window.__gpp_addFrame("__gppLocator"));
`)),
				h.Script(h.Async(), h.Src("https://cmp.uniconsent.com/v2/de538b5a3a/cmp.js")),
				h.Script(h.Type("text/javascript"), g.Raw(`
window.googletag = window.googletag || {};
window.googletag.cmd = window.googletag.cmd || [];
window.googletag.cmd.push(function () {
	window.googletag.pubads().setTargeting('pageid', '1');
    window.googletag.pubads().setTargeting('country', 'VN');
	window.googletag.pubads().setTargeting('lpe', 'low');
	window.googletag.pubads().setTargeting('keyw', 'Samsung');
    window.googletag.pubads().setTargeting('visitqos', '0');
    window.googletag.pubads().enableAsyncRendering();
    window.googletag.pubads().disableInitialLoad();
});
(adsbygoogle = window.adsbygoogle || []).pauseAdRequests = 1;
`)),
				h.Script(g.Raw(`
(function waitGEO() {
    var readyGEO;
    if (window['UnicI'] && window['UnicI'].geo && window['UnicI'].geo !== '-' ) {
        readyGEO = true;
        console.log(window['UnicI'].geo);
        if (window['UnicI'].geo === 'EU') {
            if(document.getElementById("unic-gdpr")) {
              document.getElementById("unic-gdpr").style.display = 'inline';
            }
        }
        if (window['UnicI'].geo === 'CA') {
            if(document.getElementById("unic-ccpa")) {
              document.getElementById("unic-ccpa").style.display = 'inline';
            }
        }
    }
    if (!readyGEO) {
        setTimeout(waitGEO, 200);
    }
})();
`)),
				h.Script(g.Raw(`
__tcfapi("addEventListener", 2, function(tcData, success) {
    if (success && tcData.unicLoad  === true) {
        if(!window._initAds) {
            window._initAds = true;
            var script = document.createElement('script');
            script.async = true;
            script.src = '//dsh7ky7308k4b.cloudfront.net/publishers/gsmarenacom_new_d.min.js';
            document.head.appendChild(script);
        
	    var script = document.createElement('script');
            script.async = true;
            script.src = '//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js';
            document.head.appendChild(script);		

        var s = document.createElement('script');
			s.async = true;
			s.src = '//btloader.com/tag?o=5184339635601408&upapi=true';
			document.head.appendChild(s);    
	}
    }
});
`)),
				// Google tag (gtag.js)
				h.Script(h.Async(), h.Src("https://www.googletagmanager.com/gtag/js?id=G-WECNNBCHQE")),
				h.Script(g.Raw(`
	window.dataLayer = window.dataLayer || [];
	function gtag(){dataLayer.push(arguments);}
	gtag('js', new Date());

	gtag('set', {
		'c_pagetype': '1'
	});

	gtag('config', 'G-WECNNBCHQE');
`)),
				h.Link(h.Rel("stylesheet"), h.Href("https://fdn.gsmarena.com/vv/assets12/css/index.css?v=15")),
				h.Meta(h.Name("description"), h.Content("GSMArena.com - The ultimate resource for GSM handset information")),
				h.Meta(h.Name("keywords"), h.Content("GSM,mobile,phone,Nokia,Sony,Apple,iPhone,Motorola,Alcatel,Xiaomi,Samsung,Oppo,Oneplus,cellphone,specifications,information,info,opinion,review,pictures,photos")),
				h.Link(h.Rel("alternate"), h.Type("application/rss+xml"), h.Title("GSMArena.com RSS feed"), h.Href("https://www.gsmarena.com/rss-news-reviews.php3")),
				h.Link(h.Rel("canonical"), h.Href("https://www.gsmarena.com")),
				h.Link(h.Rel("alternate"), g.Attr("media", "only screen and (max-width: 640px)"), h.Href("https://m.gsmarena.com")),
			),
			h.Body(
				h.Script(h.Type("text/javascript"), h.Src("https://fdn.gsmarena.com/vv/assets12/js/misc.js?v=128")),

				Header(),

				children,

				Footer(),

				h.Script(h.Type("text/javascript"), h.Src("https://fdn.gsmarena.com/vv/assets12/js/autocomplete.js?v=16")),
				h.Script(h.Type("text/javascript"), h.Lang("javascript"), g.Raw(`
AUTOCOMPLETE_LIST_URL = "/quicksearch-81833.jpg";
$gsm.addEventListener(document, "DOMContentLoaded", function() 
{
    new Autocomplete( "topsearch-text", "topsearch", true );
}
)
`)),
				h.Script(h.Type("text/javascript"), h.Src("https://fdn.gsmarena.com/vv/assets12/js/infinite-scroll.js?v=50")),
			),
		),
	)
}
