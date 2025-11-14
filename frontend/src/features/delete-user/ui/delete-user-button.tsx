'use client';

import { DropdownMenuItem } from '@/shared/ui/dropdown-menu';
import { Loader } from 'lucide-react';
import { useDeleteUser } from '../model/use-delete-user';

interface DeleteUserDropdownItemProps {
  userId: string;
}

export const DeleteUserDropdownItem = ({
  userId,
}: DeleteUserDropdownItemProps) => {
  const { mutate, isPending } = useDeleteUser();

  return (
    <DropdownMenuItem
      variant='destructive'
      onClick={() => mutate(userId)}
      disabled={isPending}
    >
      {isPending ? <Loader /> : null}
      Delete
    </DropdownMenuItem>
  );
};
