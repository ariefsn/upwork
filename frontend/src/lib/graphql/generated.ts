import { z } from 'zod'
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Map: { input: any; output: any; }
  Time: { input: any; output: any; }
  Upload: { input: any; output: any; }
};

export type Audit = {
  __typename?: 'Audit';
  createdAt: Scalars['Time']['output'];
  createdBy?: Maybe<Scalars['String']['output']>;
  publishedAt?: Maybe<Scalars['Time']['output']>;
  publishedBy?: Maybe<Scalars['String']['output']>;
  updatedAt?: Maybe<Scalars['Time']['output']>;
  updatedBy?: Maybe<Scalars['String']['output']>;
};

export type DeleteUserInput = {
  code: Scalars['String']['input'];
  id: Scalars['String']['input'];
};

export enum EarningType {
  FixedPrice = 'FixedPrice',
  Hourly = 'Hourly'
}

export type EarningsData = {
  __typename?: 'EarningsData';
  amount: Scalars['Float']['output'];
  day: Scalars['Int']['output'];
  description: Scalars['String']['output'];
  fee: Scalars['Float']['output'];
  id: Scalars['String']['output'];
  month: Scalars['Int']['output'];
  refID: Scalars['String']['output'];
  team: Scalars['String']['output'];
  type: EarningType;
  userID: Scalars['String']['output'];
  year: Scalars['Int']['output'];
};

export type EarningsDataMonthly = {
  __typename?: 'EarningsDataMonthly';
  items: Array<EarningsDataMonthlyItem>;
  month: Scalars['Int']['output'];
  totalAmount: Scalars['Float']['output'];
  totalFee: Scalars['Float']['output'];
  userID: Scalars['String']['output'];
  year: Scalars['Int']['output'];
};

export type EarningsDataMonthlyItem = {
  __typename?: 'EarningsDataMonthlyItem';
  amount: Scalars['Float']['output'];
  fee: Scalars['Float']['output'];
  type: EarningType;
};

export type EarningsInput = {
  email: Scalars['String']['input'];
  file: Scalars['Upload']['input'];
  userID: Scalars['String']['input'];
};

export type EarningsUserPerYear = {
  __typename?: 'EarningsUserPerYear';
  amount: Scalars['Float']['output'];
  fee: Scalars['Float']['output'];
  user: UserData;
};

export type EarningsUserPerYearInput = {
  userID: Scalars['String']['input'];
  year: Scalars['Int']['input'];
};

export type Mutation = {
  __typename?: 'Mutation';
  deleteUser: UserData;
  resendDeleteToken: Scalars['String']['output'];
  uploadEarnings: Array<EarningsData>;
};


export type MutationDeleteUserArgs = {
  input?: InputMaybe<DeleteUserInput>;
};


export type MutationResendDeleteTokenArgs = {
  input: Scalars['String']['input'];
};


export type MutationUploadEarningsArgs = {
  input: EarningsInput;
};

export type Query = {
  __typename?: 'Query';
  getEarnings: Array<Maybe<EarningsDataMonthly>>;
  getEarningsUsers: Array<Maybe<EarningsUserPerYear>>;
  getEarningsYears: Array<Scalars['Int']['output']>;
  getUser: UserData;
  getUserIds: Array<Scalars['String']['output']>;
};


export type QueryGetEarningsArgs = {
  input: EarningsUserPerYearInput;
};


export type QueryGetEarningsUsersArgs = {
  input: Scalars['Int']['input'];
};


export type QueryGetEarningsYearsArgs = {
  input: Scalars['String']['input'];
};


export type QueryGetUserArgs = {
  input: Scalars['String']['input'];
};

export type Subscription = {
  __typename?: 'Subscription';
  subEarningUsers: Array<EarningsUserPerYear>;
  subEarnings: Array<Maybe<EarningsDataMonthly>>;
  subOnEarningUpdated: Array<Maybe<EarningsDataMonthly>>;
};


export type SubscriptionSubEarningUsersArgs = {
  input: Scalars['Int']['input'];
};


export type SubscriptionSubEarningsArgs = {
  input: EarningsUserPerYearInput;
};


export type SubscriptionSubOnEarningUpdatedArgs = {
  input: EarningsUserPerYearInput;
};

