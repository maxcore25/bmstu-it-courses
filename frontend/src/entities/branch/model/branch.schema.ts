import { z } from 'zod';

export const branchSchema = z.object({
  id: z.uuid(),
  address: z.string(),
  rooms: z.string(),
  createdAt: z.iso.datetime(),
  updatedAt: z.iso.datetime(),
});

export type Branch = z.infer<typeof branchSchema>;

export const branchesSchema = z.array(branchSchema);
