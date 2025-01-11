# Movie API - Golang + REST

Essa é uma API REST para gerenciamento de filmes desenvolvida com **Golang**, criada com o objetivo de exercitar conceitos de desenvolvimento de APIs RESTful, estruturação de projetos e uso de bibliotecas como **Gin** e **MongoDB Driver**.

---

## **Recursos da API**

- **Criar um filme**: `POST /movies`
- **Listar todos os filmes**: `GET /movies`
- **Obter um filme por ID**: `GET /movies/:id`
- **Atualizar um filme por ID**: `PUT /movies/:id`
- **Excluir um filme por ID**: `DELETE /movies/:id`

---

## **Tecnologias e Bibliotecas Utilizadas**

- **Golang (1.20+)**
- **[Gin](https://github.com/gin-gonic/gin)**: Framework para criação de rotas e middlewares.
- **[MongoDB Driver](https://github.com/mongodb/mongo-go-driver)**: Conexão e interação com o banco de dados MongoDB.
- **[Docker](https://www.docker.com/)**: Contêinerização da aplicação e banco de dados.
- **[Postman](https://www.postman.com/)**: Teste dos endpoints da API.

---

## **Como Rodar a Aplicação**

### **Pré-requisitos**
- Docker e Docker Compose instalados.

### **Como Rodar a Aplicação (continuação)**

1. **Clone o repositório**:
   ```bash
   git clone https://github.com/seu-usuario/seu-repositorio.git
   cd seu-repositorio
   ```

2. **Suba os containers com Docker Compose**:
   ```bash
   docker-compose up --build
   ```

3. **Acesse a API**:
   A API estará disponível em `http://localhost:8080`.

---

## **Exemplos de Requisições**

### **Criar um Filme**
- **POST `/movies`**
- **Body (JSON)**:
  ```json
  {
    "title": "The Matrix",
    "director": "Wachowski Sisters",
    "year": 1999,
    "genre": "Science Fiction",
    "rating": "9"
  }
  ```

---

### **Listar Todos os Filmes**
- **GET `/movies`**
- **Resposta**:
  ```json
  [
    {
      "id": "6782a3645f65691e772bbd90",
      "title": "The Matrix",
      "director": "Wachowski Sisters",
      "year": 1999,
      "genre": "Science Fiction",
      "rating": "9"
    },
    ...
  ]
  ```

---

### **Atualizar um Filme**
- **PUT `/movies/:id`**
- **Body (JSON)**:
  ```json
  {
    "title": "The Matrix Reloaded",
    "director": "Wachowski Sisters",
    "year": 2003,
    "genre": "Science Fiction",
    "rating": "8"
  }
  ```

---

### **Excluir um Filme**
- **DELETE `/movies/:id`**

---

## **Estrutura do Projeto**

```plaintext
.
├── controllers
│   └── movie_controller.go   # Controladores para cada rota
├── database
│   └── connection.go         # Configuração de conexão com o MongoDB
├── models
│   └── movie.go              # Modelo de filme
├── routes
│   └── routes.go             # Definição de rotas
├── validator
│   └── validator.go          # Validações de entrada
├── Dockerfile                # Configuração do container Docker
├── docker-compose.yml        # Configuração do Docker Compose
└── main.go                   # Entrada principal da aplicação
```
