package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type JSMenu struct {
	Menu []JSMenuItem `json:"menu"`
}

type JSMenuItem struct {
	ID string `json:"id"`
	Href string `json:"href"`
	Title string `json:"title"`
	SubMenu []JSMenuItem `json:"submenu"`
}

func (c *JSMenu) New() {
	c.Menu = make([]JSMenuItem, 0)
}

func (c *JSMenu) Add(item JSMenuItem) {
	if item.SubMenu == nil {
		item.SubMenu = make([]JSMenuItem, 0)
	}
	item.SubMenu = append(item.SubMenu, JSMenuItem{ID: fmt.Sprintf("%snew", item.ID), Href: fmt.Sprintf("%s/new", item.Href), Title: "new"})
	item.SubMenu = append(item.SubMenu, JSMenuItem{ID: fmt.Sprintf("%slist", item.ID), Href: fmt.Sprintf("%s/list", item.Href), Title: "list"})
	c.Menu = append(c.Menu, item)
}

func RegisterAPIPageRoutes(e *echo.Echo) {
	g := e.Group("/api/page")
	g.GET("/menu/:id", handleGetMenu)
	g.GET("/submenu/:id", handleGetSubmenu)
} 

func handleGetMenu(e echo.Context) error {
	GetLogger(4).Flogger("handleGetMenu called")
	jsmenu := JSMenu{}
	jsmenu.New()
	jsmenu.Add(JSMenuItem{ID: "home", Href: "/", Title: "home"})
	jsmenu.Add(JSMenuItem{ID: "prompts", Href: "/prompts", Title: "prompts"})
	jsmenu.Add(JSMenuItem{ID: "jobs", Href: "/jobs", Title: "jobs"})
	jsmenu.Add(JSMenuItem{ID: "jobruns", Href: "/jobruns", Title: "job runs"})
	jsmenu.Add(JSMenuItem{ID: "dispositions", Href: "/dispositions", Title: "dispositions"})
	jsmenu.Add(JSMenuItem{ID: "comfyuitemplates", Href: "/comfyuitemplates", Title: "ComfyUI Templates"})
	jsmenu.Add(JSMenuItem{ID: "systemprompts", Href:"/systemprompts", Title: "System Prompts"})
	jsmenu.Add(JSMenuItem{ID: "workflow", Href:"/workflows", Title: "Workflows"})
	jsmenu.Add(JSMenuItem{ID: "prompttemplates", Href:"/prompttemplates", Title: "Prompt Templates"})
	jsmenu.Add(JSMenuItem{ID: "steps", Href:"/steps", Title:"Steps"})

	return e.JSON(http.StatusOK, jsmenu)
}

func handleGetSubmenu(e echo.Context) error {
	if id := e.Param("id"); id != "" {
		msi := make(map[string]interface{})
		msi["new"] = fmt.Sprintf("/%s/new", id)
		msi["list"] = fmt.Sprintf("/%s/list", id)
		return e.JSON(http.StatusOK, msi)
	}
	return e.JSON(http.StatusBadRequest, "bad request: missing id")
}