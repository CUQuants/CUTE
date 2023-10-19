import yfinance as yf
import time
import pandas as pd

# ticker = input("Ticker: ")
ticker = "GOOG"

while(True):
    time.sleep(1)
    selected = yf.Ticker(ticker)
    finfo = selected.fast_info
    print(f'{time.time()} {selected.ticker}: {finfo.last_price}')
