import { QUERY_KEYS } from '@/shared/config';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { deleteCourse } from '../api/delete-course.api';

export const useDeleteCourse = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: [QUERY_KEYS.COURSES],
    mutationFn: deleteCourse,
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: [QUERY_KEYS.COURSES],
      });
    },
  });
};
