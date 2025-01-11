# Golang
FROM golang:1.20-alpine

# Diretório de trabalho no container
WORKDIR /app

# Copiar o gerenciador de dependências do Go para dentro do container
COPY go.mod go.sum ./

# Instalar dependências automaticamente
RUN go mod download

# Copiar o restante do código para dentro do container
COPY . .

# Compilar a aplicação
RUN go build -o main .

# Expor a porta 8080 no container
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./main"]
