import { BranchesTable } from '@/widgets/branches-table';
import { CoursesTable } from '@/widgets/courses-table';
import { SchedulesTable } from '@/widgets/schedules-table';

export default function AdminPage() {
  return (
    <>
      <BranchesTable />
      <CoursesTable />
      <SchedulesTable />
    </>
  );
}
