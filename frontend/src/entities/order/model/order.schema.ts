import { z } from 'zod';

export const orderSchema = z.object({
  id: z.uuid(),
  branchId: z.uuid(),
  clientId: z.uuid(),
  courseId: z.uuid(),
  scheduleId: z.uuid(),
  price: z.number().nonnegative(),
  createdAt: z.iso.datetime(),
  updatedAt: z.iso.datetime(),
});

export type Order = z.infer<typeof orderSchema>;

export const ordersSchema = z.array(orderSchema);
