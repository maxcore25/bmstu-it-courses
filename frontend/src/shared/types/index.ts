import { tokensSchema } from '@/shared/lib/validations';
import { z } from 'zod';

export type Tokens = z.infer<typeof tokensSchema>;
