To view/interact with the API, you'll need to run the swagger ui on the
same (local) host as app.go.

Use this url in your swagger ui 'explore' text box:

http://localhost:3000/apispec/api-docs/quote/quote.json

You can serve up the swagger ui on the mac by cloning the project, then
running starting a web service in the dist directory using
`python -m SimpleHTTPServer 8000`

You can also host the swagger document in the swagger dist itself by adding a host entry that points to the 
host and port of the endpoint you want. Obviously you still need the CORS spec to enable the communication.

For example:

<pre>
{
  "swagger": "2.0",
  "host":"localhost:3000",
  "info": {
    "version": "0.0.0",
    "title": "Stock Quote"
  },
</pre>


