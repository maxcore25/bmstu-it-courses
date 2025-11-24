import { QUERY_KEYS } from '@/shared/config';
import { useQuery } from '@tanstack/react-query';
import { getOrdersMetadata } from '../api/get-orders-metadata.api';

export function useGetOrdersMetadata() {
  return useQuery({
    queryKey: [QUERY_KEYS.ORDERS_METADATA],
    queryFn: getOrdersMetadata,
  });
}
