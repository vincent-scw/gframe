package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	corsmid "github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"

	e "github.com/vincent-scw/gframe/contracts"
	"github.com/vincent-scw/gframe/game_svc/auth"
	"github.com/vincent-scw/gframe/game_svc/config"
	_ "github.com/vincent-scw/gframe/game_svc/docs" // docs is generated by Swag CLI
)

func main() {
	log.Println("Starting game service...")

	context, cancel := context.WithCancel(context.Background())

	wrapper := newBrokerRPCWrapper()
	defer wrapper.Close()
	client := wrapper.client

	app := iris.Default()

	cors := corsmid.New(corsmid.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "HEAD"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	app.Use(cors)
	app.AllowMethods(iris.MethodOptions)

	app.Get("/health", func(ctx iris.Context) {
		ctx.Text("I am good.")
	})

	srv := startWebsocket(func(conn *websocket.Conn, user *e.User) error {
		_, err := client.Checkin(context, user)
		return err
	},
		func(conn *websocket.Conn, user *e.User) {
			_, err := client.Checkout(context, user)
			if err != nil {
				log.Println(err)
			}
		})
	app.Get("/console", auth.WSJwtHandler.Serve, websocket.Handler(srv))

	api := app.Party("/api")
	player := api.Party("/user")
	player.Use(auth.JwtHandler.Serve)
	{
		player.Post("/in", func(ctx iris.Context) {
			user, err := auth.GetUserFromToken(ctx, e.User_In)
			if err != nil {
				ctx.StatusCode(iris.StatusForbidden)
			}
			status := handleUserReception(context, client, user)
			ctx.StatusCode(status)
		})
		player.Post("/out", func(ctx iris.Context) {
			user, err := auth.GetUserFromToken(ctx, e.User_Out)
			if err != nil {
				ctx.StatusCode(iris.StatusForbidden)
			}
			status := handleUserReception(context, client, user)
			ctx.StatusCode(status)
		})
	}

	swConfig := &swagger.Config{
		URL: "http://localhost:8080/swagger/doc.json",
	}
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(swConfig, swaggerFiles.Handler))

	app.Run(iris.Addr(config.GetPort()))
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-context.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	cancel()
}

func handleUserReception(ctx context.Context, c e.UserReceptionClient, user *e.User) int {
	var err error
	switch user.Status {
	case e.User_In:
		_, err = c.Checkin(ctx, user)
	case e.User_Out:
		_, err = c.Checkout(ctx, user)
	}

	if err != nil {
		log.Println(err)
		return iris.StatusInternalServerError
	}
	return iris.StatusNoContent
}