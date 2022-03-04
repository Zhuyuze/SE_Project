var sign_up = require('./login.js');
var expect = require('chai').expect;

describe('login test', function() {
it('login succuess', function() {
expect(sign_up("Roman", "11")).to.be.equal("username already exists");
});
it('login fail', function() {
expect(sign_up("roman", "dsafasd")).to.be.equal("Success");
});
});