import { axiosInstance } from '@/shared/api';
import { CreateCourseValues } from '../model/create-course.schema';

export const createCourse = async (payload: CreateCourseValues) => {
  const { data } = await axiosInstance.post('/courses', payload);
  return data;
};
