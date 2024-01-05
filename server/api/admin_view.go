package api

import (
	"net/http"
	"text/template"

	"github.com/auula/vasedb/clog"
)

const (
	// 默认的 HTML 文件文本
	loginHtml     = "text/template"
	dashboardHtml = "text/template"
)

// AdminTemplates 结构体用于存储所有后台模板
type AdminTemplates struct {
	Login     *template.Template
	Dashboard *template.Template
}

// 能被渲染全局管理员
var templates AdminTemplates

func init() {
	// 根据 html 文件来构造后台 view 的模版
	templates.Login = template.Must(template.
		New("login").
		Parse(loginHtml))
	templates.Dashboard = template.Must(template.
		New("dashboard").
		Parse(dashboardHtml))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// 使用 Login 渲染登录页面
	data := map[string]interface{}{
		"Msg": "使用 Login 渲染登录页面",
	}
	err := templates.Login.Execute(w, data)
	if err != nil {
		clog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	// 使用 Dashboard 渲染仪表盘页面
	data := map[string]interface{}{
		"Msg": "使用 Dashboard 渲染仪表盘页面",
	}
	err := templates.Dashboard.Execute(w, data)
	if err != nil {
		clog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
