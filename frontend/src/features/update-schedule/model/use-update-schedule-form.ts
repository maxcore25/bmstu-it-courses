import { zodResolver } from '@hookform/resolvers/zod';
import { useEffect } from 'react';
import { useForm } from 'react-hook-form';
import { toast } from 'sonner';
import {
  updateScheduleSchema,
  UpdateScheduleValues,
} from './update-schedule.schema';
import { useUpdateSchedule } from './use-update-schedule';

export const useUpdateScheduleForm = (
  scheduleId: string,
  initialData?: UpdateScheduleValues
) => {
  const form = useForm<UpdateScheduleValues>({
    resolver: zodResolver(updateScheduleSchema),
    defaultValues: {
      branchId: initialData?.branchId || '',
      capacity: initialData?.capacity || 0,
      courseId: initialData?.courseId || '',
      endAt: initialData?.endAt || '',
      startAt: initialData?.startAt || '',
    },
  });

  useEffect(() => {
    if (initialData) {
      form.reset({
        branchId: initialData.branchId || '',
        capacity: initialData.capacity || 0,
        courseId: initialData.courseId || '',
        endAt: initialData.endAt || '',
        startAt: initialData.startAt || '',
      });
    }
  }, [initialData, form]);

  const { error, isSuccess, isError, isPending, mutate } = useUpdateSchedule();

  useEffect(() => {
    if (isSuccess) {
      toast.success('Schedule updated successfully');
      form.reset();
    }
  }, [isSuccess, form]);

  useEffect(() => {
    if (isError) {
      toast.error(error?.message || 'Failed to update schedule', {
        description: 'Please try again.',
        action: { label: 'Close', onClick: () => null },
      });
      console.error(error);
    }
  }, [isError, error]);

  function onSubmit(values: UpdateScheduleValues) {
    mutate({ id: scheduleId, payload: values });
  }

  function handleCancel() {
    form.reset();
  }

  return {
    form,
    onSubmit,
    handleCancel,
    ...form,
    isPending,
  };
};
