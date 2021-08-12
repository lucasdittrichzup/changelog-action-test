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
