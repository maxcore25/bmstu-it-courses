'use client';

import { useGetBranches } from '@/entities/branch';
import { useGetCourses } from '@/entities/course';
import { useIsMobile } from '@/shared/lib/hooks';
import { Button } from '@/shared/ui/button';
import {
  Drawer,
  DrawerClose,
  DrawerContent,
  DrawerDescription,
  DrawerFooter,
  DrawerHeader,
  DrawerTitle,
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
import { DatetimePicker } from '../../../shared/ui/datetime-picker';
import { UpdateScheduleValues } from '../model/update-schedule.schema';
import { useUpdateScheduleForm } from '../model/use-update-schedule-form';

interface UpdateScheduleDrawerProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  scheduleId: string;
  initialData: UpdateScheduleValues;
}

export const UpdateScheduleDrawer = ({
  open,
  onOpenChange,
  scheduleId,
  initialData,
}: UpdateScheduleDrawerProps) => {
  const isMobile = useIsMobile();
  const { form, onSubmit, handleCancel, isPending } = useUpdateScheduleForm(
    scheduleId,
    initialData
  );
  const { data: branches, isLoading: isLoadingBranches } = useGetBranches();
  const { data: courses, isLoading: isLoadingCourses } = useGetCourses();

  const handleFormSubmit = (values: UpdateScheduleValues) => {
    onSubmit(values);
  };

  const handleCancelClick = () => {
    handleCancel();
  };

  return (
    <Drawer
      open={open}
      onOpenChange={onOpenChange}
      direction={isMobile ? 'bottom' : 'right'}
    >
      <DrawerContent>
        <DrawerHeader className='gap-1'>
          <DrawerTitle>Редактировать расписание</DrawerTitle>
          <DrawerDescription>
            Обновите информацию о расписании ниже.
          </DrawerDescription>
        </DrawerHeader>
        <div className='flex flex-col gap-4 overflow-y-auto px-4 text-sm'>
          <Form {...form}>
            <form
              onSubmit={form.handleSubmit(handleFormSubmit)}
              className='space-y-4'
            >
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
                name='capacity'
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Вместимость</FormLabel>
                    <FormControl>
                      <Input
                        type='number'
                        {...field}
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
                name='startAt'
                render={({ field }) => (
                  <FormItem>
                    <FormControl>
                      <DatetimePicker
                        value={field.value}
                        onChange={field.onChange}
                        label='Дата начала'
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
                    <FormControl>
                      <DatetimePicker
                        value={field.value}
                        onChange={field.onChange}
                        label='Дата окончания'
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
                Обновить расписание
              </Button>
            </form>
          </Form>
        </div>
        <DrawerFooter>
          <DrawerClose asChild>
            <Button variant='outline' onClick={handleCancelClick}>
              Отмена
            </Button>
          </DrawerClose>
        </DrawerFooter>
      </DrawerContent>
    </Drawer>
  );
};
