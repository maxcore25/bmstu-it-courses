import { BASE_API_URL } from '@/shared/config';
import { coursesSchema } from '../model/course.schema';

export async function getCoursesSSR(expand?: string[]) {
  let url = `${BASE_API_URL}/courses`;
  if (expand && expand.length > 0) {
    const expandParam = expand.map(e => encodeURIComponent(e)).join(',');
    url += `?expand=${expandParam}`;
  }

  const res = await fetch(url, {
    next: { revalidate: 60 },
  });

  if (!res.ok) throw new Error('Failed to fetch courses');

  const data = await res.json();
  return coursesSchema.parse(data);
}
