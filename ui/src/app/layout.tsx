import type { Metadata } from 'next';
import './globals.css';

export const metadata: Metadata = {
  title: 'CM-Beetle Portal',
  description: 'Computing Infrastructure Migration Dashboard for Cloud-Barista',
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
