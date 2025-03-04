import { gql } from '@urql/svelte';

export const EARNINGS_USERS = gql`
  query EarningsUsers($input: Int!) {
    getEarningsUsers(input: $input) {
      user {
        id
        fullName
        title
        city
        country
        url
      }
      amount
      fee
    }
  }
`
export const EARNINGS_USER_YEARLY = gql`
  query EarningsUserYearly($userID: String!, $year: Int!) {
    getEarningsYears(input: $userID)
    getUser(input: $userID) {
      id
      fullName
      title
      city
      country
      url
    }
    getEarnings(input: { userID: $userID, year: $year }) {
      userID
      month
      year
      totalAmount
      totalFee
      items {
        type
        amount
        fee
      }
    }
  }
`

export const EARNINGS_UPLOAD = gql`
  mutation EarningsUpload($input: EarningsInput!) {
    uploadEarnings(input: $input) {
      id
      userID
      day
      month
      year
      refID
      type
      description
      team
      amount
      fee
    }
  }
`

export const EARNINGS_USER_YEARLY_SUBS = gql`
  subscription EarningsUserYearlySubs($input: EarningsUserPerYearInput!) {
    subEarnings(input: $input) {
      userID
      month
      year
      totalAmount
      totalFee
      items {
        type
        amount
        fee
      }
    }
  }
`

export const EARNINGS_USERS_SUBS = gql`
  subscription EarningsUsersSubs($input: Int!) {
    subEarningUsers(input: $input) {
      user {
        id
        fullName
        title
        city
        country
      }
      amount
      fee
    }
  }
`