import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import Container from '@mui/material/Container';
import Register from './Register';
import Login from './Login';
import Home from './Home';
import Header from './Header';
import Footer from './Footer';
import TextList from './TextList';
import TextDetail from './TextDetail';
import ChapterList from './ChapterList';
import ChapterDetail from './ChapterDetail';
import FeedbackForm from './FeedbackForm';

const theme = createTheme();

const App: React.FC = () => {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Router>
        <div className="app">
          <Header />
          <Container maxWidth="lg" sx={{ mt: 4 }}>
            <Routes>
                <Route path="/register" element={<Register />} />
                <Route path="/login" element={<Login />} />
                <Route path="/" element={<Home />} />
                <Route path="/texts" element={<TextList />} />
                <Route path="/texts/:id" element={<TextDetail />} />
                <Route path="/chapters" element={<ChapterList />} />
                <Route path="/chapters/:id" element={<ChapterDetail />} />
                <Route path="/feedback" element={<FeedbackForm />} />
            </Routes>
          </Container>
          <Footer />
          <ToastContainer />
        </div>
      </Router>
    </ThemeProvider>
  );
};

export default App;