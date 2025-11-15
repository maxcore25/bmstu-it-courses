import { z } from 'zod';

export const createOrderSchema = z.object({
  branchId: z.uuid(),
  courseId: z.uuid(),
  scheduleId: z.uuid(),
});

export type CreateOrderValues = z.infer<typeof createOrderSchema>;
