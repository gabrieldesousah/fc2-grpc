# Grpc com GO

## Introdução

- Go é uma linguagem com retrocompatibilidade, então tudo que é possível fazer na versão 1.15, é possível fazer na 1.16
- Go não aceita ";". Todas as linhas simplismente terminam, sem pontuação.

## Configurando o ambiente com Docker

Dockerfile:

```Dockerfile
FROM golang:1.18-alpine

WORKDIR /go/src

CMD ["tail","-f","/dev/null"]
```

Obs.: Esta é a configuração básica de um container GO. Além disso, precisamos executar os pacotes para a instalação do protobuffer etc.

Docker-compose

```yml
version: "3"

services:
  app:
    build: .
    ports:
      - 50051:50051
    volumes:
      - .:/go/src
```

## Iniciando um módulo GO

É comum que o nome do módulo em GO seja o caminho do código que será disponibilizado:  
`go mod init github.com/gabrieldesousah/fc2-grpc`

## Instalando o Protoc e Protocolbuffer

Na imagem do GO, precisa instar o protobuffer

```sh
apk update && apk add protobuf
```

Em seguida, é preciso instalar as dependências do projeto GO:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get google.golang.org/grpc
```

Obs.: Para que as dependências instaladas fiquem no arquivo go.mod (afim de ter uma gestão delas), utilize o comando `go get` em cada dependência (ex.: `go get google.golang.org/protobuf/cmd/protoc-gen-go@latest`)

## Criando os protofiles

Crie um diretório chamado "proto/", onde estarão os protofiles, e "pb/", onde estarão o protobuffer. Crie um arquivo de exemplo proto/user.proto e coloque o seguinte código para dar início:

```proto
syntax = "proto3";
package pb;
option go_package = "./pb";

message User {
  // Obs.: Os valores numéricos são os índices
  optional string id = 1;
  string name = 2;
  string email = 3;
}

service UserService {
  rpc AddUser (User) returns (User);
}
```

### Gerando os arquivos do protobuffer

Precisamos gerar dois arquivos, um do protobuffer e outro do grpc. Para isso, execute:

```sh
# protoc --proto_path=proto proto/*.proto --go_out=pb
# protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb

protoc --proto_path=proto/ proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=. --go_out=.
```

## Implementando servidor gRPC

Criamos um arquivo "cmd/server/server.go", com o seguinte conteúdo:

```go
package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
)

func main () {
  lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
```

## Implementando serviços ao servidor gRPC

Criemos uma pasta chamada services/ com o arquivo user.go
## Fazendo chamadas gRPC

Instalando um cliente gRPC
`brew install evans`

O evans precisa do servidor gRPC no modo "reflection", portanto, é preciso adicionar a seguinte linha no código do server:
`
Iniciando o client:
`evans -r repl --host localhost --port 50051`

No evans, para selecionar o serviço do gRPC: `service UserService`  
Para Fazer a chamada de uma função: `call AddUser`

