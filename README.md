# fullcycle-auction

## Portas Utilizadas
- Servidor web: 8080

## Como rodar o projeto

1. Rodar o comando para iniciar o banco de dados , rodar migration e iniciar o servidor
``` shell
docker-compose up
```

## Envs
AUCTION_INTERVAL (Define o tempo de fechamento do leilão)
MONGODB_URL      (Url do MongoDB)


## Como testar o projeto

1. Para testar o fechamento do leilão basta rodar o teste automatizado no arquivo [create_auction_test.go](internal/infra/database/auction/create_auction_test.go)

