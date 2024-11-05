import requests

class CUTE:
    def __init__(self, url: str = 'http://localhost:3000', step_interval_seconds: int = 60*5, start_time_unix: int = 1673560800):
        assert step_interval_seconds >= 60
        self.version = '0.0.1'
        self.url = url
        self.step_interval_seconds = step_interval_seconds
        self.start_time_unix = start_time_unix
        data = {
            "clientVersion": self.version,
            "currentTimeUnix": start_time_unix,
            "stepIntervalSeconds": step_interval_seconds
        }
        try:
            response = requests.post(self.url + '/init', json=data)
            if response.status_code != 200:
                print(response.text)
        except Exception as e:
            print(e)
            print('Make sure CUTE is running!')
            quit()

    def buy(self, symbol: str, quantity: int):
        assert quantity > 0
        data = {
            "symbol": symbol,
            "quantity": quantity,
        }
        response = requests.post(self.url + '/buy', json=data)
        if response.status_code != 200:
            print(response.text)
        return response

    def sell(self, symbol: str, quantity: int):
        assert quantity > 0
        data = {
            "symbol": symbol,
            "quantity": quantity,
        }
        response = requests.post(self.url + '/sell', json=data)
        if response.status_code != 200:
            print(response.text)
        return response

    def step(self):
        response = requests.post(self.url + '/step')
        if response.status_code != 200:
            print(response.text)
        return response


    def fetch_stock_data(self, symbol:str, interval:str, n: int):
        assert n > 0
        print('Warning: fetch_stock_data is not fully implemented')
        data = {
            "symbol": symbol,
            "interval": interval,
            "n":'n'
        }
        response = requests.post(self.url + '/fetch-stock-data', json=data)
        if response.status_code != 200:
            print(response.text)
        return response

    def get_portfolio(self):
        print('Warning: get-portfolio is not fully implemented')
        response = requests.post(self.url + '/get-portfolio', json={})
        if response.status_code != 200:
            print(response.text)
        return response
