var MongoClient = require('mongodb').MongoClient;

MongoClient.connect('mongodb://127.0.0.1:27017/weather', function(err, db) {
    if(err) throw err;


    var data = db.collection('data');

    var query = {'State' : 1, 'Temperature': 1, '_id' : 0 };
    // 'title' : { '$regex' : 'Microsoft' }
    var projection = { 'State' : 1, 'Temperature': 1, '_id' : 0 };

    var cursor = data.find(query);
    // query
    // cursor.skip(0);
    cursor.limit(10000);
    cursor.sort([['State', 1], ['Temperature', -1]])

    cursor.each(function(err, doc) {
        if(err) throw err;
        if(doc == null) {
            return db.close();
        }
        // while cursor.hasNext(){
        //     console.dir(doc);
        // }
        console.dir(doc);
    });

    // var query = { 'media.oembed.type' : 'video' };

    // var projection = { 'media.oembed.url' : 1, '_id' : 0 };

    // db.collection('things').find(query, projection).each(function(err, doc) {
    //     if(err) throw err;

    //     if(doc == null) {
    //         return db.close();
    //     }

    //     console.dir(doc);
    // });
});