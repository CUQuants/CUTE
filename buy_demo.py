import yfinance as yf
import time

import json
import os
import time

def empty_ticker_history():
    return {
        'curr_qty': 0,
        'curr_avg_price': 0,
        'history': []
    }



data_file = 'Data/data.json'
data = {}


def write_data():
    with open(data_file, "w") as write_file:
        json.dump(data, write_file)


if os.path.exists(data_file):
    data = json.load(open(data_file))
else:
    print("Data file not found, creating")
    write_data()

while (True):
    ticker = input("Ticker: ")
    buy_qty = int(input("Buy Quantity: "))

    if ticker not in data:
        data[ticker] = empty_ticker_history()

    selected = yf.Ticker(ticker)
    finfo = selected.fast_info
    new_price = finfo.last_price

    #calucluate curr_avg_price using a weighted moving average
    data[ticker]['curr_avg_price'] = (data[ticker]['curr_avg_price'] * data[ticker][
        'curr_qty'] + buy_qty * new_price) / (data[ticker]['curr_qty'] + buy_qty)
    data[ticker]['curr_qty'] += buy_qty

    trade = {
        'size': buy_qty,
        'price': new_price,
        'timestampt': time.time()
    }

    data[ticker]['history'].append(trade)

    write_data()
