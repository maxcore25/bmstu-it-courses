import { Code2 } from 'lucide-react';

export const Footer = () => {
  const footerSections = [
    {
      title: 'Программы',
      links: [
        'Веб-разработка',
        'Наука о данных',
        'Мобильная разработка',
        'Облачные технологии',
        'Кибербезопасность',
        'Все курсы',
      ],
    },
    {
      title: 'О компании',
      links: [
        'О нас',
        'Наша команда',
        'Вакансии',
        'Пресса',
        'Партнеры',
        'Контакты',
      ],
    },
    {
      title: 'Ресурсы',
      links: ['Блог', 'Туториалы', 'Сообщество', 'События', 'FAQ'],
    },
    {
      title: 'Поддержка',
      links: ['Центр помощи', 'Личный кабинет', 'Условия пользования'],
    },
  ];

  return (
    <footer className='bg-muted/50 border-t'>
      <div className='container mx-auto px-4 py-16 md:py-20'>
        <div className='mb-12 grid grid-cols-1 gap-12 md:grid-cols-2 lg:grid-cols-6'>
          <div className='lg:col-span-2'>
            <a href='#' className='group mb-4 flex items-center gap-2'>
              <div className='bg-primary flex h-10 w-10 items-center justify-center rounded-lg transition-transform group-hover:scale-110'>
                <Code2 className='text-primary-foreground h-6 w-6' />
              </div>
              <span className='text-xl font-bold'>CodeCraft</span>
            </a>
            <p className='text-muted-foreground mb-6 max-w-sm'>
              Мы помогаем следующему поколению IT-специалистов получить
              качественное образование и практический опыт.
            </p>
          </div>

          {footerSections.map((section, index) => (
            <div key={index}>
              <h3 className='mb-4 font-semibold'>{section.title}</h3>
              <ul className='space-y-3'>
                {section.links.map((link, linkIndex) => (
                  <li key={linkIndex}>
                    <a
                      href='#'
                      className='text-muted-foreground hover:text-foreground text-sm transition-colors'
                    >
                      {link}
                    </a>
                  </li>
                ))}
              </ul>
            </div>
          ))}
        </div>

        <div className='border-t pt-8'>
          <div className='flex flex-col items-center justify-between gap-4 md:flex-row'>
            <p className='text-muted-foreground text-center text-sm md:text-left'>
              &copy; 2025 CodeCraft. Все права защищены.
            </p>
            <div className='flex flex-wrap justify-center gap-6 text-sm'>
              <a
                href='#'
                className='text-muted-foreground hover:text-foreground transition-colors'
              >
                Политика конфиденциальности
              </a>
              <a
                href='#'
                className='text-muted-foreground hover:text-foreground transition-colors'
              >
                Условия пользования
              </a>
              <a
                href='#'
                className='text-muted-foreground hover:text-foreground transition-colors'
              >
                Политика использования файлов Cookies
              </a>
              <a
                href='#'
                className='text-muted-foreground hover:text-foreground transition-colors'
              >
                Доступность
              </a>
            </div>
          </div>
        </div>
      </div>
    </footer>
  );
};
