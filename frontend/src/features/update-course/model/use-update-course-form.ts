import { zodResolver } from '@hookform/resolvers/zod';
import { useEffect } from 'react';
import { useForm } from 'react-hook-form';
import { toast } from 'sonner';
import { updateCourseSchema, UpdateCourseValues } from './update-course.schema';
import { useUpdateCourse } from './use-update-course';

export const useUpdateCourseForm = (
  courseId: string,
  initialData?: UpdateCourseValues
) => {
  const form = useForm<UpdateCourseValues>({
    resolver: zodResolver(updateCourseSchema),
    defaultValues: {
      name: initialData?.name || '',
      authorId: initialData?.authorId || '',
      difficulty: initialData?.difficulty || undefined,
      duration: initialData?.duration || '',
      format: initialData?.format || undefined,
      price: initialData?.price || 0,
    },
  });

  useEffect(() => {
    if (initialData) {
      form.reset({
        name: initialData.name || '',
        authorId: initialData.authorId || '',
        difficulty: initialData.difficulty || undefined,
        duration: initialData.duration || '',
        format: initialData.format || undefined,
        price: initialData.price || 0,
      });
    }
  }, [initialData, form]);

  const { error, isSuccess, isError, isPending, mutate } = useUpdateCourse();

  useEffect(() => {
    if (isSuccess) {
      toast.success('Course updated successfully');
      form.reset();
    }
  }, [isSuccess, form]);

  useEffect(() => {
    if (isError) {
      toast.error(error?.message || 'Failed to update course', {
        description: 'Please try again.',
        action: { label: 'Close', onClick: () => null },
      });
      console.error(error);
    }
  }, [isError, error]);

  function onSubmit(values: UpdateCourseValues) {
    mutate({ id: courseId, payload: values });
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
