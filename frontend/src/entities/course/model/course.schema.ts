import { userSchema } from '@/entities/user';
import { courseFormats, levels } from '@/shared/config';
import { z } from 'zod';

export const courseSchema = z.object({
  id: z.uuid(),
  authorId: z.uuid(),
  name: z.string(),
  price: z.number(),
  difficulty: z.enum(levels),
  duration: z.string(),
  format: z.enum(courseFormats),
  createdAt: z.string(),
  updatedAt: z.string(),
  author: userSchema.optional(),
});

export type Course = z.infer<typeof courseSchema>;

export const coursesSchema = z.array(courseSchema);
