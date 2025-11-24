import { Badge, type BadgeProps, type BadgeVariants } from '@/shared/ui/badge';
import { Level } from '../config';
import { cn } from '../lib/utils';

type LevelColor = BadgeVariants['variant'];

type LevelConfig = Record<Level, { color: LevelColor; label: string }>;

const levelConfig: LevelConfig = {
  beginner: { color: 'success', label: 'Начальный' },
  intermediate: { color: 'warn', label: 'Средний' },
  advanced: { color: 'destructive-secondary', label: 'Продвинутый' },
} as const;

type LevelBadgeProps = {
  level: Level;
  tone?: 'filled' | 'outline';
} & BadgeProps;

export function LevelBadge(props: LevelBadgeProps) {
  const {
    level,
    tone = 'filled',
    variant,
    className,
    children,
    ...rest
  } = props;

  const config = levelConfig[level];

  let computedVariant: BadgeVariants['variant'];
  if (variant) {
    computedVariant = variant;
  } else {
    const base = config.color;
    computedVariant =
      tone === 'outline'
        ? (`${base}-outline` as BadgeVariants['variant'])
        : (base as BadgeVariants['variant']);
  }

  return (
    <Badge
      variant={computedVariant}
      className={cn('h-fit', className)}
      {...rest}
    >
      {children ?? config.label}
    </Badge>
  );
}
