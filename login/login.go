package login

import (
	"honnef.co/go/js/dom"
)

var lfid string

func Create(id string, fn func()) *dom.DocumentFragment {

	lfid = id

	window := dom.GetWindow().(*dom.Window)
	doc := window.Document().(*dom.Document)
	f := doc.CreateDocumentFragment().(*dom.DocumentFragment)

	// basic container and framework for the form
	div := doc.CreateElement("div").(*dom.HTMLDivElement)
	div.SetClass("container")
	div.SetID(id)
	f.AppendChild(div)

	row := doc.CreateElement("div").(*dom.HTMLDivElement)
	row.SetClass("row")
	div.AppendChild(row)

	col := doc.CreateElement("div").(*dom.HTMLDivElement)
	col.SetClass("col s6 offset-s3")
	row.AppendChild(col)

	h3 := doc.CreateElement("h3").(*dom.HTMLHeadingElement)
	h3.SetTextContent("Login")
	h3.SetClass("center-align")
	col.AppendChild(h3)

	// username
	userRow := doc.CreateElement("div").(*dom.HTMLDivElement)
	userRow.SetClass("row")
	userDiv := doc.CreateElement("div").(*dom.HTMLDivElement)
	userDiv.SetClass("input-field col s12")
	userInput := doc.CreateElement("input").(*dom.HTMLInputElement)
	userInput.Placeholder = "User Name"
	userInput.Type = "text"
	userInput.ID(id + "-username")
	userInput.SetClass("validate")
	userLabel := doc.CreateElement("label").(*dom.HTMLLabelElement)
	userLabel.For = id + "-username"
	userLabel.SetTextContent("User Name")
	userDiv.AppendChild(userInput)
	userDiv.AppendChild(userLabel)
	userRow.AppendChild(userDiv)
	col.AppendChild(userRow)

	// passwd
	pwRow := doc.CreateElement("div").(*dom.HTMLDivElement)
	pwRow.SetClass("row")
	pwDiv := doc.CreateElement("div").(*dom.HTMLDivElement)
	pwDiv.SetClass("input-field col s12")
	pwInput := doc.CreateElement("input").(*dom.HTMLInputElement)
	pwInput.Placeholder = "Enter Your Password"
	pwInput.Type = "password"
	pwInput.ID(id + "-pw")
	pwInput.SetClass("validate")
	pwLabel := doc.CreateElement("label").(*dom.HTMLLabelElement)
	pwLabel.For = id + "-pw"
	pwLabel.SetTextContent("Password")
	pwDiv.AppendChild(pwInput)
	pwDiv.AppendChild(pwLabel)
	pwRow.AppendChild(pwDiv)
	col.AppendChild(pwRow)

	// remember me && submit btn on single row
	remRow := doc.CreateElement("div").(*dom.HTMLDivElement)
	remRow.SetClass("row")
	remDiv := doc.CreateElement("div").(*dom.HTMLDivElement)
	remDiv.SetClass("input-field col s6")
	remInput := doc.CreateElement("input").(*dom.HTMLInputElement)
	remInput.Type = "checkbox"
	remInput.ID(id + "-remember")
	remLabel := doc.CreateElement("label").(*dom.HTMLLabelElement)
	remLabel.For = id + "-remember"
	remLabel.SetTextContent("Remember Me ?")
	remDiv.AppendChild(remInput)
	remDiv.AppendChild(remLabel)
	remRow.AppendChild(remDiv)
	subDiv := doc.CreateElement("div").(*dom.HTMLDivElement)
	subDiv.SetClass("col s6")
	subBtn := doc.CreateElement("button").(*dom.HTMLButtonElement)
	subBtn.SetClass("btn btn-large waves-effect waves-light")
	subBtn.Name = "Login"
	subBtn.Type = "button"
	subBtn.SetTextContent("Login")
	subDiv.AppendChild(subBtn)
	subBtn.AddEventListener("click", false, func(event dom.Event) {
		go fn()
	})
	col.AppendChild(remRow)

	// Hide it for now
	div.Style().SetProperty("display", "none", "")
	return div
}

func Show() {

	window := dom.GetWindow().(*dom.Window)
	doc := window.Document().(*dom.Document)

	// hide splash and show login
	sp := doc.GetElementByID("splash").(*dom.HTMLDivElement)
	sp.Style().SetProperty("display", "none", "")
	lf := doc.GetElementByID(lfid).(*dom.HTMLDivElement)
	lf.Style().SetProperty("display", "inline", "")
	lb := doc.GetElementByID("nav-login-btn")
	lb.SetTextContent("Login")

	// It the menu exists, remove it
	sb := doc.GetElementByID("sidebar-menu")
	if sb != nil {
		layout := doc.GetElementByID("layout")
		layout.RemoveChild(sb)
		print("removing menu")
	} else {
		print("no menu to remove")
	}
}

func Hide() {

	window := dom.GetWindow().(*dom.Window)
	doc := window.Document().(*dom.Document)

	// hide login, and change to logout btn
	lf := doc.GetElementByID(lfid).(*dom.HTMLDivElement)
	lf.Style().SetProperty("display", "none", "")
	lb := doc.GetElementByID("nav-login-btn")
	lb.SetTextContent("Logout")
}
