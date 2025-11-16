'use client';

import { Button } from '@/shared/ui/button';
import { Calendar } from '@/shared/ui/calendar';
import { Input } from '@/shared/ui/input';
import { Label } from '@/shared/ui/label';
import { Popover, PopoverContent, PopoverTrigger } from '@/shared/ui/popover';
import { enUS, Locale, ru } from 'date-fns/locale';
import { ChevronDownIcon } from 'lucide-react';
import { useMemo, useState } from 'react';

const localeMap: Record<string, Locale> = {
  'ru-RU': ru,
  ru: ru,
  'en-US': enUS,
  en: enUS,
};

interface DatetimePickerProps {
  value?: string; // ISO string
  onChange: (value: string) => void;
  label?: string;
}

export function DatetimePicker({
  value,
  onChange,
  label,
}: DatetimePickerProps) {
  const [open, setOpen] = useState(false);

  const userLocale =
    (typeof navigator !== 'undefined' && navigator.language) || 'en-US';

  const dfnsLocale = useMemo(() => localeMap[userLocale] || enUS, [userLocale]);

  // --- 1. Derive date & time from the prop (no effects needed)
  const derived = useMemo(() => {
    if (!value) return { date: undefined, time: '10:30:00' };

    const d = new Date(value);

    const hh = String(d.getHours()).padStart(2, '0');
    const mm = String(d.getMinutes()).padStart(2, '0');
    const ss = String(d.getSeconds()).padStart(2, '0');

    return {
      date: d,
      time: `${hh}:${mm}:${ss}`,
    };
  }, [value]);

  // --- 2. Local overrides when the user interacts
  const [localDate, setLocalDate] = useState<Date | undefined>(undefined);
  const [localTime, setLocalTime] = useState<string | undefined>(undefined);

  const date = localDate ?? derived.date;
  const time = localTime ?? derived.time;

  // --- 3. Emit ISO string
  const update = (d: Date | undefined, t: string) => {
    if (!d) return;

    const [hh, mm, ss] = t.split(':').map(Number);
    const result = new Date(d);
    result.setHours(hh, mm, ss);

    onChange(result.toISOString());
  };

  return (
    <div className='flex gap-4'>
      <div className='flex flex-col gap-3'>
        {label ? <Label className='px-1'>{label}</Label> : null}

        <Popover open={open} onOpenChange={setOpen}>
          <PopoverTrigger asChild>
            <Button
              variant='outline'
              className='w-36 justify-between font-normal'
            >
              {date ? date.toLocaleDateString() : 'Выберите дату'}
              <ChevronDownIcon />
            </Button>
          </PopoverTrigger>

          <PopoverContent className='w-auto p-0' align='start'>
            <Calendar
              locale={dfnsLocale}
              mode='single'
              selected={date}
              captionLayout='dropdown'
              onSelect={d => {
                if (!d) return;
                setLocalDate(d);
                update(d, time!);
                setOpen(false);
              }}
            />
          </PopoverContent>
        </Popover>
      </div>

      {/* Time */}
      <div className='flex flex-col gap-3'>
        <Label className='px-1'>Время</Label>
        <Input
          type='time'
          step='1'
          value={time}
          className='bg-background appearance-none [&::-webkit-calendar-picker-indicator]:hidden'
          onChange={e => {
            const newTime = e.target.value;
            setLocalTime(newTime);
            update(date!, newTime);
          }}
        />
      </div>
    </div>
  );
}
