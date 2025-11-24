'use client';

import { useGetBranches } from '@/entities/branch';
import { useGetCourses } from '@/entities/course';
import { useGetSchedules } from '@/entities/schedule';
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
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/shared/ui/select';
import { Spinner } from '@/shared/ui/spinner';
import { useCreateOrderButton } from '../model/use-create-order-button';

export const CreateOrderButton = () => {
  const { form, onSubmit, handleCancel, isPending, isMobile } =
    useCreateOrderButton();

  const { data: branches, isLoading: isLoadingBranches } = useGetBranches();
  const { data: courses, isLoading: isLoadingCourses } = useGetCourses();
  const { data: schedules, isLoading: isLoadingSchedules } = useGetSchedules();

  // Find selected course object for price display
  const selectedCourseId = form.watch('courseId');
  const selectedCourse =
    courses?.find(course => course.id === selectedCourseId) ?? null;

  return (
    <Drawer direction={isMobile ? 'bottom' : 'right'} onClose={handleCancel}>
      <DrawerTrigger asChild>
        <Button>Создать заказ</Button>
      </DrawerTrigger>
      <DrawerContent>
        <DrawerHeader className='gap-1'>
          <DrawerTitle>Создать заказ</DrawerTitle>
          <DrawerDescription>
            Укажите филиал, клиента, курс и расписание для создания заказа.
          </DrawerDescription>
        </DrawerHeader>
        <div className='flex flex-col gap-4 overflow-y-auto px-4 text-sm'>
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className='space-y-4'>
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
                name='scheduleId'
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Расписание</FormLabel>
                    <Select
                      onValueChange={field.onChange}
                      defaultValue={field.value}
                    >
                      <FormControl>
                        <SelectTrigger className='h-auto! w-full py-3'>
                          <SelectValue placeholder='Выберите расписание' />
                        </SelectTrigger>
                      </FormControl>
                      <SelectContent>
                        {isLoadingSchedules ? (
                          <SelectItem disabled value=''>
                            Загрузка...
                          </SelectItem>
                        ) : (
                          schedules?.map(schedule => (
                            <SelectItem key={schedule.id} value={schedule.id}>
                              {new Date(schedule.startAt).toLocaleString()} -{' '}
                              {new Date(schedule.endAt).toLocaleString()}
                            </SelectItem>
                          ))
                        )}
                      </SelectContent>
                    </Select>
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
                Создать
                {selectedCourse ? (
                  <div className='text-muted-foreground'>
                    {`(${selectedCourse.price?.toLocaleString?.('ru-RU') ?? selectedCourse.price} ₽)`}
                  </div>
                ) : null}
              </Button>
            </form>
          </Form>
        </div>
        <DrawerFooter>
          <DrawerClose asChild>
            <Button variant='outline'>Отмена</Button>
          </DrawerClose>
        </DrawerFooter>
      </DrawerContent>
    </Drawer>
  );
};
