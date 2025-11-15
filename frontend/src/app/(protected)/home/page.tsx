'use client';

import { OrdersTable } from '@/widgets/orders-table';
import { useGetMyOrders } from '@/entities/order';

export default function HomePage() {
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
}
