import { axiosInstance } from '@/shared/api';
import { UpdateScheduleValues } from '../model/update-schedule.schema';

export const updateSchedule = async (
  id: string,
  payload: UpdateScheduleValues
) => {
  const { data } = await axiosInstance.patch(`/schedules/${id}`, payload);
  return data;
};
