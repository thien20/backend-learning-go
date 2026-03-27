import http from 'k6/http';

export const options = {
  stages: [
    { duration: '10s', target: 10 },
    { duration: '10s', target: 100 },
    { duration: '10s', target: 10 },
  ],
};

const baseURL = __ENV.BASE_URL || 'http://localhost:8082';

export default function () {
  http.get(`${baseURL}/items?view=detailed`);
}
