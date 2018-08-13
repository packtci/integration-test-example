'use strict';

const request = require('supertest');
const assert = require('assert');

const CIRCLECI_API = {
    // List of all the projects you're following on CircleCI, with build information organized by branch
    getProjects: 'https://circleci.com/api/v1.1'
};

describe('Testing CircleCI API Endpoints', function() {
    it('the /projects endpoints should return 200 with a body', function() {
        return request(CIRCLECI_API.getProjects)
            .get(`/projects?circle-token=${process.env.CIRCLECI_API_TOKEN_GITHUB}`)
            .set('Accept', 'application/json')
            .expect(200)
            .then(response => {
                assert.ok(response.body.length > 0, "Body have information")
                assert.equal(response.body[0].oss, true);
            });
    });
});