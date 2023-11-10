import os
import json
import time
import yfinance as yf

class AssetManager:
    def __init__(self, balance, data_file = 'Data/data.json'):
        self.balance = balance
        self.data_file = 'Data/data.json'
        self.data = {}

        if os.path.exists(data_file):
            data = json.load(open(data_file))
        else:
            print("Data file not found, creating")
            self.write_data()
    def create_code(self, res, message=''):
        return {
            'res': res,
            'message': message
        }
    def empty_ticker_history(self):
        return {
            'curr_qty': 0,
            'curr_avg_price': 0,
            'history': []
        }

    def write_data(self):
        with open(self.data_file, "w") as write_file:
            json.dump(self.data, write_file)

    def get_price(self, ticker):
        selected = yf.Ticker(ticker)
        finfo = selected.fast_info
        return finfo.last_price

    def modify_holdings(self, ticker, quantity, price):
        if ticker not in self.data:
            self.data[ticker] = self.empty_ticker_history()

        if self.data[ticker]['curr_qty'] + quantity < 0:
            return self.create_code(400, ("Can't sell more stocks than in possesion"))

        # calucluate curr_avg_price using a weighted moving average
        self.data[ticker]['curr_avg_price'] = (self.data[ticker]['curr_avg_price'] * self.data[ticker][
            'curr_qty'] + quantity * price) / (self.data[ticker]['curr_qty'] + quantity)
        self.data[ticker]['curr_qty'] += quantity

        trade = {
            'size': quantity,
            'price': price,
            'timestampt': time.time()
        }

        self.data[ticker]['history'].append(trade)

        self.write_data()
        self.balence += quantity*price

    def buy(self, ticker, quantity):
        price = self.get_price(ticker)

        if(self.balance < price*quantity):
            raise Exception('Not enough funds')

        self.modify_holdings(ticker, quantity, price)

        return self.create_code(200)


    def sell(self, quantity: float, ticker: str, balance: float):
        price = self.get_price(ticker)

        self.modify_holdings(ticker, -quantity, price)
    
        return self.create_code(200)
    
        # print(f"Old holdings and balance: {(holdings, balance)}")
        
        # print(f"New holdings and balance: {sell(2, 'GOOG', holdings, balance)}")
