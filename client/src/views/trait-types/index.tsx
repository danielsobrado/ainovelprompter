// material-ui
import { Typography, List, ListItem, ListItemText, Divider, Container } from '@mui/material';
import { AxiosError } from 'axios';
import { toast } from 'react-toastify';
import { useEffect, useState } from 'react';

// project imports
import MainCard from 'ui-component/cards/MainCard';
import api from '../../services/api';

// ==============================|| TRAIT TYPE LIST ||============================== //

interface TraitType {
  trait_type_id: number;
  trait_type: string;
  description: string;
}

const TraitTypeList = () => {
  const [traitTypes, setTraitTypes] = useState<TraitType[]>([]);

  useEffect(() => {
    const fetchTraitTypes = async () => {
      try {
        const response = await api.get('/trait-types');
        setTraitTypes(response.data);
      } catch (error) {
        if (error instanceof AxiosError) {
          toast.error(error.response?.data.error || 'Failed to fetch trait types');
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
          <List>
            {traitTypes.map((traitType) => (
              <div key={traitType.trait_type_id}>
                <ListItem>
                  <ListItemText primary={traitType.trait_type} secondary={traitType.description} />
                </ListItem>
                <Divider />
              </div>
            ))}
          </List>
        )}
      </Container>
    </MainCard>
  );
};

export default TraitTypeList;