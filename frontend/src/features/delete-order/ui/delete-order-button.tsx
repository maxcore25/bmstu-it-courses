'use client';

import { DropdownMenuItem } from '@/shared/ui/dropdown-menu';
import { Spinner } from '@/shared/ui/spinner';
import { useDeleteOrder } from '../model/use-delete-order';

interface DeleteOrderDropdownItemProps {
  orderId: string;
}

export const DeleteOrderDropdownItem = ({
  orderId,
}: DeleteOrderDropdownItemProps) => {
  const { mutate, isPending } = useDeleteOrder();

  return (
    <DropdownMenuItem
      variant='destructive'
      onClick={() => mutate(orderId)}
      disabled={isPending}
    >
      {isPending ? <Spinner /> : null}
      Удалить
    </DropdownMenuItem>
  );
};
