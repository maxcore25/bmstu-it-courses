import { axiosInstance } from '@/shared/api';
import { coursesSchema } from '../model/course.schema';

export async function getCourses() {
  const { data } = await axiosInstance.get('/courses');
  return coursesSchema.parse(data);
}
