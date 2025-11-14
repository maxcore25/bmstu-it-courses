import { z } from 'zod';

export const updateBranchSchema = z
  .object({
    address: z.string().min(1, 'Name is required'),
    rooms: z.number().positive(),
  })
  .partial();

export type UpdateBranchValues = z.infer<typeof updateBranchSchema>;
