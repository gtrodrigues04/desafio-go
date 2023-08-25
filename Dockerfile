# Use a imagem base do Golang
FROM golang:latest

# Defina o diretório de trabalho
WORKDIR /app

# Copie o conteúdo do diretório atual para o diretório de trabalho no contêiner
COPY . .

# Baixe as dependências e construa o executável
RUN go mod download
RUN go build -o main .

# Exponha a porta 8080
EXPOSE 8081

# Comando para iniciar a aplicação
CMD ["./main"]