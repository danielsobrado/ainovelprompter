import React, { useState, useEffect } from 'react';
import { DataGrid } from '@mui/x-data-grid';
import { Button, Dialog, DialogTitle, DialogContent, DialogActions, TextField } from '@mui/material';
import api from '../../services/api';

interface StandardPrompt {
  id: number;
  standard_name: string;
  title: string;
  prompt: string;
  created_at: string;
  version: number;
}

const columns = [
  { field: 'id', headerName: 'ID', width: 50 },
  { field: 'standard_name', headerName: 'Standard Name', flex: 1 },
  { field: 'title', headerName: 'Title', flex: 1  },
  { field: 'prompt', headerName: 'Prompt', flex: 4  },
  { field: 'created_at', headerName: 'Created At', flex: 1  },
  { field: 'version', headerName: 'Version', width: 80 },
];

const StandardPromptList: React.FC = () => {
  const [prompts, setPrompts] = useState<StandardPrompt[]>([]);
  const [selectedPrompt, setSelectedPrompt] = useState<StandardPrompt | null>(null);
  const [openDialog, setOpenDialog] = useState(false);

  useEffect(() => {
    fetchStandardPrompts();
  }, []);

  const fetchStandardPrompts = async () => {
    try {
      const response = await api.get('/standard-prompts');
      setPrompts(response.data);
    } catch (error) {
      console.error('Error fetching standard prompts:', error);
    }
  };

  const handleRowClick = (params: any) => {
    setSelectedPrompt(params.row);
    setOpenDialog(true);
  };

  const handleCloseDialog = () => {
    setSelectedPrompt(null);
    setOpenDialog(false);
  };

  const handleSavePrompt = async () => {
    if (selectedPrompt) {
      try {
        await api.put(`/standard-prompts/${selectedPrompt.id}`, selectedPrompt);
        fetchStandardPrompts();
        handleCloseDialog();
      } catch (error) {
        console.error('Error updating standard prompt:', error);
      }
    }
  };

  return (
    <div style={{ height: 400, width: '100%' }}>
      <DataGrid rows={prompts} columns={columns} onRowClick={handleRowClick} />

      <Dialog open={openDialog} onClose={handleCloseDialog}>
        <DialogTitle>Edit Standard Prompt</DialogTitle>
        <DialogContent>
          {selectedPrompt && (
            <div>
              <TextField
                label="Standard Name"
                value={selectedPrompt.standard_name}
                onChange={(e) => setSelectedPrompt({ ...selectedPrompt, standard_name: e.target.value })}
                fullWidth
                margin="normal"
              />
              <TextField
                label="Title"
                value={selectedPrompt.title}
                onChange={(e) => setSelectedPrompt({ ...selectedPrompt, title: e.target.value })}
                fullWidth
                margin="normal"
              />
              <TextField
                label="Prompt"
                value={selectedPrompt.prompt}
                onChange={(e) => setSelectedPrompt({ ...selectedPrompt, prompt: e.target.value })}
                fullWidth
                multiline
                rows={4}
                margin="normal"
              />
            </div>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseDialog}>Cancel</Button>
          <Button onClick={handleSavePrompt} color="primary">
            Save
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  );
};

export default StandardPromptList;