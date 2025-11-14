import { axiosInstance } from '@/shared/api';
import { schedulesSchema } from '../model/schedule.schema';

export async function getSchedules(expand?: string[]) {
  let url = '/schedules';
  if (expand && expand.length > 0) {
    const expandParam = expand.map(e => encodeURIComponent(e)).join(',');
    url += `?expand=${expandParam}`;
  }
  const { data } = await axiosInstance.get(url);

  return schedulesSchema.parse(data);
}
