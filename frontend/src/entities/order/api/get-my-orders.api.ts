import { axiosInstance } from '@/shared/api';
import { ordersSchema } from '../model/order.schema';

export async function getMyOrders(expand?: string[]) {
  let url = '/orders/my';
  if (expand && expand.length > 0) {
    const expandParam = expand.map(e => encodeURIComponent(e)).join(',');
    url += `?expand=${expandParam}`;
  }
  const { data } = await axiosInstance.get(url);

  return ordersSchema.parse(data);
}
