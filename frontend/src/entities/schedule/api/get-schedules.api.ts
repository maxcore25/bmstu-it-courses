import { axiosInstance } from '@/shared/api';
import { schedulesSchema } from '../model/schedule.schema';

export async function getSchedules(expand?: string[]) {
  let url = '/schedules';
  if (expand && expand.length > 0) {
    const expandParam = expand.map(e => encodeURIComponent(e)).join(',');
    url += `?expand=${expandParam}`;
  }
  const { data } = await axiosInstance.get(url);

  try {
    schedulesSchema.parse(data);
  } catch (error) {
    console.error(error);
  }

  return schedulesSchema.parse(data);
}
