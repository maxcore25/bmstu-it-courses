import { axiosInstance } from '@/shared/api';
import { coursesSchema } from '../model/course.schema';

export async function getCourses(expand?: string[]) {
  let url = '/courses';
  if (expand && expand.length > 0) {
    const expandParam = expand.map(e => encodeURIComponent(e)).join(',');
    url += `?expand=${expandParam}`;
  }

  const { data } = await axiosInstance.get(url);

  return coursesSchema.parse(data);
}
