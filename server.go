package main

//go:generate github.com/deepmap/oapi-codegen/cmd/oapi-codegen -package generated -o generated/api.gen.go api.yaml

import (
	oapiMiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	oapi "myapp/generated"
	"net/http"
)

// @impl oapi.ServerInterface
type server struct{}

func (s *server) FindPets(ctx echo.Context, params oapi.FindPetsParams) error {
	return ctx.JSON(http.StatusOK, []oapi.Pet{
		{
			Id:   1,
			Name: "dog",
		},
	})
}

func (s *server) AddPet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, oapi.Pet{
		Id:   1,
		Name: "dog",
	})
}

func (s *server) DeletePet(ctx echo.Context, id int64) error {
	return ctx.NoContent(http.StatusNoContent)
}

func (s *server) FindPetById(ctx echo.Context, id int64) error {
	return ctx.JSON(http.StatusOK, oapi.Pet{
		Id:   1,
		Name: "dog",
	})
}

func main() {
	e := echo.New()
	swagger, err := oapi.GetSwagger()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Use(oapiMiddleware.OapiRequestValidator(swagger))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := server{}
	oapi.RegisterHandlers(e, &api)

	e.Logger.Fatal(e.Start(":1323"))
}
