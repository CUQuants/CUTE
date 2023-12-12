class Metric:
    def __init__(self,balance):
        self.balance = balance

    def pnl(self,curr_price,init_price,pnl_type='DOLLAR'):
        if pnl_type.upper() == 'DOLLAR':
            return curr_price-init_price
        
        elif pnl_type.upper() == 'PERCENT':
            return (curr_price/init_price)*100
        
        print("!ERROR Invalid PNL Type For portfolio_pnl())")
        return 0
    
    def portfolio_pnl(self,init_bal=2000,bal_type = 'DOLLAR'):
        return self.pnl(self.balance,init_bal)
    
