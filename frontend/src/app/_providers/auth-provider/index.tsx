'use client';

import { useGetMe } from '@/entities/user/model/use-get-me';
import { useAuthStore } from '@/features/auth';
import { useRouter } from 'next/navigation';
import { ReactNode, useEffect } from 'react';

export function AuthProvider({ children }: { children: ReactNode }) {
  const { data: user, isError } = useGetMe();
  const { setUser, reset } = useAuthStore();
  const router = useRouter();

  useEffect(() => {
    if (user) setUser(user);
  }, [user, setUser]);

  useEffect(() => {
    if (isError) {
      reset();
      router.push('/login');
    }
  }, [isError, reset, router]);

  return <>{children}</>;
}
