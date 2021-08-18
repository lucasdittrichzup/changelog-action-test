## Roadmap

### Scaffold

- [x] Configurar golang
- [x] Adicionar pacote do github api para golang
- [x] Adicionar issues no repositório do github
- [x] Configurar variáveis de ambiente

### MVP

- [x] Definir regra de negócio para geração do changelog
    - [x] Issues
    - [x] PRs
- [x] Preencher template da documentação baseado nos valores da nova release
    - [x] Issues
    - [x] PRs
- [ ] Validações
    - [ ] Caso o arquivo não exista
    - [ ] Se não filtrar os dados corretamente, o changelog não deve ser gerado
- [ ] Acrescentar CHANGELOG da nova release no início do arquivo existente
- [ ] Transformar o script final em um action

### Improvements

- [ ] Traduzir arquivos para inglês
- [ ] Gerar log do processamento dos passos
- [ ] Gerar changelog retroativo
- [ ] Escolher nome da action

### Business rule

- [x] Para issue:
    - [x] A data de início deve ser a data de publicação da release anterior
    - [x] A data fim deve ser a data de publicação da próxima release
    - [x] Deve retornar apenas issues fechadas
    - [x] O valor que deve retornar por issue deve ser o título

- [x] Para PR:
    - [x] A data de início deve ser a data de publicação da release anterior
    - [x] A data fim deve ser a data de publicação da próxima release
    - [x] Deve retornar apenas PR mergeado
    - [x] O valor que deve retornar por PR deve ser, inicialmente, o título e, após, o valor do changelog retornado no body
