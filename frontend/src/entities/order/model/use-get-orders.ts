import { QUERY_KEYS } from '@/shared/config';
import { useQuery } from '@tanstack/react-query';
import { getOrders } from '../api/get-orders.api';

export function useGetOrders() {
  return useQuery({
    queryKey: [QUERY_KEYS.ORDERS],
    queryFn: getOrders,
  });
}
