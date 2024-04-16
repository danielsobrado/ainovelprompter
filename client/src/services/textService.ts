import api from './api';
import { AxiosError } from 'axios';

export interface Text {
  id: number;
  title: string;
  content: string;
}

export const fetchTexts = async (): Promise<Text[]> => {
  try {
    const response = await api.get('/texts');
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const fetchTextById = async (id: number): Promise<Text> => {
  try {
    const response = await api.get(`/texts/${id}`);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const createText = async (text: Partial<Text>): Promise<Text> => {
  try {
    const response = await api.post('/texts', text);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const updateText = async (id: number, text: Partial<Text>): Promise<Text> => {
  try {
    const response = await api.put(`/texts/${id}`, text);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const deleteText = async (id: number): Promise<void> => {
  try {
    await api.delete(`/texts/${id}`);
  } catch (error) {
    throw error;
  }
};