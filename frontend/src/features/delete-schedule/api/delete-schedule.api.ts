import { axiosInstance } from '@/shared/api';

export const deleteSchedule = async (id: string) => {
  const { data } = await axiosInstance.delete(`/schedules/${id}`);
  return data;
};
