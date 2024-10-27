import React from 'react';
import Header from './Header';

interface AppLayoutProps {
  children: React.ReactNode;
}

export function AppLayout({ children }: AppLayoutProps) {
  return (
    <div className="min-h-screen bg-background">
      <Header onSettingsClick={function (): void {
              throw new Error('Function not implemented.');
          } } />
      <main className="container mx-auto p-4 space-y-6">
        {children}
      </main>
    </div>
  );
}