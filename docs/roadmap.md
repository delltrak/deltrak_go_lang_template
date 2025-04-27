# Roadmap do Projeto

Este documento apresenta o roadmap para o desenvolvimento do template de Google Cloud Functions com GoLang, JWT e MySQL. O objetivo é fornecer uma visão clara das fases, marcos e entregáveis do projeto, garantindo alinhamento com os requisitos funcionais e não funcionais.

## Fases de Desenvolvimento

### Fase 1: Configuração e Estrutura Básica (MVP)
- **Duração**: 1-2 semanas
- **Objetivo**: Estabelecer a base do projeto com uma estrutura funcional mínima.
- **Entregáveis**:
  - Configuração do ambiente Google Cloud Functions com GoLang.
  - Estrutura inicial do projeto (diretórios, arquivos principais).
  - Implementação de uma função de exemplo (ex.: endpoint de saúde).
  - Documentação inicial: `stack.md`, `security.md`, `logging.md`.
- **Critérios de Sucesso**: Função de exemplo implantada e respondendo a requisições HTTP.

### Fase 2: Autenticação e Autorização
- **Duração**: 2-3 semanas
- **Objetivo**: Implementar autenticação segura com JWT e controle de acesso.
- **Entregáveis**:
  - Integração de autenticação JWT para endpoints protegidos.
  - Configuração de RBAC (Role-Based Access Control) com Google Cloud IAM.
  - Testes unitários para validação de tokens e autorização.
- **Critérios de Sucesso**: Apenas usuários autenticados com permissões adequadas podem acessar endpoints protegidos.

### Fase 3: Integração com Banco de Dados MySQL
- **Duração**: 2-3 semanas
- **Objetivo**: Conectar as funções ao banco de dados MySQL para persistência de dados.
- **Entregáveis**:
  - Configuração de conexão segura com Cloud SQL (MySQL) usando credenciais gerenciadas pelo Secret Manager.
  - Implementação de operações CRUD básicas em pelo menos um endpoint.
  - Uso de queries parametrizadas para prevenção de injeção SQL.
- **Critérios de Sucesso**: Funções conseguem ler e gravar dados no banco de dados de forma segura.

### Fase 4: Logging e Observabilidade
- **Duração**: 1-2 semanas
- **Objetivo**: Garantir que o sistema seja monitorável e rastreável.
- **Entregáveis**:
  - Implementação do esquema de logs JSON conforme definido em `logging.md`.
  - Integração com Google Cloud Logging e Trace para rastreamento de requisições.
  - Configuração de métricas customizadas e alertas no Google Cloud Monitoring.
- **Critérios de Sucesso**: Logs estruturados são gerados e visíveis no Google Cloud Logging; traces mostram o fluxo de requisições.

### Fase 5: Segurança Avançada e Testes
- **Duração**: 2-3 semanas
- **Objetivo**: Fortalecer a segurança e validar a robustez do sistema.
- **Entregáveis**:
  - Implementação de controles de segurança [MVP] conforme `security.md` (ex.: validação de entrada, TLS, gestão de segredos).
  - Testes de integração e de segurança (ex.: tentativas de injeção, autenticação inválida).
  - Relatório de cobertura de testes (mínimo 80% para código crítico).
- **Critérios de Sucesso**: Sistema resiste a ataques comuns (baseado em OWASP Top 10) e atinge cobertura de testes mínima.

### Fase 6: Otimização e Escalabilidade (Pós-MVP)
- **Duração**: 3-4 semanas
- **Objetivo**: Preparar o sistema para cargas maiores e melhorar performance.
- **Entregáveis**:
  - Configuração de caching para endpoints frequentes (ex.: usando Memorystore).
  - Ajustes de performance baseados em profiling com Google Cloud Profiler.
  - Planejamento de escalabilidade horizontal com gatilhos de auto-scaling.
- **Critérios de Sucesso**: Sistema mantém latência aceitável (<200ms) sob carga simulada.

### Fase 7: Conformidade e Auditoria (Later)
- **Duração**: 2-3 semanas
- **Objetivo**: Garantir conformidade com normas de segurança e privacidade.
- **Entregáveis**:
  - Implementação de logs de auditoria de longo prazo (1 ano de retenção).
  - Relatórios de conformidade com NIST SSDF v1.1 e GDPR (se aplicável).
  - Revisão de controles de segurança [Later] conforme `security.md`.
- **Critérios de Sucesso**: Sistema atende a requisitos de auditoria e privacidade.

## Marcos Principais

- **MVP Completo**: Final das Fases 1 a 5 (8-13 semanas a partir do início).
- **Pronto para Produção**: Final da Fase 6 (11-17 semanas).
- **Conformidade Total**: Final da Fase 7 (13-20 semanas).

## Considerações Finais

Este roadmap prioriza entregáveis críticos para o MVP nas primeiras fases, garantindo uma base sólida antes de avançar para otimizações e conformidade. Ajustes podem ser feitos com base em feedback de stakeholders ou mudanças nos requisitos. Cada fase será validada com testes e revisões para manter a qualidade do projeto.