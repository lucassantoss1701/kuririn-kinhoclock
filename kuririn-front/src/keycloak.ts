// src/keycloak.ts
import Keycloak from 'keycloak-js';

const keycloakInstance = new Keycloak({
  url: 'http://localhost:8080/', // URL do seu servidor Keycloak
  realm: 'kuririncompany',              // Nome do seu realm
  clientId: 'frontend-react',          // ID do cliente configurado no Keycloak
});

export default keycloakInstance;
