# Upwork Frontend

Built with `SvelteKit`, `Shadcn`, and `Urql`.

## How to

1. Copy `.env.example` to `.env` and set the values
2. Install the dependencies `yarn install`
3. Run the app `yarn dev --host`

## GraphQL Generator

- If there's any changes from backend side, update the schema at `src/lib/graphql/**/operations.gql.ts` then run `yarn codegen`.
- More info please check [Codegen](https://the-guild.dev/graphql/codegen).
