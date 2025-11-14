'use client';

import { DropdownMenuItem } from '@/shared/ui/dropdown-menu';
import { Loader } from 'lucide-react';
import { useDeleteSchedule } from '../model/use-delete-schedule';

interface DeleteScheduleDropdownItemProps {
  scheduleId: string;
}

export const DeleteScheduleDropdownItem = ({
  scheduleId,
}: DeleteScheduleDropdownItemProps) => {
  const { mutate, isPending } = useDeleteSchedule();

  return (
    <DropdownMenuItem
      variant='destructive'
      onClick={() => mutate(scheduleId)}
      disabled={isPending}
    >
      {isPending ? <Loader /> : null}
      Delete
    </DropdownMenuItem>
  );
};
