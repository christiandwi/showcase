package guest

import (
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/christiandwi/showcase/domain/users"
	"github.com/christiandwi/showcase/lib/event"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userRepo users.UsersRepository
	rabbitMq event.RabbitMq
}

func NewGuestService(userRepo users.UsersRepository, rabbitMq event.RabbitMq) Service {
	//consumer rabbitmq
	go func() {
		queue := rabbitMq.RabbitMqQueueOpen("user")
		rabbitMq.RabbitMqConsume(queue, func(d amqp.Delivery) {
			bodyDeliv := make(map[string]string)
			err := json.Unmarshal(d.Body, &bodyDeliv)
			if err != nil {
				panic("err unmarshal on consuming")
			}

			err = userRepo.CreateUser(bodyDeliv["email"], bodyDeliv["password"], bodyDeliv["phoneNumber"])
			if err != nil {
				panic("err create user")
			}

			log.Printf("created user %v", bodyDeliv["email"])
		})
	}()

	return &service{
		userRepo: userRepo,
		rabbitMq: rabbitMq,
	}
}

func (serv *service) LoginService(userIdentifier string, password string) (token string, err error) {
	userData, err := serv.userRepo.GetUser(userIdentifier)
	if err != nil {
		log.Print("error at getting user")
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error at hashing password, err : %s", err.Error())
		return
	}

	if bcrypt.CompareHashAndPassword(hashed, []byte(userData.Password)) != nil {
		log.Print("invalid password")
		return
	}

	token = base64.StdEncoding.EncodeToString([]byte(userData.SecureId))
	return
}

func (serv *service) PublishCreateUserRabbitMq(ctx *gin.Context, email string, password string, phoneNumber string) (err error) {
	queue := serv.rabbitMq.RabbitMqQueueOpen("user")

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error at hashing password, err : %s", err.Error())
		return
	}

	pubBody := map[string]interface{}{
		"email":       email,
		"password":    hashed,
		"phoneNumber": phoneNumber,
	}

	encoded, err := json.Marshal(pubBody)
	if err != nil {
		log.Printf("error at marshaling body, err : %s", err.Error())
		return
	}

	err = serv.rabbitMq.RabbitMqPublish(ctx, queue, encoded)
	if err != nil {
		log.Printf("error at publishing, err : %s", err.Error())
		return
	}
	return
}
