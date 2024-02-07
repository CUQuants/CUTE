import os
import json
import time
from typing import Dict, List, Any

import yfinance as yf


class Response:
    def __init__(self, res: int, message: str = ""):
        self.res: int = res
        self.message: str = message
        

class AssetManager:
    def __init__(self, data_file = 'data.json'):
        self.data_file = data_file
        self.data = {}

        if os.path.exists(data_file):
            self.data = json.load(open(data_file))
            self.balance = self.data['balance']
        else:
            self.write_data()
            self.balance = 2000
            print("Data file not found, creating new data.json with balance ", self.balance)

    def empty_ticker_history(self)-> dict[str, int | list[Any]]:
        return {
            'curr_qty': 0,
            'curr_avg_price': 0,
            'history': []
        }

    def write_data(self)->None:
        with open(self.data_file, "w") as write_file:
            json.dump(self.data, write_file)

    def get_price(self, ticker) -> float:
        selected = yf.Ticker(ticker)
        finfo = selected.fast_info
        return finfo.last_price

    def modify_holdings(self, ticker, quantity, price)->Response:
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
        return Response(200)

    def buy(self, ticker, quantity)->Response:
        price = self.get_price(ticker)

        if(self.balance < price*quantity):
            raise Exception('Not enough funds')

        self.modify_holdings(ticker, quantity, price)


        return Response(200)


    def sell(self, ticker: str, quantity: float)->Response:

        price = self.get_price(ticker)
        
        if(ticker not in self.data['stocks'] or self.data["stocks"][ticker]['curr_qty'] < quantity):
            raise Exception('Not enough holdings of', ticker, ". You have", self.data["stocks"][ticker]['curr_qty'])


        self.modify_holdings(ticker, -quantity, price)


        return Response(200)


    def printPortfolio(self)->None:
        sum = 0

        for t in self.data["stocks"]:
            stock = self.data["stocks"][t]
            if(stock['curr_qty'] > 0):
                curr_price = self.get_price(t)
                sum += curr_price * stock['curr_qty']
                print(t, stock["curr_qty"], "shares @ " + str(curr_price), "= $" + str(stock['curr_qty'] * curr_price))

        print("Total Holdings: $" + str(sum))
        print("Current Balance: $" + str(self.data['balance']))


    def printBalance(self)->None:
        print("Your current balance is $" + str(self.data['balance']))
