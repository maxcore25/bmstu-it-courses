export const BASE_URL =
  process.env.NEXT_PUBLIC_BASE_URL || 'http://localhost:8080';
export const BASE_API_URL = `${BASE_URL}/api`;

export const QUERY_KEYS = {
  USER: 'USER',
  USERS: 'USERS',
  USER_ME: 'USER_ME',
  ORDER: 'ORDER',
  ORDERS: 'ORDERS',
  BRANCH: 'BRANCH',
  BRANCHES: 'BRANCHES',
  COURSE: 'COURSE',
  COURSES: 'COURSES',
  SCHEDULE: 'SCHEDULE',
  SCHEDULES: 'SCHEDULES',
} as const;

export const LOCAL_STORAGE_KEYS = {
  ACCESS_TOKEN: 'accessToken',
} as const;
