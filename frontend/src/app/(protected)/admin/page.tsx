import { BranchesTable } from '@/widgets/branches-table';
import { CoursesTable } from '@/widgets/courses-table';
import { OrdersTable } from '@/widgets/orders-table';
import { SchedulesTable } from '@/widgets/schedules-table';

export default function AdminPage() {
  return (
    <>
      <BranchesTable />
      <CoursesTable />
      <SchedulesTable />
      <OrdersTable />
    </>
  );
}
