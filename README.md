# Rest API Sample Go/Slice
## API Restful que utiliza livros como exemplo

<p align="center">
  <img src="images/golang.png"/ alt="Golang">
</p>


A aplicação foi escrita totalmente em Go 🐹, visando utilizar o mínimo de dependências possíveis, tratar a maioria dos erros com os devidos cuidados e documentada com comentários de fácil entendimento

Pacotes utilizados

- Gorilla/Mux 🦍


## Features

- Na aplicação é possível ver os livros já adicionados previamente
- Adicionar novos livros
- Ver livros específicos
- Excluir algum livros desejado
- Atualizar as informações


## Requisitos

```sh
Golang: https://golang.org/dl/

API Client de sua preferência (O que aparece nas imagens se chama Postman)

Após ter instalado o go
Gorilla mux: go get -u github.com/gorilla/mux
```


## Na prática

Iniciando a aplicação

![](images/1-starting.png)


Acessando os livros já inseridos

![](images/2-gettingbooks.png)


Adicionando novos livros

![](images/3-addingbooks01.png)

![](images/3-addingbooks02.png)

![](images/3-addingbooks03check.png)


Atualizando o livro desejado

![](images/4-updatebooks01.png)

![](images/4-updatebooks02check.png)


Deletando o livro desejado

![](images/5-deletebook01.png)

![](images/5-deletebook02check.png)



## Utilização

Basta utilizar o comando go build e aproveitar o aplicativo! 😊
