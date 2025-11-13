import { axiosInstance } from '@/shared/api';
import { MessageResponse } from '@/shared/types';

export const logout = async () => {
  const response = await axiosInstance.post<MessageResponse>('/auth/logout');
  return response.data;
};
