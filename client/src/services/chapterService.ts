import api from './api';
import { AxiosError } from 'axios';

export interface Chapter {
  id: number;
  title: string;
  content: string;
}

export const fetchChapters = async (): Promise<Chapter[]> => {
  try {
    const response = await api.get('/chapters');
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const fetchChapterById = async (id: number): Promise<Chapter> => {
  try {
    const response = await api.get(`/chapters/${id}`);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const createChapter = async (chapter: Partial<Chapter>): Promise<Chapter> => {
  try {
    const response = await api.post('/chapters', chapter);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const updateChapter = async (id: number, chapter: Partial<Chapter>): Promise<Chapter> => {
  try {
    const response = await api.put(`/chapters/${id}`, chapter);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const deleteChapter = async (id: number): Promise<void> => {
  try {
    await api.delete(`/chapters/${id}`);
  } catch (error) {
    throw error;
  }
};