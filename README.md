# desafio-client-server-api

## Sobre a primeira correção

Consegui reproduzir o mesmo erro que o instrutor retornou:

```
2024/08/01 18:50:07 /Users/yuriespinosa/Developer/desafio-client-server-api/server/repository/awesomeapi-repository.go:52 context deadline exceeded; sql: transaction has already been committed or rolled back
[0.300ms] [rows:0] INSERT INTO `bids` (`bid`,`created_at`,`updated_at`,`deleted_at`) VALUES ("5.7518","2024-08-01 18:50:07.678","2024-08-01 18:50:07.678",NULL) RETURNING `id`,`id`
2024/08/01 18:50:07 Ocorreu um erro ao salvar o registro context deadline exceeded; sql: transaction has already been committed or rolled back 
```
#### Motivos
- Tecnicamente não houve erro lógico do programa, o erro é o timeout padrão do Context, ou seja, o tempo da inserção foi maior que 10 Milisegundos ná máquina de quem executou o server, ocasionando um retorno de erro via API e o log da transação que não foi commitada pelo gorm, devido ao timeout. 
- Testei novamente no Macos e no Linux, e obtive sucesso com o timeout de inserção de 10ms, talvez ocorreu algum ruido na maquina de quem executou, sugiro rodar o teste novamente ou verificar o desempenho do sqlite na máquina.
- Consegui reprouzir o erro somente setando um tempo de Timeout no Context absurdamente pequeno de 10 nanosegundos.


## Desafio 1 do curso Go Expert da Full Cycle.

Para executar o servidor execute o comando:

```Go
go run ./server/main.go
```

Para executar o teste com o Client http:

```Go
go run ./client/main.go
```
