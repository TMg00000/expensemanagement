---
# Expense Management API

Uma API simples e leve para gerenciar despesas pessoais e categorias. Este README é objetivo: descreve instalação, configuração, endpoints principais, autenticação e como rodar testes.

Principais conceitos
- Usuário: autenticação e identificação para operações.
- Despesa: valor, data, categoria e descrição.
- Categoria: agrupa despesas (ex.: Alimentação, Transporte).

Funcionalidades
- CRUD de despesas
- CRUD de categorias
- Autenticação (JWT)
- Paginação e filtros por data, categoria e intervalo de valores

Requisitos
- Node.js >= 16 (ou versão usada no projeto)
- npm ou yarn
- Banco de dados (ex.: PostgreSQL, SQLite para desenvolvimento)

Instalação

1. Clone o repositório

   ```bash
   git clone https://github.com/TMg00000/expensemanagement.git
   cd expensemanagement
   ```

2. Instale dependências

   ```bash
   npm install
   # ou
   yarn install
   ```

Configuração

- Crie um arquivo .env na raiz seguindo o exemplo (.env.example, se existir) e configure variáveis essenciais:
  - DATABASE_URL (ex.: postgres://user:pass@localhost:5432/expenses)
  - JWT_SECRET
  - PORT (opcional)

Uso (desenvolvimento)

```bash
npm run dev
# ou
yarn dev
```

A API geralmente ficará disponível em http://localhost:3000 (ou porta definida em PORT).

Endpoints principais

- POST /auth/login
  - Autentica usuário e retorna token JWT.
  - Body: { "email": "", "password": "" }

- POST /auth/register
  - Cria novo usuário.
  - Body: { "name": "", "email": "", "password": "" }

- GET /expenses
  - Lista despesas (suporta page, limit, fromDate, toDate, categoryId, minAmount, maxAmount)
  - Headers: Authorization: Bearer <token>

- POST /expenses
  - Cria uma despesa
  - Body: { "amount": number, "date": "YYYY-MM-DD", "categoryId": string, "description": string }
  - Headers: Authorization: Bearer <token>

- GET /expenses/:id
  - Detalha uma despesa
  - Headers: Authorization: Bearer <token>

- PUT /expenses/:id
  - Atualiza despesa
  - Headers: Authorization: Bearer <token>

- DELETE /expenses/:id
  - Remove despesa
  - Headers: Authorization: Bearer <token>

- GET /categories
  - Lista categorias
  - Headers: Authorization: Bearer <token>

- POST /categories
  - Cria categoria
  - Body: { "name": "" }
  - Headers: Authorization: Bearer <token>

Autenticação

- A API usa JWT. Inclua o header:

  ```http
  Authorization: Bearer <token>
  ```

Estrutura do projeto (exemplo)

- src/
  - controllers/
  - routes/
  - models/
  - services/
  - middleware/
  - config/

Testes

- Execute os testes com:

  ```bash
  npm test
  # ou
  yarn test
  ```

Scripts úteis

- npm run dev — modo desenvolvimento
- npm start — iniciar em produção
- npm test — rodar testes

Boas práticas

- Use migrations para o banco de dados
- Versione a API (ex.: /v1/...) se planejar breaking changes
- Valide entrada (ex.: Joi, Zod)

Contribuição

1. Fork
2. Crie branch de feature: git checkout -b feature/minha-coisa
3. Commit e push
4. Abra PR descrevendo mudanças e testes

Licença

- Adicione aqui a licença do projeto (ex.: MIT).

Contato

- Crie uma issue ou mande mensagem para o mantenedor.

---
