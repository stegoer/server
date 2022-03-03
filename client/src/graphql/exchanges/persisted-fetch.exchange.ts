import { persistedFetchExchange as urlqPersistedFetchExchange } from "@urql/exchange-persisted-fetch";

const persistedFetchExchange = urlqPersistedFetchExchange({
  preferGetForPersistedQueries: true,
});

export default persistedFetchExchange;
