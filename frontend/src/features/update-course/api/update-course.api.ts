import { axiosInstance } from '@/shared/api';
import { UpdateCourseValues } from '../model/update-course.schema';

export const updateCourse = async (id: string, payload: UpdateCourseValues) => {
  const { data } = await axiosInstance.patch(`/courses/${id}`, payload);
  return data;
};
