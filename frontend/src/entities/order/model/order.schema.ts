import { branchSchema } from '@/entities/branch';
import { courseSchema } from '@/entities/course';
import { scheduleSchema } from '@/entities/schedule';
import { userSchema } from '@/entities/user';
import { z } from 'zod';

export const orderSchema = z.object({
  id: z.uuid(),
  branchId: z.uuid(),
  clientId: z.uuid(),
  courseId: z.uuid(),
  scheduleId: z.uuid(),
  price: z.number().nonnegative(),
  createdAt: z.iso.datetime({ offset: true }),
  updatedAt: z.iso.datetime({ offset: true }),
  course: courseSchema.optional(),
  branch: branchSchema.optional(),
  client: userSchema.optional(),
  schedule: scheduleSchema.optional(),
});

export type Order = z.infer<typeof orderSchema>;

export const ordersSchema = z.array(orderSchema);

export const ordersMetadataSchema = z.object({
  count: z.number().int().nonnegative(),
  totalSum: z.number().nonnegative(),
});

export type OrdersMetadata = z.infer<typeof ordersMetadataSchema>;
