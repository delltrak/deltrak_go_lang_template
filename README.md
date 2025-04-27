# Google Cloud Functions Template com GoLang

Este é um template para criar funções no Google Cloud Functions usando GoLang, com suporte a autenticação JWT e integração com MySQL.

## Pré-requisitos

- Go 1.21 ou superior
- Conta no Google Cloud Platform com o SDK configurado
- MySQL instalado e configurado

## Configuração

1. **Clone este repositório** ou copie os arquivos para o seu projeto.

2. **Configure as credenciais do banco de dados** no arquivo `main.go`, ajustando as constantes `dbUser`, `dbPassword`, `dbHost`, `dbPort` e `dbName` conforme necessário.

3. **Configure a chave secreta JWT** no arquivo `main.go`, alterando a variável `jwtSecret` para uma chave segura (em produção, use variáveis de ambiente ou um gerenciador de segredos).

4. **Instale as dependências** executando:
   ```bash
   go mod tidy
   ```

## Estrutura do Projeto

- `main.go`: Contém a função principal do Google Cloud Functions, com autenticação JWT, conexão ao MySQL e um endpoint de exemplo `/animals`.
- `go.mod`: Arquivo de módulo Go com as dependências necessárias.

## Uso

### Teste Local

Para testar localmente, use o framework de funções do Google Cloud:

```bash
functions-framework --target=HandleFunction
```

Acesse `http://localhost:8080/animals` com um token JWT válido no cabeçalho `Authorization`.

### Implantação no Google Cloud

Para implantar a função no Google Cloud, use o comando:

```bash
gcloud functions deploy HandleFunction --runtime go121 --trigger-http --allow-unauthenticated
```

**Nota:** Remova `--allow-unauthenticated` se quiser que a função exija autenticação.

## Autenticação JWT

Envie um token JWT no cabeçalho `Authorization` no formato `Bearer <token>`. O token será validado contra a chave secreta definida no código.

## Endpoint de Exemplo

- **GET /animals**: Lista todos os animais do banco de dados com suporte a paginação. Parâmetros:
  - `page`: Número da página (padrão: 1)
  - `limit`: Número de itens por página (padrão: 10)

## Depuração

- Verifique os logs para erros de conexão com o banco de dados ou validação de token.
- Certifique-se de que o banco de dados MySQL está acessível e que as credenciais estão corretas.
- Para problemas com JWT, confirme se o token enviado é válido e se a chave secreta está correta.

## Licença

Este template é fornecido sob a licença MIT. Sinta-se à vontade para usá-lo e modificá-lo conforme necessário. 