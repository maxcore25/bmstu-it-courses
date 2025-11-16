import { HomePage } from '@/_pages/home';
import { Metadata } from 'next';

export const metadata: Metadata = {
  title: 'Личный кабинет | CodeCraft',
};

export default function HomeRoute() {
  return <HomePage />;
}
