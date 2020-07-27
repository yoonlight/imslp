import json 
import pandas as pd

with open('tchaikovsky.json') as json_file: 
	data = json.load(json_file) 

csvd = pd.json_normalize(data)

data_file = open('data_file.csv', 'w') 
data_file.write(csvd.to_csv())
data_file.close() 
