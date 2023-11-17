from AssetManager import AssetManager

if __name__ == '__main__':
    aman = AssetManager(2000)
    print(aman.balance)

    aman.buy('GOOG', 10)
    print(aman.balance)
    aman.sell('GOOG', 5)
    print(aman.balance)
    # aman.sell('GS', 10)

