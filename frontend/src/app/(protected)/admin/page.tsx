import { BranchesTable } from '@/widgets/branches-table';
import { CoursesTable } from '@/widgets/courses-table';
import { OrdersTable } from '@/widgets/orders-table';
import { SchedulesTable } from '@/widgets/schedules-table';
import { UsersTable } from '@/widgets/users-table';

export default function AdminPage() {
  return (
    <>
      <BranchesTable />
      <CoursesTable />
      <SchedulesTable />
      <UsersTable />
      <OrdersTable />
    </>
  );
}
