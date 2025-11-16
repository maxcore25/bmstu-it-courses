'use client';

import { useGetMyOrders } from '@/entities/order';
import { OrdersTable } from '@/widgets/orders-table';

export const HomePage = () => {
  const { data, isLoading } = useGetMyOrders([
    'client',
    'course',
    'branch',
    'schedule',
  ]);

  return (
    <>
      <OrdersTable orders={data} isLoading={isLoading} />
    </>
  );
};
