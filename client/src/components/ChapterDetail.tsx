import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import api from '../services/api';
import { AxiosError } from 'axios';
import { toast } from 'react-toastify';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import Paper from '@mui/material/Paper';

interface Chapter {
  id: number;
  title: string;
  content: string;
}

const ChapterDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [chapter, setChapter] = useState<Chapter | null>(null);

  useEffect(() => {
    const fetchChapter = async () => {
      try {
        const response = await api.get(`/chapters/${id}`);
        setChapter(response.data);
      } catch (error) {
        if (error instanceof AxiosError) {
          toast.error(error.response?.data.error || 'Failed to fetch chapter');
        } else {
          toast.error('An error occurred');
        }
      }
    };

    fetchChapter();
  }, [id]);

  if (!chapter) {
    return <Typography variant="body1">Loading...</Typography>;
  }

  return (
    <Container maxWidth="md">
      <Paper elevation={3} sx={{ p: 4 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          {chapter.title}
        </Typography>
        <Typography variant="body1">{chapter.content}</Typography>
      </Paper>
    </Container>
  );
};

export default ChapterDetail;