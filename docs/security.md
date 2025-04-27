# Segurança do Projeto

Este documento descreve a abordagem de segurança para o desenvolvimento do template de Google Cloud Functions com GoLang, JWT e MySQL. As ameaças são mapeadas conforme o **OWASP Top 10 (2021)**, e as mitigações seguem as diretrizes do **NIST SSDF v1.1**. Cada controle de segurança é marcado como [MVP] (para implementação imediata) ou [Later] (para fases posteriores).

## Mapeamento de Ameaças e Mitigações

### 1. Broken Access Control (OWASP A01:2021)
- **Ameaça**: Falhas no controle de acesso podem permitir que usuários não autorizados acessem funções ou dados sensíveis.
- **Mitigação**: Implementar autenticação baseada em JWT com validação rigorosa de tokens e políticas de autorização baseadas em papéis (RBAC) no Google Cloud IAM. [MVP]
- **NIST SSDF**: PO.2 (Definir políticas de segurança) e PS.2 (Proteger sistemas com autenticação forte).

### 2. Cryptographic Failures (OWASP A02:2021)
- **Ameaça**: Uso de criptografia fraca ou armazenamento inadequado de segredos pode expor dados sensíveis.
- **Mitigação**: Utilizar Google Cloud Secret Manager para armazenar chaves JWT e credenciais de banco de dados; garantir que todas as comunicações usem TLS 1.2 ou superior. [MVP]
- **NIST SSDF**: PS.3 (Proteger dados sensíveis) e RV.1 (Revisar vulnerabilidades conhecidas).

### 3. Injection (OWASP A03:2021)
- **Ameaça**: Injeção de SQL ou outros comandos maliciosos via entradas de usuário no MySQL.
- **Mitigação**: Usar consultas parametrizadas e bibliotecas seguras de banco de dados em GoLang (como `database/sql`); implementar validação de entrada rigorosa. [MVP]
- **NIST SSDF**: PW.6 (Desenvolver software seguro contra injeções).

### 4. Insecure Design (OWASP A04:2021)
- **Ameaça**: Falhas de design podem levar a vulnerabilidades que não são corrigidas por controles técnicos.
- **Mitigação**: Adotar modelagem de ameaças (STRIDE) durante o design da arquitetura e revisar com pares antes da implementação. [Later]
- **NIST SSDF**: PO.1 (Definir práticas de segurança no ciclo de vida).

### 5. Security Misconfiguration (OWASP A05:2021)
- **Ameaça**: Configurações padrão ou permissões excessivas no Google Cloud Functions ou Cloud SQL.
- **Mitigação**: Aplicar o princípio do menor privilégio (least privilege) nas permissões do IAM e desativar opções inseguras por padrão; usar ferramentas de verificação de configuração do Google Cloud. [MVP]
- **NIST SSDF**: PS.1 (Proteger componentes com configurações seguras).

### 6. Vulnerable and Outdated Components (OWASP A06:2021)
- **Ameaça**: Dependências desatualizadas em GoLang ou no ambiente do Google Cloud podem conter vulnerabilidades conhecidas.
- **Mitigação**: Utilizar Go Modules para gerenciar dependências e configurar alertas de segurança no Google Cloud para componentes vulneráveis; gerar um SBOM (Software Bill of Materials). [MVP]
- **NIST SSDF**: RV.2 (Monitorar e corrigir vulnerabilidades).

### 7. Identification and Authentication Failures (OWASP A07:2021)
- **Ameaça**: Falhas na autenticação JWT podem permitir acesso não autorizado.
- **Mitigação**: Implementar expiração de tokens, renovação segura e validação de assinatura com bibliotecas confiáveis (como `github.com/golang-jwt/jwt`). [MVP]
- **NIST SSDF**: PW.7 (Implementar autenticação segura).

### 8. Software and Data Integrity Failures (OWASP A08:2021)
- **Ameaça**: Falhas na integridade de código ou dados durante CI/CD ou deployment.
- **Mitigação**: Usar Google Cloud Build com verificações de integridade e assinatura de artefatos; implementar testes automatizados de integridade. [Later]
- **NIST SSDF**: PW.8 (Garantir integridade no ciclo de desenvolvimento).

### 9. Security Logging and Monitoring Failures (OWASP A09:2021)
- **Ameaça**: Falta de logs adequados pode dificultar a detecção de incidentes.
- **Mitigação**: Configurar logs estruturados (JSON) no Google Cloud Logging com campos como timestamp, level, service, trace_id e user_id; ativar alertas no Google Cloud Monitoring. [MVP]
- **NIST SSDF**: RV.3 (Monitorar e responder a incidentes).

### 10. Server-Side Request Forgery (SSRF) (OWASP A10:2021)
- **Ameaça**: Funções podem ser manipuladas para acessar recursos internos não autorizados.
- **Mitigação**: Restringir chamadas de rede externas nas políticas de VPC Service Controls do Google Cloud; validar URLs de entrada. [Later]
- **NIST SSDF**: PW.6 (Proteger contra abusos de funcionalidades).

## Checklist Secure-by-Design

- [x] Autenticação forte com JWT e IAM (A01, A07) [MVP]
- [x] Criptografia de dados e segredos com Secret Manager (A02) [MVP]
- [x] Proteção contra injeção com consultas parametrizadas (A03) [MVP]
- [x] Configurações seguras e menor privilégio (A05) [MVP]
- [x] Gerenciamento de dependências e SBOM (A06) [MVP]
- [x] Logging estruturado e monitoramento (A09) [MVP]
- [ ] Modelagem de ameaças e revisão de design (A04) [Later]
- [ ] Verificações de integridade no CI/CD (A08) [Later]
- [ ] Proteção contra SSRF com VPC Service Controls (A10) [Later]

## Considerações Finais

A abordagem de segurança prioriza controles críticos para o MVP, mitigando as ameaças mais relevantes do OWASP Top 10. Controles adicionais serão implementados em fases posteriores para garantir conformidade contínua com NIST SSDF v1.1. 