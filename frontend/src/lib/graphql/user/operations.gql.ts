import { gql } from '@urql/svelte';

export const USER_IDS = gql`
	query UserIds {
		getUserIds
	}
`

export const USER_DELETE = gql`
	mutation UserDelete($input: DeleteUserInput!) {
		deleteUser(input: $input) {
			id
		}
	}
`

export const USER_RESEND_DELETE_TOKEN = gql`
	mutation UserResendDeleteToken($input: String!) {
		resendDeleteToken(input: $input)
	}
`

export const USER_DETAILS = gql`
	query UserDetails($input: String!) {
		getUser(input: $input) {
			id
			fullName
			title
			city
			country
			url
		}
	}
`