### Very simple exchange rate CMD line tool

CurrenC uses the [Fixer.io](http://fixer.io/) foreign exchange rates and currency conversion API to quickly display currency exchange rates from the command line.

```
$ currenc -h
Usage of currenc:
  -amount float
    	Amount of base currency to convert. (default 1)
  -from string
    	Currency base to be quoted against. (default "USD")
  -to string
    	Currency to convert to. (default "EUR")
```

Example usage:

```
$ currenc -to AUD             # Convert default 1 USD to AUD
$ currenc -amount 25 -to AUD  # Convert 25 USD to AUD
$ currenc -from EUR -to MXN   # Convert 1 EUR to MXN
```
