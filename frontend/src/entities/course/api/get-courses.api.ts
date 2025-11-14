import { axiosInstance } from '@/shared/api';
import { coursesSchema } from '../model/course.schema';

export async function getCourses(expand?: string[]) {
  const { data } = await axiosInstance.get('/courses', {
    params: { expand },
  });
  return coursesSchema.parse(data);
}
