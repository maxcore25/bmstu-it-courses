export const HeroSection = () => {
  return (
    <section className='mt-8 grid place-content-center gap-6 text-center'>
      <div className='bg-primary/10 border-primary/20 slide-in-from-bottom-4 mx-auto inline-flex w-fit items-center gap-2 rounded-full border px-4 py-2 duration-700'>
        <span className='relative flex h-2 w-2'>
          <span className='bg-primary absolute inline-flex h-full w-full animate-ping rounded-full opacity-75'></span>
          <span className='bg-primary relative inline-flex h-2 w-2 rounded-full'></span>
        </span>
        <span className='text-sm font-medium'>Влетай с нами прямо сейчас!</span>
      </div>

      <h1 className='croll-m-20 text-4xl font-extrabold tracking-tight text-balance'>
        Школа IT-курсов CodeCraft
      </h1>
      <p className='text-muted-foreground max-w-xl text-xl'>
        Современная IT-школа CodeCraft — учитесь у экспертов и стройте карьеру
        будущего!
      </p>
    </section>
  );
};
