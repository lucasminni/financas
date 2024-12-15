package v1

import (
	debtV1 "finances/pkg/routes/v1/debt"
	incomeV1 "finances/pkg/routes/v1/income"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	incomeV1.Register(r)
	debtV1.Register(r)
}