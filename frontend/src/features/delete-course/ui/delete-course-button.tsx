'use client';

import { DropdownMenuItem } from '@/shared/ui/dropdown-menu';
import { Spinner } from '@/shared/ui/spinner';
import { useDeleteCourse } from '../model/use-delete-course';

interface DeleteCourseDropdownItemProps {
  courseId: string;
}

export const DeleteCourseDropdownItem = ({
  courseId,
}: DeleteCourseDropdownItemProps) => {
  const { mutate, isPending } = useDeleteCourse();

  return (
    <DropdownMenuItem
      variant='destructive'
      onClick={() => mutate(courseId)}
      disabled={isPending}
    >
      {isPending ? <Spinner /> : null}
      Delete
    </DropdownMenuItem>
  );
};
