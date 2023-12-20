
# Personal Showcase

A personal project for exploring my skill on Golang




## Libraries

I use this library for my personal project showcase
 - [Gin](https://github.com/gin-gonic/gin)
 - [Viper](https://github.com/spf13/viper)
 - [Gorm](https://gorm.io)
 - [AMQP RabbitMQ](https://github.com/rabbitmq/amqp091-go)


## Deployment

To deploy this project run

```bash
  docker compose up --build
```


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`POSTGRES_USER`

`POSTGRES_PASSWORD`

`POSTGRES_DB`

`RABBITMQ_USER`

`RABBITMQ_PASS`

## Features

- DDD architecture
- Clean architecture
- Docker implementation
- Gin framework
- Message broker implementation on register (RabbitMQ)
- Consumer & Publisher on 1 apps
- Support MySQL & postgresql
- Gorm on database ORM
- Auto table migration on first setup
