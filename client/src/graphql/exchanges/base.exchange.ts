import authExchange from "@graphql/exchanges/auth.exchange";
import errorExchange from "@graphql/exchanges/error.exchange";
import graphcacheExchange from "@graphql/exchanges/graphcache.exchange";
import persistedFetchExchange from "@graphql/exchanges/persisted-fetch.exchange";
import refocusExchange from "@graphql/exchanges/refocus.exchange";
import requestPolicyExchange from "@graphql/exchanges/request-policy.exchange";
import retryExchange from "@graphql/exchanges/retry.exchange";
import scalarsExchange from "@graphql/exchanges/scalars.exchange";

import { multipartFetchExchange } from "@urql/exchange-multipart-fetch";
import { dedupExchange } from "urql";

import type { Exchange } from "urql";

const Exchanges: Exchange[] = [
  dedupExchange,
  requestPolicyExchange,
  refocusExchange,
  scalarsExchange,
  graphcacheExchange,
  errorExchange,
  authExchange,
  retryExchange,
  persistedFetchExchange,
  multipartFetchExchange,
];

export default Exchanges;
