class Metric:
  def __init___(self, ticker):
    self.ticker = ticker

  def get_pnl(self, intil_price, curr_price):
    return curr_price/initl_price
