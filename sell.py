import yfinance as yf

balance = 2000

holdings = 20

def sell(quantity: float, ticker: str, holdings: float, balance: float):
    yf_ticker = yf.Ticker(ticker)
    data = yf_ticker.history()
    last_quote = data['Close'].iloc[-1]
    print(ticker, last_quote)

    new_holdings = holdings - quantity

    if (new_holdings < 0):
        return holdings, balance
    
    balance += quantity * last_quote

    return new_holdings, balance

print(f"Old holdings and balance: {(holdings, balance)}")

print(f"New holdings and balance: {sell(2, 'GOOG', holdings, balance)}")

