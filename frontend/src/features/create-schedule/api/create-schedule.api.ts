import { axiosInstance } from '@/shared/api';
import { CreateScheduleValues } from '../model/create-schedule.schema';

export const createSchedule = async (payload: CreateScheduleValues) => {
  const { data } = await axiosInstance.post('/schedules', payload);
  return data;
};
