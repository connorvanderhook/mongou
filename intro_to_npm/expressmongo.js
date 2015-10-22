var express = require('express'),
    app = express(),
    cons = require('consolidate'),
    MongoClient = require('mongodb').MongoClient,
    Server = require('mongodb').Server;

app.engine('html', cons.swig);
app.set('view engine', 'html');
app.set('views', __dirname + '/views');

var mongoclient = new MongoClient(new Server("127.0.0.1", 27017,
                                                    { 'native_parser' : true}));
var db = mongoclient.db('test');

app.get('/', function(req, res){

    // Find one document in our collection
    db.collection('names').findOne({}, function(err, doc) {

        if(err) throw err;

        res.render('hello', doc);
    });
});

app.get('*', function(req, res){
    res.send('Page Not Found', 404);
});

mongoclient.open(function(err, mongoclient) {

    if(err) throw err;

    app.listen(8080);
    console.log('Express server started on port 8080');
});
