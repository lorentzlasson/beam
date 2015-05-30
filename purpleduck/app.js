var express 	= require('express'),
	app     	= express(),
	ibmbluemix 	= require('ibmbluemix'),
	ibmpush		= require('ibmpush'),
	ibmdata     = require('ibmdata');

var masconfig = require('./mas.json');

ibmbluemix.initialize(masconfig);
var logger = ibmbluemix.getLogger();

app.get('/', function(req, res){
	res.sendfile('public/index.html');
});

app.use(function(req, res, next) {
    req.data = ibmdata.initializeService(req);
   	req.ibmpush = ibmpush.initializeService(req);
    req.logger = logger;
    console.log("init")
    next();
});


app.use(require('./lib/setup'));

//uncomment below code to protect endpoints created afterwards by MAS
//var mas = require('ibmsecurity')();
//app.use(mas);

var ibmconfig = ibmbluemix.getConfig();

logger.info('mbaas context root: '+ibmconfig.getContextRoot());

app.use(ibmconfig.getContextRoot(), require('./lib/accounts'));
app.use(ibmconfig.getContextRoot(), require('./lib/staticfile'));

app.get(ibmconfig.getContextRoot(), function(req, res){
	var instance_count = JSON.parse(process.env.INSTANCE_COUNT || 0);
	res.send({count:instance_count});
})

app.listen(ibmconfig.getPort());
logger.info('Server started at port: '+ibmconfig.getPort());