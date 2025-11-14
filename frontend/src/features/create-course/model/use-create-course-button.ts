import { useIsMobile } from '@/shared/lib/hooks';
import { zodResolver } from '@hookform/resolvers/zod';
import { useEffect } from 'react';
import { useForm } from 'react-hook-form';
import { toast } from 'sonner';
import { createCourseSchema, CreateCourseValues } from './create-course.schema';
import { useCreateCourse } from './use-create-course';

export const useCreateCourseButton = () => {
  const form = useForm<CreateCourseValues>({
    resolver: zodResolver(createCourseSchema),
    defaultValues: {
      name: '',
      authorId: '',
      difficulty: 'beginner',
      duration: '',
      format: 'group',
      price: 0,
    },
  });
  const { error, isSuccess, isError, isPending, mutate } = useCreateCourse();
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

  function onSubmit(values: CreateCourseValues) {
    mutate(values);
  }

  function handleCancel() {
    form.reset();
  }

  return { form, onSubmit, handleCancel, ...form, isPending, isMobile };
};
