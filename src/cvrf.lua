return {
	get = function(a)
		local http = require("http")
		local json = require("json")
		local options = {
			timeout = "30s",
			headers = { Accept = "application/json" },
		}
		local url = ("https://api.msrc.microsoft.com/cvrf/v2.0/cvrf/%s"):format(a.id)
		local req, err = http.get(url, options)
		if not req or not req.status == 200 then
			return nil, err or "cvrf: Did not return a valid response."
		end
		return json.decode(req.body)
	end
}
