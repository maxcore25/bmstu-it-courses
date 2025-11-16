import { Toaster } from '@/shared/ui/sonner';
import type { Metadata } from 'next';
import { ReactQueryProvider } from './_providers/react-query-provider';
import { ThemeProvider } from './_providers/theme-provider';
import './globals.css';

export const metadata: Metadata = {
  title: 'CodeCraft',
  description: 'Современная школа IT-курсов.',
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang='en' suppressHydrationWarning>
      <body className={`antialiased`}>
        <ReactQueryProvider>
          <ThemeProvider
            attribute='class'
            defaultTheme='system'
            enableSystem
            disableTransitionOnChange
          >
            {children}
            <Toaster />
          </ThemeProvider>
        </ReactQueryProvider>
      </body>
    </html>
  );
}
