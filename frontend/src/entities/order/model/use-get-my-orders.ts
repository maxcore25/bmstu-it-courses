import { QUERY_KEYS } from '@/shared/config';
import { useQuery } from '@tanstack/react-query';
import { getOrders } from '../api/get-orders.api';

export function useMyGetOrders(expand?: string[]) {
  return useQuery({
    queryKey: [QUERY_KEYS.ORDERS, { expand }],
    queryFn: () => getOrders(expand),
  });
}
