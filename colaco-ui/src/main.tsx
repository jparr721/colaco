import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App.tsx';
import './index.css';
import { ThemeProvider } from 'styled-components';
import theme from './theme.ts';
import { QueryClientProvider } from '@tanstack/react-query';
import queryClient from './query.ts';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <React.Suspense>
      <ThemeProvider theme={theme}>
        <QueryClientProvider client={queryClient}>
          <App />
        </QueryClientProvider>
      </ThemeProvider>
    </React.Suspense>
  </React.StrictMode>,
);