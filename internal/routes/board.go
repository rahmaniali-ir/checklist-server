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
		{
			Name: "createBoard",
			Path: "/board",
			Method: http.MethodPost,
			Handler: internalHttp.Handle(boardHandler.Create),
		},
		{
			Name: "deleteBoard",
			Path: "/board/{uid}",
			Method: http.MethodDelete,
			Handler: internalHttp.Handle(boardHandler.Delete),
		},
		{
			Name: "updateBoard",
			Path: "/board",
			Method: http.MethodPut,
			Handler: internalHttp.Handle(boardHandler.Update),
		},
	}
}
