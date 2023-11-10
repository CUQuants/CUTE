from AssetManager import AssetManager

if __name__ == '__main__':
    aman = AssetManager(2000)
    aman.buy('GOOG', 10)
    print(aman.balance)