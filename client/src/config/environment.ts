export const IS_PRODUCTION = process.env.NODE_ENV === `production`;

export const SERVER_URI = process.env.NEXT_PUBLIC_SERVER_URI as string;
export const SERVER_GRAPHQL_ENDPOINT = `${SERVER_URI}/graphql`;

export const GA_MEASUREMENT_ID = process.env
  .NEXT_PUBLIC_GA_MEASUREMENT_ID as string;
