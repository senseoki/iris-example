package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// PageController is ...
type PageController struct {
	Ctx iris.Context
}

// BeforeActivation is ...
func (c *PageController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/", "DashBoard")
	b.Handle("GET", "/report/criteria", "CriteriaReport")
}

// DashBoard is ...
func (c *PageController) DashBoard() mvc.Result {
	return mvc.View{
		Layout: "enliple/ads.layout.html",
		Name:   "enliple/ads.yield.html",
		Data: iris.Map{
			"content": "enliple/dashboard.html",
		},
	}
}

// CriteriaReport is ...
func (c *PageController) CriteriaReport() mvc.Result {
	return mvc.View{
		Layout: "enliple/ads.layout.html",
		Name:   "enliple/ads.yield.html",
		Data: iris.Map{
			"content": "enliple/criteria_report.html",
			"js": []string{
				"js/enliple/criteria_report.js",
			},
		},
	}
}
