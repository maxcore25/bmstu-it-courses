import { getBranchesSSR } from '@/entities/branch/api/get-branches-ssr.api';
import { Badge } from '@/shared/ui/badge';
import { Card, CardContent, CardHeader, CardTitle } from '@/shared/ui/card';
import { Clock, Mail, MapPin, Phone } from 'lucide-react';
import Image from 'next/image';

export const BranchesSection = async () => {
  const branches = await getBranchesSSR();

  return (
    <section id='branches' className='py-20 md:py-32'>
      <div className='container mx-auto px-4'>
        <div className='mx-auto mb-16 max-w-3xl text-center'>
          <Badge className='mb-4'>Our Locations</Badge>
          <h2 className='mb-4 text-3xl font-bold md:text-5xl'>
            Find A Campus Near You
          </h2>
          <p className='text-muted-foreground text-lg'>
            Visit any of our modern learning spaces equipped with the latest
            technology.
          </p>
        </div>

        <div className='grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3'>
          {branches.map((branch, index) => (
            <Card
              key={index}
              className='group overflow-hidden transition-all duration-300 hover:-translate-y-1 hover:shadow-xl'
            >
              <CardHeader className='p-0'>
                <div className='relative aspect-video overflow-hidden'>
                  <Image
                    src={'/vercel.svg'}
                    alt={branch.address}
                    className='h-full w-full object-cover transition-transform duration-300 group-hover:scale-110'
                    sizes='(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw'
                    fill
                  />
                  <div className='absolute top-4 right-4'>
                    <Badge className='bg-background/90 backdrop-blur'>
                      {branch.rooms}
                    </Badge>
                  </div>
                </div>
              </CardHeader>
              <CardContent className='p-6'>
                <CardTitle className='mb-4 text-xl'>{branch.address}</CardTitle>
                <div className='space-y-3'>
                  <div className='flex gap-3 text-sm'>
                    <MapPin className='text-muted-foreground mt-0.5 h-4 w-4 shrink-0' />
                    <span className='text-muted-foreground'>
                      {branch.address}
                    </span>
                  </div>
                  <div className='flex gap-3 text-sm'>
                    <Phone className='text-muted-foreground h-4 w-4 shrink-0' />
                    <span className='text-muted-foreground'>
                      {branch.address}
                    </span>
                  </div>
                  <div className='flex gap-3 text-sm'>
                    <Mail className='text-muted-foreground h-4 w-4 shrink-0' />
                    <span className='text-muted-foreground'>admin@mail.ru</span>
                  </div>
                  <div className='flex gap-3 text-sm'>
                    <Clock className='text-muted-foreground h-4 w-4 shrink-0' />
                    <span className='text-muted-foreground'>10:00-18:00</span>
                  </div>
                </div>
              </CardContent>
            </Card>
          ))}
        </div>
      </div>
    </section>
  );
};
