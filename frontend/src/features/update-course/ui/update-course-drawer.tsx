'use client';

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
import { Spinner } from '@/shared/ui/spinner';
import { UpdateCourseValues } from '../model/update-course.schema';
import { useUpdateCourseForm } from '../model/use-update-course-form';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/shared/ui/select';
import { useGetTutors } from '@/entities/user';

interface UpdateCourseDrawerProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  courseId: string;
  initialData: UpdateCourseValues;
}

export const UpdateCourseDrawer = ({
  open,
  onOpenChange,
  courseId,
  initialData,
}: UpdateCourseDrawerProps) => {
  const isMobile = useIsMobile();
  const { form, onSubmit, handleCancel, isPending } = useUpdateCourseForm(
    courseId,
    initialData
  );
  const { data: tutors, isLoading: isLoadingTutors } = useGetTutors();

  const handleFormSubmit = (values: UpdateCourseValues) => {
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
          <DrawerTitle>Edit Course</DrawerTitle>
          <DrawerDescription>
            Update the course information below.
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
                name='authorId'
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Автор</FormLabel>
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
                        {isLoadingTutors ? (
                          <SelectItem disabled value=''>
                            Загрузка...
                          </SelectItem>
                        ) : (
                          tutors?.map(tutor => (
                            <SelectItem key={tutor.id} value={tutor.id}>
                              {tutor.lastName} {tutor.firstName} ({tutor.email})
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
                name='difficulty'
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Difficulty</FormLabel>
                    <Select
                      onValueChange={field.onChange}
                      defaultValue={field.value}
                    >
                      <FormControl>
                        <SelectTrigger className='h-auto! w-full py-3'>
                          <SelectValue placeholder='Выберите уровень' />
                        </SelectTrigger>
                      </FormControl>
                      <SelectContent>
                        <SelectItem value='beginner'>Начальный</SelectItem>
                        <SelectItem value='intermediate'>Средний</SelectItem>
                        <SelectItem value='advanced'>Продвинутый</SelectItem>
                      </SelectContent>
                    </Select>
                    <FormMessage className='h-[20px]' />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name='duration'
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Duration</FormLabel>
                    <FormControl>
                      <Input {...field} className='h-auto py-3' />
                    </FormControl>
                    <FormMessage className='h-[20px]' />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name='format'
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Format</FormLabel>
                    <Select
                      onValueChange={field.onChange}
                      defaultValue={field.value}
                    >
                      <FormControl>
                        <SelectTrigger className='h-auto! w-full py-3'>
                          <SelectValue placeholder='Выберите уровень' />
                        </SelectTrigger>
                      </FormControl>
                      <SelectContent>
                        <SelectItem value='group'>Группа</SelectItem>
                        <SelectItem value='individual'>
                          Индивидуально
                        </SelectItem>
                        <SelectItem value='intensive'>Интенсив</SelectItem>
                      </SelectContent>
                    </Select>
                    <FormMessage className='h-[20px]' />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name='price'
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Price</FormLabel>
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
              <Button
                type='submit'
                className='mt-6! h-auto w-full gap-2 py-3'
                disabled={isPending}
              >
                {isPending ? <Spinner /> : null}
                Update Course
              </Button>
            </form>
          </Form>
        </div>
        <DrawerFooter>
          <DrawerClose asChild>
            <Button variant='outline' onClick={handleCancelClick}>
              Cancel
            </Button>
          </DrawerClose>
        </DrawerFooter>
      </DrawerContent>
    </Drawer>
  );
};
