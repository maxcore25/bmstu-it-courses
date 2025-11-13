import { z } from 'zod';

export const RoleSchema = z.object({
  id: z.uuid(),
  name: z.string().min(1),
});

export const userSchema = z.object({
  id: z.uuid(),
  email: z.email(),
  name: z.string().min(1),
  roles: z.array(RoleSchema).nonempty(),
});

export type User = z.infer<typeof userSchema>;
