'use client';

import { useGetOrders } from '@/entities/order';
import { BranchesTable } from '@/widgets/branches-table';
import { CoursesTable } from '@/widgets/courses-table';
import { OrdersTable } from '@/widgets/orders-table';
import { SchedulesTable } from '@/widgets/schedules-table';
import { UsersTable } from '@/widgets/users-table';

export default function AdminHomePage() {
  const { data, isLoading } = useGetOrders(['client', 'course', 'branch']);

  return (
    <>
      <BranchesTable />
      <CoursesTable />
      <SchedulesTable />
      <UsersTable />
      <OrdersTable orders={data} isLoading={isLoading} />
    </>
  );
}
