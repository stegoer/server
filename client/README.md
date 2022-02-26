# stegoer - client

- Client is using TypeScript, NextJS and GraphQL.

## Installation

### Install dependencies

```sh
npm install
```

### Create the `.env.local` file

Create a `.env.local` file and copy the contents of `.env.local.example` file into the `.env.local` file

```sh
cp .env.local.example .env.local
```

## Development

### Dev server

```sh
npm run dev
```

### GraphQL

[GraphQL Code Generator](https://www.graphql-code-generator.com/)
is used to generate type definitions and hooks for queries and mutations. See `client/src/graphql/codegen.yml` for
configuration options.

#### Codegen

```sh
npm run gen
```

To add a new query or mutation head to `client/src/graphql/user` or
`client/src/graphql/image` and add a new file.

To add a new fragment head to the `client/src/graphql/fragments` folder.
