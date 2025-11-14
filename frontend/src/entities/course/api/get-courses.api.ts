import { axiosInstance } from '@/shared/api';
import { coursesSchema } from '../model/course.schema';

export async function getCourses(expand?: string[]) {
  let url = '/courses';
  if (expand && expand.length > 0) {
    // Encode each part for safety
    const expandParam = expand.map(e => encodeURIComponent(e)).join(',');
    url += `?expand=${expandParam}`;
  }
  const { data } = await axiosInstance.get(url);

  try {
    coursesSchema.parse(data);
  } catch (error) {
    console.error(error);
  }

  return coursesSchema.parse(data);
}
