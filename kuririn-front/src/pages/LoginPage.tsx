import React from 'react';
import { useKeycloak } from '@react-keycloak/web';
import { Navigate } from 'react-router-dom';

const LoginPage: React.FC = () => {
  const { keycloak } = useKeycloak();

  // Se o usuário já estiver autenticado, redirecione para o Dashboard
  if (keycloak.authenticated) {
    return <Navigate to="/dashboard" />;
  }

  const handleLogin = () => {
    keycloak.login(); // Inicia o processo de login 
  };

  return (
    <div>
       <h1>React com Keycloak</h1>
      {!keycloak.authenticated ? (
        <button onClick={handleLogin}>Login</button>
      ) : (
        <>
          <p>Bem-vindo, {keycloak.tokenParsed?.preferred_username}!</p>
          <button onClick={() => keycloak.logout()}>Logout</button>
        </>
      )}
  
    </div>
    
  );
};

export default LoginPage;
