# Expense Management API 💰

Esta é uma API REST desenvolvida em **Go (Golang)** para o gerenciamento de despesas. A aplicação utiliza **MongoDB** para persistência de dados e segue boas práticas de organização, validação de dados e separação de responsabilidades.

---

## 🚀 Tecnologias Utilizadas

* **Linguagem:** Go (Golang)
* **Roteamento:** Gorilla Mux
* **Banco de Dados:** MongoDB (Driver oficial)
* **Configuração de Ambiente:** godotenv, envconfig
* **Manipulação de JSON:** encoding/json

---

## 🛠️ Instalação e Execução

1. **Clone o repositório**
```bash
git clone https://github.com/seu-usuario/expensemanagement
cd expensemanagement
Certifique-se de que o MongoDB está rodando

A API espera uma instância local em:

mongodb://localhost:27017
Configure as variáveis de ambiente

Crie um arquivo .env na raiz do projeto:

MONGO_URI=mongodb://localhost:27017
EXPENSES_COL=expenses
Instale as dependências

go mod tidy
Execute a aplicação

go run main.go
A API será iniciada na porta 9437.

🛣️ Endpoints da API
Método	Rota	Descrição
POST	/expenses	Cria uma nova despesa
GET	/expenses	Lista todas as despesas
PUT	/expenses/{id}	Atualiza uma despesa existente
DELETE	/expenses/{id}	Remove uma despesa por ID
DELETE	/expenses	Remove todas as despesas
📦 Estrutura do JSON (Exemplo)
{
  "name": "Internet",
  "description": "Conta mensal de internet",
  "value": 120.50,
  "duedate": "2026-02-10T00:00:00Z"
}
✅ Regras de Validação
Name
Não pode ser vazio

Não pode iniciar ou terminar com espaços

Deve conter entre 3 e 20 caracteres

Description
Não pode iniciar ou terminar com espaços

Máximo de 150 caracteres

Value
Deve ser maior que 0

DueDate
Deve ser maior que ontem

Quando alguma validação falha, a API retorna 400 Bad Request com uma lista de mensagens de erro.

🗄️ Banco de Dados
Database: expensesdb

Collection: definida pela variável EXPENSES_COL

Conexão validada no início da aplicação com Ping

📌 Observações Gerais
Arquitetura em camadas (handler, validation, services e repository)

Validações centralizadas antes da persistência

Mensagens de erro padronizadas

Projeto indicado para estudos, prática com Go e base para evolução

