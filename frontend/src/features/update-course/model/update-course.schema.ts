import { courseFormats, levels } from '@/shared/config';
import { z } from 'zod';

export const updateCourseSchema = z
  .object({
    name: z.string(),
    authorId: z.uuid(),
    difficulty: z.enum(levels),
    duration: z.string(),
    format: z.enum(courseFormats),
    price: z.number().nonnegative(),
  })
  .partial();

export type UpdateCourseValues = z.infer<typeof updateCourseSchema>;
