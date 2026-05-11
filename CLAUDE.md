# CLAUDE.md — Everest

## Visão Geral

**Módulo:** `github.com/elyosemite/everest`
**Go:** 1.24.0

Projeto iniciado há ~5 anos como caderno de estudos de Go. Está sendo ativamente transformado num **sistema bancário simples** com arquitetura limpa. Código de aprendizado antigo ainda existe em `src/` e `src/fundamentals/` — não remover sem perguntar, mas não é o foco atual.

---

## Exigências
- Nunca edita um arquivo se não for pedido;
- Se eu pedir algum exemplo de código sempre escreva no temrinal o código
- Por enquanto, ao gerar código não use recursos avançados tais como sync, gorountines, concorrência, etc

## Arquitetura

O projeto segue **Clean Architecture** com três camadas:

```
internal/domain/        ← entidades e regras de negócio puras
internal/application/   ← casos de uso (ainda sendo criados)
internal/infra/         ← implementações concretas (repositórios, etc.)
src/bank/               ← camada CLI / UI (menu interativo)
```

### Pacotes do domínio

| Pacote | Arquivo | Responsabilidade |
|--------|---------|-----------------|
| `account` | `internal/domain/account/account.go` | Entidade Account: saldo, depósito, saque |
| `account` | `internal/domain/account/account_type.go` | Enum AccountType (Checking / Savings) — **a mover do src/bank** |
| `user` | `internal/domain/user/user.go` | Entidade User com slice de accounts e addresses |
| `address` | `internal/domain/address/address.go` | Endereço no formato brasileiro (CEP, UF) |

### src/bank (CLI)

| Arquivo | Responsabilidade |
|---------|-----------------|
| `UIApplication.go` | Menu interativo, ponto de entrada via `StartUp()` |
| `customer.go` | Struct Customer — incompleta, a ser integrada ao domínio |
| `account_type.go` | Enum AccountType — a ser migrado para `internal/domain/account/` |

---

## Convenções Go obrigatórias

- **Construtores** sempre `NewX(...) (*X, error)` — nunca instanciar structs diretamente fora do pacote
- **Campos** sempre unexported; acesso via getters exportados
- **Erros** com `fmt.Errorf("contexto: %w", err)` para wrapping; `errors.New` para folha
- **Interfaces** pequenas e focadas — preferir poucos métodos e que definam bem o que se propõem a fazer; definir no pacote *consumidor*, não no produtor
- **Testes** com `t.Run` e table-driven quando houver múltiplos casos similares
- **Sem frameworks externos** — apenas stdlib (`testing`, `errors`, `fmt`, `regexp`, `time`)
- **Sem comentários óbvios** — só comentar WHY não-óbvio

---

## Comandos

```bash
# Rodar todos os testes
go test ./...

# Rodar testes com saída detalhada
go test -v ./...

# Rodar testes de um pacote específico
go test ./tests/domain/...
go test ./internal/...

# Build
go build ./...

# Rodar a aplicação
go run main.go
```

---

## Feature em andamento: Create Account

Plano completo em `docs/features/create-account.md`.

**Passos e status:**

| # | Descrição | Status |
|---|-----------|--------|
| 1 | Mover `AccountType` para o domínio; adicionar campo na entidade `Account` | Pendente |
| 2 | Interfaces de repositório (`UserRepository`, `AccountRepository`) | Pendente |
| 3 | Implementações in-memory dos repositórios | Pendente |
| 4 | Caso de uso `CreateAccount` | Pendente |
| 5 | Completar `Customer` + caso de uso `CreateUser` | Pendente |
| 6 | Wiring no CLI (`UIApplication.go` + DI em `main.go`) | Pendente |

**Regra do fluxo:** implementar um passo por vez; aguardar okay do usuário antes de avançar.

---

## Contexto do domínio bancário (Brasil)

- Documento do usuário: CPF (11 dígitos) ou CNPJ (14 dígitos)
- CEP: 8 dígitos sem traço no storage, formatado como `XXXXX-XXX` na exibição
- UF: 2 letras maiúsculas
- Tipos de conta: Conta Corrente (`Checking`) e Conta Poupança (`Savings`)
- Textos do menu CLI estão em português

---

## O que NÃO mexer sem perguntar

- `src/` e `src/fundamentals/` — código de estudo, pode ter valor histórico para o dono
- `src/lyingAround/` — dead code, mas deixar até confirmação de remoção
- Lógica de `Depositar` / `GetSaldo` em `UIApplication.go` — ainda usada ativamente no menu
