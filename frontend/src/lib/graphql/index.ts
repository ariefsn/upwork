import { env } from '$env/dynamic/public';
import { cacheExchange, Client, fetchExchange } from '@urql/svelte';

export const url = (isWs?: boolean) => {
	if (isWs) {
		const wsUrl = env.PUBLIC_ORIGIN.replace('https://', 'wss://').replace('http://', 'ws://')
		return wsUrl + '/api/graphql'
	}

	return env.PUBLIC_ORIGIN + '/api/graphql'
}

export const client = () =>
	new Client({
		url: url(),
		exchanges: [
			cacheExchange,
			fetchExchange,
		]
	});

export * from './earnings';
export * from './user';

