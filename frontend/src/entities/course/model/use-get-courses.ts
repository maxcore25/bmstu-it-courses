import { QUERY_KEYS } from '@/shared/config';
import { useQuery } from '@tanstack/react-query';
import { getCourses } from '../api/get-courses.api';

export function useGetCourses() {
  return useQuery({
    queryKey: [QUERY_KEYS.COURSES],
    queryFn: getCourses,
  });
}
