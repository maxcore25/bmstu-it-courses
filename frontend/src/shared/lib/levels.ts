export const levels = ['beginner', 'intermediate', 'advanced'] as const;
export type Level = (typeof levels)[number];
