import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import api from '../services/api';
import { logout } from '../services/authService';
import { AxiosError } from 'axios';
import { toast } from 'react-toastify';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';

const Home: React.FC = () => {
  const [user, setUser] = useState<{ username: string; email: string } | null>(null);

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const response = await api.get('/users/me');
        setUser(response.data);
      } catch (error) {
        if (error instanceof AxiosError) {
          toast.error(error.response?.data.error || 'Failed to fetch user data');
        } else {
          toast.error('An error occurred');
        }
      }
    };

    fetchUserData();
  }, []);

  const handleLogout = () => {
    logout();
    toast.success('Logged out successfully');
    setUser(null);
  };

  return (
    <Container maxWidth="sm">
      <Typography variant="h2" align="center" gutterBottom>
        Welcome to AI Novel Prompter
      </Typography>
      {user ? (
        <div>
          <Typography variant="h4" align="center" gutterBottom>
            Welcome, {user.username}!
          </Typography>
          <Button variant="contained" color="primary" onClick={handleLogout} fullWidth>
            Logout
          </Button>
        </div>
      ) : (
        <div>
          <Button component={Link} to="/login" variant="contained" color="primary" fullWidth sx={{ mb: 2 }}>
            Login
          </Button>
          <Button component={Link} to="/register" variant="contained" color="secondary" fullWidth>
            Register
          </Button>
        </div>
      )}
    </Container>
  );
};

export default Home;