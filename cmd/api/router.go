package main

import (
	"net/http"

	"github.com/LeoCosta17/SocialMedia/internal/responses"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func mount(app *application) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {

		data := map[string]any{
			"api_port":          app.config.addr,
			"api_db_port":       app.config.dbConfig.addr,
			"api_db_max_conns":  app.config.dbConfig.maxOpenConns,
			"api_db_idle_conns": app.config.dbConfig.maxIdleConns,
			"api_status":        "OK",
		}

		responses.WriteJSON(w, http.StatusOK, data)
	})

	r.Get("/", app.handler.Feed.GetFeed)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", app.handler.Users.Create)
		r.Route("/{user_id}", func(r chi.Router) {
			r.Get("/", app.handler.Users.GetByID)
			r.Put("/follow", app.handler.Users.Follow)
			r.Put("/unfollow", app.handler.Users.Unfollow)
		})
	})

	r.Route("/posts", func(r chi.Router) {
		r.Post("/", app.handler.Posts.Create)

		r.Route("/{post_id}", func(r chi.Router) {
			r.Get("/", app.handler.Posts.GetByID)
			r.Patch("/", app.handler.Posts.Update)
			r.Delete("/", app.handler.Posts.Delete)

			r.Route("/comments", func(r chi.Router) {
				r.Post("/", app.handler.Comments.Create)
				r.Get("/", app.handler.Comments.GetAll)
			})
		})
	})

	return r
}
