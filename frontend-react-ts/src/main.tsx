import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter } from 'react-router'; // single-page-application
import {
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query'; // data-fetching
import App from './App.tsx';
import { AuthProvider } from './context/AuthContext'; // auth-provider

const queryClient = new QueryClient();

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <AuthProvider>
      <BrowserRouter>
        <QueryClientProvider client={queryClient}>
          <App />
        </QueryClientProvider>
      </BrowserRouter>
    </AuthProvider>
  </StrictMode>,
)
