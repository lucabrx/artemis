import { superValidate } from 'sveltekit-superforms';
import { valibot } from 'sveltekit-superforms/adapters';
import { signinSchema } from './schema';
import { fail, redirect } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';
import { api } from '$lib';
import { message } from 'sveltekit-superforms/server';

export const load: PageServerLoad = async () => {
	const form = await superValidate(valibot(signinSchema));

	return { form };
};

export const actions: Actions = {
	default: async ({ request, cookies }) => {
		const form = await superValidate(request, valibot(signinSchema));

		if (!form.valid) {
			return fail(400, { form });
		}

		try {
			const response = await api
				.post('auth/login', {
					json: form.data
				})
				.json<LoginResponse>();

			cookies.set('access_token', response.access_token, {
				path: '/',
				httpOnly: true,
				sameSite: 'strict',
				secure: import.meta.env.PROD,
				expires: new Date(response.access_token_expires_at * 1000)
			});

			cookies.set('refresh_token', response.refresh_token, {
				path: '/',
				httpOnly: true,
				sameSite: 'strict',
				secure: import.meta.env.PROD,
				expires: new Date(response.refresh_token_expires_at * 1000)
			});
		} catch (error) {
			console.error(error);
			const errorMessage = error instanceof Error ? error.message : 'Invalid credentials';

			return message(form, errorMessage, {
				status: 401
			});
		}

		throw redirect(303, '/');
	}
};
