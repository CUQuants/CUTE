import os
import json
import time
import yfinance as yf

class AssetManager():
    def __int__(self, balance, data_file = 'Data/data.json'):
        self.balance = balance
        self.data_file = 'Data/data.json'
        self.data = {}

        if os.path.exists(data_file):
            data = json.load(open(data_file))
        else:
            print("Data file not found, creating")
            self.write_data()

    def empty_ticker_history(self):
        return {
            'curr_qty': 0,
            'curr_avg_price': 0,
            'history': []
        }

    def write_data(self):
        with open(self.data_file, "w") as write_file:
            json.dump(self.data, write_file)

    def buy(self, ticker, quantity):
        if ticker not in self.data:
            self.data[ticker] = self.empty_ticker_history()

        selected = yf.Ticker(ticker)
        finfo = selected.fast_info
        new_price = finfo.last_price

        # calucluate curr_avg_price using a weighted moving average
        self.data[ticker]['curr_avg_price'] = (self.data[ticker]['curr_avg_price'] * self.data[ticker][
            'curr_qty'] + quantity * new_price) / (self.data[ticker]['curr_qty'] + quantity)
        self.data[ticker]['curr_qty'] += quantity

        trade = {
            'size': quantity,
            'price': new_price,
            'timestampt': time.time()
        }

        self.data[ticker]['history'].append(trade)

        self.write_data()

    def sell(self, quantity: float, ticker: str, holdings: float, balance: float):
        yf_ticker = yf.Ticker(ticker)
        data = yf_ticker.history()
        last_quote = data['Close'].iloc[-1]
        #print(ticker, last_quote)
    
        new_holdings = holdings - quantity
    
        if (new_holdings < 0):
            return holdings, balance
        
        balance += quantity * last_quote
    
        return new_holdings, balance
    
        # print(f"Old holdings and balance: {(holdings, balance)}")
        
        # print(f"New holdings and balance: {sell(2, 'GOOG', holdings, balance)}")
