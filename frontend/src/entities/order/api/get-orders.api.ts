import { axiosInstance } from '@/shared/api';
import { ordersSchema } from '../model/order.schema';

export async function getOrders() {
  const { data } = await axiosInstance.get('/orders');
  return ordersSchema.parse(data);
}
