import { z } from 'zod';

export const scheduleSchema = z.object({
  id: z.uuid(),
  branchId: z.uuid(),
  courseId: z.uuid(),
  capacity: z.number().nonnegative(),
  reserved: z.number().nonnegative(),
  startAt: z.iso.datetime(),
  endAt: z.iso.datetime(),
  createdAt: z.iso.datetime(),
  updatedAt: z.iso.datetime(),
});

export type Schedule = z.infer<typeof scheduleSchema>;

export const schedulesSchema = z.array(scheduleSchema);
