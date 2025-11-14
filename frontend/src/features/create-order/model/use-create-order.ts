import { QUERY_KEYS } from '@/shared/config';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { createOrder } from '../api/create-order.api';

export const useCreateOrder = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: [QUERY_KEYS.ORDERS],
    mutationFn: createOrder,
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: [QUERY_KEYS.ORDERS],
      });
    },
  });
};
