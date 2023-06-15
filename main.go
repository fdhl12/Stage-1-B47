package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
	"personalWeb/connection"
	"regexp"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)
type Project struct {
	Id				int
	Name			string
	StartDate   	time.Time
	EndDate     	time.Time
	Duration    	string
	Description 	string
	Html    		bool
	Css     		bool
	Reactjs  		bool
	Javascript		bool
	Image			string
	StartDateTime	string
	EndDateTime		string
}
type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

type SessionData struct {
	IsLogin bool
	Name    string
}

var userData = SessionData{}

type Template struct {
	templates *template.Template
}
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// fmt.Print(t.templates.ExecuteTemplate(w, name, data))
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	connection.DatabaseConnect()
	e := echo.New()

	e.Static("/public", "public")

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	e.GET("/", home)
	e.GET("/project", project)
	e.GET("/contact", contact)
	e.GET("/project/:id", detailproject)
	e.GET("/testimonials", testimonials)
	e.POST("/add-project", addNewProject)
	e.GET("formAddProject/:id", formAddProject)
	e.POST("/deleteProject/:id", deleteProject)
	e.POST("/edit-project/:id", ressEditProject)
	e.GET("/edit-project/:id", editProject)

	//SIGNUP
	e.GET("/register", signupForm)
	e.POST("/signup", singup)

	//LOGIN
	e.GET("/login", loginForm)
	e.POST("/login", login)

	//LOGOuT
	e.POST("/logout", logout)

	e.Logger.Fatal(e.Start("localhost:5000"))
	//berjalan di localhost 5000
}

func home(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT * FROM tb_projects")
	
	var projectData []Project
	
	
	for data.Next() {
		var each = Project{}
		// fmt.Println(each)
		err := data.Scan(&each.Id, &each.Name, &each.StartDate, &each.EndDate, &each.Duration, &each.Description,  &each.Image , &each.Html, &each.Css, &each.Reactjs, &each.Javascript)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}
		each.Duration = countDuration(each.StartDate, each.EndDate)
		each.StartDateTime = each.StartDate.Format("2006-01-02")
		each.EndDateTime = each.EndDate.Format("2006-01-02")

		projectData = append(projectData, each)
	}
		
	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

		
		projects := map[string]interface{}{
		"Projects": projectData,
		"DataSession": userData,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), projects)
}
func project(c echo.Context) error {
	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	projects := map[string]interface{}{
		"DataSession": userData,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())
	var tmpl, err = template.ParseFiles("views/my-project.html")
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	
	return tmpl.Execute(c.Response(), projects)
}

func contact(c echo.Context) error {

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	projects := map[string]interface{}{
		"DataSession": userData,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/form-contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), projects)
}
func testimonials(c echo.Context) error {

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	projects := map[string]interface{}{
		"DataSession": userData,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/testimonials.html")
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	
	return tmpl.Execute(c.Response(), projects)
}
func detailproject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{} // pemanggialn struct interface

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, name, start_date, end_date, duration, description, image, html, css, react, javascript FROM tb_projects WHERE id=$1", id).Scan(
		&ProjectDetail.Id, &ProjectDetail.Name, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Duration, &ProjectDetail.Description , &ProjectDetail.Image, &ProjectDetail.Html, &ProjectDetail.Css, &ProjectDetail.Reactjs, &ProjectDetail.Javascript )

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{}{
		"Project":   ProjectDetail,
		"StartDate": getDateString(ProjectDetail.StartDate.Format("2006-01-02")),
		"EndDate":   getDateString(ProjectDetail.EndDate.Format("2006-01-02")),
		"DataSession": userData,
	}

	var tmpl, errTemplate = template.ParseFiles("views/blog-detail.html")
	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errTemplate.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func formAddProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/my-project")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func addNewProject(c echo.Context) error {
	name := c.FormValue("name-title")
	startDate := c.FormValue("StartDate")
	endDate := c.FormValue("EndDate")
	description := c.FormValue("Description")
	// html := c.FormValue("html")
	// css := c.FormValue("css")
	// javascript := c.FormValue("javascript")
	// reactjs := c.FormValue("rjs")
	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)

	duration := (countDuration(start, end))


	// konversi value cekbox, type data string menjadi boolean
	var html bool
	if c.FormValue("html") == "html"{
		html = true
	}
	
	var css bool
	if c.FormValue("css") == "css"{
		css = true
	}
	
	var reactjs bool
	if c.FormValue("react") == "react"{
		reactjs = true
	}
	
	var javascript bool
	if c.FormValue("javascript") == "javascript"{
		javascript = true
	}

	// var arrTeks = "{"+ strings.Join(techs[], " , ")+ "}"

	image := "/public/image/image-project.png"
	// fmt.Println(arrTeks)
	_, err := connection.Conn.Exec(context.Background(), 
	 
	"INSERT INTO tb_projects (name, description, start_date, end_date, duration, image, html, css, react, javascript) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", name, description, startDate ,endDate, duration, image, html, css, reactjs, javascript)
	

	if err != nil {
		fmt.Print(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func editProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var ProjectDetail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_projects WHERE id=$1", id).Scan(
		&ProjectDetail.Id, &ProjectDetail.Name, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Description, &ProjectDetail.Description, &ProjectDetail.Image, &ProjectDetail.Html, &ProjectDetail.Css, &ProjectDetail.Reactjs, &ProjectDetail.Javascript)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	

	// memanggil fungsi isAvailable 

	start := ProjectDetail.StartDate.Format("2006-01-02") //format default date
	end := ProjectDetail.EndDate.Format("2006-01-02")
	data := map[string]interface{}{
		"Project":   ProjectDetail,
		"StartDate": start,
		"EndDate":   end,
	}	

	var tmpl, errTemplate = template.ParseFiles("views/edit-project.html")


	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	 return tmpl.Execute(c.Response(), data)
}


