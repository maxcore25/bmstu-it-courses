import { QUERY_KEYS } from '@/shared/config';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { updateSchedule } from '../api/update-schedule.api';
import { UpdateScheduleValues } from './update-schedule.schema';

export const useUpdateSchedule = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: [QUERY_KEYS.SCHEDULES],
    mutationFn: ({
      id,
      payload,
    }: {
      id: string;
      payload: UpdateScheduleValues;
    }) => updateSchedule(id, payload),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: [QUERY_KEYS.SCHEDULES],
      });
    },
  });
};
