var jwt = require('jsonwebtoken');
var commandLineArgs = require('command-line-args');

function getClaimsFromArgs(claimsArgs) {
    var claims = []
    var claimValues = []
    for(i = 0; i < claimsArgs.length;i++) {
        if(i % 2 == 0) {
            claims.push(claimsArgs[i]);
        } else {
            claimValues.push(claimsArgs[i]);
        }
    }

    var tokenContext = {}
    for(i = 0; i < claims.length; i++) {
        if(i < claimValues.length) {
            tokenContext[claims[i]] = claimValues[i];
        } else {
            tokenContext[claims[i]] = '';
        }
    }

    return tokenContext;
}

function checkArgs(options,usage) {
    if(options.secret == undefined) {
        console.log(usage);
        process.exit(1);
    }
}

function makeTokenWithClaims(options) {
    claimsObj = getClaimsFromArgs(options.claims)
    var decodedSecret = new Buffer(options.secret, 'base64')
    var token = jwt.sign(claimsObj, decodedSecret);
    return token;
}


var cli = commandLineArgs([
  { name: 'secret', alias: 's', type: String },
  { name: 'claims', alias: 'c', type: String, multiple: true },
])

var options = cli.parse();
var usage = cli.getUsage();

checkArgs(options,usage);

var token = makeTokenWithClaims(options);
console.log(token);