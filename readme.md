# Grpc com GO

## Introdução

- Go é uma linguagem com retrocompatibilidade, então tudo que é possível fazer na versão 1.15, é possível fazer na 1.16

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
      - 50021:50021
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
