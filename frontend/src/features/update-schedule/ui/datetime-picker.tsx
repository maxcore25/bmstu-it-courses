'use client';

import { useState, useEffect } from 'react';
import { Button } from '@/shared/ui/button';
import { Calendar } from '@/shared/ui/calendar';
import { Input } from '@/shared/ui/input';
import { Label } from '@/shared/ui/label';
import { Popover, PopoverContent, PopoverTrigger } from '@/shared/ui/popover';
import { ChevronDownIcon } from 'lucide-react';

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

  // local UI state
  const [date, setDate] = useState<Date | undefined>(undefined);
  const [time, setTime] = useState('10:30:00');

  // Sync form → local UI when editing existing schedule
  useEffect(() => {
    if (!value) return;
    const d = new Date(value);
    setDate(d);

    const hh = d.getHours().toString().padStart(2, '0');
    const mm = d.getMinutes().toString().padStart(2, '0');
    const ss = d.getSeconds().toString().padStart(2, '0');
    setTime(`${hh}:${mm}:${ss}`);
  }, [value]);

  // Combine date + time → ISO string
  const updateDatetime = (newDate = date, newTime = time) => {
    if (!newDate) return;

    const [hh, mm, ss] = newTime.split(':').map(Number);

    const dt = new Date(newDate);
    dt.setHours(hh, mm, ss);

    onChange(dt.toISOString());
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
              mode='single'
              selected={date}
              captionLayout='dropdown'
              onSelect={d => {
                setDate(d);
                if (d) updateDatetime(d, time);
                setOpen(false);
              }}
            />
          </PopoverContent>
        </Popover>
      </div>

      <div className='flex flex-col gap-3'>
        <Label className='px-1'>Время</Label>
        <Input
          type='time'
          step='1'
          value={time}
          className='bg-background appearance-none [&::-webkit-calendar-picker-indicator]:hidden'
          onChange={e => {
            const newTime = e.target.value;
            setTime(newTime);
            updateDatetime(date, newTime);
          }}
        />
      </div>
    </div>
  );
}
