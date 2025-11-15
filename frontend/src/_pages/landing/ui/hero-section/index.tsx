import { Button } from '@/shared/ui/button';
import { GridPatternLinearGradient } from '../grid-pattern/linear-grid';
import Link from 'next/link';

export const HeroSection = () => {
  return (
    <section className='relative grid place-content-center overflow-hidden text-center md:py-32'>
      <GridPatternLinearGradient />
      <div className='from-primary/5 via-background to-background absolute inset-0 bg-linear-to-br' />
      <div className='relative z-10 grid gap-6'>
        <div className='bg-primary/10 border-primary/20 slide-in-from-bottom-4 mx-auto inline-flex w-fit items-center gap-2 rounded-full border px-4 py-2 duration-700'>
          <span className='relative flex h-2 w-2'>
            <span className='bg-primary absolute inline-flex h-full w-full animate-ping rounded-full opacity-75'></span>
            <span className='bg-primary relative inline-flex h-2 w-2 rounded-full'></span>
          </span>
          <span className='text-sm font-medium'>
            Влетайте с нами прямо сейчас!
          </span>
        </div>

        <h1 className='croll-m-20 text-6xl font-extrabold tracking-tight text-balance'>
          Школа IT-курсов CodeCraft
        </h1>
        <p className='text-muted-foreground mx-auto max-w-xl text-xl'>
          Современная IT-школа CodeCraft — учитесь у экспертов и стройте карьеру
          будущего!
        </p>
        <div className='mx-auto flex flex-col items-center justify-center gap-4 sm:flex-row'>
          <Button
            asChild
            size='lg'
            className='h-auto px-8 py-2 text-lg font-semibold'
          >
            <Link href='/login'>Записаться на курс</Link>
          </Button>
          <Button
            asChild
            variant='secondary'
            size='lg'
            className='h-auto px-8 py-2 text-lg font-semibold'
          >
            <a href='#courses'>Смотреть программы</a>
          </Button>
        </div>
      </div>
    </section>
  );
};