export type UserData = {
  __typename?: 'UserData';
  city: Scalars['String']['output'];
  country: Scalars['String']['output'];
  fullName: Scalars['String']['output'];
  id: Scalars['String']['output'];
  title: Scalars['String']['output'];
  url: Scalars['String']['output'];
};

export type EarningsUsersQueryVariables = Exact<{
  input: Scalars['Int']['input'];
}>;


export type EarningsUsersQuery = { __typename?: 'Query', getEarningsUsers: Array<{ __typename?: 'EarningsUserPerYear', amount: number, fee: number, user: { __typename?: 'UserData', id: string, fullName: string, title: string, city: string, country: string, url: string } } | null> };

export type EarningsUserYearlyQueryVariables = Exact<{
  userID: Scalars['String']['input'];
  year: Scalars['Int']['input'];
}>;


export type EarningsUserYearlyQuery = { __typename?: 'Query', getEarningsYears: Array<number>, getUser: { __typename?: 'UserData', id: string, fullName: string, title: string, city: string, country: string, url: string }, getEarnings: Array<{ __typename?: 'EarningsDataMonthly', userID: string, month: number, year: number, totalAmount: number, totalFee: number, items: Array<{ __typename?: 'EarningsDataMonthlyItem', type: EarningType, amount: number, fee: number }> } | null> };

export type EarningsUploadMutationVariables = Exact<{
  input: EarningsInput;
}>;


export type EarningsUploadMutation = { __typename?: 'Mutation', uploadEarnings: Array<{ __typename?: 'EarningsData', id: string, userID: string, day: number, month: number, year: number, refID: string, type: EarningType, description: string, team: string, amount: number, fee: number }> };

export type EarningsUserYearlySubsSubscriptionVariables = Exact<{
  input: EarningsUserPerYearInput;
}>;


export type EarningsUserYearlySubsSubscription = { __typename?: 'Subscription', subEarnings: Array<{ __typename?: 'EarningsDataMonthly', userID: string, month: number, year: number, totalAmount: number, totalFee: number, items: Array<{ __typename?: 'EarningsDataMonthlyItem', type: EarningType, amount: number, fee: number }> } | null> };

export type EarningsUsersSubsSubscriptionVariables = Exact<{
  input: Scalars['Int']['input'];
}>;


export type EarningsUsersSubsSubscription = { __typename?: 'Subscription', subEarningUsers: Array<{ __typename?: 'EarningsUserPerYear', amount: number, fee: number, user: { __typename?: 'UserData', id: string, fullName: string, title: string, city: string, country: string } }> };

export type UserIdsQueryVariables = Exact<{ [key: string]: never; }>;


export type UserIdsQuery = { __typename?: 'Query', getUserIds: Array<string> };

export type UserDeleteMutationVariables = Exact<{
  input: DeleteUserInput;
}>;


export type UserDeleteMutation = { __typename?: 'Mutation', deleteUser: { __typename?: 'UserData', id: string } };

export type UserResendDeleteTokenMutationVariables = Exact<{
  input: Scalars['String']['input'];
}>;


export type UserResendDeleteTokenMutation = { __typename?: 'Mutation', resendDeleteToken: string };

export type UserDetailsQueryVariables = Exact<{
  input: Scalars['String']['input'];
}>;


export type UserDetailsQuery = { __typename?: 'Query', getUser: { __typename?: 'UserData', id: string, fullName: string, title: string, city: string, country: string, url: string } };


type Properties<T> = Required<{
  [K in keyof T]: z.ZodType<T[K], any, T[K]>;
}>;

type definedNonNullAny = {};

export const isDefinedNonNullAny = (v: any): v is definedNonNullAny => v !== undefined && v !== null;

export const definedNonNullAnySchema = z.any().refine((v) => isDefinedNonNullAny(v));

export const EarningTypeSchema = z.nativeEnum(EarningType);

export function DeleteUserInputSchema(): z.ZodObject<Properties<DeleteUserInput>> {
  return z.object({
    code: z.string(),
    id: z.string()
  })
}

export function EarningsInputSchema(): z.ZodObject<Properties<EarningsInput>> {
  return z.object({
    email: z.string(),
    file: definedNonNullAnySchema,
    userID: z.string()
  })
}

export function EarningsUserPerYearInputSchema(): z.ZodObject<Properties<EarningsUserPerYearInput>> {
  return z.object({
    userID: z.string(),
    year: z.number()
  })
}
