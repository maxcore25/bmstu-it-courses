'use client';

import { Button } from '@/shared/ui/button';
import {
  Drawer,
  DrawerClose,
  DrawerContent,
  DrawerDescription,
  DrawerFooter,
  DrawerHeader,
  DrawerTitle,
  DrawerTrigger,
} from '@/shared/ui/drawer';
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/shared/ui/form';
import { Input } from '@/shared/ui/input';
import { Spinner } from '@/shared/ui/spinner';
import { useCreateUserButton } from '../model/use-create-user-button';

export const CreateUserButton = () => {
  const { form, onSubmit, handleCancel, isPending, isMobile } =
    useCreateUserButton();

  return (
    <Drawer direction={isMobile ? 'bottom' : 'right'} onClose={handleCancel}>
      <DrawerTrigger asChild>
        <Button>Create User</Button>
      </DrawerTrigger>
      <DrawerContent>
        <DrawerHeader className='gap-1'>
          <DrawerTitle>Create User</DrawerTitle>
          <DrawerDescription>
            Enter the user details to create a new user.
          </DrawerDescription>
        </DrawerHeader>
        <div className='flex flex-col gap-4 overflow-y-auto px-4 text-sm'>
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className='space-y-2'>
              <FormField
                control={form.control}
                name='firstName'
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>First Name</FormLabel>
                    <FormControl>
                      <Input autoFocus {...field} className='h-auto py-3' />
                    </FormControl>
                    <FormMessage className='h-[20px]' />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name='lastName'
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Last Name</FormLabel>
                    <FormControl>
                      <Input {...field} className='h-auto py-3' />
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
              {/* add more fields as needed */}
              <Button
                type='submit'
                className='mt-6! h-auto w-full gap-2 py-3'
                disabled={isPending}
              >
                {isPending ? <Spinner /> : null}
                Create
              </Button>
            </form>
          </Form>
        </div>
        <DrawerFooter>
          <DrawerClose asChild>
            <Button variant='outline'>Cancel</Button>
          </DrawerClose>
        </DrawerFooter>
      </DrawerContent>
    </Drawer>
  );
};
