import { Badge } from '@/shared/ui/badge';
import {
  Award,
  Building2,
  Globe,
  GraduationCap,
  Star,
  Trophy,
  Users,
  Zap,
} from 'lucide-react';

const stats = [
  {
    icon: Users,
    value: '2 500+',
    label: 'Активных студентов',
    description: 'Учится у нас',
  },
  {
    icon: GraduationCap,
    value: '1 800+',
    label: 'Выпускников',
    description: 'Успешно трудоустроены',
  },
  {
    icon: Award,
    value: '98%',
    label: 'Успешно',
    description: 'Завершили курсы',
  },
  {
    icon: Building2,
    value: '150+',
    label: 'Партнеров',
    description: 'Трудоустраивают выпускников',
  },
  {
    icon: Globe,
    value: '45+',
    label: 'Городов',
    description: 'География студентов',
  },
  {
    icon: Trophy,
    value: '25+',
    label: 'Получено наград',
    description: 'Признание отрасли',
  },
  {
    icon: Star,
    value: '4.9/5',
    label: 'Рейтинг',
    description: 'От наших студентов',
  },
  {
    icon: Zap,
    value: '50+',
    label: 'Курсов',
    description: 'Доступных программ',
  },
];

export const StatsSection = () => {
  return (
    <section id='stats' className='relative overflow-hidden py-20 md:py-32'>
      <div className='from-primary/10 via-background to-background absolute inset-0 bg-linear-to-br' />
      <div className='bg-grid-slate-900/[0.04] absolute inset-0' />

      <div className='relative z-10 container mx-auto px-4'>
        <div className='mx-auto mb-16 max-w-3xl text-center'>
          <Badge className='mb-4'>Наши достижения</Badge>
          <h2 className='mb-4 text-3xl font-bold md:text-5xl'>
            Нам доверяют тысячи студентов
          </h2>
          <p className='text-muted-foreground text-lg'>
            Цифры, которые показывают нашу приверженность качеству и успеху
            студентов.
          </p>
        </div>

        <div className='grid grid-cols-2 gap-6 md:grid-cols-4 md:gap-8'>
          {stats.map((stat, index) => {
            const Icon = stat.icon;
            return (
              <div
                key={index}
                className='group bg-card hover:border-primary/50 relative rounded-2xl border p-6 transition-all duration-300 hover:-translate-y-1 hover:shadow-lg md:p-8'
              >
                <div className='from-primary/5 absolute inset-0 rounded-2xl bg-linear-to-br to-transparent opacity-0 transition-opacity group-hover:opacity-100' />
                <div className='relative'>
                  <div className='bg-primary/10 group-hover:bg-primary/20 mb-4 flex h-12 w-12 items-center justify-center rounded-xl transition-colors md:h-14 md:w-14'>
                    <Icon className='text-primary h-6 w-6 md:h-7 md:w-7' />
                  </div>
                  <div className='from-foreground to-foreground/70 mb-1 bg-linear-to-br bg-clip-text text-3xl font-bold md:text-4xl'>
                    {stat.value}
                  </div>
                  <div className='mb-1 text-sm font-semibold md:text-base'>
                    {stat.label}
                  </div>
                  <div className='text-muted-foreground text-xs md:text-sm'>
                    {stat.description}
                  </div>
                </div>
              </div>
            );
          })}
        </div>

        <div className='from-primary to-primary/80 text-primary-foreground mt-16 rounded-3xl bg-linear-to-br p-8 text-center md:p-12'>
          <h3 className='mb-4 text-2xl font-bold md:text-4xl'>
            Готовы начать свой путь?
          </h3>
          <p className='mx-auto mb-8 max-w-2xl text-lg opacity-90'>
            Присоединяйтесь к тысячам студентов, которые уже изменили свою
            карьеру с нашими курсами.
          </p>
          <div className='flex flex-col justify-center gap-4 sm:flex-row'>
            <button className='bg-background text-foreground hover:bg-background/90 rounded-lg px-8 py-4 font-semibold shadow-lg transition-colors'>
              Записаться
            </button>
            <button className='bg-primary-foreground/10 hover:bg-primary-foreground/20 border-primary-foreground/20 rounded-lg border px-8 py-4 font-semibold backdrop-blur transition-colors'>
              Запланировать визит
            </button>
          </div>
        </div>
      </div>
    </section>
  );
};
