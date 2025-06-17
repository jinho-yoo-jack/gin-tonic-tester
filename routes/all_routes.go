package routes

import "github.com/gin-gonic/gin"

type Routable interface {
	RegisterRoutes(engine *gin.Engine, authorize gin.HandlerFunc)
}

type AllRouters struct {
	routers []Routable
}

func NewAllOfRouters(
	userRouter *UserRouters,
	minioRouter *MinioRouters,
) *AllRouters {
	return &AllRouters{
		routers: []Routable{
			userRouter, minioRouter,
		},
	}
}

func (ar *AllRouters) RegisterAllRoutes(engine *gin.Engine, authorize gin.HandlerFunc) {
	for _, router := range ar.routers {
		router.RegisterRoutes(engine, authorize)
	}
}
