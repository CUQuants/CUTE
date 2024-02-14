# CUTE
CU Quantitative Trading Emporium

Portfolio Manager For Algo Trading


## Creating a Strategy
```python
Class STRATEGY_NAME
    def step(manager: AssetManager):
        #...
        # reponse is a json with the the structure json = {
        #     'buy': {
        #         'TICKER':QUANTITY
        #     },
        #     'sell': {
        #         'TICKER':QUANTITY
        #     },
        # }
        return response
```
