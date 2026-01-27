import { api } from '$lib';
import type { Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
	const accessToken = event.cookies.get('access_token');
	const refreshToken = event.cookies.get('refresh_token');

	if (!accessToken && refreshToken) {
		try {
			const response = await api
				.post('auth/refresh', {
					json: { refresh_token: refreshToken }
				})
				.json<TokenResponse>();

			event.cookies.set('access_token', response.access_token, {
				path: '/',
				httpOnly: true,
				sameSite: 'strict',
				secure: import.meta.env.PROD,
				expires: new Date(response.access_token_expires_at * 1000)
			});

			event.cookies.set('refresh_token', response.refresh_token, {
				path: '/',
				httpOnly: true,
				sameSite: 'strict',
				secure: import.meta.env.PROD,
				expires: new Date(response.refresh_token_expires_at * 1000)
			});
		} catch (error) {
			console.error('Failed to refresh token:', error);
			event.cookies.delete('access_token', { path: '/' });
			event.cookies.delete('refresh_token', { path: '/' });
		}
	}

	return resolve(event);
};
