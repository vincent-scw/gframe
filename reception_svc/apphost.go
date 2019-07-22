package main

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	corsmid "github.com/iris-contrib/middleware/cors"
	jwtmid "github.com/iris-contrib/middleware/jwt"
	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/kataras/iris"

	"github.com/vincent-scw/gframe/events"
	k "github.com/vincent-scw/gframe/kafkactl"
	_ "github.com/vincent-scw/gframe/reception_svc/docs" // docs is generated by Swag CLI
	e "github.com/vincent-scw/gframe/reception_svc/event"
)

func main() {
	app := iris.Default()

	jwtHandler := jwtmid.New(jwtmid.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("00000000"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

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

	p := k.NewProducer()
	defer p.Dispose()
	fmt.Println("Kafka producer initilized...")

	player := app.Party("/user")
	player.Use(jwtHandler.Serve)
	{
		player.Post("/in", func(ctx iris.Context) {
			authToken := jwtHandler.Get(ctx)
			status := handleUserReception(p, authToken, events.UserEventIn)
			ctx.StatusCode(status)
		})
		player.Post("/out", func(ctx iris.Context) {
			authToken := jwtHandler.Get(ctx)
			status := handleUserReception(p, authToken, events.UserEventOut)
			ctx.StatusCode(status)
		})
	}

	config := &swagger.Config{
		URL: "http://localhost:8080/swagger/doc.json",
	}
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(config, swaggerFiles.Handler))

	fmt.Println("Start player api...")
	app.Run(iris.Addr(":8080"))
}

func handleUserReception(p *k.Producer, authToken *jwt.Token, t events.UserStatus) int {
	user, err := e.NewEvent(authToken, t)
	if err != nil {
		return iris.StatusForbidden
	}

	if err = p.Emit(user); err != nil {
		log.Fatalln(err)
		return iris.StatusInternalServerError
	}
	return iris.StatusNoContent
}
