import { PublicHeader } from '@/_pages/landing/ui/public-header';

export const LandingPage = () => {
  return (
    <>
      <PublicHeader />
      <main className='container mx-auto'>
        <div className='grid place-content-center gap-6 text-center'>
          <h1 className='croll-m-20 mt-8 text-4xl font-extrabold tracking-tight text-balance'>
            Школа IT-курсов CodeCraft
          </h1>
          <p className='text-muted-foreground max-w-xl text-xl'>
            Современная IT-школа CodeCraft — учитесь у экспертов и стройте
            карьеру будущего!
          </p>
        </div>
      </main>
    </>
  );
};
