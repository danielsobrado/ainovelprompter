import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { Typography, List, ListItem, ListItemText, Divider, Container } from '@mui/material';
import { AxiosError } from 'axios';
import { toast } from 'react-toastify';
import api from '../services/api';

interface Chapter {
  id: number;
  title: string;
}

const ChapterList: React.FC = () => {
  const [chapters, setChapters] = useState<Chapter[]>([]);

  useEffect(() => {
    const fetchChapters = async () => {
      try {
        const response = await api.get('/chapters');
        setChapters(response.data);
      } catch (error) {
        if (error instanceof AxiosError) {
          toast.error(error.response?.data.error || 'Failed to fetch chapters');
        } else {
          toast.error('An error occurred');
        }
      }
    };

    fetchChapters();
  }, []);

  return (
    <Container maxWidth="md">
      <Typography variant="h4" component="h1" gutterBottom>
        Chapters
      </Typography>
      {chapters.length === 0 ? (
        <Typography variant="body1">No chapters found.</Typography>
      ) : (
        <List>
          {chapters.map((chapter) => (
            <React.Fragment key={chapter.id}>
              <ListItem component={Link} to={`/chapters/${chapter.id}`}>
                <ListItemText primary={chapter.title} />
              </ListItem>
              <Divider />
            </React.Fragment>
          ))}
        </List>
      )}
    </Container>
  );
};

export default ChapterList;