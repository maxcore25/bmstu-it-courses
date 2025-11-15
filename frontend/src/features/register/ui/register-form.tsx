'use client';

import { cn } from '@/shared/lib/utils';
import { Button } from '@/shared/ui/button';
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/shared/ui/form';
import { Input } from '@/shared/ui/input';
import { PasswordInput } from '@/shared/ui/password-input';
import { Spinner } from '@/shared/ui/spinner';
import { useRegisterForm } from '../model/use-register-form';

export function RegisterForm({
  className,
  ...props
}: React.ComponentProps<'form'>) {
  const { form, onSubmit, isPending } = useRegisterForm();

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className={cn('space-y-2', className)}
        {...props}
      >
        <FormField
          control={form.control}
          name='name'
          render={({ field }) => (
            <FormItem>
              <FormLabel>Name</FormLabel>
              <FormControl>
                <Input autoFocus {...field} className='h-auto py-3' />
              </FormControl>
              <FormMessage className='h-[20px]' />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name='email'
          render={({ field }) => (
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input type='email' {...field} className='h-auto py-3' />
              </FormControl>
              <FormMessage className='h-[20px]' />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name='password'
          render={({ field }) => (
            <FormItem>
              <FormLabel>Password</FormLabel>
              <FormControl>
                <PasswordInput {...field} className='h-auto py-3' />
              </FormControl>
              <FormMessage className='h-[20px]' />
            </FormItem>
          )}
        />
        <Button
          type='submit'
          className='mt-6! h-auto w-full gap-2 py-3'
          disabled={isPending}
        >
          {isPending ? <Spinner /> : null}
          Register
        </Button>
      </form>
    </Form>
  );
}
