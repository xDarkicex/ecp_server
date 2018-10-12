package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/pure"
	"github.com/xDarkicex/easycare_server/app/controllers"
	"github.com/xDarkicex/easycare_server/helpers"
)

func LoadServer() *pure.Mux {
	p := pure.New()
	loadRoutes(p)
	return p
}

func loadRoutes(p *pure.Mux) {
	p.Get("/", func(w http.ResponseWriter, r *http.Request) {
		c, _ := r.Cookie("User-Language")
		if c != nil {
			lang := controllers.LoadLanguage(c.Value)
			helpers.Render(w, r, "index", map[string]interface{}{
				"Translations": lang,
			})
		} else {
			helpers.Render(w, r, "index", map[string]interface{}{
				"Translations": controllers.Translated("english"),
			})
		}
	})

	p.Get("/store-locator", func(w http.ResponseWriter, r *http.Request) {
		T := controllers.Translated("english")
		helpers.Render(w, r, "store-locator", map[string]interface{}{
			"Translations": T,
		})
	})

	p.Get("/about-us", func(w http.ResponseWriter, r *http.Request) {
		helpers.Render(w, r, "about-us", map[string]interface{}{
			"Translations": controllers.Translated("english"),
		})
	})
	p.Get("/resources", func(w http.ResponseWriter, r *http.Request) {
		helpers.Render(w, r, "resources", map[string]interface{}{
			"Translations": controllers.Translated("english"),
		})
	})
	p.Get("/staff", func(w http.ResponseWriter, r *http.Request) {
		helpers.Render(w, r, "staff", map[string]interface{}{
			"Translations": controllers.Translated("english"),
		})
	})
	p.Get("/contact", func(w http.ResponseWriter, r *http.Request) {
		helpers.Render(w, r, "contact", map[string]interface{}{
			"Translations": controllers.Translated("english"),
		})
	})
	p.Get("/set-cookie", func(w http.ResponseWriter, r *http.Request) {
		expiration := time.Now().Add(365 * 24 * time.Hour)
		fmt.Println(r.FormValue("language"))
		cookie := http.Cookie{Name: "User-Language", Value: r.FormValue("language"), Expires: expiration}
		http.SetCookie(w, &cookie)
		r.AddCookie(&cookie)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	})
	p.Get("/testing/", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("User-Language")
		if err != nil {
			fmt.Println(err)
		}
		w.Write([]byte(cookie.Value))
	})
	p.Get("/products/pooltec", func(w http.ResponseWriter, r *http.Request) {
		helpers.Render(w, r, "/products/pooltec", map[string]interface{}{
			"Translations": controllers.Translated("english"),
		})
	})
	p.Get("/products/algatec", func(w http.ResponseWriter, r *http.Request) {
		helpers.Render(w, r, "/products/algatec", map[string]interface{}{
			"Translations": controllers.Translated("english"),
		})
	})
	p.Get("/products/pooltec-winter", func(w http.ResponseWriter, r *http.Request) {
		helpers.Render(w, r, "/products/pooltec-winter", map[string]interface{}{
			"Translations": controllers.Translated("english"),
		})
	})
	p.Get("/products/beautec", func(w http.ResponseWriter, r *http.Request) {
		helpers.Render(w, r, "/products/beautec", map[string]interface{}{
			"Translations": controllers.Translated("english"),
		})
	})
	p.Get("/products/scaletec", func(w http.ResponseWriter, r *http.Request) {
		helpers.Render(w, r, "/products/scaletec", map[string]interface{}{
			"Translations": controllers.Translated("english"),
		})
	})
	p.Get("/products/protec", func(w http.ResponseWriter, r *http.Request) {
		helpers.Render(w, r, "/products/protec", map[string]interface{}{
			"Translations": controllers.Translated("english"),
		})
	})
	p.Get("/products/fountec", func(w http.ResponseWriter, r *http.Request) {
		helpers.Render(w, r, "/products/fountec", map[string]interface{}{
			"Translations": controllers.Translated("english"),
		})
	})
	p.Get("/products/startup-tec", func(w http.ResponseWriter, r *http.Request) {
		helpers.Render(w, r, "/products/startup-tec", map[string]interface{}{
			"Translations": controllers.Translated("english"),
		})
	})
	p.Get("/french", func(w http.ResponseWriter, r *http.Request) {
		helpers.Render(w, r, "french", map[string]interface{}{
			"Translations": controllers.Translated("english"),
		})
	})
	p.Get("/spanish", func(w http.ResponseWriter, r *http.Request) {
		helpers.Render(w, r, "spanish", map[string]interface{}{
			"Translations": controllers.Translated("english"),
		})
	})
	p.Get("/italian", func(w http.ResponseWriter, r *http.Request) {
		helpers.Render(w, r, "italian", map[string]interface{}{
			"Translations": controllers.Translated("english"),
		})
	})
	p.Get("/portugese", func(w http.ResponseWriter, r *http.Request) {
		helpers.Render(w, r, "portugese", map[string]interface{}{
			"Translations": controllers.Translated("english"),
		})
	})
	p.Get("/russian", func(w http.ResponseWriter, r *http.Request) {
		helpers.Render(w, r, "spanish", map[string]interface{}{
			"Translations": controllers.Translated("english"),
		})
	})
	p.Get("/greek", func(w http.ResponseWriter, r *http.Request) {
		helpers.Render(w, r, "greek", map[string]interface{}{
			"Translations": controllers.Translated("english"),
		})
	})
	p.Get("/german", func(w http.ResponseWriter, r *http.Request) {
		helpers.Render(w, r, "german", map[string]interface{}{
			"Translations": controllers.Translated("english"),
		})
	})
	p.Post("/contact-us", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", 301)
		return
	})
	p.Get("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))).ServeHTTP)
}
