# Go REST API

API RESTful simples em Go para gerenciar tarefas com SQLite.

## Configuração

O arquivo `config.toml` contém as configurações da API e do banco de dados:

```toml
[api]
port = "9000"

[database]
filepath = "data.db"
```

## Endpoints

- POST /: Criar tarefa
- GET /: Listar tarefas
- GET /{id}: Obter tarefa por ID
- PUT /{id}: Atualizar tarefa
- DELETE /{id}: Deletar tarefa

## Como rodar

1. Clone o repositório:

    ```bash
    git clone <URL_DO_REPOSITORIO>
    cd go-rest-api
    ```

2. Instale as dependências:

    ```bash
    go mod tidy
    ```

3. Execute a aplicação:

    ```bash
    go run main.go
    ```

A API estará disponível em `http://localhost:9000`.
