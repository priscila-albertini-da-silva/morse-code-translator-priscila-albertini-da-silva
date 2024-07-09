# Morse Code Translator Priscila Albertini da Silva

## Índice
- [Descrição](#descrição)
- [Descrição Técnica](#descrição-técnica)
- [Pré-requisitos](#pré-requisitos)
- [Instruções para Rodar o Programa](#instruções-para-rodar-o-programa)
  - [Compilar o Programa](#compilar-o-programa)
  - [Executar o Programa](#executar-o-programa)
  - [Inserir o Código Morse](#inserir-o-código-morse)
- [Instruções para Rodar os Testes](#instruções-para-rodar-os-testes)
  - [Executar os Testes Unitários](#executar-os-testes-unitários)
  - [Estrutura dos Testes](#estrutura-dos-testes)

## Descrição
Este projeto é um programa de linha de comando (CLI) em Go que traduz mensagens em código Morse para texto. O programa aceita uma string de entrada representando o código Morse e retorna a tradução correspondente em texto.

## Descrição Técnica
O programa consiste em:

- Um mapa (map[string]string) que associa cada sequência de código Morse ao seu respectivo caractere em texto.
- Uma função TranslateMorseToText que:
    - Remove espaços extras da entrada.
    - Separa a entrada em palavras usando três espaços como delimitador.
    - Separa cada palavra em caracteres Morse usando um único espaço como delimitador.
    - Traduz cada caractere Morse usando o mapa e constrói as palavras traduzidas.
    - Junta as palavras traduzidas em uma frase e a retorna.

## Pré-requisitos

- Go 1.22 ou superior.

## Instruções para Rodar o Programa

### Compilar o Programa

``` sh
go build main.go
```

### Executar o Programa

```sh
go run main.go
```

### Inserir o Código Morse
Digite ou cole o código Morse quando solicitado e pressione Enter.

**Exemplo de Uso**

Entrada:

```sh
.-.. . -   .. -   -... .
```	

Saída:

```sh
LET IT BE
``` 

## Instruções para Rodar os Testes

### Executar os Testes Unitários

Para garantir a cobertura completa do código, você pode executar os testes unitários utilizando o comando:

```sh
go test ./...
```

ou para informações detalhadas sobre a cobertura:

```sh
go test -v -cover ./...
```

### Estrutura dos Testes

Os testes unitários estão definidos no arquivo translator_test.go e cobrem os seguintes casos:

- Tradução de frases completas em código Morse.
- Tradução de palavras individuais.
- Tradução de números de 0 a 9.
- Tratamento de espaços extras.
- Tratamento de entradas vazias.
- Tratamento de entradas inválidas.