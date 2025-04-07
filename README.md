# IPfinder

IPfinder é uma ferramenta de linha de comando escrita em Go que permite buscar informações sobre endereços IP e servidores na web.

## Funcionalidades

- **IPfinder**: Busca endereços de IP associados a um domínio
- **ServerFinder**: Encontra os servidores DNS associados a um domínio

## Pré-requisitos

- Go 1.16 ou superior
- Git (opcional, para clonar o repositório)

## Instalação

1. Clone o repositório:
```bash
git clone [https://github.com/jasz157/GoLang-Study]
cd IPfinder
```

2. Instale as dependências:
```bash
go mod download
```

3. Compile o projeto:
```bash
go build
```

## Como usar

### Buscar IPs de um domínio
```bash
./IPfinder IPfinder --host google.com
```

### Buscar servidores de um domínio
```bash
./IPfinder ServerFinder --host google.com
```

### Parâmetros

- `--host`: Domínio que você deseja pesquisar (padrão: www.google.com)

## Exemplos

Buscar IPs do Google:
```bash
./IPfinder IPfinder --host google.com
```

Buscar servidores do GitHub:
```bash
./IPfinder ServerFinder --host github.com
```

## Contribuição

Sinta-se à vontade para abrir issues e pull requests para melhorar o projeto.

## Licença

Este projeto está sob a licença MIT.
