import json 
import pandas as pd

with open('sibelius.json', encoding='UTF8') as json_file: 
	data = json.load(json_file) 

for element in data:
    element.pop('instrument', None)
csvd = pd.json_normalize(data)

data_file = open('sibelius.csv', 'w') 
data_file.write(csvd.to_csv())
data_file.close() 
