import { axiosInstance } from '@/shared/api';
import { CreateOrderValues } from '../model/create-order.schema';

export const createOrder = async (payload: CreateOrderValues) => {
  const { data } = await axiosInstance.post('/orders', payload);
  return data;
};
