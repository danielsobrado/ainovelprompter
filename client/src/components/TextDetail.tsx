import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import api from '../services/api';
import { AxiosError } from 'axios';
import { toast } from 'react-toastify';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import Paper from '@mui/material/Paper';

interface Text {
  id: number;
  title: string;
  content: string;
}

const TextDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [text, setText] = useState<Text | null>(null);

  useEffect(() => {
    const fetchText = async () => {
      try {
        const response = await api.get(`/texts/${id}`);
        setText(response.data);
      } catch (error) {
        if (error instanceof AxiosError) {
          toast.error(error.response?.data.error || 'Failed to fetch text');
        } else {
          toast.error('An error occurred');
        }
      }
    };

    fetchText();
  }, [id]);

  if (!text) {
    return <Typography variant="body1">Loading...</Typography>;
  }

  return (
    <Container maxWidth="md">
      <Paper elevation={3} sx={{ p: 4 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          {text.title}
        </Typography>
        <Typography variant="body1">{text.content}</Typography>
      </Paper>
    </Container>
  );
};

export default TextDetail;