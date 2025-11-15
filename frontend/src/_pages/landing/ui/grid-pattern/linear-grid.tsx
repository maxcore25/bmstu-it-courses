import { cn } from '@/shared/lib/utils';
import { GridPattern } from '@/shared/ui/grid-pattern';

export const GridPatternLinearGradient = () => {
  return (
    <GridPattern
      className={cn(
        'mask-[linear-gradient(to_bottom_right,white,transparent,transparent)]'
      )}
    />
  );
};
