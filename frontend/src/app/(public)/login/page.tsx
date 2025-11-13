import { LoginForm } from '@/features/login';
import { Metadata } from 'next';

export const metadata: Metadata = {
  title: 'Login | CodeCraft',
};

export default function LoginPage() {
  return (
    <div className='flex min-h-svh w-full items-center justify-center p-6 md:p-10'>
      <div className='w-full max-w-sm'>
        <LoginForm />
      </div>
    </div>
  );
}
