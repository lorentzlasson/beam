import os
import psycopg2 as psycopg2
import simplejson as json

try:
  from SimpleHTTPServer import SimpleHTTPRequestHandler as Handler
  from SocketServer import TCPServer as Server
except ImportError:
  from http.server import SimpleHTTPRequestHandler as Handler
  from http.server import HTTPServer as Server

# Read port selected by the cloud for our application
PORT = int(os.getenv('VCAP_APP_PORT', 8000))
# Change current directory to avoid exposure of control files
os.chdir('static')


httpd = Server(("", PORT), Handler)


def connectToDb():
    db_url = ""
    if 'VCAP_SERVICES' in os.environ:
        data = json.loads(os.environ['VCAP_SERVICES'])
        db_url = data["elephantsql"][0]["credentials"]["uri"]
        print (db_url)
    if db_url is "":
        with open('VCAP_SERVICES.json') as data_file:
            data = json.load(data_file)
            db_url = data["VCAP_SERVICES"]["elephantsql"][0]["credentials"]["uri"]
            print (db_url)
    try:
        conn = psycopg2.connect(db_url)
        print "connected"
    except:
        print "I am unable to connect to the database"
    try:
        cur = conn.cursor()
        cur.execute("SELECT * from \"user\"")
        rows = cur.fetchall()
        print "\nShow me the databases:\n"
        for row in rows:
            print "   ", row[0]
    except:
        print "Can't display result"


try:
  print("Start serving at port %i" % PORT)
  connectToDb()
  httpd.serve_forever()
except KeyboardInterrupt:
  pass
httpd.server_close()


