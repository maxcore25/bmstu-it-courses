import { axiosInstance } from '@/shared/api';
import { ordersMetadataSchema } from '../model/order.schema';

export async function getOrdersMetadata() {
  const { data } = await axiosInstance.get('/orders/metadata');
  return ordersMetadataSchema.parse(data);
}
