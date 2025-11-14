import { z } from 'zod';

export const updateScheduleSchema = z
  .object({
    branchId: z.uuid(),
    capacity: z.number().nonnegative(),
    courseId: z.uuid(),
    endAt: z.iso.datetime(),
    startAt: z.iso.datetime(),
  })
  .partial();

export type UpdateScheduleValues = z.infer<typeof updateScheduleSchema>;
