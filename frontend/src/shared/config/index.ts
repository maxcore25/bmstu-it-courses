export const BASE_API_URL = 'http://localhost:8000/api';

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
