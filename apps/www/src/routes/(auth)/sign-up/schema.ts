import * as v from 'valibot';

export const signupSchema = v.object({
	name: v.pipe(v.string(), v.minLength(1, 'Name is required')),
	email: v.pipe(v.string(), v.email('Please enter a valid email')),
	password: v.pipe(v.string(), v.minLength(8, 'Password must be at least 8 characters')),
	confirmPassword: v.pipe(v.string(), v.minLength(1, 'Please confirm your password'))
});

export type SignupSchema = typeof signupSchema;
