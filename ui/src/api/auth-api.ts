import { ZodiosEndpointDefinition } from '@zodios/core';
import { z } from 'zod';

export const Signup: ZodiosEndpointDefinition = {
	method: 'post',
	path: 'v1/auth/signin',
	alias: 'login',
	parameters: [
		{
			name: 'body',
			type: 'Body',
			schema: z.object({
				identifier: z.string(),
				password: z.string(),
			}),
		},
	],
	response: z.object({
		data: z.object({
			message: z.string(),
			token: z.string(),
		}),
	}),
	errors: [
		{
			status: 401,
			schema: z.object({
				data: z.object({
					message: z.string(),
				}),
			}),
		},
	],
};

export const Signin: ZodiosEndpointDefinition = {
	method: 'post',
	path: 'v1/auth/signup',
	alias: 'register',
	parameters: [
		{
			name: 'body',
			type: 'Body',
			schema: z.object({
				username: z.string(),
				password: z.string(),
				email: z.string().email({ message: 'Invalid input for Email' }),
				gender: z.string(),
				age: z.number(),
			}),
		},
	],
	response: z.object({
		data: z.object({
			message: z.object({}),
		}),
	}),
};
