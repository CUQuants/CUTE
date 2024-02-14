import yfinance as yf
import pandas as pd
import numpy as np

class PriceTracker:
    def __init__(self, period = '1y', interval = '1h', simulated_time = "2023-05-14 12:30:00-05:00"):
        self.period = period
        self.interval = interval
        self.simulated_time = simulated_time

        self.history: pd.DataFrame = None


    def get_price_history(self, ticker: str):
        if self.simulated_time == -1:
            selected = yf.Ticker(ticker)
            history = selected.history(period=self.period, interval=self.interval)
            return history
        elif self.history == None:
            selected = yf.Ticker(ticker)
            self.history = selected.history(period=self.period, interval=self.interval)

        # history = self.history.filter
        return self.history.loc[:self.simulated_time]

if __name__ == "__main__":
    a = PriceTracker()
    a.get_price_history("BLK")
