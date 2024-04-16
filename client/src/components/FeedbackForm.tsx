import React, { useState } from 'react';
import api from '../services/api';
import { AxiosError } from 'axios';
import { toast } from 'react-toastify';
import Typography from '@mui/material/Typography';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import Container from '@mui/material/Container';

const FeedbackForm: React.FC = () => {
  const [content, setContent] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await api.post('/feedback', { content });
      toast.success('Feedback submitted successfully');
      setContent('');
    } catch (error) {
      if (error instanceof AxiosError) {
        toast.error(error.response?.data.error || 'Failed to submit feedback');
      } else {
        toast.error('An error occurred');
      }
    }
  };

  return (
    <Container maxWidth="md">
      <Typography variant="h4" component="h1" gutterBottom>
        Feedback
      </Typography>
      <form onSubmit={handleSubmit}>
        <TextField
          label="Content"
          multiline
          rows={4}
          variant="outlined"
          fullWidth
          value={content}
          onChange={(e) => setContent(e.target.value)}
          sx={{ mb: 2 }}
        />
        <Button type="submit" variant="contained" color="primary">
          Submit
        </Button>
      </form>
    </Container>
  );
};

export default FeedbackForm;