func ressEditProject(c echo.Context) error {
	id := getProjectIndex(c.Response(), c.Request())

	name := c.FormValue("input-name")
	StartDate := c.FormValue("StartDate")
	EndDate := c.FormValue("EndDate")
	Description := c.FormValue("Description")
	// html := c.FormValue("html")
	// css := c.FormValue("css")
	// reactjs := c.FormValue("react")
	// javascript := c.FormValue("javascript")

	start, _ := time.Parse("2006-01-02", StartDate)
	end, _ := time.Parse("2006-01-02", EndDate)

	var html bool
	if c.FormValue("html") == "html"{
		html = true
	}
	
	var css bool
	if c.FormValue("css") == "css"{
		css = true
	}
	
	var react bool
	if c.FormValue("react") == "react"{
		react = true
	}
	
	var javascript bool
	if c.FormValue("javascript") == "javascript"{
		javascript = true
	}

	image := "/public/image/image-project.png"

	_, err := connection.Conn.Exec(
		context.Background(),
		"UPDATE tb_projects SET name=$1, start_date=$2, end_date=$3, description=$4, image=$5, html=$6, css=$7, react=$8, javascript=$9 WHERE id=$10",
		name, start, end, Description, image, html, css, react, javascript,  id,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/project/"+id)
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_projects WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func signupForm(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/form-register.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func singup(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_user(name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)
	// fmt.Println(name,email,password,passwordHash)
	if err != nil {
		redirectWithMessage(c, "Signup failed, please try again!", false, "/register")
	}

	return redirectWithMessage(c, "Signup success!", true, "/login")
}

func loginForm(c echo.Context) error {
	sess, _ := session.Get("session", c)

	flash := map[string]interface{}{
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/form-login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), flash)
}

func login(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	email := c.FormValue("email")
	password := c.FormValue("password")

	user := User{}
	fmt.Println(user)
	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_user WHERE email=$1", email).Scan(&user.Name, &user.Email, &user.Password, &user.Id)
	if err != nil {
		return redirectWithMessage(c, "Email Incorrect!", false, "/login")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return redirectWithMessage(c, "Password Incorrect", false, "/login")
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800 // 3 JAM
	sess.Values["message"] = "Login success!"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["id"] = user.Id
	sess.Values["isLogin"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func redirectWithMessage(c echo.Context, message string, status bool, path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, path)
}

func getDateString(date string) string {
	y := date[0:4]
	m, _ := strconv.Atoi(date[5:7])
	d := date[8:10]
	if string(d[0]) == "0" {
		d = string(d[1])
	}

	mon := ""
	switch m {
	case 1:
		mon = "Jan"
	case 2:
		mon = "Feb"
	case 3:
		mon = "Mar"
	case 4:
		mon = "Apr"
	case 5:
		mon = "Mei"
	case 6:
		mon = "Jun"
	case 7:
		mon = "Jul"
	case 8:
		mon = "Agu"
	case 9:
		mon = "Sep"
	case 10:
		mon = "Okt"
	case 11:
		mon = "Nov"
	case 12:
		mon = "Des"
	}

	return d + " " + mon + " " + y
}

func countDuration(d1 time.Time, d2 time.Time) string {
	diff := d2.Sub(d1)
	days := int(diff.Hours() / 24)
	weeks := days / 7
	months := days / 30

	if months > 12 {
		return strconv.Itoa(months/12) + " tahun"
	}
	if months > 0 {
		return strconv.Itoa(months) + " bulan"
	}
	if weeks > 0 {
		return strconv.Itoa(weeks) + " minggu"
	}
	return strconv.Itoa(days) + " hari"
}

func getProjectIndex(w http.ResponseWriter, r *http.Request) string {
	// to call: getProjectIndex(c.Response(), c.Request())
	// value of url: /edit-project/0?
	// returned value: 0
	url := r.URL.String()
	lastSegment := path.Base(url)
	re := regexp.MustCompile("[0-9]+")
	return (re.FindAllString(lastSegment, -1))[0]
}