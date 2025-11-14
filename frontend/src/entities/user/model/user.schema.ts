import { levels } from '@/shared/config';
import { z } from 'zod';

export const userSchema = z.object({
  id: z.uuid(),
  firstName: z.string().min(1),
  lastName: z.string().min(1),
  middleName: z.string().optional(),
  email: z.email(),
  phone: z.string().optional(),
  knowledgeLevel: z.enum(levels),
  role: z.string().min(1),
  createdAt: z.iso.datetime({ offset: true }),
  updatedAt: z.iso.datetime({ offset: true }),
});

export const tutorSchema = userSchema.extend({
  portfolio: z.string(),
  rating: z.number().nonnegative(),
  testimonialsCount: z.number().int().nonnegative(),
});

export type User = z.infer<typeof userSchema>;
export type Tutor = z.infer<typeof tutorSchema>;
