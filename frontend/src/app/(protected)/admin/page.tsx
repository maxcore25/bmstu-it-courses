import { BranchesTable } from '@/widgets/branches-table';
import { CoursesTable } from '@/widgets/courses-table';

export default function AdminPage() {
  return (
    <>
      <BranchesTable />
      <CoursesTable />
    </>
  );
}
