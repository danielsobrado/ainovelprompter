import React from 'react';
import { Link } from 'react-router-dom';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import Container from '@mui/material/Container';

const Header: React.FC = () => {
  return (
    <AppBar position="static">
      <Container maxWidth="lg">
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            AI Novel Prompter
          </Typography>
          <Button component={Link} to="/" color="inherit">
            Home
          </Button>
          <Button component={Link} to="/texts" color="inherit">
            Texts
          </Button>
          <Button component={Link} to="/chapters" color="inherit">
            Chapters
          </Button>
          <Button component={Link} to="/feedback" color="inherit">
            Feedback
          </Button>
        </Toolbar>
      </Container>
    </AppBar>
  );
};

export default Header;