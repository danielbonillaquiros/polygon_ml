# Ticker Model

A struct used for managing the results from the Polygon API Aggregates.

Example response from API:
```
{
    "ticker": "TSLA",
    "queryCount": 504,
    "resultsCount": 504,
    "adjusted": true,
    "results": [
        {
            "v": 4.9937375e+07,
            "vw": 164.4279,
            "o": 164.4347,
            "c": 163.376,
            "h": 166.356,
            "l": 162.4,
            "t": 1590120000000,
            "n": 191846
        }, ...
    ],
    "status": "DELAYED",
    "request_id": "9171f4f161d5360601c6c3de60ab6ef3",
    "count": 504
}
```

From [Polygon API Documentation](https://polygon.io/docs/stocks/get_v2_aggs_ticker__stocksticker__range__multiplier___timespan___from___to):
- v: The trading volume of the symbol in the given time period.
- vw: The volume weighted average price.
- o: The open price for the symbol in the given time period.
- c: The close price for the symbol in the given time period.
- h: The highest price for the symbol in the given time period.
- l: The lowest price for the symbol in the given time period.
- t: The Unix Msec timestamp for the start of the aggregate window.
- n: The number of transactions in the aggregate window.

This relates to the model this way:
- v: Volume
- vw: AveragePrice
- o: OpenPrice
- c: ClosePrice
- h: HighestPrice
- l: LowestPrice
- t: Timestamp
- n: Transactions