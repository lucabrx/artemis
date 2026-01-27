import * as v from 'valibot';

export const signinSchema = v.object({
	email: v.pipe(v.string(), v.email('Please enter a valid email')),
	password: v.pipe(v.string(), v.minLength(8, 'Password must be at least 8 characters'))
});

export type SigninSchema = typeof signinSchema;
