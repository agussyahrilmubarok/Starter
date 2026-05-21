import React, { createContext, useState, useEffect, type ReactNode } from 'react';
import Cookies from 'js-cookie';

// Define the type of the context value
interface AuthContextType {
    isAuthenticated: boolean;
    setIsAuthenticated: React.Dispatch<React.SetStateAction<boolean>>;
}

// Create context with undefined as the default value
export const AuthContext = createContext<AuthContextType | undefined>(undefined);

// Define props type for AuthProvider
interface AuthProviderProps {
    children: ReactNode;
}

// Provider component for authentication context
export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
    const [isAuthenticated, setIsAuthenticated] = useState<boolean>(!!Cookies.get('token'));

    useEffect(() => {
        const handleTokenChange = () => {
            setIsAuthenticated(!!Cookies.get('token'));
        };

        window.addEventListener('storage', handleTokenChange);
        
        return () => {
            window.removeEventListener('storage', handleTokenChange);
        };
    }, []);

    return (
        <AuthContext.Provider value={{ isAuthenticated, setIsAuthenticated }}>
            {children}
        </AuthContext.Provider>
    );
};