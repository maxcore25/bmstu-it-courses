import { courseFormats, levels } from '@/shared/config';
import { z } from 'zod';

export const createCourseSchema = z.object({
  name: z.string(),
  authorId: z.uuid(),
  difficulty: z.enum(levels),
  duration: z.string(),
  format: z.enum(courseFormats),
  price: z.number().nonnegative(),
});

export type CreateCourseValues = z.infer<typeof createCourseSchema>;
