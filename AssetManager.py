import os
import json
import time
import yfinance as yf

class AssetManager:
    def __init__(self, data_file = 'Data/data.json'):
        self.data_file = 'Data/data.json'
        self.data = {
            'balance': 2000,
            'stocks': {}
        }

        if os.path.exists(data_file):
            self.data = json.load(open(data_file))
            if ('balance' not in self.data ):
                raise Exception("No balance attribute in data file, please fix before continuing, or delete data file and rerun program.")
            self.balance = self.data['balance']
        else:
            self.write_data()
            self.balance = 2000
            print("Data file not found, creating new data.json with balance ", self.balance)

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

    def get_price(self, ticker) -> float:
        selected = yf.Ticker(ticker)
        finfo = selected.fast_info
        return finfo.last_price

    def modify_holdings(self, ticker, quantity, price):
        if ticker not in self.data['stocks']:
            self.data["stocks"][ticker] = self.empty_ticker_history()

        if self.data["stocks"][ticker]['curr_qty'] + quantity < 0:
            raise Exception("Can't sell more stocks than in possesion")

        # calucluate curr_avg_price using a weighted moving average
        # self.data["stocks"][ticker]['curr_avg_price'] = (self.data["stocks"][ticker]['curr_avg_price'] * self.data["stocks"][ticker][
        #     'curr_qty'] + quantity * price) / (self.data["stocks"][ticker]['curr_qty'] + quantity)
        self.data["stocks"][ticker]['curr_qty'] += quantity

        self.data['balance'] -= quantity*price

        trade = {
            'qty': quantity,
            'price': price,
            'timestamp': time.time()
        }

        self.data["stocks"][ticker]['history'].append(trade)

        self.write_data()

    def buy(self, ticker, quantity):
        price = self.get_price(ticker)

        if(self.balance < price*quantity):
            raise Exception('Not enough funds')

        self.modify_holdings(ticker, quantity, price)


        return self.create_code(200)


    def sell(self, ticker: str, quantity: float):

        price = self.get_price(ticker)
        
        if(ticker not in self.data['stocks'] or self.data["stocks"][ticker]['curr_qty'] < quantity):
            raise Exception('Not enough holdings of', ticker, ". You have", self.data["stocks"][ticker]['curr_qty'])


        self.modify_holdings(ticker, -quantity, price)


        return self.create_code(200)


    def printPortfolio(self):
        sum = 0

        for t in self.data["stocks"]:
            stock = self.data["stocks"][t]
            if(stock['curr_qty'] > 0):
                curr_price = self.get_price(t)
                sum += curr_price * stock['curr_qty']
                print(t, stock["curr_qty"], "shares @ " + str(curr_price), "= $" + str(stock['curr_qty'] * curr_price))

        print("Total Holdings: $" + str(sum))
        print("Current Balance: $" + str(self.data['balance']))


    def printBalance(self):
        print("Your current balance is $" + str(self.data['balance']))