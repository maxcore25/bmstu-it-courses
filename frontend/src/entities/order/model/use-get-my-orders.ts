import { QUERY_KEYS } from '@/shared/config';
import { useQuery } from '@tanstack/react-query';
import { getMyOrders } from '../api/get-my-orders.api';

export function useGetMyOrders(expand?: string[]) {
  return useQuery({
    queryKey: [QUERY_KEYS.ORDERS, { expand }],
    queryFn: () => getMyOrders(expand),
  });
}
