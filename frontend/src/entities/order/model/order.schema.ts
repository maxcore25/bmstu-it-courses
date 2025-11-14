import { branchSchema } from '@/entities/branch';
import { courseSchema } from '@/entities/course';
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
});

export type Order = z.infer<typeof orderSchema>;

export const ordersSchema = z.array(orderSchema);
