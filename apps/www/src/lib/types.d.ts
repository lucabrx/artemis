type User = {
	id: string;
	email: string;
	name: string;
	avatar_url?: string;
	created_at: string;
	updated_at: string;
};

type LoginResponse = {
	access_token: string;
	refresh_token: string;
	access_token_expires_at: number;
	refresh_token_expires_at: number;
	user: User;
};

type TokenResponse = {
	access_token: string;
	refresh_token: string;
	access_token_expires_at: number;
	refresh_token_expires_at: number;
};
