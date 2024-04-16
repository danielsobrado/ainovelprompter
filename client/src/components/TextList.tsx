import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import api from '../services/api';
import { AxiosError } from 'axios';
import { toast } from 'react-toastify';
import Typography from '@mui/material/Typography';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import Divider from '@mui/material/Divider';
import Container from '@mui/material/Container';

interface Text {
  id: number;
  title: string;
  content: string;
}

const TextList: React.FC = () => {
  const [texts, setTexts] = useState<Text[]>([]);

  useEffect(() => {
    const fetchTexts = async () => {
      try {
        const response = await api.get('/texts');
        setTexts(response.data);
      } catch (error) {
        if (error instanceof AxiosError) {
          toast.error(error.response?.data.error || 'Failed to fetch texts');
        } else {
          toast.error('An error occurred');
        }
      }
    };

    fetchTexts();
  }, []);

  return (
    <Container maxWidth="md">
      <Typography variant="h4" component="h1" gutterBottom>
        Texts
      </Typography>
      {texts.length === 0 ? (
        <Typography variant="body1">No texts found.</Typography>
      ) : (
        <List>
          {texts.map((text) => (
            <React.Fragment key={text.id}>
              <ListItem component={Link} to={`/texts/${text.id}`}>
                <ListItemText primary={text.title} />
              </ListItem>
              <Divider />
            </React.Fragment>
          ))}
        </List>
      )}
    </Container>
  );
};

export default TextList;