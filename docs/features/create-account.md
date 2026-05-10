# Feature: Create Account for a User

## Contexto

O projeto já possui as seguintes bases:
- `internal/domain/user/user.go` — entidade `User` com `AddAccount()`
- `internal/domain/account/account.go` — entidade `Account` (sem AccountType)
- `internal/domain/address/address.go` — entidade `Address`
- `src/bank/account_type.go` — enum `AccountType` (Checking / Savings)
- `src/bank/customer.go` — struct `Customer` (incompleta, desconectada do domínio)
- `src/bank/UIApplication.go` — menu CLI interativo

## Objetivo

Implementar o fluxo completo de **criação de conta bancária para um usuário**, com arquitetura limpa (Clean Architecture), seguindo as convenções da comunidade Go.

---

## Passos

### Passo 1 — Mover `AccountType` para o domínio e enriquecer a entidade `Account`

**Arquivos a criar/modificar:**
- `internal/domain/account/account_type.go` ← mover enum de `src/bank/account_type.go`
- `internal/domain/account/account.go` ← adicionar campo `accountType AccountType`

**Mudanças:**
- Mover `AccountType` para o pacote de domínio `account`
- Atualizar `NewAccount` para aceitar `accountType AccountType`
- Adicionar getter `Type() AccountType`
- Manter `src/bank/account_type.go` apenas com re-export ou remover após migração

**Testes (`tests/domain/account_test.go`):**
- `TestNewAccount_WithValidCheckingType`
- `TestNewAccount_WithValidSavingsType`
- `TestNewAccount_WithNegativeInitialBalance`
- `TestAccountType_String`

---

### Passo 2 — Definir interfaces de repositório (Ports)

**Arquivos a criar:**
- `internal/domain/user/repository.go` — interface `UserRepository`
- `internal/domain/account/repository.go` — interface `AccountRepository`

**Interfaces:**
```go
// UserRepository
type Repository interface {
    Save(user *User) error
    FindByID(id string) (*User, error)
}

// AccountRepository
type Repository interface {
    Save(account *Account) error
    FindByID(id string) (*Account, error)
    FindByUserID(userID string) ([]*Account, error)
}
```

**Testes:**
- Nenhum teste direto (interfaces puras). Os testes virão nas implementações (Passo 3).

---

### Passo 3 — Implementações in-memory dos repositórios (Adapters)

**Arquivos a criar:**
- `internal/infra/repository/user_repository.go` — `InMemoryUserRepository`
- `internal/infra/repository/account_repository.go` — `InMemoryAccountRepository`

**Testes (`internal/infra/repository/`):**
- `TestInMemoryUserRepository_Save`
- `TestInMemoryUserRepository_FindByID_NotFound`
- `TestInMemoryAccountRepository_Save`
- `TestInMemoryAccountRepository_FindByUserID`

---

### Passo 4 — Caso de uso: `CreateAccount`

**Arquivos a criar:**
- `internal/application/account/create_account.go`

**Estrutura:**
```go
type CreateAccountInput struct {
    UserID         string
    AccountType    account.AccountType
    InitialBalance float64
}

type CreateAccountOutput struct {
    AccountID string
    Balance   float64
    Type      string
}

type CreateAccountUseCase struct {
    userRepo    user.Repository
    accountRepo account.Repository
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInput) (*CreateAccountOutput, error)
```

**Testes (`internal/application/account/create_account_test.go`):**
- `TestCreateAccount_Success_Checking`
- `TestCreateAccount_Success_Savings`
- `TestCreateAccount_UserNotFound`
- `TestCreateAccount_NegativeBalance`

---

### Passo 5 — Completar `Customer` e criar caso de uso `CreateUser`

**Arquivos a modificar/criar:**
- `src/bank/customer.go` → completar com constructor `NewCustomer` e validações
- `internal/application/user/create_user.go` — caso de uso `CreateUser`

**CreateUser Input/Output:**
```go
type CreateUserInput struct {
    Name     string
    Email    string
    Document string // CPF ou CNPJ
}

type CreateUserOutput struct {
    UserID string
    Name   string
}
```

**Testes:**
- `TestCreateUser_Success`
- `TestCreateUser_InvalidEmail`
- `TestCreateUser_InvalidDocument`
- `TestNewCustomer_Valid`

---

### Passo 6 — Wiring: integrar casos de uso ao menu CLI

**Arquivos a modificar:**
- `src/bank/UIApplication.go` — atualizar o menu para chamar `CreateUser` e `CreateAccount`
- `main.go` — injetar dependências (repositórios e casos de uso)

**Fluxo no menu:**
1. Usuário escolhe "Criar Conta"
2. CLI coleta: nome, email, CPF, tipo de conta, saldo inicial
3. Chama `CreateUser` → obtém `UserID`
4. Chama `CreateAccount` com `UserID` e demais dados
5. Exibe confirmação com ID da conta

**Testes (integração):**
- `TestStartUp_CreateAccount_Flow` (teste de integração simples com repositórios in-memory)

---

## Arquitetura Final

```
everest/
├── internal/
│   ├── domain/
│   │   ├── account/
│   │   │   ├── account.go          (entidade + AccountType)
│   │   │   ├── account_type.go     (enum movido do src/bank)
│   │   │   └── repository.go       (interface AccountRepository)
│   │   ├── user/
│   │   │   ├── user.go
│   │   │   └── repository.go       (interface UserRepository)
│   │   └── address/
│   │       └── address.go
│   ├── application/
│   │   ├── account/
│   │   │   └── create_account.go   (use case)
│   │   └── user/
│   │       └── create_user.go      (use case)
│   └── infra/
│       └── repository/
│           ├── user_repository.go
│           └── account_repository.go
├── src/bank/
│   ├── UIApplication.go            (atualizado)
│   ├── customer.go                 (completado)
│   └── account_type.go             (mantido ou removido)
├── tests/domain/
│   ├── user_test.go
│   └── account_test.go
└── docs/features/
    └── create-account.md           (este arquivo)
```

---

## Convenções Go aplicadas

- Construtores `NewX(...)` retornam `(*X, error)`
- Campos todos unexported, acesso via getters
- Interfaces pequenas e focadas (segregation)
- Erros com `errors.New` ou `fmt.Errorf` com `%w` para wrapping
- Testes com `t.Run` e casos em table-driven tests onde aplicável
- Nenhum framework externo — apenas `testing` da stdlib

---

## Status dos Passos

| Passo | Descrição                              | Status     |
|-------|----------------------------------------|------------|
| 1     | AccountType no domínio + Account       | [ ] Pendente |
| 2     | Interfaces de repositório              | [ ] Pendente |
| 3     | Implementações in-memory               | [ ] Pendente |
| 4     | Caso de uso CreateAccount              | [ ] Pendente |
| 5     | Customer completo + CreateUser         | [ ] Pendente |
| 6     | Wiring no CLI                          | [ ] Pendente |
