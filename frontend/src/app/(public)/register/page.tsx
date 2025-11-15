import { RegisterForm } from '@/features/register';
import { Metadata } from 'next';

export const metadata: Metadata = {
  title: 'Register | CodeCraft',
};

export default function RegisterPage() {
  return (
    <div className='flex min-h-svh w-full items-center justify-center p-6 md:p-10'>
      <div className='w-full max-w-sm'>
        <RegisterForm />
      </div>
    </div>
  );
}
