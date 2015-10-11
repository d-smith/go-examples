Running the lambda sample

* cross compile the exe, e.g. `GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build`
* create a zip file that includes the executable (lambda) and index.js
* upload the code via the AWS console. Here I pretty much just used the defaults. Key was to have the index.js and lambda
in the root of the zip, and note the alignment between the Handler value (index.handler) and the js file (index.js)

The function can be installed via the AWS console or the AWS command line.

<pre>
aws lambda create-function --function-name gopher-cli --runtime nodejs --role arn:aws:iam::xxxx:role/lambda_basic_execution --handler index.handler --zip-file fileb://Archive.zip
</pre>

The command line can also be used to invoke the function.

<pre>
aws lambda invoke --function-name gopher-cli --payload '{"s1":"aaaaa"}' --log-type Tail outfile
</pre>

