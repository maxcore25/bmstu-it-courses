'use client';

import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarRail,
} from '@/shared/ui/sidebar';
import { NavMain } from '@/widgets/app-sidebar/ui/nav-main';
import { NavUser } from '@/widgets/app-sidebar/ui/nav-user';
import { TeamSwitcher } from '@/widgets/app-sidebar/ui/team-switcher';
import { BookOpen, Home, Pizza, Settings } from 'lucide-react';
import * as React from 'react';

const data = {
  team: {
    name: 'DDD Pizza',
    logo: Pizza,
  },
  navMain: [
    {
      title: 'Home',
      url: '#',
      icon: Home,
      isActive: true,
      items: [
        {
          title: 'Main',
          url: '/',
        },
      ],
    },
    {
      title: 'Documentation',
      url: '#',
      icon: BookOpen,
      items: [
        {
          title: 'Introduction',
          url: '#',
        },
        {
          title: 'Get Started',
          url: '#',
        },
        {
          title: 'Tutorials',
          url: '#',
        },
        {
          title: 'Changelog',
          url: '#',
        },
      ],
    },
    {
      title: 'Settings',
      url: '#',
      icon: Settings,
      items: [
        {
          title: 'General',
          url: '/settings/general',
        },
      ],
    },
  ],
};

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  return (
    <Sidebar collapsible='icon' {...props}>
      <SidebarHeader>
        <TeamSwitcher team={data.team} />
      </SidebarHeader>
      <SidebarContent>
        <NavMain items={data.navMain} />
      </SidebarContent>
      <SidebarFooter>
        <NavUser />
      </SidebarFooter>
      <SidebarRail />
    </Sidebar>
  );
}
