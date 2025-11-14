import { QUERY_KEYS } from '@/shared/config';
import { useQuery } from '@tanstack/react-query';
import { getSchedules } from '../api/get-schedules.api';

export function useGetSchedules() {
  return useQuery({
    queryKey: [QUERY_KEYS.SCHEDULES],
    queryFn: getSchedules,
  });
}
