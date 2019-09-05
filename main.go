package main

import (
	"github.com/eandreani/go-platform/web"
	"github.com/eandreani/go-template-repository/api"
)

func main() {
	// change go-tamplate-repository en .gitignore
	server := web.NewServer(web.ReadConfigFromEnv())
	server.AddMetrics()
	server.AddCorsAllOrigins()
	server.AddHealth(web.HealthAlwaysUp)
	server.AddApiDocs()
	// exec$ swag init
	api.SetupRouter(server.GetRouter())
	server.ListenAndServe()
}

// export GO111MODULE=on
