### Arquivo go.sum

Um módulo pode ter um arquivo de texto chamado go.sum em seu diretório raiz, junto com seu arquivo go.mod.  O arquivo go.sum contém hashes criptográficos das dependências diretas e indiretas do módulo.  Quando o comando go baixa um arquivo .mod ou .zip do módulo no cache do módulo, ele calcula um hash e verifica se o hash corresponde ao hash correspondente no arquivo go.sum do módulo principal.  go.sum pode estar vazio ou ausente se o módulo não tiver dependências ou se todas as dependências forem substituídas por diretórios locais usando diretivas de substituição.

See more: https://go.dev/ref/mod#go-sum-files