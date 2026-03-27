import http from 'k6/http';
import { check } from 'k6';

export const options = {
  vus: 1,
  iterations: 5,
};

const baseURL = __ENV.BASE_URL || 'http://localhost:8082';

export default function () {
  const res = http.get(`${baseURL}/items/item-1`);
  check(res, {
    'status is 200': (r) => r.status === 200,
  });
}
