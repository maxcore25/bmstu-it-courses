import { QUERY_KEYS } from '@/shared/config';
import { useQuery } from '@tanstack/react-query';
import { getCourses } from '../api/get-courses.api';

export function useGetCourses(expand?: string[]) {
  return useQuery({
    queryKey: [QUERY_KEYS.COURSES, { expand }],
    queryFn: () => getCourses(expand),
  });
}
