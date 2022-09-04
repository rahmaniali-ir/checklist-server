package routes

import (
	"net/http"

	boardHandler "github.com/rahmaniali-ir/checklist-server/internal/handlers/board"
	internalHttp "github.com/rahmaniali-ir/checklist-server/internal/http"
	"github.com/rahmaniali-ir/checklist-server/internal/router"
)

func BoardRoutes(boardHandler boardHandler.IHandler) []router.Route {
	return []router.Route{
		{
			Name: "getBoards",
			Path: "/boards",
			Method: http.MethodGet,
			Handler: internalHttp.Handle(boardHandler.List),
		},
	}
}
