import { axiosInstance } from '@/shared/api';

export const deleteCourse = async (id: string) => {
  const { data } = await axiosInstance.delete(`/courses/${id}`);
  return data;
};
