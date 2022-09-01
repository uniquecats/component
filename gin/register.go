package gin

import (
	"github.com/gin-gonic/gin"
)

type Register interface {
	RegisterHandler(r gin.IRouter)
}
