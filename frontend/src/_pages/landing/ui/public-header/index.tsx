import { Button } from '@/shared/ui/button';
import { GraduationCap } from 'lucide-react';
import Link from 'next/link';

export const PublicHeader = () => {
  return (
    <header className='bg-background sticky top-0 z-50 flex h-16 shrink-0 items-center gap-2 border-b px-4 transition-[width,height] ease-linear group-has-data-[collapsible=icon]/sidebar-wrapper:h-12'>
      <div className='container mx-auto flex w-full items-center justify-between'>
        <div className='flex items-center gap-2'>
          <div className='bg-sidebar-primary text-sidebar-primary-foreground flex aspect-square size-8 items-center justify-center rounded-lg'>
            <GraduationCap className='size-5' />
          </div>
          <div className='grid flex-1 text-left text-sm leading-tight'>
            <span className='truncate text-xl font-medium'>CodeCraft</span>
          </div>
        </div>
        <div className='flex items-center gap-6'>
          <Button asChild variant='link'>
            <a href='#courses'>Courses</a>
          </Button>
          <Button asChild variant='link'>
            <a href='#tutors'>Tutors</a>
          </Button>
          <Button asChild variant='link'>
            <a href='#branches'>Branches</a>
          </Button>
        </div>
        <Button asChild>
          <Link href='/register'>Register</Link>
        </Button>
      </div>
    </header>
  );
};
