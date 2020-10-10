
go get github.com/urfave/negroni

Instalar o curl for windows
https://curl.haxx.se/windows/

https://github.com/gorilla/mux

https://www.gorillatoolkit.org/pkg/mux
go get -u github.com/gorilla/mux

Mongo
go get gopkg.in/mgo.v2

go run main.go --> Iniciar o mongo antes

Iniciar o mongodb

mongod –-dbpath=C:\data\db 

ex: mongod --dbpath=C:\restgo\data\db 
Caso tenha dificuldade, apague o c:\data\db (ou onde estiver os seus dados do mongodb) e inicie com o diretório vazio

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
----------------------------------------------------------------------------------------
é o cliente

show collections

db.customers.find()

db.customers.insert({ nome: "Luiz", idade: 29 })

Consulta básica:
SELECT * FROM clientes LIMIT 10;
db.clientes.find({}).limit(10);
SELECT id, nome FROM clientes ORDER BY id;
db.clientes.find({},{“_id”: 1, “nome”: 1}).sort({“_id”: 1});
SELECT id, nome FROM clientes ORDER BY id DESC;
db.clientes.find({},{“_id”: 1, “nome”: 1}).sort({“_id”: -1});
----------------------------------------------------------------------------------------

TESTES COM O CURL: (NOTA: USE o cmd em vez do powershell)
veja: https://stackoverflow.com/questions/45183355/invoke-webrequest-cannot-bind-parameter-headers
Consulta:
curl http://127.0.0.1:8721/tasks

Consultar pelo ID (curl -v GET http://127.0.0.1:8701/tasks/{id})
curl -v GET http://127.0.0.1:8721/tasks/5ea883358699e13ed085d5b2

Criar:
curl -d {"""name""":"""TESTE AAA"""} http://127.0.0.1:8721/tasks

Update
curl -v -X PUT -d {"""name""":"""Rico Torres"""} http://127.0.0.1:8721/tasks/5f80fc8c8699e148347d7e12

Delete
curl -v -X DELETE http://127.0.0.1:8721/tasks/5ea8894d8699e12c7c51960f


