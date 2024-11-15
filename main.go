package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"go-sessions/database"
	"go-sessions/templates"
)

func main() {
	e := echo.New()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.GET("/", func(c echo.Context) error {
		conn := database.Connect()
		session, err := session.Get("session", c)
		if err != nil {
			return Render(c, http.StatusOK, templates.Login())
		}

		userID, ok := session.Values["userID"].(int)
		if !ok {
			return Render(c, http.StatusOK, templates.Login())
		}

		user := database.GetUserByID(conn, userID)
		if user.Username == "" {
			return Render(c, http.StatusOK, templates.Login())
		}

		return c.Redirect(http.StatusSeeOther, "/dashboard")
	})

	e.POST("/login", func(c echo.Context) error {
		conn := database.Connect()
		username := c.FormValue("username")
		password := c.FormValue("password")

		user := database.AuthenticateUser(conn, username, password)

		if user.Username == "" {
			return c.Redirect(http.StatusOK, "/")
		}

		session, _ := session.Get("session", c)
		session.Options = &sessions.Options{
			Path: "/",
			MaxAge: 86400,
			HttpOnly: true,
		}
		session.Values["userID"] = user.ID
		_ = session.Save(c.Request(), c.Response().Writer)

		return c.Redirect(http.StatusSeeOther, "/dashboard")
	})

	e.POST("/logout", func(c echo.Context) error {
		session, _ := session.Get("session", c)
		session.Options = &sessions.Options{
			MaxAge: -1,
		}
		_ = session.Save(c.Request(), c.Response().Writer)

		return c.Redirect(http.StatusSeeOther, "/")
	})

	e.GET("/dashboard", func(c echo.Context) error {
		// get some kind of cookie, grab id, check if user exists
		session, _ := session.Get("session", c)
		userId, ok := session.Values["userID"].(int)
		if !ok {
			return c.Redirect(http.StatusSeeOther, "/")
		}
		conn := database.Connect()
		user := database.GetUserByID(conn, userId)
		return Render(c, http.StatusOK, templates.Dashboard(user))
	})

	e.Logger.Error(e.Start(":8080"))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
