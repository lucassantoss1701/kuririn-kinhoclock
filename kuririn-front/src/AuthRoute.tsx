import { useKeycloak } from '@react-keycloak/web';
import { Navigate } from 'react-router-dom';
import { ReactNode } from 'react';

interface Props {
  children: ReactNode;
}
const PrivateRoute = ({ children }: Props) => {
  const { keycloak } = useKeycloak();

  return keycloak.authenticated ? children : <Navigate to="/login" />;
};

export default PrivateRoute;
