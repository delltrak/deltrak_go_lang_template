# Pilha Tecnológica

Este documento detalha a seleção de tecnologias e ferramentas para o desenvolvimento do template de Google Cloud Functions utilizando GoLang, com autenticação JWT e integração com MySQL. A escolha de cada componente foi avaliada com base em critérios como adequação ao negócio, expertise da equipe, comunidade, maturidade, riscos de segurança (OWASP Top 10 e NIST SSDF v1.1) e licença.

## Seleção de Tecnologias

### Linguagem: GoLang
- **Justificativa**: GoLang é ideal para funções serverless devido à sua performance, compilação rápida e suporte nativo a concorrência. É amplamente utilizado em ambientes de nuvem como Google Cloud.

### Plataforma Serverless: Google Cloud Functions
- **Justificativa**: Oferece escalabilidade automática, integração nativa com outros serviços do Google Cloud (como Cloud SQL) e suporte a GoLang.

### Banco de Dados: MySQL (Cloud SQL)
- **Justificativa**: MySQL é um banco de dados relacional maduro, com suporte robusto no Google Cloud via Cloud SQL, garantindo alta disponibilidade e segurança.

### Autenticação: JWT (JSON Web Tokens)
- **Justificativa**: Padrão amplamente adotado para autenticação stateless, compatível com arquiteturas serverless e com bibliotecas maduras em GoLang.

### Logging e Observabilidade: Google Cloud Logging e Monitoring
- **Justificativa**: Ferramentas nativas do Google Cloud que facilitam a integração com Cloud Functions, permitindo rastreamento de logs estruturados e métricas customizadas.

### CI/CD: Google Cloud Build
- **Justificativa**: Integração direta com Google Cloud Functions para automação de builds e deploys, com suporte a pipelines seguras.

## Matriz de Pontuação (1-5)

| **Tecnologia**            | **Adequação ao Negócio** | **Expertise da Equipe** | **Comunidade** | **Maturidade** | **Riscos OWASP/NIST** | **Licença** | **Pontuação Total** |
|---------------------------|--------------------------|-------------------------|----------------|----------------|-----------------------|-------------|---------------------|
| GoLang                   | 5                        | 4                       | 5              | 5              | 4                     | 5           | 28/30              |
| Google Cloud Functions   | 5                        | 4                       | 4              | 4              | 4                     | 5           | 26/30              |
| MySQL (Cloud SQL)        | 5                        | 5                       | 5              | 5              | 4                     | 5           | 29/30              |
| JWT                      | 5                        | 4                       | 5              | 5              | 4                     | 5           | 28/30              |
| Google Cloud Logging     | 5                        | 4                       | 4              | 4              | 4                     | 5           | 26/30              |
| Google Cloud Build (CI/CD) | 5                      | 4                       | 4              | 4              | 4                     | 5           | 26/30              |

## Considerações Finais

A pilha tecnológica escolhida prioriza a integração nativa com o ecossistema Google Cloud, garantindo escalabilidade, segurança e facilidade de manutenção. Cada componente foi selecionado para mitigar riscos de segurança conforme OWASP Top 10 e NIST SSDF v1.1, com pontuações altas em maturidade e suporte da comunidade. 