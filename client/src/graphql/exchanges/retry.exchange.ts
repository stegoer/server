import { retryExchange as urqlRetryExchange } from "@urql/exchange-retry";

const retryExchange = urqlRetryExchange({});

export default retryExchange;
