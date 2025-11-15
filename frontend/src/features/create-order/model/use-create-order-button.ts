import { useIsMobile } from '@/shared/lib/hooks';
import { zodResolver } from '@hookform/resolvers/zod';
import { useEffect } from 'react';
import { useForm } from 'react-hook-form';
import { toast } from 'sonner';
import { createOrderSchema, CreateOrderValues } from './create-order.schema';
import { useCreateOrder } from './use-create-order';

export const useCreateOrderButton = () => {
  const form = useForm<CreateOrderValues>({
    resolver: zodResolver(createOrderSchema),
    defaultValues: {
      branchId: '',
      courseId: '',
      scheduleId: '',
    },
  });
  const { error, isSuccess, isError, isPending, mutate } = useCreateOrder();
  const isMobile = useIsMobile();

  useEffect(() => {
    if (isSuccess) {
      form.reset();
    }
  }, [isSuccess, form]);

  useEffect(() => {
    if (isError) {
      toast.error(error.message || 'Something went wrong', {
        description: 'Please try again.',
        action: {
          label: 'Close',
          onClick: () => null,
        },
      });

      console.error(error);
    }
  }, [isError, error]);

  function onSubmit(values: CreateOrderValues) {
    mutate(values);
  }

  function handleCancel() {
    form.reset();
  }

  return { form, onSubmit, handleCancel, ...form, isPending, isMobile };
};
