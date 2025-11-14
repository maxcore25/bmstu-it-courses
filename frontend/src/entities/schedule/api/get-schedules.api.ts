import { axiosInstance } from '@/shared/api';
import { schedulesSchema } from '../model/schedule.schema';

export async function getSchedules() {
  const { data } = await axiosInstance.get('/schedules');
  return schedulesSchema.parse(data);
}
