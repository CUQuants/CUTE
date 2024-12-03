
# Usage Guide

## Getting Started

To start using CUTE, first import the package and create an instance of the `CUTE` class. You’ll need to set parameters like the interval for each trading step and the start time for the simulation.

### Example Program

Below is an example program that demonstrates how to simulate random buying and selling of Apple (AAPL) stock every 5 minutes.

```python
from CUTE import CUTE
import random
import time

# Create an instance of the CUTE class
cute = CUTE(
    step_interval_seconds=5*60,  # Each "step" advances the virtual clock by 5 minutes
    start_time_unix=1673560800   # Unix timestamp for Jan 12, 2023, 22:00:00 UTC
)

while True:
    action = random.randint(0, 2)
    if action == 0:
        print('Buying 1 share of AAPL')
        cute.buy('AAPL', 1)
    elif action == 1:
        print('Selling 1 share of AAPL')
        cute.sell('AAPL', 1)

    # Proceed to the next step in time (5 minutes later)
    cute.step()

    # Sleep for a short period to simulate time pass and respect API rate limits
    time.sleep(4)
```

### Key Concepts

- **Step Interval**: Controls how much virtual time progresses with each step. It’s best set to multiples of 60 seconds (e.g., 300 for 5 minutes) for optimal results.

- **Start Time**: Specifies the initial time using a Unix timestamp. Use tools like [unixtimestamp.com](https://www.unixtimestamp.com/) to find specific timestamps.

### Methods Overview

#### `buy(symbol: str, quantity: int)`

Executes a simulated purchase of the specified quantity of stock for the given symbol.

#### `sell(symbol: str, quantity: int)`

Executes a simulated sale of the specified quantity of stock for the given symbol.

#### `step()`

Advances the simulation's internal clock by the defined step interval.

#### `fetch_stock_data(symbol: str, interval: str, n: int)`

Fetches recent stock data for a specified symbol. This method requires the stock symbol, the data interval (e.g., '1min'), and the number of data points (`n`) to retrieve. Here's an example of fetching AAPL data:

```python
response = cute.fetch_stock_data('AAPL', '1min', 5)
print(response.text)
```

**Example JSON Response:**

```json
{
  "meta": {
    "symbol": "AAPL",
    "interval": "1min",
    "currency": "USD",
    "exchange_timezone": "America/New_York",
    "exchange": "NASDAQ",
    "mic_code": "XNGS",
    "type": "Common Stock"
  },
  "values": [
    {
      "datetime": "2023-01-12 15:00:00",
      "open": "133.355",
      "high": "133.36",
      "low": "133.23",
      "close": "133.27",
      "volume": "91833"
    },
    {
      "datetime": "2023-01-12 14:59:00",
      "open": "133.28",
      "high": "133.37",
      "low": "133.2501",
      "close": "133.3501",
      "volume": "96452"
    },
    {
      "datetime": "2023-01-12 14:58:00",
      "open": "133.16",
      "high": "133.3",
      "low": "133.155",
      "close": "133.295",
      "volume": "66312"
    },
    {
      "datetime": "2023-01-12 14:57:00",
      "open": "133.18",
      "high": "133.2601",
      "low": "133.1",
      "close": "133.14",
      "volume": "95919"
    },
    {
      "datetime": "2023-01-12 14:56:00",
      "open": "133.15",
      "high": "133.25",
      "low": "133.08",
      "close": "133.19",
      "volume": "116493"
    }
  ],
  "code": 0,
  "message": "",
  "status": "ok"
}
```

#### `get_portfolio()`

Retrieve your current virtual portfolio. The `get_portfolio` method returns a JSON object that includes your account balance and the details of your holdings.

```python
response = cute.get_portfolio()
print(response.text)
```

**Example JSON Response:**

```json
{
  "balance": -133.17,
  "portfolio": {
    "AAPL": 1
  }
}
```

- **balance**: Represents the virtual cash balance, which can go negative if the value of purchased stocks exceeds deposits.
- **portfolio**: A dictionary where keys are stock symbols and values are the quantity of shares owned.
