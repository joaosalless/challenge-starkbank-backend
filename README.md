# Challenge StarkBank Backend

Este repositório contém a implementação do desafio proposto pela StarkBank para desenvolvedores backend. A aplicação realiza a emissão de invoices e processa webhooks de recebimento, efetuando a transferência do valor recebido para uma conta especificada.

## Funcionalidades Implementadas

- [x] Emissão de invoices a cada 3 horas para destinatários aleatórios.
- [x] Processamento de webhooks para invoices pagas.
- [x] Cálculo do valor a ser transferido com base nas informações da invoice.
- [x] Transferência do valor para a conta configurada após o pagamento da invoice.
- [x] Implementação das camadas:
    - [x] Handler
        - [ ] Testes unitários
    - [x] Controller
        - [ ] Testes unitários
    - [x] Service
        - [x] Testes unitários

## Estrutura do Projeto

A estrutura do projeto segue uma arquitetura baseada em camadas, visando manter a separação de responsabilidades e facilitar a manutenção do código.

## Makefile

O projeto utiliza um `Makefile` para facilitar a execução das principais tarefas durante o desenvolvimento. Abaixo estão os comandos disponíveis:

### Comandos

- `make all`: Compila o projeto.
- `make api`: Inicia a API.
- `make schedule`: Inicia o agendador.
- `make test`: Executa os testes unitários.
- `make coverage`: Gera relatório de cobertura de testes unitários.
- `make mocks`: Gera os mocks das interfaces.
- `make clean`: Limpa arquivos gerados.
- `make docker-build`: Constrói a imagem Docker do projeto.
- `make docker-up`: Sobe os containers Docker.
- `make docker-down`: Derruba os containers Docker.

## Executando a Aplicação

Para rodar a aplicação, utilize os comandos do `Makefile` seguindo os passos abaixo:

### 1. Clone o projeto usando o git

```sh
git clone https://github.com/joaosalless/challenge-starkbank-joaosalless.git
cd challenge-starkbank-joaosalless
```

### 2. Instale o Go (se ainda não estiver instalado)

Siga as instruções específicas para o seu sistema operacional a partir do site oficial do Go: https://golang.org/doc/install

2.1 **Verifique se o Go foi instalado corretamente**

```sh
go version
```

### 3. Baixe as dependências do projeto listadas no go.mod

```sh
go mod tidy
```

### 4. Instale o MockGen (gerador de mocks). Opcional

```sh
go install github.com/golang/mock/mockgen@v1.6.0
```

4.1 **Verifique se o MockGen foi instalado corretamente**

```sh
mockgen --version
```

### 5. (Opcional) Atualize o PATH do Go para incluir o MockGen, se necessário

```sh
export PATH=$PATH:$(go env GOPATH)/bin
```

### 6. Instale o Docker (se ainda não estiver instalado)

Siga as instruções específicas para o seu sistema operacional a partir do site oficial do Docker:
https://docs.docker.com/get-docker/

6.1. Verifique se o Docker foi instalado corretamente

```sh
docker --version
```

### 7. Instale o Docker Compose (se necessário)

Algumas distribuições podem já incluir o Docker Compose com o Docker, caso contrário:
 https://docs.docker.com/compose/install/

7.1. Verifique se o Docker Compose foi instalado corretamente

```sh
docker-compose --version
```

### 8. Construa a imagem Docker do projeto (se necessário)

```sh
make docker-build
```

### 9. Instale o direnv (se necessário)

Siga as instruções específicas para o seu sistema operacional a partir do site oficial do direnv:
https://direnv.net

9.1 Verifique se o direnv foi instalado corretamente

```sh
direnv --version
```

### 10. Defina as variáveis de ambient

Crie um arquivo chamado `.env` na raiz do projeto e defina as variáveis abaixo

```sh
# APP
APP_DEBUG=true
APP_ENV=development
APP_DEBUG=true

# API
API_PORT=8080

# Timezone
CLOCK_LOCATION="UTC"

# Invoice
INVOICE_EXPIRATION_DAYS=
INVOICE_RANDOM_INVOICES_NUMBER_MIN=
INVOICE_RANDOM_INVOICES_NUMBER_MAX=

# Starkbank
BANK_ACCOUNT_NAME=
BANK_ACCOUNT_TAX_ID=
BANK_ACCOUNT_BANK_CODE=
BANK_ACCOUNT_BRANCH_CODE=
BANK_ACCOUNT_ACCOUNT_NUMBER=
BANK_ACCOUNT_ACCOUNT_TYPE=

# Schedule
SCHEDULER_ENABLED=true
INVOICE_CREATE_SCHEDULED_TIME='@every 3h'

# Bank Gateway
STARKBANK_PROJECT_ID=
STARKBANK_ENVIRONMENT=
STARKBANK_PRIVATE_KEY=
STARKBANK_PUBLIC_KEY=
```

10.1 Carregue as variáveis de ambiente no direnv

```sh
direnv allow
```

### 11. Execute os containers Docker

11.1 Construa a imagem Docker do projeto

```sh
make docker-build
```

11.2 Suba os containers Docker

```sh
make docker-up
```

11.3 Para encerrar a aplicação, derrube os containers docker

```sh
make docker-down
```

