package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	//Static for Access Folder
	e.Static("assets/", "assets")

	//Routing
	e.GET("/", home)                       // localhost:5000/
	e.GET("/contactMe", contactMe)         // localhost:5000/contactMe
	e.GET("/addProject", addProject)       // localhost:5000/addProject
	e.GET("/projectDetail", projectDetail) // localhost:5000/projectDetail
	e.POST("/addProject", formAddProject)  //localhost:5000/formAddProject

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	template, err := template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
	}

	return template.Execute(c.Response().Writer, nil)
}

func contactMe(c echo.Context) error {
	template, err := template.ParseFiles("views/contact-me.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
	}

	return template.Execute(c.Response().Writer, nil)
}

func addProject(c echo.Context) error {
	template, err := template.ParseFiles("views/add-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
	}

	return template.Execute(c.Response().Writer, nil)
}

func projectDetail(c echo.Context) error {
	template, err := template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
	}

	return template.Execute(c.Response(), nil)
}

func formAddProject(c echo.Context) error {
	title := c.FormValue("TitleProject")
	content := c.FormValue("ContentProject")

	fmt.Println(title)
	fmt.Println(content)

	return c.Redirect(http.StatusMovedPermanently, "/")
}
