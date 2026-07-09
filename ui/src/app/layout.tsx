import type { Metadata } from 'next';
import './globals.css';

export const metadata: Metadata = {
  title: 'Beetle Lab',
  description: 'Beetle Lab — CM-Beetle Demo & Testing Dashboard for Cloud-Barista',
  icons: {
    icon: '/favicon.svg',
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="light" suppressHydrationWarning>
      <body suppressHydrationWarning>{children}</body>
    </html>
  );
}
