'use client';

import { DropdownMenuItem } from '@/shared/ui/dropdown-menu';
import { Spinner } from '@/shared/ui/spinner';
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
      {isPending ? <Spinner /> : null}
      Delete
    </DropdownMenuItem>
  );
};
