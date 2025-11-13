'use client';

import { clearAuthTokens } from '@/shared/api';
import { DropdownMenuItem } from '@/shared/ui/dropdown-menu';
import { LogOut } from 'lucide-react';
import { useRouter } from 'next/navigation';
import { useLogout } from '../model/use-logout';

export const LogoutDropdownItem = () => {
  const router = useRouter();
  const { mutate } = useLogout();

  const handleLogout = () => {
    mutate(undefined, {
      onSuccess: () => {
        clearAuthTokens();
        router.push('/login');
      },
    });
  };

  return (
    <DropdownMenuItem onClick={handleLogout}>
      <LogOut />
      Log out
    </DropdownMenuItem>
  );
};
