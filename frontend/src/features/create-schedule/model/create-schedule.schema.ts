import { z } from 'zod';

export const createScheduleSchema = z.object({
  branchId: z.uuid(),
  capacity: z.number().nonnegative(),
  courseId: z.uuid(),
  endAt: z.iso.datetime(),
  startAt: z.iso.datetime(),
});

export type CreateScheduleValues = z.infer<typeof createScheduleSchema>;
