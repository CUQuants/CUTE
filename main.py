from AssetManager import AssetManager

def main():
    manager = AssetManager()
    while True:
        cmd = input("Command: ")
        if "BUY" == cmd.upper():
            ticker = input("Ticker: ").upper()
            qty = int(input("Quantity: "))

            res = manager.buy(ticker, qty)
            if res['res'] == 200:
                 manager.printBalance()
        if "SELL" == cmd.upper():
            ticker = input("Ticker: ").upper()
            qty = int(input("Quantity: "))
            res = manager.sell(ticker, qty)
            if res['res'] == 200:
                 manager.printBalance()
        if "PORTFOLIO" == cmd.upper():
            sum = 0
            manager.printPortfolio()

if __name__ == "__main__":
    main()
