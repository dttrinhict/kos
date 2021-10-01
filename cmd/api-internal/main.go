package main

import (
	"errors"
serverConfig "kos/cmd/api-internal/config"
	handler "kos/cmd/api-internal/handler"
	"kos/delivery/rest"
	k8s "kos/pkg/k8s"
	k8sUseCase "kos/usecases/k8s"
	"log"
	"net/http"
)

func main() {
	//defer shutdown.SigtermHandler().Wait()
	//registerErrMap()
	//registerZapFieldExtractors()
	server := initServer()
	server.Run()
}

type server struct {
	httpServer *http.Server
}

func (s *server) Run() {
	//shutdown.SigtermHandler().RegisterErrorFuncContext(context.Background(), s.httpServer.Shutdown)
	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Panic("Server listen and serve error", err)
	}
}

func initServer() *server {
	conf := loadConfig()

	domain := k8s.NewK8s()
	usecase := k8sUseCase.NewK8sUseCase(&domain)
	k8sOpsRest := rest.NewK8sOpsRest(usecase)
	k8sOps := handler.NewK8sOps(k8sOpsRest)
	handler := handler.NewServer(k8sOps)
	return &server{
		httpServer: &http.Server{
			Addr:    conf.HTTPServer.Addr,
			Handler: handler,
		},
	}
}

func loadConfig() *serverConfig.Config {
	conf, err := serverConfig.Load()
	if err != nil {
		log.Printf("load config error: %v", err.Error())
	}
	return conf
}
