import { QUERY_KEYS } from '@/shared/config';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { updateCourse } from '../api/update-course.api';
import { UpdateCourseValues } from './update-course.schema';

export const useUpdateCourse = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: [QUERY_KEYS.COURSES],
    mutationFn: ({
      id,
      payload,
    }: {
      id: string;
      payload: UpdateCourseValues;
    }) => updateCourse(id, payload),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: [QUERY_KEYS.COURSES],
      });
    },
  });
};
