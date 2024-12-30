import { makeApi, Zodios } from '@zodios/core';
import z from 'zod';

export const apis = makeApi([
	{
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
	},
	{
		method: 'get',
		path: 'v1/user',
		alias: 'getUsers',
		description: 'Get all users',
		response: z.object({
			data: z.array(
				z.object({
					ID: z.number(),
					CreatedAt: z.string(),
					UpdatedAt: z.string(),
					DeletedAt: z.null(),
					username: z.string(),
					email: z.string(),
					gender: z.string(),
					age: z.number(),
					Posts: z.null(),
				})
			),
		}),
		status: 200,
		headers: {
			Authorization: `Bearer ${localStorage.getItem('token')}`,
		},
	},
]);

const apiClient = new Zodios('http://localhost:8080', apis);
export default apiClient;
