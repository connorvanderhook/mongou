db.clients.find({}).forEach( function(client) 
	{ print(client["_id"]); 
	client["postJobs"].forEach( function(pj) {
			if(!pj["enabled"]) { 
					print(pj["enabled"]); 
				}, 
			}
		});  
})