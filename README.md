# Criptoanálise em Go

Projeto em Go que realiza análise de hashes utilizando wordlists para encontrar correspondências entre entradas pré-definidas e senhas conhecidas.

## Descrição

O programa lê uma lista de hashes a partir de um arquivo e, para cada hash, percorre uma wordlist de senhas possíveis. Cada senha é processada pela função de análise para identificar possíveis correspondências.

## Estrutura do Projeto

```text
.
├── main.go
├── cryptoanalysis/
│   └── validate/
├── wordlists/
│   ├── hash.txt
│   └── wordlist.txt
└── README.md
```

## Funcionamento

1. O arquivo `hash.txt` é carregado com os hashes.
2. Cada hash é processado individualmente.
3. O arquivo `wordlist.txt` é carregado com possíveis senhas.
4. Cada senha é enviada para análise em conjunto com o hash atual.
5. A lógica de análise é executada no pacote `validate`.

## Organização do Código

- `main.go`: responsável por ler os arquivos e controlar o fluxo
- `cryptoanalysis/validate`: contém a lógica de análise de hashes
- `wordlists/`: arquivos de entrada (hashes e senhas)

## Licença

Projeto para fins de estudo e pesquisa em Go e processamento de hashes.

![Go](https://img.shields.io/badge/Go-1.26.3-blue)
