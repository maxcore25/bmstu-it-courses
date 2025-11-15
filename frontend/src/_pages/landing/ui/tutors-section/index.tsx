import { getTutorsSSR } from '@/entities/user/api/get-tutors-ssr.api';
import { Badge } from '@/shared/ui/badge';
import { Card, CardContent, CardHeader } from '@/shared/ui/card';
import { Star, User } from 'lucide-react';

export const TutorsSection = async () => {
  const tutors = await getTutorsSSR();

  return (
    <section id='tutors' className='pt-32'>
      <div className='container mx-auto px-4'>
        <div className='mx-auto mb-16 max-w-3xl text-center'>
          <Badge className='mb-4'>Наши эксперты</Badge>
          <h2 className='mb-4 text-3xl font-bold md:text-5xl'>
            Учитесь у лучших
          </h2>
          <p className='text-muted-foreground text-lg'>
            Наши преподаватели — профессионалы отрасли с реальным опытом,
            который они привносят в каждое занятие.
          </p>
        </div>

        <div className='grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-4'>
          {tutors.map((tutor, index) => (
            <Card
              key={index}
              className='group gap-0 overflow-hidden py-0 transition-all duration-300 hover:-translate-y-2 hover:shadow-xl'
            >
              <CardHeader className='gap-0 p-0'>
                <div className='bg-muted relative grid aspect-square place-content-center overflow-hidden'>
                  {/* <Image
                    src={'/vercel.svg'}
                    alt={`${tutor.firstName} ${tutor.lastName}`}
                    className='h-full w-full object-cover transition-transform duration-500 group-hover:scale-110'
                    sizes='(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw'
                    fill
                  /> */}
                  <span className='grid w-fit place-content-center rounded-full bg-gray-200 p-6'>
                    <User className='size-16 text-gray-400' />
                  </span>
                  <div className='absolute inset-0 bg-linear-to-t from-black/60 via-transparent to-transparent opacity-0 transition-opacity duration-300 group-hover:opacity-100' />
                </div>
              </CardHeader>
              <CardContent className='p-6'>
                <h3 className='mb-1 flex gap-1 text-lg font-bold'>
                  {`${tutor.firstName} ${tutor.lastName}`}
                  <Badge variant='outline'>{tutor.knowledgeLevel}</Badge>
                </h3>
                <div className='mb-3 flex items-center gap-1'>
                  <Star className='size-4' />
                  <p className='text-primary text-sm font-medium'>
                    {tutor.rating}
                  </p>
                </div>
                <p className='text-muted-foreground line-clamp-2 text-sm'>
                  {tutor.portfolio}
                </p>
              </CardContent>
            </Card>
          ))}
        </div>
      </div>
    </section>
  );
};
