import os
import psycopg2 as psycopg2
import simplejson as json
import urlparse # import urllib.parse for python 3+
from flask import Flask, jsonify, render_template, request, redirect, url_for, send_from_directory
from psycopg2.extras import RealDictCursor
# import pyodbc

app = Flask(__name__, static_url_path = "")
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
@app.route('/')
def index():
    return app.send_static_file('index.html')

@app.route('/user/<username>')
def show_user_profile(username):
    # show the user profile for that user
    return 'User %s' % username

@app.route('/post/<int:post_id>')
def show_post(post_id):
    # show the post with the given id, the id is an integer
    return 'Post %d' % post_id

@app.route('/getUsers')
def APIgetUsers():
    conn = connectToDb()
    data = getDataFromDB(conn, getAllFromUser)
    # return jsonify({'users': data})
    return data
@app.route('/getBeacons')
def APIgetBeacons():
    conn = connectToDb()
    data = getDataFromDB(conn, getAllFromBeacon)
    # return jsonify({'beacons': data})
    return data

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
        cur = conn.cursor(cursor_factory=RealDictCursor)
        cur.execute(sql())
        colnames = [desc[0] for desc in cur.description]
        print colnames
        return json.dumps(cur.fetchall(), indent=2)
        rows = cur.fetchall()
        return rows
    except:
        print "Can't display result"

# RUN APPLICATION #

if __name__ == '__main__':
    print environment
    if environment=='local':
        app.run(debug=True)
    if environment=='bluemix':
        app.run(host='0.0.0.0', port=PORT)
