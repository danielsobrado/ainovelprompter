import React from 'react';
import { useState } from 'react';
import { Box, Button, TextField, Typography, CircularProgress } from '@mui/material';
import api from '../../../../services/api';

const OllamaChat = () => {
  const [prompt, setPrompt] = useState('');
  const [response, setResponse] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setLoading(true);

    try {
      const res = await api.post('/ollama/generate', {
        model: 'llama3',
        prompt: prompt,
        stream: false
      });
      setResponse(res.data.response);
    } catch (error) {
      console.error('Error:', error);
      setResponse('An error occurred while fetching the response.');
    }

    setLoading(false);
  };

  return (
    <Box sx={{ p: 3 }}>
      <Typography variant="h4" gutterBottom>
        Ollama Chat
      </Typography>
      <Box component="form" onSubmit={handleSubmit}>
        <TextField
          fullWidth
          label="Enter your prompt"
          multiline
          rows={4}
          value={prompt}
          onChange={(e) => setPrompt(e.target.value)}
          variant="outlined"
          sx={{ mb: 2 }}
        />
        <Button variant="contained" color="primary" type="submit" disabled={loading}>
          {loading ? <CircularProgress size={24} /> : 'Send'}
        </Button>
      </Box>
      {response && (
        <Box sx={{ mt: 4 }}>
          <Typography variant="h6" gutterBottom>
            Response:
          </Typography>
          <Typography variant="body1" sx={{ whiteSpace: 'pre-line' }}>
            {response}
          </Typography>
        </Box>
      )}
    </Box>
  );
};

export default OllamaChat;