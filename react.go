package main

import (
	"fmt"
	"html/template"

	"github.com/zserge/webview"
)

var uiFrameworkName = "ReactJS+Babel"

func loadUIFramework(w webview.WebView) {
	// Inject React and Babel
	w.Eval(string(MustAsset("assets/react/vendor/babel.min.js")))
	w.Eval(string(MustAsset("assets/react/vendor/preact.min.js")))

	// Inject our app code
	w.Eval(fmt.Sprintf(`(function(){
		var n=document.createElement('script');
		n.setAttribute('type', 'text/babel');
		n.appendChild(document.createTextNode("%s"));
		document.body.appendChild(n);
	})()`, template.JSEscapeString(string(MustAsset("assets/react/app.jsx")))))

	// Process our code with Babel
	w.Eval(`Babel.transformScriptTags()`)
}
