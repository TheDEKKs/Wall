package transport

import (
	"net/http"
	"thedekk/WWT/internal/domains/comments"
	"thedekk/WWT/internal/domains/users"
	"thedekk/WWT/internal/domains/walls"
	"thedekk/WWT/internal/transport/handlers"
	"thedekk/WWT/internal/transport/middlewares"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewService(conn *pgxpool.Pool) error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	configAPI := huma.DefaultConfig("My API", "1.0.0")
	api := humachi.New(router, configAPI)
	api.UseMiddleware(middlewares.MyMiddleware)

	wallService := walls.NewWallService(conn)
	userService := users.NewUserService(conn, wallService)
	commentService := comments.NewCommentService(conn, userService)

	userHandler := handlers.NewUserHandler(userService)
	commentHandler := handlers.NewCommentHandler(commentService)
	wallHandler := handlers.NewWallHandler(wallService, commentService)

	huma.Register(api, huma.Operation{
		OperationID: "registration",
		Method:      http.MethodPost,
		Path:        "/registration",
		Summary:     "registration",
	}, userHandler.RegistrationUser)

	huma.Register(api, huma.Operation{
		OperationID: "login",
		Method:      http.MethodPost,
		Path:        "/login",
		Summary:     "login",
	}, userHandler.LoginUser)

	huma.Register(api, huma.Operation{
		OperationID: "new-comment",
		Method:      http.MethodPost,
		Path:        "/comment/new/{wall}",
		Summary:     "comment",
	}, commentHandler.NewComment)

	huma.Register(api, huma.Operation{
		OperationID: "wall-comments",
		Method:      http.MethodPost,
		Path:        "/{wall}/comments",
		Summary:     "wall-Comments",
	}, wallHandler.GetCommentsWall)

	// Start the server!
	http.ListenAndServe("127.0.0.1:8888", router)

	return nil
}
