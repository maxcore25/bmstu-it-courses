import { getCoursesSSR } from '@/entities/course/api/get-courses-ssr.api';
import { Badge } from '@/shared/ui/badge';
import { Button } from '@/shared/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/shared/ui/card';
import { Clock, TrendingUp, Users } from 'lucide-react';
import Image from 'next/image';

import { readdirSync } from 'fs';
import { join } from 'path';

let filesCache: string[] | null = null;

export function getRandomPublicImage() {
  if (!filesCache) {
    const folder = join(process.cwd(), 'public', 'img');
    filesCache = readdirSync(folder).filter(f =>
      /\.(png|jpe?g|webp|gif|svg)$/i.test(f)
    );
  }

  const random = Math.floor(Math.random() * filesCache.length);
  return '/img/' + filesCache[random];
}

export const CoursesSection = async () => {
  const courses = await getCoursesSSR(['author']);

  return (
    <section id='courses' className='py-20 md:py-32'>
      <div className='container mx-auto px-4'>
        <div className='mx-auto mb-16 max-w-3xl text-center'>
          <Badge className='mb-4'>Учебная программа</Badge>
          <h2 className='mb-4 text-3xl font-bold md:text-5xl'>
            Изучайте наши курсы
          </h2>
          <p className='text-muted-foreground text-lg'>
            Выбирайте из широкого спектра курсов, подобранных для ваших
            карьерных целей и уровня навыков.
          </p>
        </div>

        <div className='grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3'>
          {courses.map((course, index) => {
            const randomImage = getRandomPublicImage();

            return (
              <Card
                key={index}
                className='group overflow-hidden pt-0 transition-all duration-300 hover:-translate-y-1 hover:shadow-lg'
              >
                <div className='relative aspect-video overflow-hidden'>
                  <Image
                    src={randomImage}
                    alt={course.name}
                    className='h-full w-full object-cover transition-transform duration-300 group-hover:scale-110'
                    sizes='(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw'
                    fill
                  />
                  <div className='absolute top-4 right-4'>
                    <Badge variant='secondary'>{course.format}</Badge>
                  </div>
                </div>
                <CardHeader>
                  <CardTitle className='text-xl'>{course.name}</CardTitle>
                  <CardDescription className='line-clamp-2'>
                    Сложность: {course.difficulty}
                  </CardDescription>
                </CardHeader>
                <CardContent>
                  <div className='text-muted-foreground flex items-center gap-4 text-sm'>
                    <div className='flex items-center gap-1'>
                      <Clock className='h-4 w-4' />
                      <span>{course.duration}</span>
                    </div>
                    <div className='flex items-center gap-1'>
                      <Users className='h-4 w-4' />
                      {(course.format === 'group' ||
                        course.format === 'intensive') && (
                        <span>от 5 человек</span>
                      )}
                      {course.format === 'individual' && <span>1 человек</span>}
                    </div>
                  </div>
                </CardContent>
                <CardFooter>
                  <Button className='group/btn w-full'>
                    Начать учиться
                    <TrendingUp className='ml-2 h-4 w-4 transition-transform group-hover/btn:translate-x-1' />
                  </Button>
                </CardFooter>
              </Card>
            );
          })}
        </div>
      </div>
    </section>
  );
};
