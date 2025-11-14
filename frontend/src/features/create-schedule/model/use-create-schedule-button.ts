import { useIsMobile } from '@/shared/lib/hooks';
import { zodResolver } from '@hookform/resolvers/zod';
import { useEffect } from 'react';
import { useForm } from 'react-hook-form';
import { toast } from 'sonner';
import {
  createScheduleSchema,
  CreateScheduleValues,
} from './create-schedule.schema';
import { useCreateSchedule } from './use-create-schedule';

export const useCreateScheduleButton = () => {
  const form = useForm<CreateScheduleValues>({
    resolver: zodResolver(createScheduleSchema),
    defaultValues: {
      branchId: '',
      capacity: 0,
      courseId: '',
      endAt: '',
      startAt: '',
    },
  });
  const { error, isSuccess, isError, isPending, mutate } = useCreateSchedule();
  const isMobile = useIsMobile();

  useEffect(() => {
    if (isSuccess) {
      form.reset();
    }
  }, [isSuccess, form]);

  useEffect(() => {
    if (isError) {
      toast.error(error?.message || 'Something went wrong', {
        description: 'Please try again.',
        action: {
          label: 'Close',
          onClick: () => null,
        },
      });

      console.error(error);
    }
  }, [isError, error]);

  function onSubmit(values: CreateScheduleValues) {
    mutate(values);
  }

  function handleCancel() {
    form.reset();
  }

  return { form, onSubmit, handleCancel, ...form, isPending, isMobile };
};
