import os
import psycopg2 as psycopg2
import simplejson as json
import urlparse # import urllib.parse for python 3+
from flask import Flask, jsonify
app = Flask(__name__)
environment = ""

# INIT #
if 'VCAP_SERVICES' in os.environ:
    environment = "bluemix"
else:
    environment = "local"    
# Read port selected by the cloud for our application
PORT = int(os.getenv('VCAP_APP_PORT', 8000))

# SQL #
def getAllFromUser():
     return "SELECT * from \"user\""   
def getAllFromBeacon():
     return "SELECT * from \"beacon\""   

# API #
@app.route('/getUsers')
def APIgetUsers():
    conn = connectToDb()
    data = getDataFromDB(conn, getAllFromUser)
    return jsonify({'users': data})
@app.route('/getBeacons')
def APIgetBeacons():
    conn = connectToDb()
    data = getDataFromDB(conn, getAllFromBeacon)
    return jsonify({'beacons': data})

# METHODS #
def connectToDb():
    if environment=='bluemix':
        data = json.loads(os.environ['VCAP_SERVICES'])
        db_url = data["elephantsql"][0]["credentials"]["uri"]
    elif environment=='local':
        with open('VCAP_SERVICES.json') as data_file:
            data = json.load(data_file)
            db_url = data["VCAP_SERVICES"]["elephantsql"][0]["credentials"]["uri"]
    result = urlparse.urlparse(db_url)
    username = result.username
    password = result.password
    database = result.path[1:]
    hostname = result.hostname
    try:
        conn = psycopg2.connect(
        database = database,
        user = username,
        password = password,
        host = hostname
        )
        print "connected"
        return conn
    except Exception as ex:
        print "general exception:   ", ex.message
    except psycopg2.Error as e:
        print "psycopg2 error code: ", e.pgcode
        print "psycopg2 error msg: ", e.pgerror  
        print "Unable to connect to the database"

def getDataFromDB(conn, sql):
    try:
        cur = conn.cursor()
        cur.execute(sql())
        rows = cur.fetchall()
        dataRows = []
        for row in rows:
            dataColumns = []
            for x in xrange(0,len(row)):
                dataColumns.append(row[x])
            dataRows.append(dataColumns)
        return dataRows
    except:
        print "Can't display result"

# RUN APPLICATION #

if __name__ == '__main__':
    print environment
    if environment=='local':
        app.run(debug=True)
    if environment=='bluemix':
        app.run(host='0.0.0.0', port=PORT)
