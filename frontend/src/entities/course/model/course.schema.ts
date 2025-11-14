import { levels } from '@/shared/lib/levels';
import { z } from 'zod';

export const courseSchema = z.object({
  id: z.uuid(),
  authorId: z.uuid(),
  name: z.string(),
  price: z.number(),
  difficulty: z.enum(levels),
  duration: z.string(),
  format: z.string(),
  createdAt: z.string(),
  updatedAt: z.string(),
});

export type Course = z.infer<typeof courseSchema>;

export const coursesSchema = z.array(courseSchema);
