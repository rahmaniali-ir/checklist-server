package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	boardHandler "github.com/rahmaniali-ir/checklist-server/internal/handlers/board"
	"github.com/rahmaniali-ir/checklist-server/internal/models/board"
	"github.com/rahmaniali-ir/checklist-server/internal/router"
	"github.com/rahmaniali-ir/checklist-server/internal/routes"
	boardService "github.com/rahmaniali-ir/checklist-server/internal/services/board"
	// "github.com/rahmaniali-ir/checklist-server/pkg/session"
	// "github.com/syndtr/goleveldb/leveldb"
)

var EnvMap map[string]string
var defaultEnv = map[string]string{
	"SERVER_PORT": "8081",
}

type app struct {
	router *mux.Router
}

func New() (*http.Server, error) {
	var err error

	EnvMap, err = godotenv.Read(".env")
	if err != nil {
		EnvMap = defaultEnv
	}

	allRoutes := []router.Route{}

	// dbPath := EnvMap["DB_PATH"]
	// db, err := leveldb.OpenFile(dbPath, nil)
	// if err != nil {
	// 	panic(err)
	// }

	// board routes
	bm := board.New()
	bs := boardService.New(bm)
	allRoutes = append(allRoutes, routes.BoardRoutes(boardHandler.New(bs))...)

	// session manager
	// session.Init(db)

	newApp := app{}
	err = newApp.createResources(allRoutes...)

	if err != nil {
		return nil, err
	}

	return newApp.server(), nil
}

func (a *app) createResources(rs ...router.Route) error {
	a.router = mux.NewRouter().StrictSlash(true)

	for _, r := range rs {
		err := a.router.Name(r.Name).Path(r.Path).Methods(r.Method, http.MethodOptions).HandlerFunc(r.Handler).GetError()

		if err != nil {
			return err
		}
	}

	return nil
}

func (a *app) server() *http.Server {
	port := EnvMap["SERVER_PORT"]

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: a.router,
	}
	
	fmt.Printf("Listening on port: %v\n", port)
	return server
}
