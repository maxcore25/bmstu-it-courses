import { axiosInstance } from '@/shared/api';
import { LoginFormValues, Tokens } from '@/shared/types';

export const login = async (data: LoginFormValues) => {
  const response = await axiosInstance.post<Tokens>('/auth/login', data);
  return response.data;
};
