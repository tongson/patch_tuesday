return {
	get = function(a)
		local http = require("http")
		local json = require("json")
		local options = {
			timeout = "30s",
			headers = { Accept = "application/json" },
		}
		local url = ("https://api.msrc.microsoft.com/cvrf/v2.0/cvrf/%s"):format(a.id)
		local req = http.get(url, options)
		return json.decode(req.body)
	end
}
