'use client';

import { useGetBranches } from '@/entities/branch';
import { useGetCourses } from '@/entities/course';
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
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/shared/ui/select';
import { Spinner } from '@/shared/ui/spinner';
import { useCreateScheduleButton } from '../model/use-create-schedule-button';

export const CreateScheduleButton = () => {
  const { form, onSubmit, handleCancel, isPending, isMobile } =
    useCreateScheduleButton();

  const { data: branches, isLoading: isLoadingBranches } = useGetBranches();
  const { data: courses, isLoading: isLoadingCourses } = useGetCourses();

  return (
    <Drawer direction={isMobile ? 'bottom' : 'right'} onClose={handleCancel}>
      <DrawerTrigger asChild>
        <Button>Create Schedule</Button>
      </DrawerTrigger>
      <DrawerContent>
        <DrawerHeader className='gap-1'>
          <DrawerTitle>Create Schedule</DrawerTitle>
          <DrawerDescription>
            Enter the required information to create the schedule.
          </DrawerDescription>
        </DrawerHeader>
        <div className='flex flex-col gap-4 overflow-y-auto px-4 text-sm'>
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className='space-y-2'>
              <FormField
                control={form.control}
                name='branchId'
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Филиал</FormLabel>
                    <Select
                      onValueChange={field.onChange}
                      defaultValue={field.value}
                    >
                      <FormControl>
                        <SelectTrigger className='h-auto! w-full py-3'>
                          <SelectValue placeholder='Выберите филиал' />
                        </SelectTrigger>
                      </FormControl>
                      <SelectContent>
                        {isLoadingBranches ? (
                          <SelectItem disabled value=''>
                            Загрузка...
                          </SelectItem>
                        ) : (
                          branches?.map(branch => (
                            <SelectItem key={branch.id} value={branch.id}>
                              {branch.address}
                            </SelectItem>
                          ))
                        )}
                      </SelectContent>
                    </Select>
                    <FormMessage className='h-[20px]' />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name='courseId'
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Курс</FormLabel>
                    <Select
                      onValueChange={field.onChange}
                      defaultValue={field.value}
                    >
                      <FormControl>
                        <SelectTrigger className='h-auto! w-full py-3'>
                          <SelectValue placeholder='Выберите курс' />
                        </SelectTrigger>
                      </FormControl>
                      <SelectContent>
                        {isLoadingCourses ? (
                          <SelectItem disabled value=''>
                            Загрузка...
                          </SelectItem>
                        ) : (
                          courses?.map(course => (
                            <SelectItem key={course.id} value={course.id}>
                              {course.name}
                            </SelectItem>
                          ))
                        )}
                      </SelectContent>
                    </Select>
                    <FormMessage className='h-[20px]' />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name='capacity'
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Capacity</FormLabel>
                    <FormControl>
                      <Input
                        type='number'
                        step='1'
                        min='0'
                        value={
                          typeof field.value === 'number'
                            ? String(field.value)
                            : field.value
                        }
                        onChange={e => field.onChange(Number(e.target.value))}
                        className='h-auto py-3'
                      />
                    </FormControl>
                    <FormMessage className='h-[20px]' />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name='startAt'
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Start At</FormLabel>
                    <FormControl>
                      <Input
                        type='datetime-local'
                        {...field}
                        className='h-auto py-3'
                      />
                    </FormControl>
                    <FormMessage className='h-[20px]' />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name='endAt'
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>End At</FormLabel>
                    <FormControl>
                      <Input
                        type='datetime-local'
                        {...field}
                        className='h-auto py-3'
                      />
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
