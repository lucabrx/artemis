import { superValidate } from 'sveltekit-superforms';
import { valibot } from 'sveltekit-superforms/adapters';
import { signupSchema } from './schema';
import { fail, redirect } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';
import { api } from '$lib';

export const load: PageServerLoad = async () => {
	const form = await superValidate(valibot(signupSchema));

	return { form };
};

export const actions: Actions = {
	default: async ({ request }) => {
		const form = await superValidate(request, valibot(signupSchema));

		if (!form.valid) {
			return fail(400, { form });
		}

		console.log('Valid Form Data:', form.data);
		const payload = {
			email: form.data.email,
			password: form.data.password,
			name: form.data.name
		};

		try {
			api.post('auth/register', {
				json: payload
			});
		} catch (error) {
			console.error('Error during registration:', error);
			return fail(500, {
				form,
				message: 'An error occurred during registration. Please try again.'
			});
		}

		throw redirect(307, '/sign-in');
	}
};
