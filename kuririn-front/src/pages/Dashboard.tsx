import React, { useEffect, useState } from 'react';
import api from '../api';
import { useKeycloak } from '@react-keycloak/web';
import styles from './Dashboard.module.css'; // Importando o CSS Module


const Dashboard: React.FC = () => {
  const [data, setData] = useState<string | null>(null);
  const { keycloak } = useKeycloak();


  useEffect(() => {
    console.log("aaaa ", keycloak.token)
    const fetchData = async () => {
      try {
        const response = await api.get('/dashboard', {
          headers: {
            Authorization: `Bearer ${keycloak.token}`,
          },
        });
        setData(response.data);
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    fetchData();
  }, [keycloak.token]);

  const handleLogout = () => {
    keycloak.logout();
  };

  return (
    <div className={styles.dashboard}>
      <h1 className={styles.title}>Welcome to the Dashboard!</h1>
      <button className={styles.logoutButton} onClick={handleLogout}>Logout</button>
      <p>{data}</p>
    </div>
  );
};

export default Dashboard;
