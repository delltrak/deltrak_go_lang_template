# Logging e Observabilidade

Este documento define as práticas de logging e observabilidade para o template de Google Cloud Functions com GoLang, JWT e MySQL. O objetivo é garantir rastreabilidade, depuração eficiente e conformidade com requisitos de segurança.

## Esquema de Logs (JSON)

Os logs serão estruturados em formato JSON para facilitar a análise automatizada e integração com o Google Cloud Logging. O esquema inclui os seguintes campos obrigatórios:

- **timestamp**: Data e hora do evento no formato ISO 8601 (ex.: `2023-10-01T14:30:00Z`).
- **level**: Nível de severidade do log (`DEBUG`, `INFO`, `WARN`, `ERROR`, `FATAL`).
- **service**: Nome do serviço ou função que gerou o log (ex.: `auth-function`).
- **trace_id**: Identificador único para rastrear uma solicitação através de múltiplos serviços (gerado pelo Google Cloud Trace ou manualmente).
- **span_id**: Identificador de uma operação específica dentro de um trace.
- **user_id**: Identificador do usuário associado à solicitação (anonimizado se necessário para conformidade com privacidade).
- **msg**: Mensagem descritiva do evento (ex.: `User authentication successful`).
- **meta**: Objeto JSON contendo metadados adicionais contextuais (ex.: `{ "http_method": "POST", "endpoint": "/login" }`).

## Regras de Correlação-ID

- Cada solicitação HTTP recebida por uma Google Cloud Function deve gerar ou receber um `trace_id` único.
- O `trace_id` será propagado através de cabeçalhos HTTP (como `X-Cloud-Trace-Context`) para serviços downstream.
- O `span_id` será gerado para cada operação significativa dentro da função (ex.: autenticação, consulta ao banco).
- Bibliotecas como `cloud.google.com/go/trace` serão usadas para integração nativa com Google Cloud Trace.

## Retenção de Logs

- **Período de Retenção**: Logs serão retidos por 30 dias no Google Cloud Logging para análise de curto prazo (conformidade com MVP).
- **Logs de Auditoria**: Logs críticos de segurança (ex.: falhas de autenticação, erros de autorização) serão exportados para um bucket do Google Cloud Storage com retenção de 1 ano (implementação [Later]).

## Limpeza de PII (Informações Pessoalmente Identificáveis)

- Dados sensíveis como senhas, tokens JWT completos ou informações de usuário (ex.: CPF, endereço) nunca serão registrados nos logs.
- Campos como `user_id` serão anonimizados (ex.: hash do valor real) para evitar exposição de PII.
- Mensagens de log serão filtradas para remover qualquer dado sensível antes de serem gravadas, utilizando bibliotecas como `github.com/securego/gosec` para validação.

## Integração com Observabilidade

- **Google Cloud Logging**: Todos os logs serão enviados para o Google Cloud Logging para centralização e análise.
- **Google Cloud Monitoring**: Métricas customizadas serão definidas para monitorar latência das funções, taxa de erros (5xx) e tentativas de autenticação falhas, com alertas configurados para thresholds críticos.
- **Trace e Profiling**: Google Cloud Trace será usado para rastrear latência entre operações; profiling será ativado periodicamente para identificar gargalos de performance.

## Exemplo de Linhas de Log

Abaixo estão exemplos de logs estruturados em JSON que seguem o esquema definido:

```json
{
  "timestamp": "2023-10-01T14:30:00Z",
  "level": "INFO",
  "service": "auth-function",
  "trace_id": "1234567890abcdef1234567890abcdef",
  "span_id": "span-001",
  "user_id": "hashed-user-123",
  "msg": "User authentication successful",
  "meta": {
    "http_method": "POST",
    "endpoint": "/login",
    "response_time_ms": 120
  }
}
```

```json
{
  "timestamp": "2023-10-01T14:30:05Z",
  "level": "ERROR",
  "service": "auth-function",
  "trace_id": "1234567890abcdef1234567890abcdef",
  "span_id": "span-002",
  "user_id": "hashed-user-123",
  "msg": "Database connection failed",
  "meta": {
    "http_method": "POST",
    "endpoint": "/login",
    "error_code": "ECONNREFUSED"
  }
}
```

## Considerações Finais

As práticas de logging e observabilidade descritas garantem que o sistema seja rastreável e monitorável desde o MVP, com extensões planejadas para auditoria de longo prazo e conformidade com privacidade. A integração com ferramentas nativas do Google Cloud reduz a complexidade operacional e melhora a capacidade de resposta a incidentes. 