# Python program to convert 
# JSON file to CSV 


import json 
# import csv 
import pandas as pd
# Opening JSON file and loading the data 
# into the variable data 
with open('tchaikovsky.json') as json_file: 
	data = json.load(json_file) 

csvd = pd.json_normalize(data)
# print(csvd.to_csv())

# # now we will open a file for writing 
data_file = open('data_file.csv', 'w') 
data_file.write(csvd.to_csv())

# # create the csv writer object 
# csv_writer = csv.writer(data_file) 

# # Counter variable used for writing 
# # headers to the CSV file 
# count = 0

# for emp in data: 
# 	header = emp["Instrs"]
# 	brass = emp["Instrs"]["brass"]
# 	if count == 0: 

# 		# Writing headers of CSV file 

# 		csv_writer.writerows(header.keys(),brass.keys())
# 		count += 1

# 	# Writing data of CSV file 
# 	csv_writer.writerows(header.values(),brass.values())

data_file.close() 
