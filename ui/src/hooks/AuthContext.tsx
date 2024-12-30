import React, { createContext, useState, useEffect } from 'react';
import apiClient from '../api/client';

type Props = {
	children?: React.ReactNode;
};

export enum AuthStatus {
	Loading,
	SignedIn,
	SignedOut,
}

export interface IAuthContext {
	authenticated: boolean;
	login: (email: string, password: string) => void;
	logout: () => void;
}

const initialValue: IAuthContext = {
	authenticated: false,
	login: () => {},
	logout: () => {},
};

const AuthContext = createContext<IAuthContext>(initialValue);

export const AuthProvider = ({ children }: Props) => {
	const [token, setToken] = useState(localStorage.getItem('token') || null);
	const [authenticated, setAuthenticated] = useState(!!token);

	const login = async (email: string, password: string) => {
		try {
			console.log('here');
			const { data } = await apiClient.login({
				identifier: email,
				password,
			});
			const { token } = data;
			setToken(token);
			setAuthenticated(true);
			localStorage.setItem('token', token);
		} catch (error) {
			console.error('Login failed:', error);
			throw error;
		}
	};

	// const register = async (email: string, password: string) => {
	// 	try {
	// 		const response = await apiClient.('/api/register', {
	// 			email,
	// 			password,
	// 		});
	// 		const { token } = response.data;
	// 		setToken('token');
	// 		localStorage.setItem('token', 'token');
	// 		setIsAuthenticated(true);
	// 	} catch (error) {
	// 		console.error('Registration failed:', error);
	// 		throw error;
	// 	}
	// };

	const logout = () => {
		setToken(null);
		setAuthenticated(false);
		localStorage.removeItem('token');
	};

	useEffect(() => {
		if (token) {
			setAuthenticated(true);
		}
	}, [token]);

	return (
		<AuthContext.Provider value={{ authenticated, login, logout }}>
			{children}
		</AuthContext.Provider>
	);
};

export { AuthContext };
