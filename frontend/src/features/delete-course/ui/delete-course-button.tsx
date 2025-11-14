'use client';

import { DropdownMenuItem } from '@/shared/ui/dropdown-menu';
import { Loader } from 'lucide-react';
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
      {isPending ? <Loader /> : null}
      Delete
    </DropdownMenuItem>
  );
};
