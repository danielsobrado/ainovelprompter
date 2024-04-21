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
  { field: 'id', headerName: 'ID', width: 50 },
  { field: 'traitType', headerName: 'Trait Type', flex: 1 },
  { field: 'description', headerName: 'Description', flex: 2 },
  { field: 'triggerText', headerName: 'Trigger Text', flex: 2 }
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
    <MainCard title="Trait Types" >
      <Container maxWidth="lg">
        {traitTypes.length === 0 ? (
          <Typography variant="body1">No trait types found.</Typography>
        ) : (
          <Paper style={{ height: '100%', width: '100%', overflow: 'hidden' }}>
            <DataGrid
              rows={traitTypes}
              columns={columns}
              checkboxSelection
              autoHeight
            />
          </Paper>
        )}
      </Container>
    </MainCard>
  );
};

export default TraitTypeList;
