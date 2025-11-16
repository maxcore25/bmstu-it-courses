import { AdminHomePage } from '@/_pages/admin';
import { Metadata } from 'next';

export const metadata: Metadata = {
  title: 'Админ-панель | CodeCraft',
};

export default function AdminHomeRoute() {
  return <AdminHomePage />;
}
