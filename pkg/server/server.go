package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/cqbqdd11519/apiserver-test/internal/utils"
	"github.com/cqbqdd11519/apiserver-test/internal/wrapper"
	"github.com/cqbqdd11519/apiserver-test/pkg/apis"
)

var log = logf.Log.WithName("approve-server")

type Server struct {
	Wrapper *wrapper.RouterWrapper
}

func New() *Server {
	server := &Server{}
	server.Wrapper = wrapper.New("/", nil, server.rootHandler)
	server.Wrapper.Router = mux.NewRouter()
	server.Wrapper.Router.HandleFunc("/", server.rootHandler)

	if err := apis.AddApis(server.Wrapper); err != nil {
		os.Exit(1)
	}

	return server
}

func (s *Server) Start() {

	addr := "0.0.0.0:80"
	log.Info(fmt.Sprintf("Server is running on %s", addr))
	if err := http.ListenAndServe(addr, s.Wrapper.Router); err != nil {
		log.Error(err, "cannot launch server")
		os.Exit(1)
	}
}

func (s *Server) rootHandler(w http.ResponseWriter, _ *http.Request) {
	paths := metav1.RootPaths{}

	addPath(&paths.Paths, s.Wrapper)

	_ = utils.RespondJSON(w, paths)
}

func addPath(paths *[]string, w *wrapper.RouterWrapper){
	if w.Handler != nil {
		*paths = append(*paths, w.FullPath())
	}

	for _, c := range w.Children {
		addPath(paths, c)
	}
}
