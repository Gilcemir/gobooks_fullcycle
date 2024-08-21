# GoBooks - Gerenciador de Livros

GoBooks é uma aplicação escrita em Go que permite gerenciar uma coleção de livros. Ela oferece uma interface web para operações CRUD (Create, Read, Update, Delete) e também suporta uma interface de linha de comando (CLI) para busca e simulação de leitura de múltiplos livros.

## Funcionalidades

### Web API
A aplicação oferece uma API REST com as seguintes rotas:

- **GET /books**: Retorna uma lista de todos os livros.
- **POST /books**: Cria um novo livro.
- **GET /books/{id}**: Retorna os detalhes de um livro específico pelo seu ID.
- **PUT /books/{id}**: Atualiza as informações de um livro específico.
- **DELETE /books/{id}**: Deleta um livro específico pelo seu ID.
- **GET /books/search**: Pesquisa livros por título.
- **POST /books/simulate**: Simula a leitura de múltiplos livros.

### CLI
A aplicação também pode ser executada via linha de comando com as seguintes funcionalidades:

- **search**: Busca livros pelo título.
    - **Exemplo**: `gobooks search "Nome do Livro"`
- **simulate**: Simula a leitura de múltiplos livros, passando os IDs dos livros como argumento.
    - **Exemplo**: `gobooks simulate 1 2 3 4`