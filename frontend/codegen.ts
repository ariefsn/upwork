import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  schema: 'http://localhost:3100/graphql',
  documents: ['src/lib/graphql/**/*.gql.{ts,tsx}'],
  generates: {
    'src/lib/graphql/generated.ts': {
      plugins: ['typescript', 'typescript-operations', 'typescript-validation-schema'],
      config: {
        schema: 'zod'
      }
    }
  }
};

export default config;
