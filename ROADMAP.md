## Roadmap

### Scaffold

- [ ] Configurar golang
- [ ] Adicionar pacote do github api para golang
- [ ] Adicionar issues no repositório do github

### MVP

- [ ] Definir regra de negócio para geração do changelog
    - [ ] Issues
    - [ ] PRs
- [ ] Preencher template da documentação baseado nos valores da nova release
    - [ ] Issues
    - [ ] PRs
- [ ] Acrescentar CHANGELOG da nova release no início do arquivo existente
- [ ] Transformar o script final em um action

### Improvements

- [ ] Traduzir arquivos para inglês
- [ ] Gerar log do processamento dos passos

### Business rule

- [ ] Para issue:
    - [ ] A data de início deve ser a data (tag)
    - [ ] A data fim deve ser a data (tag)
    - [ ] Deve retornar apenas issues fechadas
    - [ ] O valor que deve retornar por issue deve ser o título

- [ ] Para PR:
    - [ ] A data de início deve ser a data (tag)
    - [ ] A data fim deve ser a data (tag)
    - [ ] Deve retornar apenas PR mergeado
    - [ ] O valor que deve retornar por PR deve ser, inicialmente, o título e, após, o valor do changelog retornado no body
