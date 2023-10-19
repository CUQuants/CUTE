import yfinance as yf
import time

import json
import pandas as pd

# ticker = input("Ticker: ")
ticker = "GOOG"

data = json.load(open('Data/data.json'))


while(True):
    time.sleep(1)
    selected = yf.Ticker(ticker)
    finfo = selected.fast_info
    print(f'{time.time()} {selected.ticker}: {finfo.last_price}')


