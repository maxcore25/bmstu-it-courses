import { branchSchema } from '@/entities/branch';
import { courseSchema } from '@/entities/course';
import { z } from 'zod';

export const scheduleSchema = z.object({
  id: z.uuid(),
  branchId: z.uuid(),
  courseId: z.uuid(),
  capacity: z.number().nonnegative(),
  reserved: z.number().nonnegative(),
  startAt: z.iso.datetime({ offset: true }),
  endAt: z.iso.datetime({ offset: true }),
  createdAt: z.iso.datetime({ offset: true }),
  updatedAt: z.iso.datetime({ offset: true }),
  course: courseSchema.optional(),
  branch: branchSchema.optional(),
});

export type Schedule = z.infer<typeof scheduleSchema>;

export const schedulesSchema = z.array(scheduleSchema);
