import React, { useState, useEffect } from 'react';
import api from '../services/api';
import { AxiosError } from 'axios';
import { toast } from 'react-toastify';
import Typography from '@mui/material/Typography';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import Container from '@mui/material/Container';
import Select from '@mui/material/Select';
import MenuItem from '@mui/material/MenuItem';
import Checkbox from '@mui/material/Checkbox';
import ListItemText from '@mui/material/ListItemText';
import Chip from '@mui/material/Chip';
import Box from '@mui/material/Box';

interface Trait {
  key: string;
  triggerText: string;
}

const PromptGenerator: React.FC = () => {
  const [traits, setTraits] = useState<Trait[]>([]);
  const [selectedTraits, setSelectedTraits] = useState<string[]>([]);
  const [chapterLength, setChapterLength] = useState(1000);
  const [responseFormat, setResponseFormat] = useState('markdown');
  const [generatedPrompt, setGeneratedPrompt] = useState('');

  useEffect(() => {
    const fetchTraits = async () => {
      try {
        const response = await api.get('/traits');
        setTraits(response.data);
      } catch (error) {
        if (error instanceof AxiosError) {
          toast.error(error.response?.data.error || 'Failed to fetch traits');
        } else {
          toast.error('An error occurred');
        }
      }
    };

    fetchTraits();
  }, []);

  const handleTraitToggle = (traitKey: string) => {
    const currentIndex = selectedTraits.indexOf(traitKey);
    const newSelectedTraits = [...selectedTraits];

    if (currentIndex === -1) {
      newSelectedTraits.push(traitKey);
    } else {
      newSelectedTraits.splice(currentIndex, 1);
    }

    setSelectedTraits(newSelectedTraits);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const response = await api.post('/generate-prompt', {
        trait_keys: selectedTraits,
        chapter_length: chapterLength,
        response_format: responseFormat,
      });
      const prompt = response.data.prompt;
      setGeneratedPrompt(prompt);
    } catch (error) {
      if (error instanceof AxiosError) {
        toast.error(error.response?.data.error || 'Failed to generate prompt');
      } else {
        toast.error('An error occurred');
      }
    }
  };

  return (
    <Container maxWidth="md">
      <Typography variant="h4" component="h1" gutterBottom>
        Prompt Generator
      </Typography>
      <form onSubmit={handleSubmit}>
        <Select
          multiple
          value={selectedTraits}
          onChange={(e) => setSelectedTraits(e.target.value as string[])}
          renderValue={(selected) => (
            <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
              {selected.map((traitKey) => (
                <Chip key={traitKey} label={traitKey} />
              ))}
            </Box>
          )}
          fullWidth
          sx={{ mb: 2 }}
        >
          {traits.map((trait) => (
            <MenuItem key={trait.key} value={trait.key}>
              <Checkbox checked={selectedTraits.indexOf(trait.key) > -1} />
              <ListItemText primary={trait.triggerText} />
            </MenuItem>
          ))}
        </Select>
        <TextField
          label="Chapter Length"
          type="number"
          variant="outlined"
          fullWidth
          value={chapterLength}
          onChange={(e) => setChapterLength(parseInt(e.target.value))}
          sx={{ mb: 2 }}
        />
        <Select
          value={responseFormat}
          onChange={(e) => setResponseFormat(e.target.value as string)}
          fullWidth
          sx={{ mb: 2 }}
        >
          <MenuItem value="markdown">Markdown</MenuItem>
          <MenuItem value="json">JSON</MenuItem>
        </Select>
        <Button type="submit" variant="contained" color="primary">
          Generate Prompt
        </Button>
      </form>
      {generatedPrompt && (
        <Box sx={{ mt: 4 }}>
          <Typography variant="h6" gutterBottom>
            Generated Prompt:
          </Typography>
          <Typography variant="body1" whiteSpace="pre-wrap">
            {generatedPrompt}
          </Typography>
        </Box>
      )}
    </Container>
  );
};

export default PromptGenerator;