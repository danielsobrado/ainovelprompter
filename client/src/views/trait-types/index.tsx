import { Typography, Container, Paper } from '@mui/material';
import { DataGrid } from '@mui/x-data-grid'; // Updated import
import { AxiosError } from 'axios';
import { toast } from 'react-toastify';
import { useEffect, useState } from 'react';

// project imports
import MainCard from 'ui-component/cards/MainCard';
import api from '../../services/api';

// ==============================|| TRAIT TYPE LIST ||============================== //

interface TraitType {
  id: number;
  traitType: string;
  description: string;
}

const columns = [
  { field: 'id', headerName: 'ID', width: 100 },
  { field: 'traitType', headerName: 'Trait Type', width: 150 },
  { field: 'description', headerName: 'Description', width: 250 },
  { field: 'triggerText', headerName: 'Trigger Text', width: 250 }
];

const TraitTypeList = () => {
  const [traitTypes, setTraitTypes] = useState<TraitType[]>([]);

  useEffect(() => {
    const fetchTraitTypes = async () => {
      try {
        const response = await api.get('/trait-types');
        setTraitTypes(response.data);
      } catch (error) {
        if (error instanceof AxiosError) {
          const errorMessage = error.response?.data.error || 'Failed to fetch trait types';
          toast.error(errorMessage);
        } else {
          toast.error('An error occurred');
        }
      }
    };

    fetchTraitTypes();
  }, []);

  return (
    <MainCard title="Trait Types">
      <Container maxWidth="md">
        {traitTypes.length === 0 ? (
          <Typography variant="body1">No trait types found.</Typography>
        ) : (
          <Paper style={{ height: 400, width: '100%' }}>
            <DataGrid
              rows={traitTypes}
              columns={columns}
              checkboxSelection
            />
          </Paper>
        )}
      </Container>
    </MainCard>
  );
};

export default TraitTypeList;
