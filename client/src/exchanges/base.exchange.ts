import authExchange from "@exchanges/auth.exchange";
import errorExchange from "@exchanges/error.exchange";
import graphcacheExchange from "@exchanges/graphcache.exchange";
import refocusExchange from "@exchanges/refocus.exchange";
import requestPolicyExchange from "@exchanges/request-policy.exchange";
import retryExchange from "@exchanges/retry.exchange";
import scalarsExchange from "@exchanges/scalars.exchange";

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
  multipartFetchExchange,
];

export default Exchanges;
