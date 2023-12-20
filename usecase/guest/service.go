package guest

import "github.com/gin-gonic/gin"

type Service interface {
	LoginService(username string, password string) (token string, err error)
	PublishCreateUserRabbitMq(ctx *gin.Context, email string, password string, phoneNumber string) (err error)
}
