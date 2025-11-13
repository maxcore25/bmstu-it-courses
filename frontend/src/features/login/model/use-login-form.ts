import { LOCAL_STORAGE_KEYS } from '@/shared/config';
import { zodResolver } from '@hookform/resolvers/zod';
import { useRouter } from 'next/navigation';
import { useEffect } from 'react';
import { useForm } from 'react-hook-form';
import { toast } from 'sonner';
import { loginFormSchema, LoginFormValues } from './login.schema';
import { useLogin } from './use-login';

export const useLoginForm = () => {
  const router = useRouter();
  const form = useForm<LoginFormValues>({
    resolver: zodResolver(loginFormSchema),
    defaultValues: {
      email: '',
      password: '',
    },
  });
  const { data, error, isSuccess, isError, isPending, mutate } = useLogin();

  useEffect(() => {
    if (isSuccess) {
      localStorage.setItem(LOCAL_STORAGE_KEYS.ACCESS_TOKEN, data.accessToken);
      form.reset();
      router.push('/');
    }
  }, [isSuccess]);

  useEffect(() => {
    if (isError) {
      toast.error(error.message || 'Something went wrong', {
        description: 'Please try again.',
        action: {
          label: 'Close',
          onClick: () => null,
        },
      });

      console.error(error);
    }
  }, [isError]);

  function onSubmit(values: LoginFormValues) {
    mutate(values);
  }

  return { form, onSubmit, ...form, isPending };
};
