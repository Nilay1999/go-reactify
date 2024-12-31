import React, { createContext, useState, useEffect, useCallback } from 'react';
import apiClient, { apis } from '../api/client';
import { isErrorFromAlias } from '@zodios/core';

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
	login: (email: string, password: string) => Promise<void>;
	logout: () => void;
}

const initialValue: IAuthContext = {
	authenticated: false,
	login: async () => {},
	logout: () => {},
};

const AuthContext = createContext<IAuthContext>(initialValue);

export const AuthProvider = ({ children }: Props) => {
	const [token, setToken] = useState(() => localStorage.getItem('token'));
	const [authenticated, setAuthenticated] = useState(() => !!token);

	const login = useCallback(async (email: string, password: string) => {
		try {
			const response = await apiClient.login({
				identifier: email,
				password,
			});
			const { token } = response?.data;
			setToken(token);
			setAuthenticated(true);
			localStorage.setItem('token', token);
		} catch (error) {
			if (
				isErrorFromAlias(apis, 'login', error) &&
				error.response?.status === 401
			) {
				throw error.response.data;
			}
			throw new Error('Something went wrong');
		}
	}, []);

	const logout = useCallback(() => {
		setToken(null);
		setAuthenticated(false);
		localStorage.removeItem('token');
	}, []);

	useEffect(() => {
		setAuthenticated(!!token);
	}, [token]);

	return (
		<AuthContext.Provider value={{ authenticated, login, logout }}>
			{children}
		</AuthContext.Provider>
	);
};

export { AuthContext };
