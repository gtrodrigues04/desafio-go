# desafio-go

Informações do desafio
Neste desafio, você deve criar uma aplicação Golang com Docker que rode na porta 8080.

Esta aplicação precisa expor 2 rotas de API Rest:


Listar routes - POST /api/routes

Criar routes - GET /api/routes


Uma rota tem os seguintes dados:


id (gerado automaticamente pelo MySQL)

name (nome da rota)

source (campo JSON que contém lat e lng)

destination (campo JSON que contém lat e lng)


Use o pacote database/sql Golang para trabalhar com o MySQL e o pacote CHI para montar as rotas WEB


O banco de dados precisa ser o MySQL, image: mysql:8.0.30-debian


Crie o arquivo api.http para fazer as chamadas HTTP. Ao rodar o docker compose up já precisa subir logo de cara o projeto com o Golang rodando + o MySQL.
