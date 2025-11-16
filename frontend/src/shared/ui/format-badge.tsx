import { Badge, type BadgeProps, type BadgeVariants } from '@/shared/ui/badge';
import { CourseFormat } from '../config';

type FormatColor = BadgeVariants['variant'];

type FormatConfig = Record<CourseFormat, { color: FormatColor; label: string }>;

const formatConfig: FormatConfig = {
  group: { color: 'secondary', label: 'Группа' },
  individual: { color: 'info', label: 'Индивидуально' },
  intensive: { color: 'info-secondary', label: 'Интенсив' },
} as const;

type FormatBadgeProps = {
  format: CourseFormat;
  tone?: 'filled' | 'outline';
} & BadgeProps;

export function FormatBadge(props: FormatBadgeProps) {
  const {
    format,
    tone = 'filled',
    variant,
    className,
    children,
    ...rest
  } = props;

  const config = formatConfig[format];

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
    <Badge variant={computedVariant} className={className} {...rest}>
      {children ?? config.label}
    </Badge>
  );
}
