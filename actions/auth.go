package actions

import (
	"fmt"
	"log"
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/gplus"
)

func init() {
	gothic.Store = App().SessionStore
	log.Println("The app Host: ", app.Host)
	goth.UseProviders(
		gplus.New(os.Getenv("GPLUS_KEY"), os.Getenv("GPLUS_SECRET"), fmt.Sprintf("%s%s", app.Host, "/auth/gplus/callback")),
	)
}

func AuthCallback(c buffalo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(401, err)
	}

	s := c.Session()
	s.Set("uid", user.UserID)
	s.Set("avatar", user.AvatarURL)
	s.Set("username", user.Name)
	s.Save()
	return c.Redirect(301, "/")
	// Do something with the user, maybe register them/sign them in
	//return c.Render(200, r.JSON(user))
}
