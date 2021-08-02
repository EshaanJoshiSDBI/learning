import mysql.connector
from datetime import date
import datetime
db = mysql.connector.connect(
    host = '192.168.0.110',
    user = 'eshu',
    passwd = 'Local#1dbMySQL',
    database = 'sample'
)
mycursor = db.cursor()
mycursor.execute('USE employees')
mycursor.execute('SELECT * FROM departments')
print(mycursor.fetchall()) 