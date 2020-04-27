
go get github.com/urfave/negroni

Instalar o curl for windows
https://curl.haxx.se/windows/

https://github.com/gorilla/mux

https://www.gorillatoolkit.org/pkg/mux
go get -u github.com/gorilla/mux

Mongo
go get gopkg.in/mgo.v2

go run main.go

Iniciar o mongodb

mongod –-dbpath=C:\data\db 

ex: mongod --dbpath=C:\valdir\private\coders_go\restgo\data\db 

Comandos uteis

db.help()
Lista todos os comandos que podem ser aplicados no banco de dados e contém também  uma explicação (guarde este comando no coração);

show dbs
Mostra todos os bancos de dados existentes no servidor.

db
Mostra qual banco de dados está sendo usado. Se você acabou de instalar o MongoDB, e executar este comando verá que o banco de dados usado é o test, banco de dados default e para testes do MongoDB

use nomeBD
Este comando acessa um determinado banco de dados, quando ele existe. Se o banco não existe ele é criado automaticamente.

mongo
é o cliente

show collections

db.customers.find()

db.customers.insert({ nome: "Luiz", idade: 29 })

TESTES COM O CURL:
Consulta:
curl -v -X GET http://127.0.0.1:8701/tasks

Criar:
curl -v -X POST -d {\"name\":\"Rico\"} http://127.0.0.1:8701/tasks

Update
curl -v -X PUT -d {\"name\":\"RicoTorres\"} http://127.0.0.1:8701/tasks/5d6e9c1487cb4f2870c1b7d9
curl -v -X PUT -d {\"name\":\"Rico%20Torres\"} http://127.0.0.1:8701/tasks/5d6e9c1487cb4f2870c1b7d9

Delete
curl -v -X DELETE http://127.0.0.1:8701/tasks/5d6e9c1487cb4f2870c1b7d9


