var sign_in = require('./login.js');
var expect = require('chai').expect;

describe('login test', function() {
it('login succuess', function() {
expect(sign_in("Roman", "11")).to.be.equal("Success");
});
it('login fail', function() {
expect(sign_in("roman", "dsafasd")).to.be.equal("Wrong password");
});
});