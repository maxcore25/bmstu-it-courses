export const BASE_URL =
  process.env.NEXT_PUBLIC_BASE_URL || 'http://localhost:8080';
export const BASE_API_URL = `${BASE_URL}/api`;

export const QUERY_KEYS = {
  CATEGORY: 'CATEGORY',
  CATEGORIES: 'CATEGORIES',
  PRODUCT: 'PRODUCT',
  PRODUCTS: 'PRODUCTS',
  TOPPING: 'TOPPING',
  TOPPINGS: 'TOPPINGS',
  USER_ME: 'USER_ME',
  ORDER: 'ORDER',
  ORDERS: 'ORDERS',
} as const;

export const LOCAL_STORAGE_KEYS = {
  ACCESS_TOKEN: 'accessToken',
} as const;
