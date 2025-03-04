import { client } from '..';
import type { EarningsInput, EarningsUploadMutation, EarningsUploadMutationVariables, EarningsUserPerYearInput, EarningsUsersQuery, EarningsUsersQueryVariables, EarningsUsersSubsSubscriptionVariables, EarningsUserYearlyQuery, EarningsUserYearlyQueryVariables, EarningsUserYearlySubsSubscription, EarningsUserYearlySubsSubscriptionVariables } from '../generated';
import {
  EARNINGS_UPLOAD,
  EARNINGS_USER_YEARLY,
  EARNINGS_USER_YEARLY_SUBS,
  EARNINGS_USERS,
  EARNINGS_USERS_SUBS
} from './operations.gql';

export const earningsUsers = async (input: number) => {
  return client().query<EarningsUsersQuery, EarningsUsersQueryVariables>(EARNINGS_USERS, { input });
};

export const earningsUpload = async (input: EarningsInput) => {
  return client().mutation<EarningsUploadMutation, EarningsUploadMutationVariables>(EARNINGS_UPLOAD, { input });
};

export const earningsUserYearly = async (userID: string, year: number) => {
  return client().query<EarningsUserYearlyQuery, EarningsUserYearlyQueryVariables>(EARNINGS_USER_YEARLY, { userID, year });
};

export const earningsUserYearlySubs = (input: EarningsUserPerYearInput) => {
  return client().subscription<EarningsUserYearlySubsSubscription, EarningsUserYearlySubsSubscriptionVariables>(EARNINGS_USER_YEARLY_SUBS, { input })
}

export const earningsUsersSubs = (input: number) => {
  return client().subscription<EarningsUserYearlySubsSubscription, EarningsUsersSubsSubscriptionVariables>(EARNINGS_USERS_SUBS, { input })
}