import { z } from 'zod';

export const branchSchema = z.object({
  id: z.uuid(),
  address: z.string(),
  rooms: z.string(),
  createdAt: z.string(),
  updatedAt: z.string(),
});

export type Branch = z.infer<typeof branchSchema>;

export const branchesSchema = z.array(branchSchema);
