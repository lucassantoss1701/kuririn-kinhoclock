// src/main.tsx
import ReactDOM from 'react-dom/client';
import { ReactKeycloakProvider } from '@react-keycloak/web';
import keycloak from './keycloak'; // Importe a instância criada
import App from './App';
import './index.css';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <>
    <ReactKeycloakProvider authClient={keycloak}>
      <App />
    </ReactKeycloakProvider>
  </>
);
