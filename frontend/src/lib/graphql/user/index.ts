import { client } from '..';
import type { DeleteUserInput, UserDeleteMutation, UserDeleteMutationVariables, UserIdsQuery, UserIdsQueryVariables, UserResendDeleteTokenMutation, UserResendDeleteTokenMutationVariables } from '../generated';
import {
	USER_DELETE,
	USER_IDS,
	USER_RESEND_DELETE_TOKEN
} from './operations.gql';

export const userDelete = async (
	input: DeleteUserInput
) => {
	return client().mutation<UserDeleteMutation, UserDeleteMutationVariables>(USER_DELETE, { input });
};

export const userResendDeleteToken = async (
	input: string
) => {
	return client().mutation<UserResendDeleteTokenMutation, UserResendDeleteTokenMutationVariables>(USER_RESEND_DELETE_TOKEN, { input });
};

export const userIds = async () => client().query<UserIdsQuery, UserIdsQueryVariables>(USER_IDS, {});