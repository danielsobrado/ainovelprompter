import api from './api';
import { AxiosError } from 'axios';

export interface Feedback {
  id: number;
  content: string;
  rating: number;
}

export const fetchFeedback = async (): Promise<Feedback[]> => {
  try {
    const response = await api.get('/feedback');
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const createFeedback = async (feedback: Partial<Feedback>): Promise<Feedback> => {
  try {
    const response = await api.post('/feedback', feedback);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const updateFeedback = async (id: number, feedback: Partial<Feedback>): Promise<Feedback> => {
  try {
    const response = await api.put(`/feedback/${id}`, feedback);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const deleteFeedback = async (id: number): Promise<void> => {
  try {
    await api.delete(`/feedback/${id}`);
  } catch (error) {
    throw error;
  }
};