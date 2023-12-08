from AssetManager import AssetManager
from Metric import Metric

def main():
    manager = AssetManager()
    metric = Metric()
    
    while True:
        cmd = input("Command: ")
        if "BUY" == cmd.upper():
            ticker = input("Ticker: ").upper()
            qty = int(input("Quantity: "))

            res = manager.buy(ticker, qty)
            if res['res'] == 200:
                 manager.printBalance()
        elif "SELL" == cmd.upper():
            ticker = input("Ticker: ").upper()
            qty = int(input("Quantity: "))
            res = manager.sell(ticker, qty)
            if res['res'] == 200:
                 manager.printBalance()
        elif "PORTFOLIO" == cmd.upper():
            sum = 0
            manager.printPortfolio()

        else:
            print("!ERROR! Command Not Found.")

if __name__ == "__main__":
    main()
