import { axiosInstance } from '@/shared/api';

export const deleteOrder = async (id: string) => {
  const { data } = await axiosInstance.delete(`/orders/${id}`);
  return data;
};
