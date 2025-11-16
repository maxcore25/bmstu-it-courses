import { Badge, type BadgeProps, type BadgeVariants } from '@/shared/ui/badge';
import { Role } from '../config';

type RoleColor = BadgeVariants['variant'];

type RoleConfig = Record<Role, { color: RoleColor; label: string }>;

const roleConfig: RoleConfig = {
  admin: { color: 'outline', label: 'Админ' },
  client: { color: 'info-outline', label: 'Клиент' },
  tutor: { color: 'success-outline', label: 'Преподаватель' },
} as const;

type RoleBadgeProps = {
  role: Role;
  tone?: 'filled' | 'outline';
} & BadgeProps;

export function RoleBadge(props: RoleBadgeProps) {
  const {
    role,
    tone = 'filled',
    variant,
    className,
    children,
    ...rest
  } = props;

  const config = roleConfig[role];

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
